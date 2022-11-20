package seeder

import (
	"go-learn/db"
	model "go-learn/models"
	"strconv"
)

// create seeder 5 student
func CreateSubject() {
	db := db.DB

	for i := 0; i < 5; i++ {
		subject := model.Subject{
			Name: "Subject " + strconv.Itoa(i),
		}
		db.Create(&subject)
	}

}
