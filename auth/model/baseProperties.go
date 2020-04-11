package model

import "time"

type BaseProperties struct {
	Id        int
	IsDeleted bool //default value false
	CreatedBy string
	UpdatedBy string
	CreatedAt time.Time
	UpdatedAt time.Time
}
