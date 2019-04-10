package main

import "fmt"

func main(){
	var countryCapitalMap map[string]string
	countryCapitalMap = make(map[string]string)

	countryCapitalMap ["France"] = "Paris"
	countryCapitalMap ["Italy"] = "罗马"
	countryCapitalMap ["Japan"] = "东京"
	countryCapitalMap ["India"] = "新德里"

	for country := range countryCapitalMap	{
		fmt.Println(country, "首都是", countryCapitalMap [country])
	}

	capital, ok := countryCapitalMap [ "美国" ]
	fmt.Println(capital)
	fmt.Println(ok)
	if (ok)	{
		fmt.Println("美国的首都是", capital)
	} else {
		fmt.Println("美国的首都不存在")
	}
}