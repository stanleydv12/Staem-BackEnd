package model

import (
	"github.com/stanleydv12/gqlgen-todos/database"
	"time"
)

type User struct {
	ID       int64 `gorm:"primaryKey"`
	Name     string
	Email    string `gorm:"uniqueIndex"`
	Password string
	Wallet   float64
	Image    string
	Country  string
	Level    int
	Summary  string
}

type Friend struct {
	UserID        int  `gorm:"primaryKey"`
	FriendID      int  `gorm:"primaryKey"`
	FriendAccount User `gorm:"foreignKey:FriendID"`
}

type Wishlist struct {
	ID           int  `gorm:"primaryKey"`
	GameID       int  `gorm:"primaryKey"`
	WishlistGame Game `gorm:"foreignKey:GameID"`
	UserID       int  `gorm:"primaryKey"`
	WishlistUser User `gorm:"foreignKey:UserID"`
}

type PaymentMethod struct {
	UserID      int `gorm:"primaryKey"`
	Card        string
	CardNumber  string
	Date        time.Time
	Name        string
	Address     string
	PostalCode  string
	PhoneNumber string
	Country     string
}

func init() {
	db := database.GetInstance()
	db.DropTableIfExists(&User{})
	db.AutoMigrate(&User{})

	db.Create(&User{
		Name:     "user",
		Email:    "user@user.com",
		Password: "user",
		Wallet:   999,
		Image:    "./assets/icons/user.png",
		Country:  "Indonesia",
		Level:    2,
		Summary:  "Jangan kepo",
	})

	db.Create(&User{
		Name:     "admin",
		Email:    "admin@admin.com",
		Password: "admin",
		Wallet:   0,
		Image:    "./assets/icons/user.png",
		Country:  "Indonesia",
		Level:    3,
		Summary:  "Dont kepo",
	})

	seedWishlist()
}

func seedWishlist() {
	db := database.GetInstance()
	db.DropTableIfExists(&Wishlist{})
	db.AutoMigrate(&Wishlist{})

	db.Create(&Wishlist{
		GameID: 1,
		UserID: 1,
	})

	db.Create(&Wishlist{
		GameID: 2,
		UserID: 1,
	})
}

func seedFriend() {
	db := database.GetInstance()
	db.DropTableIfExists(&Friend{})
	db.AutoMigrate(&Friend{})

	db.Create(&Friend{
		UserID:   1,
		FriendID: 2,
	})

	db.Create(&Friend{
		UserID:   2,
		FriendID: 1,
	})
}
