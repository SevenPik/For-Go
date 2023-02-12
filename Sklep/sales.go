package main

import (
	"fmt"
	"math/rand"
	"time"
)

func Menu(produkt string) (float32, string) {
	var wybórProduktu string
	var cenaProduktów float32

	switch produkt {
	case "Bułka":
		wybórProduktu = "Bułka"
		cenaProduktów = cenaProduktów + 3.99
	default:
		wybórProduktu = "Nic"
	}
	fmt.Println(cenaProduktów, wybórProduktu)
	return cenaProduktów, wybórProduktu
}
func main() {
	rand.Seed(time.Now().UnixNano())
	promka := rand.Intn(100)
	czaspromki := rand.Intn(100)
	promka = promka * -1
	fmt.Println(promka)
	fmt.Println(czaspromki)

	promocja := 100
	promocja = promka / promocja
	fmt.Println(promka)
	fmt.Println("Co chcesz kupić? ")

	var zachcianka string
	fmt.Println(promocja)
	fmt.Scanln(&zachcianka)
	var menucena float32
	var menunazwa string
	menucena, menunazwa = Menu(zachcianka)
	menucena = menucena - menucena*float32(promocja)

	fmt.Println(menucena, menunazwa)

}
