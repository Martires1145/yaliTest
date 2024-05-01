package server

import (
	"cmdTest/internal/dto/model"
	"cmdTest/internal/dto/mysql"
	"cmdTest/pkg/util"
	"fmt"
	"mime/multipart"
	"strconv"
	"sync"
)

type InFunction func(int, int) (*model.RangeData, error)
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
	tp, pp := hMap[id].TrueDataPath, hMap[id].PDataPath
	fromChan, toChan := make(chan int), make(chan int)
	rangeDataChan, errorChan := make(chan *model.RangeData), make(chan error)
	closeChan := make(chan int)

	getInFunc := func() (from int, to int, c int) {
		from = <-fromChan
		to = <-toChan
		c = <-closeChan
		return
	}

	go func() {
		_, _, err := getHistoryData(tp, pp)
		if err != nil {
			fmt.Println(err.Error())
		}
		for {
			_, _, c := getInFunc()
			if c == 1 {
				break
			}
			rangeDataChan <- nil
			errorChan <- nil
		}
	}()

	return func(from int, to int) (*model.RangeData, error) {
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

func getHistoryData(t string, p string) ([]model.Data, []model.Data, error) {
	// todo
	panic("todo")
}

func GetRangeData(id, from, to string) (*model.RangeData, error) {
	f, _ := strconv.Atoi(from)
	t, _ := strconv.Atoi(to)

	inF := inMap[id]
	mutex.Lock()
	data, err := inF(f, t)
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
