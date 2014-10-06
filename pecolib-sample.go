package main

import (
	"fmt"

	"github.com/reiki4040/peco"
)

type Something struct {
	buf         string
	Name        string
	Description string
	Price       int
}

func NewSomething(buf, name, description string, price int) *Something {
	return &Something{
		buf:         buf,
		Name:        name,
		Description: description,
		Price:       price,
	}
}

func (o *Something) Okay() string {
	return "okay! " + o.Name
}

func main() {
	thing1 := "iPhone6"
	thing2 := "iPhone6 plus"
	thing3 := "iPhone5s"

	obj_map := map[string]*Something{
		thing1: NewSomething("iPhone6 Buffer", thing1, "new model iPhone 2014/9", 67800),
		thing2: NewSomething("iPhone6 plus Buffer", thing2, "new big model iPhone 2014/9", 79800),
		thing3: NewSomething("iPhone5s Buffer", thing3, "previous model iPhone 2013/9", 57800),
	}

	// TODO what is boolean
	matches := []peco.Match{
		peco.NewNoMatch(thing1, true),
		peco.NewNoMatch(thing2, true),
		peco.NewNoMatch(thing3, true),
	}

	result, err := peco.Pecolib(matches, "which do you choice the model? >")
	if err != nil {
		fmt.Printf("pecolib error: %v", err)
	}

	result2, err := peco.Pecolib(matches, "which do you choice the model if sold out that you previous choose one? >")

	if err != nil {
		fmt.Printf("pecolib error: %v", err)
	}

	for _, r := range result {
		orig := r.Output()
		v, ok := obj_map[orig]
		if ok {
			fmt.Printf("you choose first: %s\n", v.Okay())
		} else {
			fmt.Println("Fail!")
		}
	}

	for _, r := range result2 {
		orig := r.Output()
		v, ok := obj_map[orig]
		if ok {
			fmt.Printf("you choose second: %s\n", v.Okay())
		} else {
			fmt.Println("Fail!")
		}
	}

}
