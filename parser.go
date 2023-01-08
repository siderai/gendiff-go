package main

import (
	"fmt"
	"reflect"
)

type NodeCompared struct {
	value  string
	status NodeComparisonStatus
}

type NodeComparisonStatus string

const (
	Equal   NodeComparisonStatus = "Equal"
	Deleted                      = "Deleted"
	Added                        = "Added"
	Nested                       = "Nested"
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
	makeDiff(iMap, anotherMap)
}

func makeDiff(firstMap, secondMap map[string]interface{}) map[NodeCompared]interface{} {
	// extract keys
	keys_first := make([]string, 0, len(firstMap))
	for k := range firstMap {
		keys_first = append(keys_first, k)
	}
	keys_second := make([]string, 0, len(secondMap))
	for k := range secondMap {
		keys_second = append(keys_second, k)
	}
	fmt.Println("First only", diffKeys(keys_first, keys_second))
	fmt.Println("Second only", diffKeys(keys_second, keys_first))
	fmt.Println("Common", findKeysIntersection(keys_second, keys_first))

	// group keys

	// firstOnlyKeys := diffKeys(keys_first, keys_second)
	// secondOnlyKeys := diffKeys(keys_second, keys_first)
	commonKeys := findKeysIntersection(keys_second, keys_first)

	var diff map[NodeCompared]interface{}
	fmt.Println("diff init", diff)

	for _, key := range commonKeys {

		valueInFirst := firstMap[key]
		valueInSecond := secondMap[key]

		var res NodeCompared
		if valueInFirst == valueInSecond {
			res = NodeCompared{status: Equal, value: key}
			diff[res] = valueInFirst
		} else if isMap(valueInFirst) && isMap(valueInSecond) {
			res = NodeCompared{status: Nested, value: makeDiff(valueInFirst, valueInSecond)}
		} else {
			diff[NodeCompared{status: Deleted, value: key}] = valueInFirst
			diff[NodeCompared{status: Added, value: key}] = valueInSecond
		}
	}

	fmt.Println("diff internal representation", diff)
	return diff
}

func isMap(node interface{}) bool {
	return reflect.ValueOf(node).Kind() == reflect.Map
}

func diffKeys(a, b []string) []string {
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

func findKeysIntersection(a, b []string) []string {
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

// store results in struct. Create

// TODO method to translate any value into string
func (kcr *NodeComparisonResult) toString() {
}
