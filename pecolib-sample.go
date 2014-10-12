package main

import (
	"fmt"

	"github.com/reiki4040/peco"
)

type Something struct {
	Name        string
	Description string
	Price       int
}

func NewSomething(name, description string, price int) *Something {
	return &Something{
		Name:        name,
		Description: description,
		Price:       price,
	}
}

func (o *Something) Info() string {
	return fmt.Sprintf("%s : %s price: %d", o.Name, o.Description, o.Price)
}

func (o *Something) Choice() string {
	return o.Name
}

func main() {
	choices := []peco.Choosable{
		NewSomething("iPhone6", "new model iPhone 2014/9", 67800),
		NewSomething("iPhone6 plus", "new big model iPhone 2014/9", 79800),
		NewSomething("iPhone5s", "previous model iPhone 2013/9", 57800),
	}

	result, err := peco.PecolibWithPrompt(choices, "which do you choose the model? >")
	if err != nil {
		fmt.Printf("pecolib error: %v", err)
	}

	result2, err := peco.PecolibWithPrompt(choices, "which do you choose the model if sold out that you previous chosen one? >")
	if err != nil {
		fmt.Printf("pecolib error: %v", err)
	}

	for _, r := range result {
		if t, ok := r.(*Something); ok {
			fmt.Printf("you choose first: %s\n", t.Info())
		} else {
			fmt.Println("invalid struct returns from pecolib")
		}
	}

	for _, r := range result2 {
		if t, ok := r.(*Something); ok {
			fmt.Printf("you choose second: %s\n", t.Info())
		} else {
			fmt.Println("invalid struct returns from pecolib")
		}
	}
}
