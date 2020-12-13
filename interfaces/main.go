package main

import "fmt"

// Pizza : Let´s make some Pizza
type Pizza interface{ Ingredients() }

// MeatLoverPizza : A type of Pizza for meat lovers with a type of meat for our pizza
type MeatLoverPizza struct {
	meatType string
	base     string
}

// HawaianPizza : A type of Pizza that always has pineapple on it
type HawaianPizza struct {
	base string
}

// Ingredients : Our meat pizza has a type of meat
func (c MeatLoverPizza) Ingredients() {
	fmt.Println("A pizza with " + c.meatType + " on a " + c.base + " base.")
}

// Ingredients : Our Hawaian pizza has Pineapple, of course
func (h HawaianPizza) Ingredients() {
	fmt.Println("A pizza with Pineapple on a " + h.base + " base.")
}

func main() {
	var p1 Pizza
	c1 := MeatLoverPizza{"Peperoni", "Thick"}
	p1 = c1
	p1.Ingredients()

	h1 := HawaianPizza{"Thin"}
	h1.Ingredients()

}
