package entity

type Product struct {
	Base
	Name     string `json:"name"`
	Price    uint   `json:"price"`
	SellerID uint   `json:"sellerID"`
	Seller   *User  `gorm:"foreignKey:SellerID;references:ID" json:"seller"`
}
