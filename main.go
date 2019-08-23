package main

func main() {
	//declare slice with square bracket and type of data then curly braces
	//creates new deck of cards
	cards := newDeck()
	cards.shuffle()
	cards.print()

}
