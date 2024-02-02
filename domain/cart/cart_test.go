package cart

import (
	"checkoutCaseGoProject/common"
	"checkoutCaseGoProject/domain/item"
	"testing"
)

func TestAddItem(t *testing.T) {

	cart := NewCart()
	for i := 0; i < common.CartItemLimit; i++ {
		testItem := item.Item{
			ItemID:     i % 10,
			Price:      100.0,
			Quantity:   1,
			CategoryID: 1,
			SellerID:   4,
		}
		response := cart.AddItem(testItem)
		if response != "item_added" {
			t.Errorf("Expected item_added, got %s instead", response)
		}

		if cart.TotalNumberOfItems() != i+1 {
			t.Errorf("Expected %d items in cart, found %d", i+1, cart.TotalNumberOfItems())
		}
	}

	extraItem := item.Item{
		ItemID:     111,
		Price:      100.0,
		Quantity:   1,
		CategoryID: 1,
		SellerID:   4,
	}

	// Test adding one more item than the limit
	response := cart.AddItem(extraItem)
	if response != "total number of item exceed limit" {
		t.Errorf("Expected total number of item exceed limit, got %s instead", response)
	}

	if cart.TotalNumberOfItems() != common.CartItemLimit {
		t.Errorf("Cart items should not exceed %d, found %d", common.CartItemLimit, cart.TotalNumberOfItems())
	}
}

func TestAddItemUniqueItemLimitExceed(t *testing.T) {

	cart := NewCart()
	for i := 0; i < common.CartUniqueItemLimit; i++ {
		testItem := item.Item{
			ItemID:     i % 10,
			Price:      100.0,
			Quantity:   1,
			CategoryID: 1,
			SellerID:   4,
		}
		response := cart.AddItem(testItem)
		if response != "item_added" {
			t.Errorf("Expected item_added, got %s instead", response)
		}

		if cart.TotalNumberOfItems() != i+1 {
			t.Errorf("Expected %d items in cart, found %d", i+1, cart.TotalNumberOfItems())
		}
	}

	extraItem := item.Item{
		ItemID:     10,
		Price:      100.0,
		Quantity:   1,
		CategoryID: 1,
		SellerID:   4,
	}

	// Test adding one more item than the limit
	response := cart.AddItem(extraItem)
	if response != "unique number of item exceed limit" {
		t.Errorf("Expected unique number of item exceed limit, got %s instead", response)
	}

	if cart.TotalNumberOfItems() != common.CartUniqueItemLimit {
		t.Errorf("Cart items should not exceed %d, found %d", common.CartUniqueItemLimit, cart.TotalNumberOfItems())
	}
}

func TestDigitalItemLimitExceed(t *testing.T) {

	cart := NewCart()
	for i := 0; i < common.DigitalItemQuantityLimit; i++ {
		testItem := item.Item{
			ItemID:     i % 10,
			Price:      100.0,
			Quantity:   1,
			CategoryID: 7889,
			SellerID:   4,
		}
		response := cart.AddItem(testItem)
		if response != "item_added" {
			t.Errorf("Expected item_added, got %s instead", response)
		}

		if cart.TotalNumberOfItems() != i+1 {
			t.Errorf("Expected %d items in cart, found %d", i+1, cart.TotalNumberOfItems())
		}
	}

	extraItem := item.Item{
		ItemID:     0,
		Price:      100.0,
		Quantity:   5,
		CategoryID: 7889,
		SellerID:   4,
	}

	// Test adding one more item than the limit
	response := cart.AddItem(extraItem)
	if response != "digital number of item quantity exceed limit" {
		t.Errorf("Expected digital number of item quantity exceed limit, got %s instead", response)
	}

	if cart.TotalNumberOfItems() != common.DigitalItemQuantityLimit {
		t.Errorf("Cart items should not exceed %d, found %d", common.CartUniqueItemLimit, cart.TotalNumberOfItems())
	}
}

func TestAddItemTotalPriceLimitExceed(t *testing.T) {

	cart := NewCart()
	for i := 0; i < common.CartUniqueItemLimit; i++ {
		testItem := item.Item{
			ItemID:     i % 10,
			Price:      49000.0,
			Quantity:   1,
			CategoryID: 1,
			SellerID:   4,
		}
		response := cart.AddItem(testItem)
		if response != "item_added" {
			t.Errorf("Expected item_added, got %s instead", response)
		}

		if cart.TotalNumberOfItems() != i+1 {
			t.Errorf("Expected %d items in cart, found %d", i+1, cart.TotalNumberOfItems())
		}
	}

	extraItem := item.Item{
		ItemID:     0,
		Price:      11000.0,
		Quantity:   1,
		CategoryID: 1,
		SellerID:   4,
	}

	// Test adding one more item than the limit
	response := cart.AddItem(extraItem)
	if response != "total price of item exceed limit" {
		t.Errorf("Expected total price of item exceed limit, got %s instead", response)
	}

	if cart.TotalNumberOfItems() != common.CartUniqueItemLimit {
		t.Errorf("Cart items should not exceed %d, found %d", common.CartUniqueItemLimit, cart.TotalNumberOfItems())
	}
}

