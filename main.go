package main

import (
	"checkoutCaseGoProject/application/command"
	cartService "checkoutCaseGoProject/domain/cart"
	"encoding/json"
	"fmt"
	"os"
)

type CommandList struct {
	Commands []Command `json:"commands"`
}

type Command struct {
	Command string                 `json:"command"`
	Payload map[string]interface{} `json:"payload"`
}

type Payload struct {
	ItemId        int `json:"itemId,omitempty"`
	CategoryId    int `json:"categoryId,omitempty"`
	SellerId      int `json:"sellerId,omitempty"`
	Price         int `json:"price,omitempty"`
	Quantity      int `json:"quantity,omitempty"`
	VasItemId     int `json:"vasItemId,omitempty"`
	VasCategoryId int `json:"vasCategoryId,omitempty"`
	VasSellerId   int `json:"vasSellerId,omitempty"`
}

func main() {
	// Dosya okuma
	content, err := os.ReadFile("resources/commands.json")
	if err != nil {
		fmt.Println("Error reading file:", err)
		return
	}

	var cl CommandList
	err = json.Unmarshal([]byte(content), &cl)

	if err != nil {
		fmt.Println("Error parsing JSON:", err)
		return
	}

	cart := cartService.NewCart()

	for _, cmd := range cl.Commands {
		processCommand(cmd, cart)
	}
}

func processCommand(cmd Command, cart *cartService.Cart) {
	// Komut işleme
	switch cmd.Command {
	case "addItem":
		// addItem işlemleri
		item := command.PayloadToItem(cmd.Payload)
		fmt.Println(cart.AddItem(item))

		fmt.Println("Adding item:", cmd.Payload)
	case "addVasItemToItem":
		// addVasItemToItem işlemleri
		item := command.PayloadToVasItem(cmd.Payload)
		fmt.Println(cart.AddVasItemToItem(item))

		fmt.Println("Adding VAS item to item:", cmd.Payload)
	case "displayCart":
		// displayCart işlemleri
		fmt.Println(cart.DisplayCart())
		fmt.Println("Displaying cart")
	default:
		fmt.Println("Unknown command:", cmd.Command)
	}
}
