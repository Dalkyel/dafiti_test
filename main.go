package main

import (
	"errors"
	"fmt"
	"log"
	"sort"
)

// isAValidHand Checks if the hand is a valid hand.
func isAValidHand(hand []int) error {
	numberOfCards := len(hand)
	if numberOfCards < 5 || numberOfCards > 7 {
		return errors.New(("Is not a valid hand, must have between 5 and 7 cards"))
	}

	for _, card := range hand {
		if card < 1 || card > 14 {
			return errors.New(fmt.Sprintf("Is not a valid hand, the card %v does not exist", card))
		}
	}

	return nil
}

// DiscardDuplicatedCards Discards the duplicated cards in the hand
func DiscardDuplicatedCards(hand []int) []int {
	keys := make(map[int]bool)
	var newHand []int

	for _, entry := range hand {
		if _, value := keys[entry]; !value {
			keys[entry] = true
			newHand = append(newHand, entry)
		}
	}

	return newHand
}

func checkIfHasTwoAs(hand []int) []int {
	keys := make(map[int]bool)
	var newHand []int

	for _, entry := range hand {
		if _, value := keys[entry]; !value {
			keys[entry] = true
			newHand = append(newHand, entry)
		}

		if entry == 14 || entry == 1 {
			keys[14] = true
			keys[1] = true

		}
	}

	return newHand
}

// isAConsecutiveLadder Validates if the numbers of the hand are in consecutive order
func isAConsecutiveLadder(hand []int) error {
	var newHand []int
	for i := 0; i < len(hand)-1; i++ {
		if (hand[i+1]-hand[i]) != 1 && hand[i] != 2 {
			continue
		}
		newHand = append(newHand, hand[i], hand[i+1])

	}

	// Validates if the hand goes around (14 - 2)
	if hand[len(hand)-1] == 14 && hand[0] == 2 {
		newHand = append(newHand, hand[len(hand)-1], hand[0])
	}

	if newHand != nil && ((hand[len(hand)-1] - newHand[len(newHand)-1]) == 1) {
		newHand = append(newHand, hand[len(hand)-1])
	}

	newHand = DiscardDuplicatedCards(newHand)
	fmt.Println(newHand)
	// Validates if the hand has five consecutive cards
	if len(newHand) < 5 {
		return errors.New("is not a consecutive ladder")
	}

	return nil
}

func IsStraight(hand []int) (string, error) {
	hand = DiscardDuplicatedCards(hand)
	hand = checkIfHasTwoAs(hand)

	sort.Ints(hand)
	log.Println("Sorted Hand", hand)

	err := isAValidHand(hand)
	if err != nil {
		return "", err
	}

	err = isAConsecutiveLadder(hand)
	if err != nil {
		return "", err
	}

	return "Congrats!! you have a valid hand", nil
}

func main() {
	hand := []int{2, 7, 8, 5, 10, 9, 11}

	log.Println("Validating hand", hand)

	defer func() {
		if err := recover(); err != nil {
			log.Println(err)
		}
	}()

	result, err := IsStraight(hand)
	if err != nil {
		panic(err)
	}

	sort.Ints(hand)
	hand = DiscardDuplicatedCards(hand)
	log.Println(result, hand)
}
