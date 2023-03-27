package service

import (
	"ms_logs_go/data/request"
	"ms_logs_go/data/response"
)

type OkLogService interface {
	FindById(id uint) response.OkLogResponse
	FindLatest() response.OkLogResponse
	Create(log request.CreateOkLogRequest)
}
