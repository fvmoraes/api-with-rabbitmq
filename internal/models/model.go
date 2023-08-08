package models

import (
	"gorm.io/gorm"
)

type Foobar struct {
	gorm.Model
	What         string `json:"what" validate:"nonzero,max=30,regexp=^[a-zA-Z]*$"`
	Message      string `json:"message" validate:"nonzero,max=60"`
	Registration uint64 `json:"registration" validate:"min=8"`
}
