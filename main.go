package main

import (
	"encoding/json"
	"fmt"
	"os"
)

func main() {
	var i interface{}
	d := json.NewDecoder(os.Stdin)
	d.UseNumber()
	err := d.Decode(&i)
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
