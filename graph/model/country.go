package model

import "github.com/stanleydv12/gqlgen-todos/database"

type Country struct {
	ID   int64 `gorm:"primaryKey"`
	Name string
}

func init() {
	db := database.GetInstance()
	db.DropTableIfExists(&Country{})
	db.AutoMigrate(&Country{})

	db.Create(&Country{
		Name: "Indonesia",
	})
	db.Create(&Country{
		Name: "United States",
	})
	db.Create(&Country{
		Name: "Japan",
	})
	db.Create(&Country{
		Name: "China",
	})
	db.Create(&Country{
		Name: "South Korea",
	})

}
