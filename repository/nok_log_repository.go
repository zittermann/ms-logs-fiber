package repository

import "ms_logs_go/model"

type NokLogRepository interface {
	FindById(id uint) (log model.NokLog, err error)
	FindLatest() (log model.NokLog, err error)
	Save(log model.NokLog)
}
