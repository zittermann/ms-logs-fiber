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

type NokLogServiceImpl struct {
	NokLogRepository repositories.NokLogRepository
	Validate         *validator.Validate
}

func NewNokLogServiceImpl(
	nokLogRepository repositories.NokLogRepository,
	validate *validator.Validate) NokLogService {

	return &NokLogServiceImpl{
		NokLogRepository: nokLogRepository,
		Validate:         validate,
	}

}

// FindById implements service.NokLogService
func (t *NokLogServiceImpl) FindById(id uint) response.NokLogResponse{
	logData, err := t.NokLogRepository.FindById(id)
	helper.ErrorPanic(err)

	logResponse := response.NokLogResponse{
		ID:             logData.ID,
		App:            logData.App,
		Proceso:        logData.Proceso,
		ProcesoDetalle: logData.ProcesoDetalle,
		IDUsr:          logData.IDUsr,
		IDExtra:        logData.IDExtra,
		Fecha:          logData.Fecha,
		Request:        logData.Request,
		Response:       logData.Response,
		ErrorCode:      logData.ErrorCode,
	}

	return logResponse
}

// FindLatest implements NokLogService
func (t *NokLogServiceImpl) FindLatest() response.NokLogResponse {
	logData, err := t.NokLogRepository.FindLatest()
	helper.ErrorPanic(err)

	logResponse := response.NokLogResponse{
		ID:             logData.ID,
		App:            logData.App,
		Proceso:        logData.Proceso,
		ProcesoDetalle: logData.ProcesoDetalle,
		IDUsr:          logData.IDUsr,
		IDExtra:        logData.IDExtra,
		Fecha:          logData.Fecha,
		Request:        logData.Request,
		Response:       logData.Response,
		ErrorCode:      logData.ErrorCode,
	}

	return logResponse
}

// Create implements service.NokLogService
func (t *NokLogServiceImpl) Create(log request.CreateNokLogRequest) {
	err := t.Validate.Struct(log)
	helper.ErrorPanic(err)

	logModel := model.NokLog{
		App:            log.App,
		Proceso:        log.Proceso,
		ProcesoDetalle: log.ProcesoDetalle,
		IDUsr:          uint64(log.IDUsr),
		IDExtra:        log.IDExtra,

		Fecha: util.GenerateFecha(),

		Request:   log.Request,
		Response:  log.Response,
		ErrorCode: log.ErrorCode,
	}

	t.NokLogRepository.Save(logModel)
}
