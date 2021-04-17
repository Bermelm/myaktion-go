package model

type Account struct {
	//gorm.Model
	Name     string `gorm:"notNull;size:60"` //das sind letztendlich SQL Abfragen!
	BankName string `gorm:"notNull;size:40"`
	Number   string `gorm:"notNull;size:20"`
}
