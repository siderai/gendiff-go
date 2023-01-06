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
	firstFilePath := os.Args[1]
	secondFilePath := os.Args[2]

	// TODO: Validate paths

	// pathFirst, err := validate_path(ioutil.ReadFile(args[1]))
	// if err != nil {
	// 	return nil, err
	// }
	// pathSecond, err := validate_path(ioutil.ReadFile(args[2]))
	// if err != nil {
	// 	return nil, err
	// }

	contentFirst := readFile(firstFilePath)
	contentSecond := readFile(secondFilePath)

	firstMap := jsonToMap(contentFirst)
	secondMap := jsonToMap(contentSecond)

}

func readFile(path string) []byte {
	content, err := os.ReadFile(path)
	if err != nil {
		log.Fatal("Error when opening file: ", err)
	}
	fmt.Printf("Successfully opened %s", path)
	return content
}

func jsonToMap(content []byte) map[string]interface{} {
	var payload map[string]interface{}
	err := json.Unmarshal(content, &payload)
	if err != nil {
		log.Fatal("Error during Unmarshal(): ", err)
	}
	return payload
}
