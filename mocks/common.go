package mocks

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
)

func getMoackData(out interface{}, path string) {
	data, err := ioutil.ReadFile(path)
	if err != nil {
		fmt.Println(err)
	}

	err = json.Unmarshal(data, &out)

	if err != nil {
		fmt.Println("error:", err)
	}
}
