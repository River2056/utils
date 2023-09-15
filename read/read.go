package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"regexp"
	"strings"
	"utils/common"
)

func Read() {
	fmt.Println("Enter file name: ")
	var filename string
	fmt.Scanln(&filename)
	fmt.Printf("filename entered: %v\n", filename)

	if _, err := os.Stat(filename); err != nil {
		fmt.Printf("file not exists, %v\n", filename)
		return
	}

	contentBytes, err := ioutil.ReadFile(filename)
	common.CheckError(err)

	reg, _ := regexp.Compile(".*.js")
	content := string(contentBytes)
	contentArr := strings.Split(content, "\n")
	for _, line := range contentArr {
		if !reg.MatchString(line) {
			fmt.Println(line)
		}
	}
}
