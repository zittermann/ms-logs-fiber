package repository

import (
	"errors"
	"ms_logs_go/helper"
	"ms_logs_go/model"

	"gorm.io/gorm"
)

type OkLogRepositoryImpl struct {
	Db *gorm.DB
}

func NewOkLogRepositoryImpl(Db *gorm.DB) OkLogRepository {
	return &OkLogRepositoryImpl{Db: Db}
}

// FindById implements OkLogRepository
func (t *OkLogRepositoryImpl) FindById(id uint) (log model.OkLog, err error) {

	var logFound model.OkLog
	result := t.Db.Find(&logFound, id)

	if result != nil {
		return logFound, nil
	} else {
		return logFound, errors.New("Log no existe")
	}

}

// FindLatest implements OkLogRepository
func (t *OkLogRepositoryImpl) FindLatest() (log model.OkLog, err error) {
	
	var logFound model.OkLog
	result := t.Db.Last(&logFound)

	if result != nil {
		return logFound, nil
	} else {
		return logFound, errors.New(`No se ha encontrado el último log 
		registrado. Probablemente porque no existe ningún log`)
	}
	
}

// Save implements OkLogRepository
func (t *OkLogRepositoryImpl) Save(log model.OkLog) {

	result := t.Db.Create(&log)
	err := result.Error

	helper.ErrorPanic(err)
}
