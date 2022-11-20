package seeder

import (
	"go-learn/db"
	model "go-learn/models"
)

// create seeder 5 student
func CreateSubjectStudent() {
	db := db.DB

	for i := 0; i < 5; i++ {
		ss := model.StudentSubject{
			StudentID: i + 1,
			SubjectID: i + 1,
		}
		db.Create(&ss)

	}

}
