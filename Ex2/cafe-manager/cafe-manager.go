package main

import (
	"fmt"
)

type Drink struct {
	id       int
	name     string
	price    float64
	category string
}

func main() {
	cafeList := [...]Drink{}
	fmt.Println(cafeList)

}

func findItemById(id int, list []Drink) (int, Drink) {
	for index, item := range list {
		if item.id == id {
			return index, item
		}
	}
	return -1, Drink{}
}

func findAllItem(list []string) {
	for _, item := range list {
		fmt.Println(item + ", ")
	}
}

func insertList(newDrink Drink, list []Drink) []Drink {
	list = append(list, newDrink)
	return list
}

func deleteItem(id int, list []Drink) []Drink {
	index, _ := findItemById(id, list)
	if index != -1 {
		list = append(list[:index], list[index+1:]...)
	}
	return list
}

func updateItem(id int, newDrink Drink, list []Drink) []Drink {
	index, item := findItemById(id, list)
	if item != (Drink{}) {
		item.category = newDrink.category
		item.name = newDrink.name
		item.price = newDrink.price
		list[index] = item
	}
	return list
}
