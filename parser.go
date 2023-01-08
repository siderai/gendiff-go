package main

import (
	"fmt"
)

func main() {
	iMap := make(map[string]int)
	iMap["k1"] = 12
	iMap["k2"] = 13
	iMap["k3"] = 14
	iMap["k5"] = 15

	fmt.Println("iMap:", iMap)
	anotherMap := map[string]int{
		"k1": 12,
		"k2": 13,
		"k4": 14,
		"k5": 15,
	}
	for key, value := range iMap {
		fmt.Println(key, value)
	}
	for key, value := range anotherMap {
		fmt.Println(key, value)
	}

	keys_first := make([]string, 0, len(iMap))
	for k := range iMap {
		keys_first = append(keys_first, k)
	}
	keys_second := make([]string, 0, len(anotherMap))
	for k := range anotherMap {
		keys_second = append(keys_second, k)
	}
	fmt.Println("First only", makeDiff(keys_first, keys_second))
	fmt.Println("Second only", makeDiff(keys_second, keys_first))
	fmt.Println("Common", makeIntersection(keys_second, keys_first))
}

func makeDiff(a, b []string) []string {
	mb := make(map[string]struct{}, len(b))
	for _, x := range b {
		mb[x] = struct{}{}
	}
	var diff []string
	for _, x := range a {
		if _, found := mb[x]; !found {
			diff = append(diff, x)
		}
	}
	return diff
}

func makeIntersection(a, b []string) []string {
	mb := make(map[string]struct{}, len(b))
	for _, x := range b {
		mb[x] = struct{}{}
	}
	var diff []string
	for _, x := range a {
		if _, found := mb[x]; found {
			diff = append(diff, x)
		}
	}
	return diff
}
