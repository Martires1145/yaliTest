package mysql

import (
	"cmdTest/internal/dto/model"
)

type ModelDaoMysql struct{}

func (m *ModelDaoMysql) NewParams(paramsData *model.ParamsJson) (ID int64, err error) {
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

	sqlStr := `INSERT INTO params_usually(task_name, is_training, root_path,
                           data_path, data_train_path, data_vali_path,
                           data_test_path, model_id, model,
                           data, target, features,
                           seq_len, label_len, pred_len,
                           e_layers, d_layers, factor,
                           enc_in, dec_in, c_out,
                           des, itr, use_kafka,
                           scale, optim, finish)
				VALUE (:task_name, :is_training, :root_path,
					   :data_path, :data_train_path, :data_vali_path,
					   :data_test_path, :model_id, :model,
					   :data, :target, :features,
					   :seq_len, :label_len, :pred_len,
					   :e_layers, :d_layers, :factor,
					   :enc_in, :dec_in, :c_out,
					   :des, :itr, :use_kafka,
					   :scale, :optim, :finish)`

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
	}

	_, err = tx.NamedExec(sqlStr, &paramsData.PE)
	return id, err
}

func (m *ModelDaoMysql) NewModel(modelData *model.JsonModel, ParamsID int64) error {
	sqlStr := `INSERT INTO models(name, params_id, create_time)
				   VALUE (?, ?, ?)`

	_, err := db.Exec(sqlStr,
		modelData.Name,
		ParamsID,
		modelData.CreateTime,
	)

	return err
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
			   WHERE id = (SELECT params_id FROM models WHERE models.id = ?)`
	_, err = tx.Exec(sqlStr, id)
	if err != nil {
		return err
	}

	sqlStr = `DELETE FROM params_usually
			  WHERE id = (SELECT params_id FROM models WHERE models.id = ?)`
	_, err = tx.Exec(sqlStr, id)
	if err != nil {
		return err
	}

	sqlStr = `DELETE FROM models WHERE id = ?`
	_, err = tx.Exec(sqlStr, id)
	return err
}
