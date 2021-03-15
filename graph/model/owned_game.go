package model

import (
	"github.com/stanleydv12/gqlgen-todos/database"
)

type OwnedGame struct {
	GameID            int  `gorm:"primaryKey"`
	OwnedGameGame     Game `gorm:"foreignKey:GameID"`
	UserID            int  `gorm:"primaryKey"`
	OwnedGameUser     User `gorm:"foreignKey:UserID"`
	OwnedGameItemGame OwnedGameItem
	OwnedBadge        OwnedBadge
	OwnedTradingCard  OwnedTradingCard
}

type OwnedGameItem struct {
	GameItemID int
	GameItem   GameItem
	UserID     int
	User       User
}

type OwnedTradingCard struct {
	GameID        int
	TradingCardID int
	UserID        int
	User          User
	TradingCard   TradingCard
}

type OwnedBadge struct {
	GameID      int
	GameBadgeID int
	UserID      int
	User        User
	Badges      Badges
}

func init() {
	db := database.GetInstance()
	db.DropTableIfExists(&OwnedGame{})
	db.AutoMigrate(&OwnedGame{})

	db.Create(&OwnedGame{
		GameID: 1,
		UserID: 1,
	})

	db.Create(&OwnedGame{
		GameID: 2,
		UserID: 1,
	})

	db.Create(&OwnedGame{
		GameID: 3,
		UserID: 1,
	})

	db.Create(&OwnedGame{
		GameID: 4,
		UserID: 1,
	})

	db.Create(&OwnedGame{
		GameID: 5,
		UserID: 1,
	})

	// 2

	db.Create(&OwnedGame{
		GameID: 1,
		UserID: 2,
	})

	db.Create(&OwnedGame{
		GameID: 2,
		UserID: 2,
	})

	db.Create(&OwnedGame{
		GameID: 3,
		UserID: 2,
	})

	db.Create(&OwnedGame{
		GameID: 4,
		UserID: 2,
	})

	db.Create(&OwnedGame{
		GameID: 5,
		UserID: 2,
	})

	seedOwnedGameItem()
	seedOwnedBadges()
	seedOwnedCard()
}

func seedOwnedCard() {
	db := database.GetInstance()
	db.DropTableIfExists(&OwnedTradingCard{})
	db.AutoMigrate(&OwnedTradingCard{})

	db.Create(&OwnedTradingCard{
		GameID:        1,
		TradingCardID: 1,
		UserID:        1,
	})

	db.Create(&OwnedTradingCard{
		GameID:        1,
		TradingCardID: 2,
		UserID:        1,
	})

	db.Create(&OwnedTradingCard{
		GameID:        1,
		TradingCardID: 3,
		UserID:        1,
	})
}

func seedOwnedBadges() {
	db := database.GetInstance()
	db.DropTableIfExists(&OwnedBadge{})
	db.AutoMigrate(&OwnedBadge{})

	db.Create(&OwnedBadge{
		GameID:      1,
		GameBadgeID: 1,
		UserID:      1,
	})

	db.Create(&OwnedBadge{
		GameID:      1,
		GameBadgeID: 2,
		UserID:      1,
	})

	db.Create(&OwnedBadge{
		GameID:      1,
		GameBadgeID: 3,
		UserID:      1,
	})
}

func seedOwnedGameItem() {
	db := database.GetInstance()
	db.DropTableIfExists(&OwnedGameItem{})
	db.AutoMigrate(&OwnedGameItem{})

	db.Create(&OwnedGameItem{
		GameItemID: 1,
		UserID:     1,
	})

	db.Create(&OwnedGameItem{
		GameItemID: 1,
		UserID:     5,
	})

	db.Create(&OwnedGameItem{
		GameItemID: 2,
		UserID:     1,
	})

}
