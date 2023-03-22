package request

type CreateOkLogRequest struct {
	App string `validate:"max=100,min=1" json:"app"`
	Proceso string `validate:"max=100,min=1" json:"proceso"`
	ProcesoDetalle string `validate:"max=100,min=1" json:"proceso_detalle"`
	IDUsr uint `json:"id_usr"`
	IDExtra uint `json:"id_extra"`
}
