package gendiff

import (
	"encoding/json"
	"fmt"
	"log"
	"os"
)

func validate_path(path string) (validated_path string, err error) {
	return path, nil
}

func main() {
	args := os.Args
	if len(args) < 3 {
		fmt.Println("Not enough arguments")
	} else if len(args) > 3 {
		fmt.Println("Too many arguments")
	}
	firstFilePath := args[1]
	secondFilePath := args[2]

	// TODO: Validate paths

	// pathFirst, err := validate_path(ioutil.ReadFile(args[1]))
	// if err != nil {
	// 	return nil, err
	// }
	// pathSecond, err := validate_path(ioutil.ReadFile(args[2]))
	// if err != nil {
	// 	return nil, err
	// }

	contentFirst, err := readFile(firstFilePath)
	if err != nil {
		log.Fatal("Error when opening file: ", err)
	}
	contentSecond, err := readFile(secondFilePath)
	if err != nil {
		log.Fatal("Error when opening file: ", err)
	}
	firstMap, err := jsonToMap(contentFirst)
	if err != nil {
		log.Fatal("Error during Unmarshal(): ", err)
	}
	secondMap, err := jsonToMap(contentSecond)
	if err != nil {
		log.Fatal("Error during Unmarshal(): ", err)
	}

	// parse
	// replace
	// calculate diff in goroutine
}

func readFile(path string) ([]byte, error) {
	content, err := os.ReadFile(path)
	if err != nil {
		return content, err
	}
	fmt.Printf("Successfully opened %s", path)
	return content, err
}

func jsonToMap(content []byte) (map[string]interface{}, error) {
	var payload map[string]interface{}
	err := json.Unmarshal(content, &payload)
	return payload, err
}
