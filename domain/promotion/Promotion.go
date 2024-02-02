package promotion

import (
	"checkoutCaseGoProject/domain/item"
)

type Promotion interface {
	ApplyPromotion(items []item.Item, vasItems []item.VasItem) float64
	GetDiscount() float64
}
