package models

import "time"

type Product struct {
	Barcode   string    `gorm:"primaryKey;type:text" json:"barcode"`
	Name      string    `gorm:"not null;type:text" json:"name"`
	BuyPrice  float64   `gorm:"not null;default:0" json:"buy_price"`
	SellPrice float64   `gorm:"not null;default:0" json:"sell_price"`
	Stock     int       `gorm:"not null;default:0" json:"stock"`
	MinStock  int       `gorm:"not null;default:0" json:"min_stock"`
	Type      *string   `gorm:"type:text" json:"type,omitempty"`
	Image     *string   `gorm:"type:text" json:"image,omitempty"`
	CreatedAt time.Time `gorm:"not null;default:CURRENT_TIMESTAMP" json:"created_at"`
	UpdatedAt time.Time `gorm:"not null;default:CURRENT_TIMESTAMP" json:"updated_at"`

	// Relations
	TransactionItems []TransactionItem `gorm:"foreignKey:ProductBarcode;references:Barcode" json:"-"`
	StockMovements   []StockMovement   `gorm:"foreignKey:ProductBarcode;references:Barcode" json:"-"`
}
