package model

import "github.com/stanleydv12/gqlgen-todos/database"

type OwnedGame struct {
	GameID        int  `gorm:"primaryKey"`
	OwnedGameGame Game `gorm:"foreignKey:GameID"`
	UserID        int  `gorm:"primaryKey"`
	OwnedGameUser User `gorm:"foreignKey:UserID"`
}

func init() {
	db := database.GetInstance()
	db.DropTableIfExists(&OwnedGame{})
	db.AutoMigrate(&OwnedGame{})

	//db.Create(&OwnedGame{
	//	GameID:        1,
	//	UserID:        1,
	//})
}
