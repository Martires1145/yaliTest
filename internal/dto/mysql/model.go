package mysql

import (
	"cmdTest/internal/dto/model"
	"database/sql"
	"fmt"
)

type ModelDaoMysql struct{}

func (m *ModelDaoMysql) NewParams(paramsData *model.ParamsJson, modelID int64) (ID int64, err error) {
	tx, err := db.Beginx()
	if err != nil {
		return 0, err
	}
	defer func() {
		if p := recover(); p != nil {
			_ = tx.Rollback()
			panic(p)
		} else if err != nil {
			_ = tx.Rollback()
		} else {
			err := tx.Commit()
			if err != nil {
				_ = tx.Rollback()
				panic(err.Error())
			}
		}
	}()

	sqlStr := `INSERT INTO params_usually(model_id, task_name, is_training, root_path,
                           data_path, data_train_path, data_vali_path,
                           data_test_path, model,
                           data, target, features,
                           seq_len, label_len, pred_len,
                           e_layers, d_layers, factor,
                           enc_in, dec_in, c_out,
                           des, itr, use_kafka,
                           scale, optim)
				VALUE (:model_id, :task_name, :is_training, :root_path,
					   :data_path, :data_train_path, :data_vali_path,
					   :data_test_path, :model,
					   :data, :target, :features,
					   :seq_len, :label_len, :pred_len,
					   :e_layers, :d_layers, :factor,
					   :enc_in, :dec_in, :c_out,
					   :des, :itr, :use_kafka,
					   :scale, :optim)`

	paramsData.PU.ModelID = fmt.Sprintf("%d", modelID)
	exec, err := tx.NamedExec(sqlStr, &paramsData.PU)
	if err != nil {
		return 0, err
	}
	id, _ := exec.LastInsertId()

	if paramsData.UseExtra {
		paramsData.PE.ID = id

		sqlStr = `INSERT INTO params_extra(id, freq, checkpoints,
                         seasonal_patterns, mask_rate, anomaly_ratio,
                         top_k, num_kernels, d_model,
                         n_heads, d_ff, moving_avg,
                         distil, dropout, embed,
                         activation, output_attention, num_workers,
                         train_epochs, batch_size, patience,
                         learning_rate, loss, lradj,
                         use_amp, use_gpu, gpu,
                         use_multi_gpu, devices, p_hidden_dims,
                         p_hidden_layers, w_lin)
				VALUE (:id, :freq, :checkpoints,
					   :seasonal_patterns, :mask_rate, :anomaly_ratio,
					   :top_k, :num_kernels, :d_model,
					   :n_heads, :d_ff, :moving_avg,
					   :distil, :dropout, :embed,
					   :activation, :output_attention, :num_workers,
					   :train_epochs, :batch_size, :patience,
					   :learning_rate, :loss, :lradj,
					   :use_amp, :use_gpu, :gpu,
					   :use_multi_gpu, :devices,:p_hidden_dims,
					   :p_hidden_layers, :w_lin)`

		_, err = tx.NamedExec(sqlStr, &paramsData.PE)
	}

	return id, err
}

func (m *ModelDaoMysql) NewModel(modelData *model.JsonModel) (int64, error) {
	sqlStr := `INSERT INTO models(name, create_time, use_extra) VALUE (?, ?, ?)`

	exec, err := db.Exec(sqlStr,
		modelData.Name,
		modelData.CreateTime,
		modelData.Params.UseExtra,
	)

	id, err := exec.LastInsertId()
	return id, err
}

