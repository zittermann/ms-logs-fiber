package repository

import "ms_logs_go/model"

type OkLogRepository interface {
	FindById(id uint) (log model.OkLog, err error)
	FindLatest() (log model.OkLog, err error)
	Save(log model.OkLog)
}
