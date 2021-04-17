package model

import "gorm.io/gorm"

type Donation struct {
	gorm.Model
	CampaignID       uint
	Amount           float64 `gorm:"notNull;check:amount >= 1.0"`
	DonorName        string  `gorm:"notNull;size:40"`
	ReceiptRequested bool    `gorm:"notNull"`
	Status           Status  `gorm:"notNull;type:ENUM('TRANSFERRED','IN_PROCESS')"`
	Account          Account `gorm:"embedded;embeddedPrefix:account_"`
}

type Status string

const (
	StatusInProcess   Status = "IN_PROCESS"
	StatusTransferred Status = "TRANSFERRED"
)
