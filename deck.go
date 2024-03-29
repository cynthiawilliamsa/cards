package main

import (
	"fmt"
	"io/ioutil"
	"math/rand"
	"os"
	"strings"
	"time"
)

//creates a deck type (borrows behavior of type of string)
//which is a slice of strings

type deck []string

func newDeck() deck {
	//returns a list of cards
	cards := deck{}
	cardSuits := []string{"Spades", "Diamonds", "Hearts", "Clubs"}
	cardValues := []string{"Ace", "Two", "Three", "Four"}
	//underscore is used when variable declared is not used in the code later
	for _, suit := range cardSuits {
		for _, value := range cardValues {
			cards = append(cards, value+" of "+suit)
		}
	}
	return cards
}

//receiver on a function
//any variable of type deck get access to "print method"
//d in this case refers to instance of type
func (d deck) print() {
	for i, card := range d {
		fmt.Println(i, card)
	}
}

//call deal function with type deck and handsize of int
//2nd () tells go what we are returning
func deal(d deck, handSize int) (deck, deck) {
	return d[:handSize], d[handSize:]

}

func (d deck) toString() string {
	return strings.Join([]string(d), " , ")
}

func (d deck) saveToFile(filename string) error {
	return ioutil.WriteFile(filename, []byte(d.toString()), 0666)
}

func newDeckFromFile(filename string) deck {
	bs, err := ioutil.ReadFile(filename)
	if err != nil {
		fmt.Println("Error:", err)
		os.Exit(1)
	}
	s := strings.Split(string(bs), " , ")
	return deck(s)
}

func (d deck) shuffle() {
	//create randome seed for rand.Intn function
	source := rand.NewSource(time.Now().UnixNano())
	r := rand.New(source)

	for i := range d {
		newPosition := r.Intn(len(d) - 1)
		//swaps elements
		d[i], d[newPosition] = d[newPosition], d[i]
	}
}
