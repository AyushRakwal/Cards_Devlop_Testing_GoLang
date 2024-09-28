package main

import (
	"fmt"
	"strings"
	"io/ioutil"
	"math/rand"
	"os"
	"time"
)

type deck [] string

func newDeck() deck {
	cards := deck{}

	cardSuit := [] string {"Spades" , "Diamonds" , "Heart" , "Club"}
	cardValue := [] string {"Ace" , "One" , "Two" , "Three" , "Four" , "Five"}

	for _,suit := range cardSuit {
		for _,value := range cardValue {
			cards = append(cards , value + "of" + suit)
		}
	}

	return cards
}

func deal(this deck , handSize int )   (deck , deck) {  // Return multiple return statements
	return this[:handSize]  , this[handSize:]   // Divided the splice into two parts using [start(inclusive) : end(exclusive)]
}

func (this deck)print(){
	for idx, item := range this {
		fmt.Println(idx , item)
	}
}

func (this deck) toString() string {
	deckToString := []string(this)
	return strings.Join(deckToString,",")
}

func (this deck) saveToFile(filename string) error{
	return ioutil.WriteFile(filename,[]byte(this.toString()),0666)
}


func newDeckFromFile(filename string) (deck) {
	byteSlice , err := os.ReadFile(filename)

	if err != nil {
		// log out from game by displaying the error
		fmt.Println("ERROR : ",err)
		os.Exit(1)
	}

	stringSlice := strings.Split(string(byteSlice),",")
	return deck(stringSlice)
}

func (this deck) shuffle() {
	source := rand.NewSource(time.Now().UnixNano())
	r := rand.New(source)
	
	for idx := range this {
		newPos := r.Intn(len(this) - 1)

		this[idx] , this[newPos] = this[newPos] , this[idx]
	} 
}