package main

import "fmt"

type Pizza interface{ Ingredients() }
type CheesePizza struct{ cheeseType string }

func (c CheesePizza) Ingredients() {
	fmt.Println(c.cheeseType)
}

func main() {
	var p1 Pizza
	c1 := CheesePizza{"Emmental"}
	p1 = c1
	//c1.cheeseType = "Emmental"
	p1.Ingredients()

}
