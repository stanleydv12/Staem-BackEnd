package model

import "github.com/stanleydv12/gqlgen-todos/database"

type Game_Slider struct {
	ID    int64 `gorm:"primaryKey"`
	Name  string
	Image string
}

func init() {
	db := database.GetInstance()
	db.DropTableIfExists(&Game_Slider{})
	db.AutoMigrate(&Game_Slider{})

	db.Create(&Game_Slider{
		Name:  "Rust",
		Image: "./assets/image/rust.jpg",
	})

	db.Create(&Game_Slider{
		Name:  "Raft",
		Image: "./assets/image/raft.jpg",
	})

	db.Create(&Game_Slider{
		Name:  "Issac Repentace",
		Image: "./assets/image/issac-repentance.jpg",
	})
}
