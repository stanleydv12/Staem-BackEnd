package model

import (
	"github.com/stanleydv12/gqlgen-todos/database"
	"time"
)

type User struct {
	ID                    int `gorm:"primaryKey"`
	Name                  string
	Email                 string `gorm:"uniqueIndex"`
	Password              string
	Wallet                float64
	Image                 string
	Border                string
	Country               string
	Level                 int
	Status                string
	Badge                 string
	Summary               string
	Theme                 string
	Background            string
	MiniProfileBackground string
	Notification          int
	Point                 int
	Reported              int
	Suspended             bool
}

type Friend struct {
	UserID        int  `gorm:"primaryKey"`
	FriendID      int  `gorm:"primaryKey"`
	FriendAccount User `gorm:"foreignKey:FriendID"`
}

type FriendRequest struct {
	UserID        int  `gorm:"primaryKey"`
	FriendID      int  `gorm:"primaryKey"`
	FriendAccount User `gorm:"foreignKey:FriendID"`
	Status        string
}

type ProfileComment struct {
	UserID  int
	Comment string
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
		Name:       "user",
		Email:      "user@user.com",
		Password:   "user",
		Wallet:     9999,
		Image:      "./assets/icons/user.png",
		Country:    "Indonesia",
		Status:     "Online",
		Level:      2,
		Badge:      "https://cdn.cloudflare.steamstatic.com/steamcommunity/public/images/apps/105600/ef421de200597d9553957957bb16b1504a25c752.jpg",
		Summary:    "Jangan kepo",
		Theme:      "#eb7434",
		Background: "https://cdn.pixabay.com/photo/2019/05/27/14/41/background-4232859_1280.png",
		Point:      100000,
		Reported:   0,
	})

	db.Create(&User{
		Name:     "admin",
		Email:    "admin@admin.com",
		Password: "admin",
		Wallet:   0,
		Image:    "./assets/icons/user.png",
		Country:  "America",
		Status:   "Offline",
		Level:    3,
		Badge:    "https://cdn.cloudflare.steamstatic.com/steamcommunity/public/images/apps/105600/ef421de200597d9553957957bb16b1504a25c752.jpg",
		Summary:  "Dont kepo",
		Point:    100000,
		Reported: 0,
	})

	db.Create(&User{
		Name:     "Stanley",
		Email:    "stanley@email.com",
		Password: "stanley",
		Wallet:   9999,
		Image:    "./assets/icons/user.png",
		Country:  "South Korea",
		Status:   "Online",
		Level:    2,
		Badge:    "https://cdn.cloudflare.steamstatic.com/steamcommunity/public/images/apps/105600/ef421de200597d9553957957bb16b1504a25c752.jpg",
		Summary:  "Jangan kepo",
		Point:    100000,
		Reported: 0,
	})

	db.Create(&User{
		Name:     "Dave",
		Email:    "dave@email.com",
		Password: "dave",
		Wallet:   0,
		Image:    "./assets/icons/user.png",
		Country:  "Japan",
		Status:   "Offline",
		Level:    3,
		Badge:    "https://cdn.cloudflare.steamstatic.com/steamcommunity/public/images/apps/105600/ef421de200597d9553957957bb16b1504a25c752.jpg",
		Summary:  "Dont kepo",
		Point:    100000,
		Reported: 5,
	})

	db.Create(&User{
		Name:     "Teherag",
		Email:    "teherag@email.com",
		Password: "teherag",
		Wallet:   0,
		Image:    "./assets/icons/user.png",
		Country:  "China",
		Status:   "Offline",
		Level:    3,
		Badge:    "https://cdn.cloudflare.steamstatic.com/steamcommunity/public/images/apps/105600/ef421de200597d9553957957bb16b1504a25c752.jpg",
		Summary:  "Dont kepo",
		Point:    100000,
		Reported: 0,
	})

	seedWishlist()
	seedFriend()
	seedFriendRequest()
	seedProfileComment()
}

func seedProfileComment() {
	db := database.GetInstance()
	db.DropTableIfExists(&ProfileComment{})
	db.AutoMigrate(&ProfileComment{})

	db.Create(&ProfileComment{
		UserID:  1,
		Comment: "I think Steam is not good right now",
	})

	db.Create(&ProfileComment{
		UserID:  1,
		Comment: "It needs a big and good update",
	})
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

	db.Create(&Friend{
		UserID:   1,
		FriendID: 3,
	})

	db.Create(&Friend{
		UserID:   3,
		FriendID: 1,
	})
}

func seedFriendRequest() {
	db := database.GetInstance()
	db.DropTableIfExists(&FriendRequest{})
	db.AutoMigrate(&FriendRequest{})

	db.Create(&FriendRequest{
		UserID:   1,
		FriendID: 4,
	})

	db.Create(&FriendRequest{
		UserID:   4,
		FriendID: 1,
	})

	db.Create(&FriendRequest{
		UserID:   1,
		FriendID: 5,
	})

	db.Create(&FriendRequest{
		UserID:   5,
		FriendID: 1,
	})
}
