package main

import (
	"fmt"
	"strings"
)

func main() {
	cafeList := []string{}
	fmt.Println(cafeList)
	cafeList = insertList("Espresso", cafeList)
	fmt.Println(cafeList)
	findAllItem(cafeList)
}

func findItemName(name string, list []string) int {
	for i, item := range list {
		if strings.ToLower(item) == strings.ToLower(name) {
			return i
		}
	}
	return -1
}

func findAllItem(list []string) {
	for _, item := range list {
		fmt.Println(item + ", ")
	}
}

func insertList(name string, list []string) []string {
	newList := make([]string, len(list)+1)
	newList[len(list)] = name
	for i, item := range list {
		newList[i] = item
	}
	return newList
}

func deleteItem(id int, list []string) []string {
	newList := make([]string, 0, len(list)-1)
	for i, item := range list {
		if i != id {
			newList[i] = item
		}
	}
	return newList
}

func updateItem(id int, name string, list []string) []string {
	newList := make([]string, len(list))
	for i, item := range list {
		if i == id {
			newList[i] = name
		} else {
			newList[i] = item
		}
	}
	return newList
}
