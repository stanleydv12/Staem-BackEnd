package model

import "github.com/stanleydv12/gqlgen-todos/database"

type MarketGameItem struct {
	ID                 int `gorm:"primaryKey"`
	UserID             int
	GameItemID         int
	MarketGameItemUser User     `gorm:"foreignKey:UserID"`
	MarketGameItem     GameItem `gorm:"foreignKey:GameItemID"`
	Price              int
	Type               string
}

func init() {
	seedMarketGameItem()
}

func seedMarketGameItem() {
	db := database.GetInstance()
	db.DropTableIfExists(&MarketGameItem{})
	db.AutoMigrate(&MarketGameItem{})
	db.AutoMigrate(&MarketTransaction{})
	db.AutoMigrate(&MarketListing{})

	db.Create(&MarketGameItem{
		UserID:     4,
		GameItemID: 1,
		Price:      500,
		Type:       "offer",
	})

	db.Create(&MarketGameItem{
		UserID:     5,
		GameItemID: 1,
		Price:      550,
		Type:       "offer",
	})

	db.Create(&MarketGameItem{
		UserID:     2,
		GameItemID: 1,
		Price:      600,
		Type:       "offer",
	})

	db.Create(&MarketGameItem{
		UserID:     3,
		GameItemID: 1,
		Price:      400,
		Type:       "bid",
	})

	db.Create(&MarketGameItem{
		UserID:     4,
		GameItemID: 1,
		Price:      300,
		Type:       "bid",
	})
}
