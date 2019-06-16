package main

import (
	"bufio"
	"fmt"
	"github.com/weasel/server"
	"io"
	"os"
)

func main() {
	fmt.Println("weasel main start!")
	server.Run()

	//test()
}

func test() {
	f, err := os.Open("/Users/ycysvler/Documents/aba.txt")

	if err != nil {
		fmt.Println("os Open error: ", err)
		return
	}
	defer f.Close()

	br := bufio.NewReader(f)
	for {
		line, _, err := br.ReadLine()
		if err == io.EOF {
			break
		} else if err != nil {
			fmt.Println("br ReadLine error: ", err)
			return
		}

		fmt.Printf("%s\n", string(line))
		//fmt.Println("line: ", string(line))
	}

}
