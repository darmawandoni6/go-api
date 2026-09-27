package models

import "time"

type CashFlowType string

const (
	CashFlowIncome  CashFlowType = "income"
	CashFlowExpense CashFlowType = "expense"
)

type CashFlow struct {
	ID        uint         `gorm:"primaryKey;autoIncrement" json:"id"`
	Type      CashFlowType `gorm:"not null;check:type IN ('income','expense')" json:"type"`
	Amount    float64      `gorm:"not null" json:"amount"`
	Note      string       `gorm:"not null;default:'';type:text" json:"note"`
	Date      string       `gorm:"not null;type:text" json:"date"`
	CreatedAt time.Time    `gorm:"not null;default:CURRENT_TIMESTAMP" json:"created_at"`
}
