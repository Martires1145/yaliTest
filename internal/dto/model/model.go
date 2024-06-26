package model

import (
	"encoding/json"
	"github.com/spf13/viper"
	"strings"
	"time"
)

var modelPath = viper.GetString("model.path") + viper.GetString("model.name")

type Params struct {
	TaskName         string `json:"task_name" db:"task_name"`
	IsTraining       string `json:"is_training" db:"is_training"`
	ModelID          string `json:"model_id" db:"model_id"`
	Model            string `json:"model" db:"model"`
	Data             string `json:"data" db:"data"`
	RootPath         string `json:"root_path" db:"root_path"`
	DataPath         string `json:"data_path" db:"data_path"`
	DataTrainPath    string `json:"data_train_path" db:"data_train_path"`
	DataValiPath     string `json:"data_vali_path" db:"data_vali_path"`
	DataTestPath     string `json:"data_test_path" db:"data_test_path"`
	Features         string `json:"features" db:"features"`
	Target           string `json:"target" db:"target"`
	Freq             string `json:"freq" db:"freq"`
	Checkpoints      string `json:"checkpoints" db:"checkpoints"`
	SeqLen           string `json:"seq_len" db:"seq_len"`
	LabelLen         string `json:"label_len" db:"label_len"`
	PredLen          string `json:"pred_len" db:"pred_len"`
	SeasonalPatterns string `json:"seasonal_patterns" db:"seasonal_patterns"`
	MaskRate         string `json:"mask_rate" db:"mask_rate"`
	AnomalyRatio     string `json:"anomaly_ratio" db:"anomaly_ratio"`
	TopK             string `json:"top_k" db:"top_k"`
	NumKernels       string `json:"num_kernels" db:"num_kernels"`
	EncIn            string `json:"enc_in" db:"enc_in"`
	DecIn            string `json:"dec_in" db:"dec_in"`
	COut             string `json:"c_out" db:"c_out"`
	DModel           string `json:"d_model" db:"d_model"`
	NHeads           string `json:"n_heads" db:"n_heads"`
	ELayers          string `json:"e_layers" db:"e_layers"`
	DLayers          string `json:"d_layers" db:"d_layers"`
	DFF              string `json:"d_ff" db:"d_ff"`
	MovingAvg        string `json:"moving_avg" db:"moving_avg"`
	Factor           string `json:"factor" db:"factor"`
	Distil           string `json:"distil" db:"distil"`
	Dropout          string `json:"dropout" db:"dropout"`
	Embed            string `json:"embed" db:"embed"`
	Activation       string `json:"activation" db:"activation"`
	OutputAttention  string `json:"output_attention" db:"output_attention"`
	NumWorkers       string `json:"num_workers" db:"num_workers"`
	Itr              string `json:"itr" db:"itr"`
	TrainEpochs      string `json:"train_epochs" db:"train_epochs"`
	BatchSize        string `json:"batch_size" db:"batch_size"`
	Patience         string `json:"patience" db:"patience"`
	LearningRate     string `json:"learning_rate" db:"learning_rate"`
	Des              string `json:"des" db:"des"`
	Loss             string `json:"loss" db:"loss"`
	Lradj            string `json:"lradj" db:"lradj"`
	UseAMP           string `json:"use_amp" db:"use_amp"`
	UseGPU           string `json:"use_gpu" db:"use_gpu"`
	GPU              string `json:"gpu" db:"gpu"`
	UseMultiGPU      string `json:"use_multi_gpu" db:"use_multi_gpu"`
	Devices          string `json:"devices" db:"devices"`
	HiddenDims       string `json:"p_hidden_dims" db:"p_hidden_dims"`
	HiddenLayers     string `json:"p_hidden_layers" db:"p_hidden_layers"`
	WeightLin        string `json:"w_lin" db:"w_lin"`
	UseKafka         string `json:"use_kafka" db:"use_kafka"`
	Scale            string `json:"scale" db:"scale"`
	Optim            string `json:"optim" db:"optim"`
}

