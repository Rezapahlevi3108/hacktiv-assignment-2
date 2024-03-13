package models

type Item struct {
	ItemId      int64  `gorm:"primaryKey" json:"item_id"`
	ItemCode    int    `gorm:"type:integer" json:"item_code"`
	Description string `gorm:"type:text" json:"description"`
	Quantity    int    `gorm:"type:integer" json:"quantity"`
	OrderId     int64  `gorm:"index" json:"order_id"`
	Order       Order
}
