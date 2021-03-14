package model

import "github.com/stanleydv12/gqlgen-todos/database"

type UnsuspensionRequest struct {
	UserID int
	Reason string
	Status string
}

type ReportRequest struct {
	ID          int `gorm:"primaryKey"`
	ReporterID  int
	SuspectedID int
	Reason      string
}

type SuspensionList struct {
	UserID    int
	User      User
	Reason    string
	Suspended bool
}

func init() {
	db := database.GetInstance()
	db.AutoMigrate(&UnsuspensionRequest{})
	db.DropTableIfExists(&ReportRequest{})
	db.AutoMigrate(&ReportRequest{})
	db.AutoMigrate(&SuspensionList{})

	db.Create(&ReportRequest{
		ReporterID:  1,
		SuspectedID: 4,
		Reason:      "Toxic Player",
	})

	db.Create(&ReportRequest{
		ReporterID:  2,
		SuspectedID: 4,
		Reason:      "Toxic Player",
	})

	db.Create(&ReportRequest{
		ReporterID:  3,
		SuspectedID: 4,
		Reason:      "Toxic Player",
	})

	db.Create(&ReportRequest{
		ReporterID:  5,
		SuspectedID: 4,
		Reason:      "Toxic Player",
	})
}
