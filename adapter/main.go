package adapter

import "fmt"

type Banner struct {
	str string
}

func (banner *Banner) ShowWithParen() {
	fmt.Print("なんだこのやろー" + banner.str)
}

func (banner *Banner) ShowWithAster() {
	fmt.Print("タピオカパン" + banner.str)
}

func main() {

}
