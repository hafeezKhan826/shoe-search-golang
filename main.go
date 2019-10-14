package main

import (
	"fmt"
	"math/rand"
	"time"
)

// Suggestions from golang-nuts
// http://play.golang.org/p/Ctg3_AQisl
var shoes [20]int
var shoeColors = []string{"red", "navy blue", "green"}
var shoeBrands = []string{"Puma", "Reebok", "UCB"}

type shoeItem struct {
	name   string
	itemID string
	size   int
	color  string
	brand  string
}

var allShoes []shoeItem

func getShoeSize(min int, max int) int {
	rand.Seed(time.Now().UnixNano())
	return rand.Intn(max-min) + min
}

func getShoeColor() string {
	rand.Seed(time.Now().UnixNano())
	return shoeColors[rand.Intn(len(shoeColors))]
}

func getShoeBrand() string {
	rand.Seed(time.Now().UnixNano())
	return shoeBrands[rand.Intn(len(shoeBrands))]
}

func createShoe() shoeItem {
	var shoeSize int
	var shoeColor string
	var shoeBrand string

	rand.Seed(time.Now().UnixNano())

	shoeSize = getShoeSize(7, 13)
	shoeColor = getShoeColor()
	shoeBrand = getShoeBrand()

	shoe := shoeItem{color: shoeColor, brand: shoeBrand, size: shoeSize}

	return shoe
}
func checkUnique(shoe shoeItem) bool {
	for i := 0; i < len(allShoes); i++ {
		if allShoes[i].color == shoe.color && allShoes[i].size == shoe.size && allShoes[i].brand == shoe.brand {
			return false
		}
	}
	return true
}

func pushShoes(count int) {
	counter := count
	for i := 0; i < count && counter > 0; i++ {
		time.Sleep(2 * time.Nanosecond)
		newShoe := createShoe()
		isUnique := checkUnique(newShoe)
		if isUnique {
			allShoes = append(allShoes, newShoe)
			counter--
		}
	}
}

func findMyShoe(size int, color string) bool {
	found := false

	for i := 0; i < len(allShoes); i++ {

		// fmt.Println(color == allShoes[i].color, size == allShoes[i].size)
		if size == allShoes[i].size && color == allShoes[i].color {
			// if color == allShoes[i].color && size == allShoes[i].size {
			fmt.Println("------------> Shoe found <------------")
			fmt.Println("Brand :", allShoes[i].brand, "\tColor :", allShoes[i].color, "\tSize :", allShoes[i].size)
			found = true
		}
	}
	return found
}

func main() {

	fmt.Println("Stocking shoes....")
	pushShoes(5)
	toFindSize := 11
	toFindColor := "red"

	fmt.Println("------------> Searching for shoe of size: ", toFindSize, " and color ", toFindColor, " <------------")
	// fmt.Print("asdf: ", found, toFindSize, toFindColor)

	fmt.Printf("Congratulations! Your 3 second AfterFunc() timer finished.\n")
	found := findMyShoe(toFindSize, toFindColor)

	if found {
		fmt.Println("------------> Shoe found <------------")
	}
	for !found {
		fmt.Println("------------> Shoe not found <------------")
		fmt.Println("------------> Restocking <------------")
		pushShoes(5)
		found := findMyShoe(toFindSize, toFindColor)
		if found {
			fmt.Println("------------> Shoe search completed <------------")
			return
		}
	}

	fmt.Println("------------> Shoe search completed <------------")

}
