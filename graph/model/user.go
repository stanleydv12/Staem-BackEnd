package model

import (
	"github.com/stanleydv12/gqlgen-todos/database"
)

type User struct {
	ID       int64 `gorm:"primaryKey"`
	Name     string
	Email    string `gorm:"uniqueIndex"`
	Password string
	Wallet   float64
	Image    string
	Country  string
}

type Wishlist struct {
	ID           int  `gorm:"primaryKey"`
	GameID       int  `gorm:"primaryKey"`
	WishlistGame Game `gorm:"foreignKey:GameID"`
	UserID       int  `gorm:"primaryKey"`
	WishlistUser User `gorm:"foreignKey:UserID"`
}

func init() {
	db := database.GetInstance()
	db.DropTableIfExists(&User{})
	db.AutoMigrate(&User{})

	db.Create(&User{
		Name:     "user",
		Email:    "user@user.com",
		Password: "user",
		Wallet:   0,
		Image:    "./assets/icons/user.png",
		Country:  "Indonesia",
	})

	db.Create(&User{
		Name:     "admin",
		Email:    "admin@admin.com",
		Password: "admin",
		Wallet:   0,
		Image:    "./assets/icons/user.png",
		Country:  "Indonesia",
	})

	seedOwnedGame()
}

func seedOwnedGame() {
	db := database.GetInstance()
	db.DropTableIfExists(&Wishlist{})
	db.AutoMigrate(&Wishlist{})

	db.Create(&Wishlist{
		GameID: 1,
		UserID: 1,
	})
}
