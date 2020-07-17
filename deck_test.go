package main

import (
	"os"
	"testing"
)

func TestBewDeck(t *testing.T)  {
	deck := newDeck()

	if len(deck) != 40 {
		t.Errorf("Expected 40 cards in the deck, but got %v", len(deck))
	}
	if deck[0] != "Ace of Diamonds" {
		t.Errorf("Expected Ace of Diamonds, but got %v", deck[0])
	}
	if deck[len(deck) - 1] != "Ten of Clubs" {
		t.Errorf("Expected Ten of Clubs, but got %v", deck[len(deck) - 1])
	}
}

func TestSaveToFileAndNewDeckFromFile(t *testing.T)  {
	os.Remove("_test_deck")

	deck := newDeck()
	deck.saveToFile("_test_deck")

	loadedDeck := newDeckFromFile("_test_deck")
	if len(loadedDeck) != 40 {
		t.Errorf("Expected 40 cards in the deck, but got %v", len(deck))
	}
	os.Remove("_test_deck")
}