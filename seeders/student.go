package seeder

import (
	"go-learn/db"
	model "go-learn/models"
	"strconv"

	"github.com/golang-jwt/jwt"
)

// create seeder 5 student
func CreateStudent() {
	db := db.DB

	for i := 0; i < 5; i++ {
		student := model.Student{
			Name:     "Student " + strconv.Itoa(i),
			Email:    "student" + strconv.Itoa(i) + "@gmail.com",
			Password: jwt.EncodeSegment([]byte("student" + strconv.Itoa(i))),
			Dob:      "1999-01-01",
		}
		db.Create(&student)
	}

}
