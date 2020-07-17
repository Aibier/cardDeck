package main

import (
	"fmt"
	"io/ioutil"
	"math/rand"
	"os"
	"strings"
	"time"
)

type  deck []string

func newDeck() deck {
	cards := deck{}
	cardSuits := []string{"Diamonds", "Spades", "Hearts", "Clubs"}
	cardValues := []string{"Ace", "Two", "Three", "Four", "Five", "Six", "Seven", "Eight", "Nine", "Ten"}

	for _, suit :=range cardSuits {
		for _, value :=range cardValues {
			cards = append(cards, value + " of " + suit)
		}
	}
	return cards
}
func (d deck) print() {
	for index, value :=range d {
		fmt.Println(index, value)
	}
}

func deal(d deck, handSize int) (deck, deck){
	return d[:handSize], d[handSize:]
}

func (d deck) toString() string {
	return strings.Join([]string(d), ",")
}

func (d deck) saveToFile(fileName string )  error{
	return ioutil.WriteFile(fileName, []byte(d.toString()) , 0666 )
}

func newDeckFromFile(fileName string) deck {
	bs, err := ioutil.ReadFile(fileName)
	if err != nil {
		fmt.Println("Error: ", err)
		os.Exit(1)
	}
	return deck(strings.SplitAfter(string(bs), ","))
}

func (d deck) shuffle()  {
	source := rand.NewSource(time.Now().UnixNano())
	r := rand.New(source)
	for index :=range d {
		newPosition := r.Intn(len(d))
		d[index], d[newPosition] = d[newPosition], d[index]
	}
}