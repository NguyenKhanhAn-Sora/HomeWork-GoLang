package card

import "fmt"

type Card struct {
	ID          int
	CardNumber  string
	OwnerID     int
	ExpiryDate  string
	CVV         int
	Balance     float64
	InitBalance float64
}

var cards []Card

func FindCardByID(id int) *Card {
	for i := range cards {
		if cards[i].ID == id {
			return &cards[i]
		}
	}

	return nil
}

func GetCards() {
	fmt.Println("Getting all cards:")
	for _, card := range cards {
		fmt.Printf("Card ID: %d, Card Number: %s, Owner ID: %d, Expiry Date: %s, CVV: %d, Balance: %.2f, InitBalance: %.2f\n", card.ID, card.CardNumber, card.OwnerID, card.ExpiryDate, card.CVV, card.Balance, card.InitBalance)
	}
}

func AddCard(id int, card_number string, owner_id int, expiry_date string, cvv int) {

	card := Card{
		ID:          id,
		CardNumber:  card_number,
		OwnerID:     owner_id,
		ExpiryDate:  expiry_date,
		CVV:         cvv,
		Balance:     0,
		InitBalance: 50_000,
	}
	cards = append(cards, card)
}

func UpdateCard(id int, card Card) {
	existCard := FindCardByID(id)
	if existCard != nil {
		*existCard = card
	}
}

func DeleteCard(id int) {
	existCard := FindCardByID(id)
	if existCard != nil {
		for i, card := range cards {
			if card.ID == id {
				cards = append(cards[:i], cards[i+1:]...)
				break
			}
		}
	}
}

func withDraw(card *Card, amount float64) bool {
	existCard := FindCardByID(card.ID)
	if existCard == nil {
		fmt.Println("Error Card Not Found!")
		return false
	}
	if card.Balance >= amount {
		card.Balance -= amount
		return true
	}
	return false
}

func deposit(card *Card, amount float64, cardReceive *Card) bool {
	existCard := FindCardByID(card.ID)
	existCardReceive := FindCardByID(cardReceive.ID)
	if existCard == nil || existCardReceive == nil {
		fmt.Println("Error Card Not Found!")
		return false
	}
	card.Balance -= amount
	cardReceive.Balance += amount
	return true
}
