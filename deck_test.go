package main

import(
	"testing"
	"os"
)

func testNewDeck(t *testing.T){
	sampleDeck := newDeck()

	if(len(sampleDeck) != 24){
		t.Errorf("Expected Deck of length 24 , but got %v",len(sampleDeck))
	}

	if sampleDeck[0] != "AceofSpades" {
		t.Errorf("Expected first card to be Ace of spades, but got %v",sampleDeck[0])
	}

	if sampleDeck[len(sampleDeck) - 1] != "FiveofClub" {
		t.Errorf("Expected last card to be Five of clubs, but got %v",sampleDeck[len(sampleDeck) - 1])
	}
} 

func testSaveToFileAndNewDeckFromFile(t *testing.T){
	os.Remove("_deckTest")

	dec := newDeck()
	dec.saveToFile("_deckTest")

	loadedDeck := newDeckFromFile("_deckTest")

	if len(loadedDeck) != 24 {
		t.Errorf("Expected Deck of length 24 , but got %v",len(loadedDeck))
	}
	os.Remove("_deckTest")
}