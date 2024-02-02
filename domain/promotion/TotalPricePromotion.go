package promotion

import (
	"checkoutCaseGoProject/common"
	"checkoutCaseGoProject/domain/item"
)

type TotalPricePromotion struct {
	PromotionID int
	Discount    *float64
}

func NewTotalPricePromotion() Promotion {
	return &TotalPricePromotion{
		PromotionID: common.TotalPricePromotionID,
		Discount:    nil,
	}
}

func (tp *TotalPricePromotion) GetTotalPrice(items []item.Item, vasItems []item.VasItem) float64 {
	var totalPrice float64
	for _, it := range items {
		totalPrice += it.GetTotalPrice()
	}
	for _, it := range vasItems {
		totalPrice += it.GetTotalPrice()
	}
	return totalPrice
}

func (tp *TotalPricePromotion) ApplyPromotion(items []item.Item, vasItems []item.VasItem) float64 {

	totalPrice := tp.GetTotalPrice(items, vasItems)
	var discount float64
	if totalPrice >= 500 && totalPrice < 5000 {
		discount = 250
	} else if totalPrice >= 5000 && totalPrice < 10000 {
		discount = 500
	} else if totalPrice >= 10000 && totalPrice < 50000 {
		discount = 1000
	} else if totalPrice >= 50000 {
		discount = 2000
	} else {
		discount = 0
	}
	tp.Discount = &discount
	return discount
}

func (cp *TotalPricePromotion) GetDiscount() float64 {
	return *cp.Discount
}
