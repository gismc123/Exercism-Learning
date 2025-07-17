package blackjack

// ParseCard returns the integer value of a card following blackjack ruleset.
func ParseCard(card string) int {
	// panic("Please implement the ParseCard function")
	switch card {
	case "ace":
		var value = 11
		return value
	case "two":
		var value = 2
		return value
	case "three":
		var value = 3
		return value
	case "four":
		var value = 4
		return value
	case "five":
		var value = 5
		return value
	case "six":
		var value = 6
		return value
	case "seven":
		var value = 7
		return value
	case "eight":
		var value = 8
		return value
	case "nine":
		var value = 9
		return value
	case "ten":
		var value = 10
		return value
	case "jack":
		var value = 10
		return value
	case "queen":
		var value = 10
		return value
	case "king":
		var value = 10
		return value
	case "joker":
		var value = 0
		return value
	default:
		var value = 0
		return value
	}
}

// FirstTurn returns the decision for the first turn, given two cards of the
// player and one card of the dealer.
func FirstTurn(card1, card2, dealerCard string) string {
	// panic("Please implement the FirstTurn function")

	var newCard1 = ParseCard(card1)
	var newCard2 = ParseCard(card2)
	var playerHand = newCard1 + newCard2

	switch playerHand {
	case 22:
		return "P"
	case 21:
		if ParseCard(dealerCard) > 9 {
			return "S"
		} else {
			return "W"
		}
	case 20:
		return "S"
	case 19:
		return "S"
	case 18:
		return "S"
	case 17:
		return "S"
	case 16:
		if ParseCard(dealerCard) > 6 {
			return "H"
		} else {
			return "S"
		}
	case 15:
		if ParseCard(dealerCard) > 6 {
			return "H"
		} else {
			return "S"
		}
	case 14:
		if ParseCard(dealerCard) > 6 {
			return "H"
		} else {
			return "S"
		}
	case 13:
		if ParseCard(dealerCard) > 6 {
			return "H"
		} else {
			return "S"
		}
	case 12:
		if ParseCard(dealerCard) > 6 {
			return "H"
		} else {
			return "S"
		}
	case 11:
		return "H"
	case 10:
		return "H"
	case 9:
		return "H"
	case 8:
		return "H"
	case 7:
		return "H"
	case 6:
		return "H"
	case 5:
		return "H"
	case 4:
		return "H"
	case 3:
		return "H"
	default:
		return "H"
	}

}
