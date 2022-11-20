package models

import (
	"gorm.io/gorm"
)

type Subject struct {
	gorm.Model
	Name      string     `json:"name,omitempty"`
	CreatedAt string     `json:"created_at,omitempty" gorm:"column:created_at,default:CURRENT_TIMESTAMP"`
	UpdatedAt string     `json:"updated_at,omitempty" gorm:"column:updated_at,default:CURRENT_TIMESTAMP"`
	Students  []*Student `json:"students,omitempty" gorm:"many2many:student_subjects;"`
}
