package models

import "time"

type DebtStatus string

const (
	DebtStatusActive  DebtStatus = "active"
	DebtStatusPartial DebtStatus = "partial"
	DebtStatusPaid    DebtStatus = "paid"
	DebtStatusBadDebt DebtStatus = "bad_debt"
)

type Debt struct {
	ID           uint       `gorm:"primaryKey;autoIncrement" json:"id"`
	CustomerName string     `gorm:"not null;type:text" json:"customer_name"`
	Phone        *string    `gorm:"type:text" json:"phone,omitempty"`
	TotalDebt    float64    `gorm:"not null;default:0" json:"total_debt"`
	PaidAmount   float64    `gorm:"not null;default:0" json:"paid_amount"`
	Status       DebtStatus `gorm:"not null;default:'active';check:status IN ('active','partial','paid','bad_debt')" json:"status"`
	Note         *string    `gorm:"type:text" json:"note,omitempty"`
	CreatedAt    time.Time  `gorm:"not null;default:CURRENT_TIMESTAMP" json:"created_at"`
	UpdatedAt    time.Time  `gorm:"not null;default:CURRENT_TIMESTAMP" json:"updated_at"`

	// Relations
	Payments []DebtPayment `gorm:"foreignKey:DebtID;constraint:OnDelete:CASCADE" json:"payments,omitempty"`
}

type DebtPayment struct {
	ID        uint      `gorm:"primaryKey;autoIncrement" json:"id"`
	DebtID    uint      `gorm:"not null;index" json:"debt_id"`
	Amount    float64   `gorm:"not null" json:"amount"`
	Note      *string   `gorm:"type:text" json:"note,omitempty"`
	CreatedAt time.Time `gorm:"not null;default:CURRENT_TIMESTAMP" json:"created_at"`

	// Relations
	Debt Debt `gorm:"foreignKey:DebtID;constraint:OnDelete:CASCADE" json:"-"`
}
