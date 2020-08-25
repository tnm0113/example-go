package json

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"os"
	"strings"
	"time"
)

// FruitBasket : example struct
type FruitBasket struct {
    Name    string
    Fruit   []string
    ID      int64  `json:"ref"`
    private string // An unexported field is not encoded.
    Created time.Time
}

// EncodeStructToJSON : encode a struct to json data
func EncodeStructToJSON() {
	basket := FruitBasket{
		Name:    "Standard",
		Fruit:   []string{"Apple", "Banana", "Orange"},
		ID:      999,
		private: "Second-rate",
		Created: time.Now(),
	}
	
	var jsonData []byte
	// json Marshal return JSON encoding of struct basket ([]byte)
	jsonData, err := json.Marshal(basket)
	if err != nil {
		log.Println(err)
	}
	fmt.Println(string(jsonData))

	// pretty print 
	jsonDataPretty, err := json.MarshalIndent(basket, "", "    ")
	if err != nil {
		log.Println(err)
	}
	fmt.Println(string(jsonDataPretty))
}

// DecodeJSONToStruct : decode json data ([]byte) to struct
func DecodeJSONToStruct(){
	jsonData := []byte(`
	{
		"Name": "Standard",
		"Fruit": [
			"Apple",
			"Banana",
			"Orange"
		],
		"ref": 999,
		"Created": "2018-04-09T23:00:00Z"
	}`)

	var basket FruitBasket
	err := json.Unmarshal(jsonData, &basket)
	if err != nil {
		log.Println(err)
	}
	fmt.Println(basket.Name, basket.Fruit, basket.ID)
	fmt.Println(basket.Created)
}

// IterateObjectAndArray : unmarshal json data and iterate data
// The encoding/json package uses
//  - map[string]interface{} to store arbitrary JSON objects, and
//  - []interface{} to store arbitrary JSON arrays.
// It will unmarshal any valid JSON data into a plain interface{} value.
// {                                  map[string]interface{}{
//    "Name": "Eve",                      "Name": "Eve",
//    "Age": 6,                           "Age": 6,
//    "Parents": [          ===           "Parents": []interface{}{  
//        "Alice",                         "Alice",
//        "Bob"                            "Bob",
//    ]                                    },
// }                                   }
func IterateObjectAndArray() {
	jsonData := []byte(`{"Name":"Eve","Age":6,"Parents":["Alice","Bob"]}`)

	var v interface{}
	json.Unmarshal(jsonData, &v)
	data := v.(map[string]interface{})

	for k, v := range data {
		switch v := v.(type) {
		case string:
			fmt.Println(k, v, "(string)")
		case float64:
			fmt.Println(k, v, "(float64)")
		case []interface{}:
			fmt.Println(k, "(array):")
			for i, u := range v {
				fmt.Println("    ", i, u)
			}
		default:
			fmt.Println(k, v, "(unknown)")
		}
	}
}

// FileReadWrite : read stream of JSON object, modify and write objects
func FileReadWrite() {
	const jsonData = `
		{"Name": "Alice", "Age": 25}
		{"Name": "Bob", "Age": 22}
	`
	reader := strings.NewReader(jsonData)
	writer := os.Stdout

	dec := json.NewDecoder(reader)
	enc := json.NewEncoder(writer)

	for {
		// Read one JSON object and store it in a map.
		var m map[string]interface{}
		if err := dec.Decode(&m); err == io.EOF {
			break
		} else if err != nil {
			log.Fatal(err)
		}

		// Remove all key-value pairs with key == "Age" from the map.
		for k := range m {
			if k == "Age" {
				delete(m, k)
			}
		}

		// Write the map as a JSON object.
		if err := enc.Encode(&m); err != nil {
			log.Println(err)
		}
	}
}