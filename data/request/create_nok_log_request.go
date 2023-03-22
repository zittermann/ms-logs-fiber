package request

type CreateNokLogRequest struct {
	App string `validate:"max=100,min=1" json:"app"`
	Proceso string `validate:"max=100,min=1" json:"proceso"`
	ProcesoDetalle string `validate:"max=100,min=1" json:"proceso_detalle"`
	IDUsr uint `json:"id_usr"`
	IDExtra uint `json:"id_extra"`
	Request string `json:"request"`
	Response string `json:"response"`
	ErrorCode uint16 `json:"error_code"`
}
