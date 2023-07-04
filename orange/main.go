package main

import (
	"fmt"
	"orange/calculation"
)

func main() {
	fmt.Println("ini func main")
	hasil := calculation.Multiple(4, 5)
	fmt.Println(hasil)
	perkalian := calculation.Multiple(2, 2)
	fmt.Println("ini hasil perkalian", perkalian)
	fmt.Println(calculation.CheckNilai(75))

	type orang string
	var char orang = "23"
	char = "hai"
	fmt.Println(char)

	indexName := "khairul aswad"[1]

	// chechIndex := indexName[3]

	const value string = "halo dek"
	// loadIndex := string(indexName)

	var (
		lastname = "halo dek"
		past     = 89898
		fori     = 8.90
	)

	const (
		halo    = "78"
		biarlah = 78
		naomi   = 7.8
	)
	fmt.Println(halo, biarlah, naomi)
	fmt.Println(lastname, past, fori)
	var lint float32 = 10.76
	convert := float64(lint)
	fmt.Println(convert)
	fmt.Println(indexName)

	fmt.Println(value)
	var full = 6.7
	var name = "string"

	name = "maliobro"

	fmt.Println(name)

	lgiut := 9.98

	fmt.Println(full)

	fmt.Println(lgiut)

}
