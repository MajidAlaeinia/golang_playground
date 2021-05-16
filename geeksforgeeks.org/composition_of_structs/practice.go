package main

import "fmt"

type detailss struct {
	genre       string
	genreRating string
	reviews     string
}

type gamee struct {
	name  string
	price string
	detailss
}

func (d detailss) showDetails() {
	fmt.Println("Genre:", d.genre)
	fmt.Println("Genre Rating:", d.genreRating)
	fmt.Println("Reviews:", d.reviews)
}

func (g gamee) show() {
	fmt.Println("Name:", g.name)
	fmt.Println("Price:", g.price)
	g.showDetails()
}

func main() {
	gameDetails := detailss{"Action", "+18", "mostly positive"}
	newGame := gamee{"XYZ", "$125", gameDetails}

	newGame.show()
}
