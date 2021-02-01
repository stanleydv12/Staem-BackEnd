package model

import "github.com/stanleydv12/gqlgen-todos/database"

type GameReview struct {
	ID          int64 `gorm:"primaryKey"`
	GameId      int
	UserId      int
	Description string
	Rating      string
	Positive    int64
	Negative    int64
}

func init() {
	db := database.GetInstance()
	db.DropTableIfExists(&GameReview{})
	db.AutoMigrate(&GameReview{})

	db.Create(&GameReview{
		GameId:      1,
		UserId:      1,
		Description: "Ini game yang bagus",
		Rating:      "Positive",
		Positive:    1,
		Negative:    1,
	})

	db.Create(&GameReview{
		GameId:      1,
		UserId:      1,
		Description: "Ini game yang kurang bagus",
		Rating:      "Negative",
		Positive:    1,
		Negative:    1,
	})

	db.Create(&GameReview{
		GameId:      1,
		UserId:      2,
		Description: "Ini game yang bagus",
		Rating:      "Positive",
		Positive:    1,
		Negative:    1,
	})

	db.Create(&GameReview{
		GameId:      1,
		UserId:      2,
		Description: "Ini game yang kurang bagus",
		Rating:      "Negative",
		Positive:    1,
		Negative:    1,
	})
}
