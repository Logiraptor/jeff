package main

import (
	"encoding/json"
	"fmt"
	"os"
)

func main() {
	var i interface{}
	err := json.NewDecoder(os.Stdin).Decode(&i)
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	buf, err := json.MarshalIndent(i, "", "\t")
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	os.Stdout.Write(buf)
	fmt.Fprintln(os.Stdout)
}
