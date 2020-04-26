package main

import (
	"os"
	"testing"
)

func TestNewDeck(t *testing.T) {
	d := newDeck()
	if len(d) != 16 {
		t.Errorf("Expected length of deck 16 but got %v", len(d))
	}
	if d[0] != "Ace of Spades" {
		t.Errorf("Expected first position card to be Ace of Spades but got %v", d[0])
	}
	if d[len(d)-1] != "Four of Clubs" {
		t.Errorf("Expected last position card to be Four of Clubs but got %v", d[len(d)-1])
	}
}

func TestSaveToFileAndNewDeckFromFile(t *testing.T) {
	os.Remove("_decktesting")

	d := newDeck()
	//Save it
	d.saveToFile("_decktesting")
	//load it
	newDeck := newDeckFromFile("_decktesting")

	if len(newDeck) != 16 {
		t.Errorf("Expected new file with lenght 16, but got lenght %v", len(newDeck))
	}

	os.Remove("_decktesting")

}
