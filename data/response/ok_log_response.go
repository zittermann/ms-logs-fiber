package response

import "time"

type OkLogResponse struct {
	ID uint `json:"id"`
	App string `json:"app"`
	Proceso string `json:"proceso"`
	ProcesoDetalle string `json:"proceso_detalle"`
	IDUsr uint64 `json:"id_usr"`
	IDExtra uint `json:"id_extra"`

	Fecha time.Time `json:"fecha"`
}
