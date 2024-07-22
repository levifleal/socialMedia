package schemas

import (
	"gorm.io/gorm"
)

type Profile struct {
	gorm.Model
	
}

type ProfileRespose struct {

	BaseSchema
}
