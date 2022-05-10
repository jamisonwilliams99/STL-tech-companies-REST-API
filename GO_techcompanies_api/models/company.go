package models

import (
	"gorm.io/gorm"
)

type Companies struct {
	ID             uint   `gorm:"primary key:autoIncrement" json:"id"`
	Name           string `json:"name"`
	Industry       string `json:"industry"`
	Funding        string `json:"funding"`
	Employees      int    `json:"employees"`
	EmployeeGrowth string `json:"employeegrowth"`
	Revenue        string `json:"revenue"`
}

func MigrateCompanies(db *gorm.DB) error {
	err := db.AutoMigrate(&Companies{})
	return err
}
