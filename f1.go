package main

import (
	"fmt"
	"os"
)

func main() {
	os.Mkdir("danqi", 0777)
	os.MkdirAll("danqi/test1/test2", 0777)
	err := os.Remove("danqi")
	if err != nil {
		fmt.Println(err)
	}
	os.RemoveAll("danqi")

}
