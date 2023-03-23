package repository

import (
	"errors"
	"ms_logs_go/helper"
	"ms_logs_go/model"

	"gorm.io/gorm"
)

type NokLogRepositoryImpl struct {
	Db *gorm.DB
}

func NewNokLogRepositoryImpl(Db *gorm.DB) NokLogRepository {
	return &NokLogRepositoryImpl{Db: Db}
}

// FindById implements repository.NokLogRepository
func (t *NokLogRepositoryImpl) FindById(id uint) (log model.NokLog, err error) {
	var logFound model.NokLog
	result := t.Db.Find(&logFound, id) 

	if result != nil {
		return logFound, nil
	} else {
		return logFound, errors.New("Log no existe")
	}
}

// Save implements repositories.NokLogRepository
func (t *NokLogRepositoryImpl) Save(log model.NokLog) {
	result := t.Db.Create(&log)
	err := result.Error
	helper.ErrorPanic(err)
}
