package item

type Item struct {
	ItemID     int
	Price      float64
	Quantity   int
	CategoryID int
	SellerID   int
}

func NewItem(itemID, sellerID int, price float64, quantity int, categoryId int) *Item {
	return &Item{
		ItemID:     itemID,
		Price:      price,
		Quantity:   quantity,
		CategoryID: categoryId,
		SellerID:   sellerID,
	}
}

func (i *Item) GetTotalPrice() float64 {
	return i.Price * float64(i.Quantity)
}

func (i *Item) Equals(other *Item) bool {
	return i.ItemID == other.ItemID &&
		i.Price == other.Price &&
		i.CategoryID == other.CategoryID &&
		i.SellerID == other.SellerID
}
