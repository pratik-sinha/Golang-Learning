package main

import (
	"fmt"
	"io/ioutil"
	"math/rand"
	"os"
	"strings"
	"time"
)

//Creae a new type of 'deck'
//which is a slice of strings
//type is analogous to classes
type deck []string

//Receivers
// (d deck) instance of working variable
//analogous to this or self in js/python
//any variable of type deck has access to this funtion

func newDeck() deck {
	cards := deck{}
	cardSuits := []string{"Spades", "Diamonds", "Hearts", "Clubs"}
	cardValues := []string{"Ace", "Two", "Three", "Four"}

	for _, suit := range cardSuits {
		for _, value := range cardValues {
			cards = append(cards, suit+" of "+value)
		}
	}

	return cards
}

func (d deck) deal(numberOfCards int) (deck, deck) {
	return d[:numberOfCards], d[numberOfCards:]
}

func (d deck) shuffle() {
	source := rand.NewSource(time.Now().UnixNano())
	r := rand.New(source)
	for i := range d {
		newPos := r.Intn(len(d) - 1)
		d[i], d[newPos] = d[newPos], d[i]
	}
}

func (d deck) toString() string {
	return strings.Join([]string(d), ",")
}

func (d deck) saveToFile(fileName string) error {
	return ioutil.WriteFile(fileName, []byte(d.toString()), 0666)
}

func newDeckFromFile(fileName string) deck {
	byteSlice, err := ioutil.ReadFile(fileName)
	if err != nil {
		//Option #1 - log error and return call to newDeck()
		//Option #2 - log error and entirely quit the program
		fmt.Println("Error: ", err)
		os.Exit(1)
	}
	return deck(strings.Split(string(byteSlice), ","))
}

func (d deck) print() {
	for i, card := range d {
		fmt.Println(i, card)
	}
}
