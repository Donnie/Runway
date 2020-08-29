package main

import (
	"time"

	tb "gopkg.in/tucnak/telebot.v2"
	"gorm.io/gorm"
)

// Global holds fundamental items
type Global struct {
	Bot *tb.Bot
	Orm *gorm.DB
}

// Record represents one Record
type Record struct {
	ID        *int64     `json:"id" gorm:"primaryKey"`
	CreatedAt *time.Time `json:"created_at"`
	UpdatedAt *time.Time `json:"updated_at"`
	DeletedAt *time.Time `json:"-"`

	Account     *string    `json:"account"`
	Amount      *int64     `json:"amount"`
	Currency    *string    `json:"currency"`
	Date        *time.Time `json:"date"`
	Description *string    `json:"description"`
	Form        *string    `json:"form"`
	FromDate    *time.Time `json:"from_date"`
	TillDate    *time.Time `json:"till_date"`
	Type        *string    `json:"type"`
	UserID      *int64     `json:"user_id"`
}

// Convo stores a conversation context
type Convo struct {
	ID        *int64     `json:"id" gorm:"primaryKey"`
	CreatedAt *time.Time `json:"created_at"`
	UpdatedAt *time.Time `json:"updated_at"`
	DeletedAt *time.Time `json:"-"`

	UserID  *int64  `json:"user_id"`
	Context *string `json:"context"`
}

// Trans defines a money transaction
type Trans struct {
	Name   string `json:"name"`
	Amount int32  `json:"amount"`
}

// Account represents costs and earnings
type Account struct {
	Costs    []Trans `json:"costs"`
	Earnings []Trans `json:"earnings"`
	Savings  []Trans `json:"savings"`
}
