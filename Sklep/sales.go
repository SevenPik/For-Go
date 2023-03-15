package main

import (
	"fmt"
	"math/rand"
	"time"
)

func floatnaint(val float32) float32 {
	valInt := int64(val * 100)
	val = float32(valInt) / 100
	return val
}

func Menu(produkt string) (float32, string) {
	var wybórProduktu string
	var cenaProduktów float32

	switch produkt {
	case "Bułka":
		wybórProduktu = "Bułka"
		cenaProduktów = cenaProduktów + 3.99
	default:
		wybórProduktu = "Nic"
		break
	}
	fmt.Println(cenaProduktów, wybórProduktu)
	return cenaProduktów, wybórProduktu
}

func promocjewbiedrze(promka int, promocja float64) float64 {
	rand.Seed(time.Now().UnixNano())
	promka = rand.Intn(50)
	fmt.Println(promka, "% Promocja (promka)")
	promocja = 100
	promocja = float64(promka) / 100
	fmt.Println(promocja, "% Promocja (promka)")
	return float64(promocja)
}

func main() {
	var x int
	var y int
	fmt.Println("Witaj w Sklepie!", " Oto nasze produkty: Bułka - 3.99 ")
	fmt.Println("Ile chcesz kupić rzeczy?")
	fmt.Scanln(&y)
	var łącznykoszt float32
	for x = 0; x < y; x++ {
		fmt.Println("Co chcesz kupić? ")
		var zachcianka string
		fmt.Scanln(&zachcianka)
		if zachcianka == "nic" {
			break
		}
		var menucena float32
		var menunazwa string
		bydziałało := promocjewbiedrze(1, 1)

		menucena, menunazwa = Menu(zachcianka)
		menucena = menucena - menucena*float32(bydziałało)
		menucena = float32(floatnaint(menucena))
		fmt.Println(menucena, menunazwa)
		łącznykoszt = łącznykoszt + menucena
		fmt.Println(łącznykoszt, "Łączny koszt")
	}
	fmt.Println(łącznykoszt)
}
