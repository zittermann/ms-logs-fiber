package service

import (
	"ms_logs_go/data/request"
	"ms_logs_go/data/response"
	"ms_logs_go/helper"
	"ms_logs_go/model"
	repositories "ms_logs_go/repository"
	"ms_logs_go/util"

	"github.com/go-playground/validator/v10"
)

type OkLogServiceImpl struct {
	OkLogRepository repositories.OkLogRepository
	Validate        *validator.Validate
}

func NewOkLogServiceImpl(
	okLogRepository repositories.OkLogRepository,
	validate *validator.Validate) OkLogService {

	return &OkLogServiceImpl{
		OkLogRepository: okLogRepository,
		Validate:        validate,
	}
	
}

// FindById implements service.OkLogService
func (t *OkLogServiceImpl) FindById(id uint) response.OkLogResponse {
	logData, err := t.OkLogRepository.FindById(id)
	helper.ErrorPanic(err, err.Error())

	logResponse := response.OkLogResponse {
		ID: logData.ID,
		App: logData.App,
		Proceso: logData.Proceso,
		ProcesoDetalle: logData.ProcesoDetalle,
		IDUsr: logData.IDUsr,
		IDExtra: logData.IDExtra,
		Fecha: logData.Fecha,
	}

	return logResponse

}

// Create implements service.OkLogService
func (t *OkLogServiceImpl) Create(log request.CreateOkLogRequest) {
	err := t.Validate.Struct(log)
	helper.ErrorPanic(err, err.Error())

	logModel := model.OkLog {
		App: log.App,
		Proceso: log.Proceso,
		ProcesoDetalle: log.ProcesoDetalle,
		IDUsr: uint64(log.IDUsr),
		IDExtra: log.IDExtra,

		Fecha: util.GenerateFecha(),

	}

	t.OkLogRepository.Save(logModel)

}
