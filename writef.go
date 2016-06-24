package main

import (
	"fmt"
	"os"
)

func main() {
	userFile := "danqi.txt"
	fout, err := os.Create(userFile)
	if err != nil {
		fmt.Println(userFile, err)
		return
	}
	defer fout.Close()
	for i := 0; i < 10; i++ {
		fout.WriteString("Danqi is the best!\r\n")
		fout.Write([]byte("Just a test!\r\n"))
	}

}