type ParamsExtra struct {
	ID               int64  `json:"id" db:"id"`
	Freq             string `json:"freq" db:"freq"`
	Checkpoints      string `json:"checkpoints" db:"checkpoints"`
	SeasonalPatterns string `json:"seasonal_patterns" db:"seasonal_patterns"`
	MaskRate         string `json:"mask_rate" db:"mask_rate"`
	AnomalyRatio     string `json:"anomaly_ratio" db:"anomaly_ratio"`
	TopK             string `json:"top_k" db:"top_k"`
	NumKernels       string `json:"num_kernels" db:"num_kernels"`
	DModel           string `json:"d_model" db:"d_model"`
	NHeads           string `json:"n_heads" db:"n_heads"`
	DFF              string `json:"d_ff" db:"d_ff"`
	MovingAvg        string `json:"moving_avg" db:"moving_avg"`
	Distil           string `json:"distil" db:"distil"`
	Dropout          string `json:"dropout" db:"dropout"`
	Embed            string `json:"embed" db:"embed"`
	Activation       string `json:"activation" db:"activation"`
	OutputAttention  string `json:"output_attention" db:"output_attention"`
	NumWorkers       string `json:"num_workers" db:"num_workers"`
	TrainEpochs      string `json:"train_epochs" db:"train_epochs"`
	BatchSize        string `json:"batch_size" db:"batch_size"`
	Patience         string `json:"patience" db:"patience"`
	LearningRate     string `json:"learning_rate" db:"learning_rate"`
	Loss             string `json:"loss" db:"loss"`
	Lradj            string `json:"lradj" db:"lradj"`
	UseAMP           string `json:"use_amp" db:"use_amp"`
	UseGPU           string `json:"use_gpu" db:"use_gpu"`
	GPU              string `json:"gpu" db:"gpu"`
	UseMultiGPU      string `json:"use_multi_gpu" db:"use_multi_gpu"`
	Devices          string `json:"devices" db:"devices"`
	HiddenDims       string `json:"p_hidden_dims" db:"p_hidden_dims"`
	HiddenLayers     string `json:"p_hidden_layers" db:"p_hidden_layers"`
	WeightLin        string `json:"w_lin" db:"w_lin"`
}

type ParamsUsual struct {
	ID            int64  `json:"id" db:"id"`
	TaskName      string `json:"task_name" db:"task_name"`
	IsTraining    string `json:"is_training" db:"is_training"`
	ModelID       string `json:"model_id" db:"model_id"`
	Model         string `json:"model" db:"model"`
	Data          string `json:"data" db:"data"`
	RootPath      string `json:"root_path" db:"root_path"`
	DataPath      string `json:"data_path" db:"data_path"`
	DataTrainPath string `json:"data_train_path" db:"data_train_path"`
	DataValiPath  string `json:"data_vali_path" db:"data_vali_path"`
	DataTestPath  string `json:"data_test_path" db:"data_test_path"`
	Features      string `json:"features" db:"features"`
	Target        string `json:"target" db:"target"`
	SeqLen        string `json:"seq_len" db:"seq_len"`
	LabelLen      string `json:"label_len" db:"label_len"`
	PredLen       string `json:"pred_len" db:"pred_len"`
	EncIn         string `json:"enc_in" db:"enc_in"`
	DecIn         string `json:"dec_in" db:"dec_in"`
	COut          string `json:"c_out" db:"c_out"`
	ELayers       string `json:"e_layers" db:"e_layers"`
	DLayers       string `json:"d_layers" db:"d_layers"`
	Factor        string `json:"factor" db:"factor"`
	Itr           string `json:"itr" db:"itr"`
	Des           string `json:"des" db:"des"`
	UseKafka      string `json:"use_kafka" db:"use_kafka"`
	Scale         string `json:"scale" db:"scale"`
	Optim         string `json:"optim" db:"optim"`
}

type ParamsJson struct {
	PE       *ParamsExtra `json:"pe"`
	PU       *ParamsUsual `json:"pu"`
	UseExtra bool         `json:"useExtra,omitempty"`
}

