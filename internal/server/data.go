package server

import (
	"cmdTest/internal/dto/model"
	"cmdTest/internal/dto/mysql"
	"cmdTest/internal/serror"
	"cmdTest/pkg/util"
	"fmt"
	"github.com/gocarina/gocsv"
	"mime/multipart"
	"os"
	"strconv"
	"sync"
)

type InFunction func(int64, int64) (*model.RangeData, error)
type CloseFunction func() error
type useCnt struct {
	cnt int
	sync.Mutex
}

var (
	forecastPath = "./static/forecast/"
	truePath     = "./static/true/"
	dataDaoMysql = mysql.DataDaoMysql{}
	inMap        = make(map[string]InFunction)
	hMap         = make(map[string]*model.DataHistory)
	cMap         = make(map[string]CloseFunction)
	cntMap       = make(map[string]*useCnt)
	mutex        = sync.Mutex{}
)

func NewForecastData(id string, file *multipart.FileHeader) error {
	path := forecastPath + file.Filename
	err := util.SaveFile(path, file)
	if err != nil {
		return err
	}

	return dataDaoMysql.SaveForecastDataHistory(id, file.Filename)
}

func GetHistoryData() ([]model.DataHistoryJson, error) {
	return dataDaoMysql.GetHistoryData()
}

func DeleteHistoryData(id string) error {
	tp, fp, err := dataDaoMysql.DeleteHistoryData(id)
	if err != nil {
		return err
	}

	tp, fp = truePath+tp, forecastPath+fp

	err = util.DeleteFile(tp)
	if err != nil {
		return err
	}

	return util.DeleteFile(fp)
}

func DataDetailOpen(id string) error {
	if cnt, ok := cntMap[id]; ok && cnt.cnt != 0 {
		cntMap[id].Lock()
		cntMap[id].cnt++
		cntMap[id].Unlock()
		return nil
	} else if cnt.cnt == 0 {
		delete(cntMap, id)
		delete(cMap, id)
		delete(inMap, id)
		delete(hMap, id)
	}

	cntMap[id] = &useCnt{
		cnt:   1,
		Mutex: sync.Mutex{},
	}

	history, err := dataDaoMysql.GetOneHistory(id)
	if err != nil {
		return err
	}
	hMap[id] = history

	inF, closeF := DataDetailProcess(id)
	inMap[id] = inF
	cMap[id] = closeF

	return nil
}

func DataDetailProcess(id string) (InFunction, CloseFunction) {
	tp, pp := truePath+hMap[id].TrueDataPath, forecastPath+hMap[id].PDataPath
	fromChan, toChan := make(chan int64), make(chan int64)
	rangeDataChan, errorChan := make(chan *model.RangeData), make(chan error)
	closeChan := make(chan int)

	getInFunc := func() (from int64, to int64, c int) {
		from = <-fromChan
		to = <-toChan
		c = <-closeChan
		return
	}

	go func() {
		td, pd, err := getHistoryData(tp, pp)
		if err != nil {
			fmt.Println(err.Error())
		}

		ty, py, err := getHistoryYaliData(td, pd)
		if err != nil {
			fmt.Println(err.Error())
		}

		tSum, pSum := getHistoryYaliSum(ty, py)
		for {
			from, to, c := getInFunc()
			if c == 1 {
				break
			}
			fi, ti, err := getDataIndex(from, to, td)
			if err != nil {
				rangeDataChan <- nil
				errorChan <- err
				continue
			}

			rangeData, err := getRangeData(fi, ti, tSum, pSum)
			if err != nil {
				rangeDataChan <- nil
				errorChan <- err
				continue
			}

			rangeDataChan <- rangeData
			errorChan <- nil
		}
	}()

	return func(from int64, to int64) (*model.RangeData, error) {
			fromChan <- from
			toChan <- to
			closeChan <- 0
			err := <-errorChan
			if err != nil {
				return nil, err
			}
			rangeData := <-rangeDataChan
			return rangeData, err
		},
		func() error {
			fromChan <- -1
			toChan <- -1
			closeChan <- 1

			err := <-errorChan
			return err
		}
}

func getHistoryData(t string, p string) (td []model.Data, pd []model.Data, err error) {
	openT, err := os.OpenFile(t, os.O_RDWR|os.O_CREATE, os.ModePerm)
	if err != nil {
		return nil, nil, err
	}
	defer openT.Close()

	openP, err := os.OpenFile(p, os.O_RDWR|os.O_CREATE, os.ModePerm)
	if err != nil {
		return nil, nil, err
	}
	defer openP.Close()

	err = gocsv.UnmarshalFile(openT, &td)
	if err != nil {
		return nil, nil, err
	}

	err = gocsv.UnmarshalFile(openP, &pd)
	if err != nil {
		return nil, nil, err
	}

	return
}

func getRangeData(fi, ti int, tSum []float64, pSum []float64) (*model.RangeData, error) {
	// todo
	panic("todo")
}

func getHistoryYaliData(td, pd []model.Data) (ty []float64, py []float64, err error) {
	for _, d := range td {
		yaliString := d.Press
		yaliFloat, err := strconv.ParseFloat(yaliString, 64)
		if err != nil {
			return nil, nil, err
		}
		ty = append(ty, yaliFloat)
	}

	for _, d := range pd {
		yaliString := d.Press
		yaliFloat, err := strconv.ParseFloat(yaliString, 64)
		if err != nil {
			return nil, nil, err
		}
		ty = append(py, yaliFloat)
	}

	return
}

func getHistoryYaliSum(ty, py []float64) (tSum []float64, pSum []float64) {
	tSum = make([]float64, len(ty)+1)
	pSum = make([]float64, len(py)+1)

	for i := 1; i < len(tSum); i++ {
		tSum[i] = tSum[i-1] + ty[i-1]
	}

	for i := 1; i < len(pSum); i++ {
		pSum[i] = pSum[i-1] + py[i-1]
	}

	return
}

func getDataIndex(from, to int64, td []model.Data) (fi int, ti int, err error) {
	for i := range td {
		if td[i].ParseTime() == from {
			fi = i
		} else if td[i].ParseTime() == to {
			ti = i
		}
	}

	if fi <= ti {
		return -1, -1, serror.WrongRangeError
	}

	return
}

func GetRangeData(id, from, to string) (*model.RangeData, error) {
	f, _ := strconv.Atoi(from)
	t, _ := strconv.Atoi(to)

	inF := inMap[id]
	mutex.Lock()
	data, err := inF(int64(f), int64(t))
	mutex.Unlock()
	return data, err
}

func DataDetailClose(id string) error {
	cntMap[id].Lock()
	cntMap[id].cnt--
	if cntMap[id].cnt == 0 {
		c := cMap[id]
		err := c()
		return err
	}
	cntMap[id].Unlock()
	return nil
}
