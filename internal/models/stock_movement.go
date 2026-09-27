package models

import "time"

type StockMovementType string

const (
	StockMovementRestock        StockMovementType = "restock"
	StockMovementSale           StockMovementType = "sale"
	StockMovementAdjustmentLoss StockMovementType = "adjustment_loss"
	StockMovementAdjustmentGain StockMovementType = "adjustment_gain"
)

type StockMovement struct {
	ID             uint              `gorm:"primaryKey;autoIncrement" json:"id"`
	ProductBarcode string            `gorm:"not null;type:text;index" json:"product_barcode"`
	Type           StockMovementType `gorm:"not null;check:type IN ('restock','sale','adjustment_loss','adjustment_gain')" json:"type"`
	Quantity       int               `gorm:"not null" json:"quantity"`
	PreviousStock  int               `gorm:"not null" json:"previous_stock"`
	FinalStock     int               `gorm:"not null" json:"final_stock"`
	TotalCost      float64           `gorm:"default:0" json:"total_cost"`
	Note           *string           `gorm:"type:text" json:"note,omitempty"`
	CreatedAt      time.Time         `gorm:"not null;default:CURRENT_TIMESTAMP" json:"created_at"`

	// Relations
	Product Product `gorm:"foreignKey:ProductBarcode;references:Barcode;constraint:OnDelete:CASCADE" json:"-"`
}