func TestAddVasItemToItem(t *testing.T) {
	cart := NewCart()

	defaultItem := item.Item{
		ItemID:     1,
		Price:      100.0,
		Quantity:   1,
		CategoryID: common.FixedVasItemDefaultItemCategoryID_1,
	}
	cart.Items = append(cart.Items, defaultItem)

	digitalItem := item.Item{
		ItemID:     11,
		Price:      100.0,
		Quantity:   1,
		CategoryID: 11,
	}
	cart.Items = append(cart.Items, digitalItem)

	vasItem := item.VasItem{
		VasItemID:     22,
		ItemID:        1,
		Price:         10.0,
		Quantity:      1,
		VasCategoryID: common.FixedVasItemCategoryID,
		VasSellerID:   common.FixedVasItemSellerId,
	}

	response := cart.AddVasItemToItem(vasItem)
	if response != "item_added" {
		t.Errorf("Expected item_added, got %s instead", response)
	}

	vasItem.VasCategoryID = 123
	response = cart.AddVasItemToItem(vasItem)
	if response != "vas item categori idsi hatalı" {
		t.Errorf("Expected vas item categori idsi hatalı, got %s instead", response)
	}

	vasItem.VasCategoryID = common.FixedVasItemCategoryID
	vasItem.VasSellerID = 123
	response = cart.AddVasItemToItem(vasItem)
	if response != "vas item seller idsi hatalı" {
		t.Errorf("Expected vas item seller idsi hatalı, got %s instead", response)
	}

	vasItem.VasSellerID = common.FixedVasItemSellerId
	for i := 0; i < common.CartVasItemLimit; i++ {
		cart.AddVasItemToItem(vasItem)
	}
	response = cart.AddVasItemToItem(vasItem)
	if response != "3ten fazla oldu" {
		t.Errorf("Expected 3ten fazla oldu, got %s instead", response)
	}

	vasItem.ItemID = 11
	response = cart.AddVasItemToItem(vasItem)
	if response != "bu category idli ürüne eklenemez" {
		t.Errorf("Expected bu category idli ürüne eklenemez, got %s instead", response)
	}

}

func TestAddVasItem(t *testing.T) {

	cart := NewCart()
	for i := 0; i < common.CartItemLimit; i++ {
		testItem := item.Item{
			ItemID:     i % 10,
			Price:      100.0,
			Quantity:   1,
			CategoryID: common.FixedVasItemDefaultItemCategoryID_1,
			SellerID:   4,
		}
		response := cart.AddItem(testItem)
		if response != "item_added" {
			t.Errorf("Expected item_added, got %s instead", response)
		}

		if cart.TotalNumberOfItems() != i+1 {
			t.Errorf("Expected %d items in cart, found %d", i+1, cart.TotalNumberOfItems())
		}
	}

	vasItem := item.VasItem{
		VasItemID:     22,
		ItemID:        1,
		Price:         10.0,
		Quantity:      1,
		VasCategoryID: common.FixedVasItemCategoryID,
		VasSellerID:   common.FixedVasItemSellerId,
	}
	// Test adding one more item than the limit
	response := cart.AddVasItemToItem(vasItem)
	if response != "total number of item exceed limit" {
		t.Errorf("Expected total number of item exceed limit, got %s instead", response)
	}

	if cart.TotalNumberOfItems() != common.CartItemLimit {
		t.Errorf("Cart items should not exceed %d, found %d", common.CartItemLimit, cart.TotalNumberOfItems())
	}
}

func TestAddVasItemTotalPriceExceed(t *testing.T) {

	cart := NewCart()
	for i := 0; i < common.CartUniqueItemLimit; i++ {
		testItem := item.Item{
			ItemID:     i % 10,
			Price:      100.0,
			Quantity:   1,
			CategoryID: common.FixedVasItemDefaultItemCategoryID_1,
			SellerID:   4,
		}
		response := cart.AddItem(testItem)
		if response != "item_added" {
			t.Errorf("Expected item_added, got %s instead", response)
		}

		if cart.TotalNumberOfItems() != i+1 {
			t.Errorf("Expected %d items in cart, found %d", i+1, cart.TotalNumberOfItems())
		}
	}

	vasItem := item.VasItem{
		VasItemID:     3,
		ItemID:        1,
		Price:         500000.0,
		Quantity:      1,
		VasCategoryID: common.FixedVasItemCategoryID,
		VasSellerID:   common.FixedVasItemSellerId,
	}
	// Test adding one more item than the limit
	response := cart.AddVasItemToItem(vasItem)
	if response != "total price of item exceed limit" {
		t.Errorf("Expected total price of item exceed limit, got %s instead", response)
	}

	if cart.TotalNumberOfItems() != common.CartUniqueItemLimit {
		t.Errorf("Cart items should not exceed %d, found %d", common.CartUniqueItemLimit, cart.TotalNumberOfItems())
	}
}
