package service

import (
	"ms_logs_go/data/request"
	"ms_logs_go/data/response"
)

type NokLogService interface {
	FindById(id uint) response.NokLogResponse
	Create(log request.CreateNokLogRequest)
}
