package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
)

func Check(e error) {
	if e != nil {
		panic(e)
	}
}

var combine map[string]interface{}

func mergedata(map1, map2 map[string]interface{}) map[string]interface{} {
	if len(map1) == len(map2) {
		return merge(map1, map2)
	}
}

func merge(greaterdata, lessdata map[string]interface{}) map[string]interface{} {
	for k, v := range greaterdata{

	}
}

func main() {

	dat1, err := ioutil.ReadFile("json/grpc/test-data/response1.json")
	Check(err)
	dat2, err := ioutil.ReadFile("json/grpc/test-data/response2.json")
	Check(err)
	fmt.Printf("%s %s ", dat1, dat2)
	var x map[string]interface{}
	var y map[string]interface{}
	json.Unmarshal(dat1, &x)
	json.Unmarshal(dat2, &y)

	fmt.Println(x["user"].(map[string]interface{})["email"])
	for k, v := range x{
		fmt.Printf("key: %v, value: %v\n", k, v)
	}
	fmt.Println(len(x))
	cm := mergedata(x, y)
	fmt.Println(cm)

}
