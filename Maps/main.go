package main

import "fmt"

func main() {
	//iniciar map

	var myMap map[string]string
	myMap = make(map[string]string)

	//nil map
	fmt.Println(myMap)

	countryCapitalMap := make(map[string]string)

	countryCapitalMap["França"] = "Paris"
	countryCapitalMap["Italia"] = "Roma"
	countryCapitalMap["Japão"] = "Tokyo"
	// countryCapitalMap["EUA"] = "Washington"

	for country := range countryCapitalMap {
		println(country, countryCapitalMap[country])
	}

	capital, ok := countryCapitalMap["EUA"]
	if ok {
		fmt.Println(capital)
	} else {
		fmt.Println("não existe")
	}

	//deletar
	delete(countryCapitalMap, "França")
	fmt.Println("frança deletada")
	fmt.Println("novo map")
	for country := range countryCapitalMap {
		println(country, countryCapitalMap[country])
	}
}
