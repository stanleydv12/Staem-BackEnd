package model

import "github.com/stanleydv12/gqlgen-todos/database"

type CommunityContent struct {
	ID                 int `gorm:"primaryKey;autoIncrement"`
	ContentPath        string
	ContentDescription string
	ContentType        string
	UserID             int
	ContentUser        User `gorm:"foreignKey:UserID"`
	Positive           int
	Negative           int
}

type CommunityContentReview struct {
	ContentID         int
	Content           CommunityContent `gorm:"foreignKey:ContentID"`
	UserID            int
	ContentReviewUser User `gorm:"foreignKey:UserID"`
	Review            string
}

type CommunityReview struct {
	ID            int `gorm:"primaryKey"`
	ReviewContent string
	UserID        int
	ReviewUser    User `gorm:"foreignKey:UserID"`
	Rating        string
}

type CommunityReviewDetail struct {
	ReviewID         int `gorm:"primaryKey"`
	ReviewContent    string
	UserID           int
	ReviewDetailUser User `gorm:"foreignKey:UserID"`
}

type CommunityDiscussion struct {
	ID                int `gorm:"primaryKey"`
	GameID            int
	DiscussionGame    Game `gorm:"foreignKey:GameID"`
	UserID            int
	DiscussionUser    User `gorm:foreignKey:UserID`
	DiscussionContent string
}

type CommunityDiscussionDetail struct {
	GameID               int
	DiscussionID         int
	UserID               int
	DiscussionDetailUser User `gorm:foreignKey:UserID`
	DiscussionContent    string
}

func init() {
	seedCommunityContent()
	seedCommunityContentReview()
	seedCommunityReview()
	seedCommunityReviewDetail()
	seedCommunityDiscussion()
	seedCommunityDiscussionDetail()
}

func seedCommunityDiscussionDetail() {
	db := database.GetInstance()
	db.DropTableIfExists(&CommunityDiscussionDetail{})
	db.AutoMigrate(&CommunityDiscussionDetail{})

	db.Create(&CommunityDiscussionDetail{
		GameID:            1,
		DiscussionID:      1,
		UserID:            1,
		DiscussionContent: "You should try this game",
	})

	db.Create(&CommunityDiscussionDetail{
		GameID:            1,
		DiscussionID:      1,
		UserID:            1,
		DiscussionContent: "I dont think so",
	})

	db.Create(&CommunityDiscussionDetail{
		GameID:            1,
		DiscussionID:      1,
		UserID:            1,
		DiscussionContent: "You should not try this game",
	})

	db.Create(&CommunityDiscussionDetail{
		GameID:            2,
		DiscussionID:      3,
		UserID:            1,
		DiscussionContent: "You should try this game",
	})

	db.Create(&CommunityDiscussionDetail{
		GameID:            2,
		DiscussionID:      3,
		UserID:            1,
		DiscussionContent: "I dont think so",
	})

	db.Create(&CommunityDiscussionDetail{
		GameID:            2,
		DiscussionID:      3,
		UserID:            1,
		DiscussionContent: "You should not try this game",
	})
}

func seedCommunityDiscussion() {
	db := database.GetInstance()
	db.DropTableIfExists(&CommunityDiscussion{})
	db.AutoMigrate(&CommunityDiscussion{})

	db.Create(&CommunityDiscussion{
		GameID:            1,
		UserID:            1,
		DiscussionContent: "Hello, i reviewed this game",
	})

	db.Create(&CommunityDiscussion{
		GameID:            1,
		UserID:            1,
		DiscussionContent: "Hmmmm, i dont like this kind of game",
	})

	db.Create(&CommunityDiscussion{
		GameID:            1,
		UserID:            2,
		DiscussionContent: "Hello, i reviewed this game too 1",
	})

	db.Create(&CommunityDiscussion{
		GameID:            1,
		UserID:            1,
		DiscussionContent: "Hello, i reviewed this game 1",
	})

	db.Create(&CommunityDiscussion{
		GameID:            1,
		UserID:            1,
		DiscussionContent: "Hmmmm, i dont like this kind of game 1",
	})

	db.Create(&CommunityDiscussion{
		GameID:            1,
		UserID:            2,
		DiscussionContent: "Hello, i reviewed this game too",
	})

	db.Create(&CommunityDiscussion{
		GameID:            2,
		UserID:            1,
		DiscussionContent: "Hello, i reviewed this game",
	})

	db.Create(&CommunityDiscussion{
		GameID:            2,
		UserID:            1,
		DiscussionContent: "Hmmmm, i dont like this kind of game",
	})

	db.Create(&CommunityDiscussion{
		GameID:            2,
		UserID:            2,
		DiscussionContent: "Hello, i reviewed this game too",
	})
}

