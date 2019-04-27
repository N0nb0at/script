package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"math/rand"
	"time"
)

type Scheme struct {
	Name  string
	Right int64
	Cost  float64
	Count int64
}

var scheme []Scheme

func readFile(fileName string) ([]Scheme, error) {
	bytes, err := ioutil.ReadFile(fileName)
	if err != nil {
		log.Println("ReadFile error: ", err.Error())
		return nil, err
	}

	if err := json.Unmarshal(bytes, &scheme); err != nil {
		log.Println("Unmarshal error: ", err.Error())
		return nil, err
	}

	log.Printf("%+v", scheme)
	log.Println()

	return scheme, nil
}

func main() {

	schemeList, err := readFile("main/scheme.json")

	if err != nil {
		fmt.Println("readFile error: ", err.Error())
	}

	length := len(schemeList)

	random := rand.New(rand.NewSource(time.Now().UnixNano())).Intn(length)
	log.Println(schemeList[random])

	fmt.Println(random)
	fmt.Println(schemeList[random])
}
