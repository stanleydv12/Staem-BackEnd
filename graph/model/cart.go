package model

import "github.com/stanleydv12/gqlgen-todos/database"

type Cart struct {
	GameID   int  `gorm:"primaryKey"`
	UserID   int  `gorm:"primaryKey"`
	CartGame Game `gorm:"foreignKey:GameID"`
}

func init() {
	db := database.GetInstance()
	db.DropTableIfExists(&Cart{})
	db.AutoMigrate(&Cart{})

	db.Create(&Cart{
		GameID: 1,
		UserID: 1,
	})

	db.Create(&Cart{
		GameID: 2,
		UserID: 1,
	})
}
