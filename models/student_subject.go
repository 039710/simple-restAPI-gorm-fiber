package models

import (
	"gorm.io/gorm"
)

// create type StudentSubject with foreign key student_id and subject_id
type StudentSubject struct {
	gorm.Model
	// cascade delete
	StudentID int     `json:"student_id" gorm:"onDelete:CASCADE"`
	SubjectID int     `json:"subject_id" gorm:"onDelete:CASCADE"`
	CreatedAt string  `json:"created_at,omitempty" gorm:"column:created_at,default:CURRENT_TIMESTAMP"`
	UpdatedAt string  `json:"updated_at,omitempty" gorm:"column:updated_at,default:CURRENT_TIMESTAMP"`
	Student   Student `json:"student,omitempty"`
	Subject   Subject `json:"subject,omitempty"`
}
