package item

type VasItem struct {
	VasItemID     int
	VasCategoryID int
	VasSellerID   int
	ItemID        int
	Price         float64
	Quantity      int
}

func NewVasItem(itemID, vasItemID, vasCategoryID, vasSellerID int, price float64, quantity int) *VasItem {
	return &VasItem{
		ItemID:        itemID,
		Price:         price,
		Quantity:      quantity,
		VasCategoryID: vasCategoryID,
		VasSellerID:   vasSellerID,
		VasItemID:     vasItemID,
	}
}

func (i *VasItem) GetTotalPrice() float64 {
	return i.Price * float64(i.Quantity)
}