type DBModel struct {
	ID         uint   `db:"id"`
	Name       string `db:"name"`
	UseCnt     int    `db:"use_cnt"`
	UseExtra   int    `db:"use_extra"`
	CreateTime int64  `db:"create_time"`
	State      int    `db:"state"`
	ParamsID   int    `db:"params_id"`
}

type JsonModel struct {
	ID         uint        `json:"ID"`
	Name       string      `json:"name"`
	UseCnt     int         `json:"useCnt"`
	CreateTime int64       `json:"createTime"`
	Params     *ParamsJson `json:"params"`
}

var DefaultParams = ParamsExtra{
	DModel: "512",
	NHeads: "8",
	DFF:    "2048",
	Embed:  "timeF",
	Distil: "True",
}

func (p *Params) Parse() (args []string, err error) {
	args = append(args, "-u")
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

func (p *ParamsJson) GetParams() *Params {
	params := &Params{
		TaskName:      p.PU.TaskName,
		IsTraining:    p.PU.IsTraining,
		ModelID:       p.PU.ModelID,
		Model:         p.PU.Model,
		Data:          p.PU.Data,
		RootPath:      p.PU.RootPath,
		DataPath:      p.PU.DataPath,
		DataTrainPath: p.PU.DataTrainPath,
		DataValiPath:  p.PU.DataValiPath,
		DataTestPath:  p.PU.DataTestPath,
		Features:      p.PU.Features,
		Target:        p.PU.Target,
		SeqLen:        p.PU.SeqLen,
		LabelLen:      p.PU.LabelLen,
		PredLen:       p.PU.PredLen,
		EncIn:         p.PU.EncIn,
		DecIn:         p.PU.DecIn,
		COut:          p.PU.COut,
		ELayers:       p.PU.ELayers,
		DLayers:       p.PU.DLayers,
		Factor:        p.PU.Factor,
		Itr:           p.PU.Itr,
		Des:           p.PU.Des,
		UseKafka:      p.PU.UseKafka,
		Scale:         p.PU.Scale,
		Optim:         p.PU.Optim,
	}
	if p.UseExtra {
		params.Freq = p.PE.Freq
		params.Checkpoints = p.PE.Checkpoints
		params.SeasonalPatterns = p.PE.SeasonalPatterns
		params.MaskRate = p.PE.MaskRate
		params.AnomalyRatio = p.PE.AnomalyRatio
		params.TopK = p.PE.TopK
		params.NumKernels = p.PE.NumKernels
		params.DModel = p.PE.DModel
		params.NHeads = p.PE.NHeads
		params.DFF = p.PE.DFF
		params.MovingAvg = p.PE.MovingAvg
		params.Distil = p.PE.Distil
		params.Dropout = p.PE.Dropout
		params.Embed = p.PE.Embed
		params.Activation = p.PE.Activation
		params.OutputAttention = p.PE.OutputAttention
		params.NumWorkers = p.PE.NumWorkers
		params.TrainEpochs = p.PE.TrainEpochs
		params.BatchSize = p.PE.BatchSize
		params.Patience = p.PE.Patience
		params.LearningRate = p.PE.LearningRate
		params.Loss = p.PE.Loss
		params.Lradj = p.PE.Lradj
		params.UseAMP = p.PE.UseAMP
		params.UseGPU = p.PE.UseGPU
		params.GPU = p.PE.GPU
		params.UseMultiGPU = p.PE.UseMultiGPU
		params.Devices = p.PE.Devices
		params.HiddenDims = p.PE.HiddenDims
		params.HiddenLayers = p.PE.HiddenLayers
		params.WeightLin = p.PE.WeightLin
	}

	return params
}

func (p *Params) IsStream() bool {
	return p.UseKafka == "1"
}

func (j *JsonModel) SetTime() {
	j.CreateTime = time.Now().Unix()
}

func (p *ParamsJson) SetUseExtra() {
	if p.PE == nil {
		p.UseExtra = true
	}
	p.UseExtra = false
}
