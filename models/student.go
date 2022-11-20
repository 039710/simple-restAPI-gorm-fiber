package models

import (
	"gorm.io/gorm"
)

type Student struct {
	gorm.Model
	Name      string `json:"name,omitempty"`
	Email     string `json:"email,omitempty"`
	Password  string `json:"-"`
	Dob       string `json:"dob,omitempty"`
	CreatedAt string `json:"created_at,omitempty" gorm:"column:created_at,default:CURRENT_TIMESTAMP"`
	UpdatedAt string `json:"updated_at,omitempty" gorm:"column:updated_at,default:CURRENT_TIMESTAMP"`

	Subjects []*Subject `json:"subjects,omitempty" gorm:"many2many:student_subjects;"`
}
