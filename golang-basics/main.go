package main

func main() {
   cards := deck{newCard(), newCard(), "Yusuf"}
   cards = append(cards, "Sunnatullayev")

   cards.print()
}

func newCard() string {
	return "Five of Diamonds"
}