package models

import "time"

type Transaction struct {
	ID            uint      `gorm:"primaryKey;autoIncrement" json:"id"`
	TotalAmount   float64   `gorm:"not null" json:"total_amount"`
	PaymentAmount float64   `gorm:"not null" json:"payment_amount"`
	ChangeAmount  float64   `gorm:"not null;default:0" json:"change_amount"`
	Note          *string   `gorm:"type:text" json:"note,omitempty"`
	CreatedAt     time.Time `gorm:"not null;default:CURRENT_TIMESTAMP" json:"created_at"`

	// Relations
	Items []TransactionItem `gorm:"foreignKey:TransactionID;constraint:OnDelete:CASCADE" json:"items,omitempty"`
}

type TransactionItem struct {
	ID             uint    `gorm:"primaryKey;autoIncrement" json:"id"`
	TransactionID  uint    `gorm:"not null;index" json:"transaction_id"`
	ProductBarcode string  `gorm:"not null;type:text;index" json:"product_barcode"`
	ProductName    string  `gorm:"not null;type:text" json:"product_name"`
	SellPrice      float64 `gorm:"not null" json:"sell_price"`
	Quantity       int     `gorm:"not null" json:"quantity"`
	Subtotal       float64 `gorm:"not null" json:"subtotal"`

	// Relations (untuk constraint)
	Transaction Transaction `gorm:"foreignKey:TransactionID;constraint:OnDelete:CASCADE" json:"-"`
	Product     Product     `gorm:"foreignKey:ProductBarcode;references:Barcode;constraint:OnDelete:RESTRICT" json:"-"`
}
