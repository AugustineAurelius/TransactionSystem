package user

import "gorm.io/gorm"

type Exchanger struct {
	gorm.Model
	ID    int `gorm:"primaryKey; autoIncrement"`
	Money float64
}
