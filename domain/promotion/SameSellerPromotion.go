package promotion

import (
	"checkoutCaseGoProject/common"
	"checkoutCaseGoProject/domain/item"
)

type SameSellerPromotion struct {
	PromotionID int
	Discount    *float64
}

func NewSameSellerPromotion() Promotion {
	return &SameSellerPromotion{
		PromotionID: common.SameSellerPromotionID,
		Discount:    nil,
	}
}

func (tp *SameSellerPromotion) GetTotalPrice(items []item.Item, vasItems []item.VasItem) float64 {
	var totalPrice float64
	for _, it := range items {
		totalPrice += it.GetTotalPrice()
	}
	for _, it := range vasItems {
		totalPrice += it.GetTotalPrice()
	}
	return totalPrice
}

func (sp *SameSellerPromotion) ApplyPromotion(items []item.Item, vasItems []item.VasItem) float64 {
	discountedPrice := 0.0

	if sp.IsApplicable(items) {
		discountedPrice = sp.GetTotalPrice(items, vasItems) * common.DiscountRateSameSeller
	}

	sp.Discount = &discountedPrice
	return discountedPrice
}

func (sp *SameSellerPromotion) IsApplicable(items []item.Item) bool {
	if len(items) < 2 {
		return true
	}
	firstSellerId := items[0].SellerID
	for _, item := range items[1:] {
		if item.SellerID != firstSellerId {
			return false
		}
	}
	return true
}

func (cp *SameSellerPromotion) GetDiscount() float64 {
	return *cp.Discount
}
