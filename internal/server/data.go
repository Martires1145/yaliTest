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
	"time"
)

type InFunction func(int64, int64) (*model.RangeData, *model.RangeData, error)
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

func NewHistoryData(fileTrue *multipart.FileHeader, filePredict *multipart.FileHeader, history *model.DataHistoryJson) error {
	pPath := forecastPath + filePredict.Filename
	err := util.SaveFile(pPath, filePredict)
	if err != nil {
		return err
	}

	tPath := truePath + fileTrue.Filename
	err = util.SaveFile(tPath, fileTrue)
	if err != nil {
		return err
	}

	return dataDaoMysql.SaveDataHistory(&model.DataHistory{
		ModelID:       history.ModelID,
		WellID:        history.WellID,
		EngineeringID: history.EngineeringID,
		CreateTime:    time.Now().Unix(),
		TrueDataPath:  fileTrue.Filename,
		PDataPath:     filePredict.Filename,
	})
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

func DataDetailOpen(id string) (history *model.DataHistory, err error) {
	if cnt, ok := cntMap[id]; ok && cnt.cnt != 0 {
		cntMap[id].Lock()
		cntMap[id].cnt++
		cntMap[id].Unlock()
		return hMap[id], nil
	}

	cntMap[id] = &useCnt{
		cnt:   1,
		Mutex: sync.Mutex{},
	}

	history, err = dataDaoMysql.GetOneHistory(id)
	if err != nil {
		return nil, err
	}
	hMap[id] = history

	inF, closeF := DataDetailProcess(id, history.GetDay())
	inMap[id] = inF
	cMap[id] = closeF

	return history, nil
}

func DataDetailProcess(id, day string) (InFunction, CloseFunction) {
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
			return
		}

		ty, py, err := getHistoryYaliData(td, pd)
		if err != nil {
			fmt.Println(err.Error())
			return
		}

		tSum, pSum := getHistoryYaliSum(ty, py)

		tStMax, pStMax := getHistoryYaliMaxSt(ty, py)
		tStMin, pStMin := getHistoryYaliMinSt(ty, py)

		tVariance, pVariance := getHistoryYaliVariance(ty, py, tSum, pSum)
		for {
			from, to, c := getInFunc()
			if c == 1 {
				errorChan <- nil
				break
			}
			var fi, ti int
			if from == to && to == 0 {
				fi, ti = 0, len(pd)-1
			} else {
				fi, ti, err = getDataIndex(from, to, td, day)
				if err != nil {
					errorChan <- err
					continue
				}
			}
			rangeDataT, rangeDataP, err := getRangeData(
				fi, ti,
				tSum, tStMax, tStMin, tVariance,
				pSum, pStMax, pStMin, pVariance,
			)

			if err != nil {
				errorChan <- err
				continue
			}

			errorChan <- nil
			rangeDataChan <- rangeDataT
			rangeDataChan <- rangeDataP
		}
	}()

	return func(from int64, to int64) (*model.RangeData, *model.RangeData, error) {
			fromChan <- from
			toChan <- to
			closeChan <- 0
			err := <-errorChan
			if err != nil {
				return nil, nil, err
			}
			rangeDataT := <-rangeDataChan
			rangeDataP := <-rangeDataChan
			return rangeDataT, rangeDataP, err
		},
		func() error {
			fromChan <- -1
			toChan <- -1
			closeChan <- 1

			err := <-errorChan
			return err
		}
}

func getHistoryYaliMaxSt(ty, py []float64) (*util.DataRangeMax, *util.DataRangeMax) {
	return util.GetStMax(ty), util.GetStMax(py)
}

func getHistoryYaliMinSt(ty, py []float64) (*util.DataRangeMin, *util.DataRangeMin) {
	return util.GetStMin(ty), util.GetStMin(py)
}

func getHistoryYaliVariance(ty, py, tSum, pSum []float64) (*util.DataRangeVariance, *util.DataRangeVariance) {
	tSquareSum := make([]float64, len(ty)+1)
	pSquareSum := make([]float64, len(py)+1)

	for i := 1; i < len(ty); i++ {
		tSum[i] = tSum[i-1] + ty[i-1]*ty[i-1]
	}

	for i := 1; i < len(py); i++ {
		pSum[i] = pSum[i-1] + py[i-1]*py[i-1]
	}

	return &util.DataRangeVariance{
			Sum:       tSum,
			SquareSum: tSquareSum,
		}, &util.DataRangeVariance{
			Sum:       pSum,
			SquareSum: pSquareSum,
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

func getRangeData(fi, ti int, features ...util.DataRangeFeatures) (tR *model.RangeData, pR *model.RangeData, err error) {
	tR = &model.RangeData{}
	pR = &model.RangeData{}

	for i, feature := range features {
		if i < len(features)/2 {
			feature.GetFeatures(fi, ti, tR)
		} else {
			feature.GetFeatures(fi, ti, pR)
		}
	}

	return
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
		py = append(py, yaliFloat)
	}

	return
}

func getHistoryYaliSum(ty, py []float64) (tSum util.DataRangeMean, pSum util.DataRangeMean) {
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

func getDataIndex(from, to int64, td []model.Data, day string) (fi int, ti int, err error) {
	for i := range td {
		if td[i].ParseTime(day) == from {
			fi = i
		} else if td[i].ParseTime(day) == to {
			ti = i
		}
	}

	if (fi == ti && fi == 0) || fi > ti {
		return -1, -1, serror.WrongRangeError
	}

	return
}

func GetRangeData(id, from, to string) (*model.RangeData, *model.RangeData, error) {
	f, _ := strconv.Atoi(from)
	t, _ := strconv.Atoi(to)

	inF, ok := inMap[id]
	if !ok {
		return nil, nil, serror.DetailCloseError
	}
	mutex.Lock()
	dataT, dataP, err := inF(int64(f), int64(t))
	mutex.Unlock()
	return dataT, dataP, err
}

func DataDetailClose(id string) error {
	cnt, ok := cntMap[id]
	if !ok {
		return serror.DetailCloseError
	}
	cnt.Lock()
	cnt.cnt--
	if cnt.cnt == 0 {
		c := cMap[id]
		err := c()
		if err != nil {
			return err
		}
	}
	cnt.Unlock()

	if cnt.cnt == 0 {
		delete(cntMap, id)
		delete(cMap, id)
		delete(inMap, id)
		delete(hMap, id)
	}
	return nil
}
