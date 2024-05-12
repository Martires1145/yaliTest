package util

import (
	"cmdTest/internal/dto/model"
	"math"
)

type DataRangeFeatures interface {
	GetFeatures(f, t int, rangData *model.RangeData)
}

type DataRangeMean []float64

func (d DataRangeMean) GetFeatures(f, t int, rangeData *model.RangeData) {
	rangeData.Mean = (d[t+1] - d[f]) / float64(t-f+1)
}

func max(a, b float64) float64 {
	if a > b {
		return a
	}
	return b
}

func min(a, b float64) float64 {
	if a < b {
		return a
	}
	return b
}

func GetStMax(nums []float64) *DataRangeMax {
	n := len(nums)
	maxLog := int(math.Log2(float64(n)))

	log := make([]int, n+1)
	for i := 2; i <= n; i++ {
		log[i] = log[i/2] + 1
	}

	mx := make([][]float64, maxLog+1)
	for i := range mx {
		mx[i] = make([]float64, n)
	}

	for i := 0; i < n; i++ {
		mx[0][i] = nums[i]
	}

	for j := 1; j <= maxLog; j++ {
		for i := 0; i+(1<<j)-1 < n; i++ {
			mx[j][i] = max(mx[j-1][i], mx[j-1][i+(1<<(j-1))])
		}
	}

	return &DataRangeMax{
		Max:    mx,
		Log:    log,
		Length: n,
	}
}

func GetStMin(nums []float64) *DataRangeMin {
	n := len(nums)
	maxLog := int(math.Log2(float64(n)))

	log := make([]int, n+1)
	for i := 2; i <= n; i++ {
		log[i] = log[i/2] + 1
	}

	mn := make([][]float64, maxLog+1)
	for i := range mn {
		mn[i] = make([]float64, n)
	}

	for i := 0; i < n; i++ {
		mn[0][i] = nums[i]
	}

	for j := 1; j <= maxLog; j++ {
		for i := 0; i+(1<<j)-1 < n; i++ {
			mn[j][i] = min(mn[j-1][i], mn[j-1][i+(1<<(j-1))])
		}
	}

	return &DataRangeMin{
		Min:    mn,
		Log:    log,
		Length: n,
	}
}

type DataRangeMax struct {
	Max    [][]float64
	Log    []int
	Length int
}

func (d *DataRangeMax) GetFeatures(f, t int, rangData *model.RangeData) {
	k := d.Log[t-f+1]
	rangData.Max = max(d.Max[k][f], d.Max[k][t-(1<<k)+1])
}

type DataRangeMin struct {
	Min    [][]float64
	Log    []int
	Length int
}

func (d *DataRangeMin) GetFeatures(f, t int, rangData *model.RangeData) {
	k := d.Log[t-f+1]
	rangData.Min = min(d.Min[k][f], d.Min[k][t-(1<<k)+1])
}

type DataRangeVariance struct {
	Sum       []float64
	SquareSum []float64
}

func (d *DataRangeVariance) GetFeatures(f, t int, rangData *model.RangeData) {
	rangData.Variance = (d.SquareSum[t+1]-d.SquareSum[f])/float64(t-f+1) - (d.Sum[t+1]-d.Sum[f])/float64(t-f+1)
}
