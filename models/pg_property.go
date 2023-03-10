package models

type Pg struct {
	ID         uint   `json:"id" validate:"required"`
	PG_Type    string `json:"pgtype" validate:"required"`
	PropertyID uint   `json:"propertyid" validate:"required"`
}
