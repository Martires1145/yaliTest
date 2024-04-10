package model

import (
	"encoding/json"
	"github.com/spf13/viper"
	"strings"
)

var modelPath = viper.GetString("model.path")

type Params struct {
	TaskName         string `json:"task_name"`
	IsTraining       string `json:"is_training"`
	ModelID          string `json:"model_id"`
	Model            string `json:"model"`
	Data             string `json:"data"`
	RootPath         string `json:"root_path"`
	DataPath         string `json:"data_path"`
	DataTrainPath    string `json:"data_train_path"`
	DataValiPath     string `json:"data_vali_path"`
	DataTestPath     string `json:"data_test_path"`
	Features         string `json:"features"`
	Target           string `json:"target"`
	Freq             string `json:"freq"`
	Checkpoints      string `json:"checkpoints"`
	SeqLen           string `json:"seq_len"`
	LabelLen         string `json:"label_len"`
	PredLen          string `json:"pred_len"`
	SeasonalPatterns string `json:"seasonal_patterns"`
	MaskRate         string `json:"mask_rate"`
	AnomalyRatio     string `json:"anomaly_ratio"`
	TopK             string `json:"top_k"`
	NumKernels       string `json:"num_kernels"`
	EncIn            string `json:"enc_in"`
	DecIn            string `json:"dec_in"`
	COut             string `json:"c_out"`
	DModel           string `json:"d_model"`
	NHeads           string `json:"n_heads"`
	ELayers          string `json:"e_layers"`
	DLayers          string `json:"d_layers"`
	DFF              string `json:"d_ff"`
	MovingAvg        string `json:"moving_avg"`
	Factor           string `json:"factor"`
	Distil           string `json:"distil"`
	Dropout          string `json:"dropout"`
	Embed            string `json:"embed"`
	Activation       string `json:"activation"`
	OutputAttention  string `json:"output_attention"`
	NumWorkers       string `json:"num_workers"`
	Itr              string `json:"itr"`
	TrainEpochs      string `json:"train_epochs"`
	BatchSize        string `json:"batch_size"`
	Patience         string `json:"patience"`
	LearningRate     string `json:"learning_rate"`
	Des              string `json:"des"`
	Loss             string `json:"loss"`
	Lradj            string `json:"lradj"`
	UseAMP           string `json:"use_amp"`
	UseGPU           string `json:"use_gpu"`
	GPU              string `json:"gpu"`
	UseMultiGPU      string `json:"use_multi_gpu"`
	Devices          string `json:"devices"`
	HiddenDims       string `json:"p_hidden_dims"`
	HiddenLayers     string `json:"p_hidden_layers"`
	WeightLin        string `json:"w_lin"`
	UseKafka         string `json:"use_kafka"`
	Scale            string `json:"scale"`
	Optim            string `json:"optim"`
}

type Model struct {
	ID         uint    `json:"ID" db:"id"`
	Name       string  `json:"name" db:"name"`
	UseCnt     int     `json:"useCnt" db:"use_cnt"`
	FileCnt    int     `json:"fileCnt" db:"file_cnt"`
	CreateTime int64   `json:"createTime" db:"create_time"`
	Param      *Params `json:"param"`
}

func (p *Params) Parse() (args []string, err error) {
	args = append(args, "-m")
	args = append(args, modelPath)

	bytes, err := json.Marshal(p)
	if err != nil {
		return nil, err
	}
	jsonStr := string(bytes)
	pair := strings.Split(jsonStr[1:len(jsonStr)-1], ",")
	for _, p := range pair {
		kv := strings.Split(p, ":")
		if kv[1] == `""` {
			continue
		}
		args = append(args, "--"+kv[0][1:len(kv[0])-1])
		args = append(args, kv[1][1:len(kv[1])-1])
	}

	return
}

func (p *Params) IsStream() bool {
	return p.UseKafka == "1"
}
