package model

import "time"

type OkLog struct {
	
	ID uint `gorm:"primarykey" json:"id"`
	App string `gorm:"type:varchar(100)" json:"app"`
	Proceso string `gorm:"type:varchar(100)" json:"proceso"`
	ProcesoDetalle string `gorm:"type:varchar(100)" json:"proceso_detalle"`
	IDUsr uint64 `json:"id_usr"`
	IDExtra uint `json:"id_extra"`

	Fecha time.Time `json:"fecha"`
	
}