func (m *ModelDaoMysql) DeleteModel(id string) error {
	tx, err := db.Beginx()
	if err != nil {
		return err
	}
	defer func() {
		if p := recover(); p != nil {
			_ = tx.Rollback()
			panic(p)
		} else if err != nil {
			_ = tx.Rollback()
		} else {
			err := tx.Commit()
			if err != nil {
				_ = tx.Rollback()
				panic(err.Error())
			}
		}
	}()

	sqlStr := `DELETE FROM params_extra
			   WHERE id = (SELECT params_id FROM models WHERE models.id = ?) 
			   AND (
			   SELECT COUNT(1) FROM models WHERE params_id = (
			   	SELECT params_id FROM models WHERE models.id = ?
			                                )
			   ) = 1`
	_, err = tx.Exec(sqlStr, id, id)
	if err != nil {
		return err
	}

	sqlStr = `DELETE FROM params_usually
			  WHERE id = (SELECT params_id FROM models WHERE models.id = ?)
			  AND (
			   SELECT COUNT(1) FROM models WHERE params_id = (
			   	SELECT params_id FROM models WHERE models.id = ?
			                                )
			   ) = 1`
	_, err = tx.Exec(sqlStr, id, id)
	if err != nil {
		return err
	}

	sqlStr = `DELETE FROM models WHERE id = ?`
	_, err = tx.Exec(sqlStr, id)
	return err
}

func (m *ModelDaoMysql) ModifyModel(id, name, useKafka string) error {
	tx, err := db.Beginx()
	if err != nil {
		return err
	}
	defer func() {
		if p := recover(); p != nil {
			_ = tx.Rollback()
			panic(p)
		} else if err != nil {
			_ = tx.Rollback()
		} else {
			err := tx.Commit()
			if err != nil {
				_ = tx.Rollback()
				panic(err.Error())
			}
		}
	}()

	sqlStr := "UPDATE models SET name = ? WHERE id = ?"
	_, err = tx.Exec(sqlStr, name, id)
	if err != nil {
		return err
	}

	sqlStr = `UPDATE params_usually
			  SET use_kafka = ? 
			  WHERE id = (
			  	SELECT params_id FROM models WHERE models.id = ?
			  )`
	_, err = tx.Exec(sqlStr, useKafka, id)

	return err
}

func (m *ModelDaoMysql) CopyModel(id, name string) error {
	sqlStr := `INSERT INTO
    		   models(name, use_cnt, use_extra, state, params_id, create_time)
               SELECT ?, use_cnt, use_extra, state, params_id, UNIX_TIMESTAMP(NOW())
               FROM models WHERE id = ?`
	_, err := db.Exec(sqlStr, name, id)
	return err
}

func (m *ModelDaoMysql) GetAllModel() (models []model.DBModel, err error) {
	sqlStr := `SELECT * FROM models`
	err = db.Select(&models, sqlStr)
	return
}

func (m *ModelDaoMysql) GetModelParams(id string) (*model.ParamsJson, error) {
	tx, err := db.Beginx()
	if err != nil {
		return nil, err
	}
	defer func() {
		if p := recover(); p != nil {
			_ = tx.Rollback()
			panic(p)
		} else if err != nil {
			_ = tx.Rollback()
		} else {
			err := tx.Commit()
			if err != nil {
				_ = tx.Rollback()
				panic(err.Error())
			}
		}
	}()

	var params model.ParamsJson
	params.PE = &model.ParamsExtra{}
	params.PU = &model.ParamsUsual{}

	sqlStr := `SELECT * 
	           FROM params_usually 
	           WHERE id = (SELECT params_id FROM models WHERE models.id = ?)`
	err = tx.Get(params.PU, sqlStr, id)
	if err != nil {
		return nil, err
	}

	sqlStr = `SELECT * 
	           FROM params_extra 
	           WHERE id = (SELECT params_id FROM models WHERE models.id = ?)`
	err = tx.Get(params.PE, sqlStr, id)
	if err == sql.ErrNoRows {
		err = nil
		params.PE = nil
	}
	return &params, err
}

func (m *ModelDaoMysql) UpdateModelParamsID(modelID, paramsID int64) error {
	sqlStr := "UPDATE models SET params_id = ? WHERE id = ?"
	_, err := db.Exec(sqlStr, paramsID, modelID)
	return err
}

func (m *ModelDaoMysql) UseModel(id string) error {
	sqlStr := "UPDATE models SET use_cnt = use_cnt + 1 WHERE id = ?"
	_, err := db.Exec(sqlStr, id)
	return err
}