func seedCommunityReviewDetail() {
	db := database.GetInstance()
	db.DropTableIfExists(&CommunityReviewDetail{})
	db.AutoMigrate(&CommunityReviewDetail{})

	db.Create(&CommunityReviewDetail{
		ReviewID:      1,
		ReviewContent: "Gw setuju dengan komen ini",
		UserID:        1,
	})

	db.Create(&CommunityReviewDetail{
		ReviewID:      1,
		ReviewContent: "HAHAHAHAHAH",
		UserID:        2,
	})

	db.Create(&CommunityReviewDetail{
		ReviewID:      2,
		ReviewContent: "Gw setuju dengan komen ini",
		UserID:        1,
	})

	db.Create(&CommunityReviewDetail{
		ReviewID:      2,
		ReviewContent: "HAHAHAHAHAH",
		UserID:        2,
	})
}

func seedCommunityContent() {
	db := database.GetInstance()
	db.DropTableIfExists(&CommunityContent{})
	db.AutoMigrate(&CommunityContent{})

	db.Create(&CommunityContent{
		ContentPath:        "https://cdn.akamai.steamstatic.com/steam/apps/256740604/movie480.webm?t=1551460377",
		ContentDescription: "Ini game yaang lumayan oke",
		ContentType:        "vid",
		UserID:             1,
		Positive:           10,
		Negative:           10,
	})

	db.Create(&CommunityContent{
		ContentPath:        "https://cdn.akamai.steamstatic.com/steam/apps/1372810/ss_296dc7b553f50ec0f574f88d7e922de4fb016442.600x338.jpg",
		ContentDescription: "Ini game yaang lumayan oke",
		ContentType:        "img",
		UserID:             2,
		Positive:           14,
		Negative:           5,
	})
}

func seedCommunityContentReview() {
	db := database.GetInstance()
	db.DropTableIfExists(&CommunityContentReview{})
	db.AutoMigrate(&CommunityContentReview{})

	db.Create(&CommunityContentReview{
		ContentID: 1,
		UserID:    2,
		Review:    "Game ini bagus",
	})

	db.Create(&CommunityContentReview{
		ContentID: 1,
		UserID:    1,
		Review:    "Game ini bagus , tapi jelek",
	})

	db.Create(&CommunityContentReview{
		ContentID: 2,
		UserID:    2,
		Review:    "Game ini bagus",
	})

	db.Create(&CommunityContentReview{
		ContentID: 2,
		UserID:    1,
		Review:    "Game ini bagus , tapi jelek",
	})
}

func seedCommunityReview() {
	db := database.GetInstance()
	db.DropTableIfExists(&CommunityReview{})
	db.AutoMigrate(&CommunityReview{})

	db.Create(&CommunityReview{
		ReviewContent: "Lorem ipsum dolor sit amet, consectetur adipiscing elit, sed do eiusmod tempor incididunt ut labore et dolore magna aliqua. Ut enim ad minim veniam, quis nostrud exercitation ullamco laboris nisi ut aliquip ex ea commodo consequat. Duis aute irure dolor in reprehenderit in voluptate velit esse cillum dolore eu fugiat nulla pariatur",
		UserID:        1,
		Rating:        "recommended",
	})

	db.Create(&CommunityReview{
		ReviewContent: "Lorem ipsum dolor sit amet, consectetur adipiscing elit, sed do eiusmod tempor incididunt ut labore et dolore magna aliqua. Ut enim ad minim veniam, quis nostrud exercitation ullamco laboris nisi ut aliquip ex ea commodo consequat. Duis aute irure dolor in reprehenderit in voluptate velit esse cillum dolore eu fugiat nulla pariatur",
		UserID:        2,
		Rating:        "not recommended",
	})
}
