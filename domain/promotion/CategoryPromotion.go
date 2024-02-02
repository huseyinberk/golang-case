package promotion

import (
	"checkoutCaseGoProject/common"
	"checkoutCaseGoProject/domain/item"
)

type CategoryPromotion struct {
	PromotionID int
	Discount    *float64
}

func NewCategoryPromotion() Promotion {
	return &CategoryPromotion{
		PromotionID: common.CategoryPromotionID,
		Discount:    nil,
	}
}

func (tp *CategoryPromotion) GetTotalPrice(items []item.Item, vasItems []item.VasItem) float64 {
	var totalPrice float64
	for _, it := range items {
		if it.CategoryID == common.FixedCategoryPromotionCategoryId {
			totalPrice += it.GetTotalPrice()
		}
	}
	return totalPrice
}

func (cp *CategoryPromotion) ApplyPromotion(items []item.Item, vasItems []item.VasItem) float64 {

	discountedPrice := cp.GetTotalPrice(items, vasItems) * common.DiscountRateCategory

	cp.Discount = &discountedPrice

	return discountedPrice
}

func (cp *CategoryPromotion) GetDiscount() float64 {
	return *cp.Discount
}
