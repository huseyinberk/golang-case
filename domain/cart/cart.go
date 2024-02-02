package cart

import (
	"checkoutCaseGoProject/common"
	"checkoutCaseGoProject/domain/item"
	"checkoutCaseGoProject/domain/promotion"
)

type Cart struct {
	Items               []item.Item
	VasItems            []item.VasItem
	SelectedPromotion   promotion.Promotion
	AvailablePromotions []promotion.Promotion
}

func NewCart() *Cart {
	cart := &Cart{
		Items:    []item.Item{},
		VasItems: []item.VasItem{},
		AvailablePromotions: []promotion.Promotion{
			promotion.NewSameSellerPromotion(),
			promotion.NewCategoryPromotion(),
			promotion.NewTotalPricePromotion(),
		},
	}
	return cart
}

func (c *Cart) GetTotalPrice() float64 {
	var totalPrice float64
	for _, it := range c.Items {
		totalPrice += it.GetTotalPrice()
	}
	for _, it := range c.VasItems {
		totalPrice += it.GetTotalPrice()
	}
	return totalPrice
}

func (c *Cart) totalNumberOfItemsExceedLimitWithNewItem(quantity int) bool {
	count := quantity + c.TotalNumberOfItems()
	return count > common.CartItemLimit
}

func (c *Cart) TotalNumberOfItems() int {
	count := 0
	for _, it := range c.Items {
		count += it.Quantity
	}
	for _, it := range c.VasItems {
		count += it.Quantity
	}
	return count
}

func (c *Cart) totalPriceOfItemsExceedLimitWithNewItem(price float64) bool {
	return price+c.GetTotalPrice() > common.CartPriceLimit
}

func (c *Cart) uniqueNumberOfItemsExceedLimitWithNewItem() bool {
	return len(c.Items) >= common.CartUniqueItemLimit
}

func (c *Cart) isItemInTheCart(it *item.Item) *item.Item {
	for i, item := range c.Items {
		if item.Equals(it) {
			return &c.Items[i]
		}
	}
	return nil
}

func (c *Cart) isDigitalItem(id int) bool {
	return id == common.FixedCategoryIDDigitalItem
}

func (c *Cart) digitalItemQuantityExceedLimitWithNewItem(it item.Item, quantity int) bool {
	return it.Quantity+quantity > common.DigitalItemQuantityLimit
}

func (c *Cart) AddItem(it item.Item) string {
	if c.totalNumberOfItemsExceedLimitWithNewItem(it.Quantity) {
		return "total number of item exceed limit"
	}
	if c.totalPriceOfItemsExceedLimitWithNewItem(it.GetTotalPrice()) {
		return "total price of item exceed limit"
	}
	item := c.isItemInTheCart(&it)
	if item != nil {
		if c.isDigitalItem(item.CategoryID) {
			if c.digitalItemQuantityExceedLimitWithNewItem(*item, it.Quantity) {
				return "digital number of item quantity exceed limit"
			}
		}
		item.Quantity += it.Quantity
	} else {
		if c.uniqueNumberOfItemsExceedLimitWithNewItem() {
			return "unique number of item exceed limit"
		}
		c.Items = append(c.Items, it)
	}
	c.SetMostAdvantageousPromotion()
	return "item_added"
}

func (c *Cart) AddVasItemToItem(vasIt item.VasItem) string {
	if !c.IsTheCategoryIdCorrect(vasIt.VasCategoryID) {
		return "vas item categori idsi hatalı"
	}
	if !c.IsTheSellerIdCorrect(vasIt.VasSellerID) {
		return "vas item seller idsi hatalı"
	}
	if !c.IsTheDefaultItem(vasIt.ItemID) {
		return "bu category idli ürüne eklenemez"
	}
	if c.totalNumberOfItemsExceedLimitWithNewItem(vasIt.Quantity) {
		return "total number of item exceed limit"
	}
	if c.totalNumberOfVasItemsExceedLimitWithNewItem(vasIt.Quantity) {
		return "3ten fazla oldu"
	}
	if c.totalPriceOfItemsExceedLimitWithNewItem(vasIt.GetTotalPrice()) {
		return "total price of item exceed limit"
	}
	c.VasItems = append(c.VasItems, vasIt)
	c.SetMostAdvantageousPromotion()
	return "item_added"
}

func (c *Cart) IsTheDefaultItem(id int) bool {
	for _, it := range c.Items {
		if it.ItemID == id {
			if it.CategoryID == common.FixedVasItemDefaultItemCategoryID_1 || it.CategoryID == common.FixedVasItemDefaultItemCategoryID_2 {
				return true
			}
		}
	}
	return false
}

func (c *Cart) IsTheCategoryIdCorrect(id int) bool {
	if id == common.FixedVasItemCategoryID {
		return true
	}
	return false
}

func (c *Cart) IsTheSellerIdCorrect(id int) bool {
	if id == common.FixedVasItemSellerId {
		return true
	}
	return false
}

func (c *Cart) totalNumberOfVasItemsExceedLimitWithNewItem(quantity int) bool {
	count := quantity
	for _, it := range c.VasItems {
		count += it.Quantity
	}
	return count > common.CartVasItemLimit
}

func (c *Cart) SetMostAdvantageousPromotion() {
	for _, promotion := range c.AvailablePromotions {
		promotion.ApplyPromotion(c.Items, c.VasItems)
		if c.SelectedPromotion == nil {
			c.SelectedPromotion = promotion
		}
		if promotion.GetDiscount() > c.SelectedPromotion.GetDiscount() {
			c.SelectedPromotion = promotion
		}
	}

}
