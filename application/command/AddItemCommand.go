package command

import (
	"checkoutCaseGoProject/domain/item"
)

func PayloadToItem(payload map[string]interface{}) item.Item {

	item := item.Item{
		ItemID:     int(payload["itemId"].(float64)),
		CategoryID: int(payload["categoryId"].(float64)),
		SellerID:   int(payload["sellerId"].(float64)),
		Price:      payload["price"].(float64),
		Quantity:   int(payload["quantity"].(float64)),
	}
	return item
}
