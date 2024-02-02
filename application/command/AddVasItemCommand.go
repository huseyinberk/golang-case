package command

import (
	"checkoutCaseGoProject/domain/item"
)

func PayloadToVasItem(payload map[string]interface{}) item.VasItem {

	item := item.VasItem{
		VasItemID:     int(payload["vasItemId"].(float64)),
		ItemID:        int(payload["itemId"].(float64)),
		VasCategoryID: int(payload["vasCategoryId"].(float64)),
		VasSellerID:   int(payload["vasSellerId"].(float64)),
		Price:         payload["price"].(float64),
		Quantity:      int(payload["quantity"].(float64)),
	}
	return item
}
