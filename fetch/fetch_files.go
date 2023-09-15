package fetch

import (
	"fmt"
	"io/ioutil"
	"os"
	"strings"
	"utils/common"
)

func Fetch() {
	args := os.Args[1:]
	filename := args[0]
	if _, err := os.Stat(filename); err != nil {
		panic("file not exists!")
	}

	bytes, err := ioutil.ReadFile(filename)
    common.CheckError(err)

	content := strings.Split(string(bytes), "\n")
	base := "/root/excalibur/"
	for _, line := range content {
		line = strings.TrimSpace(line)
		if len(line) > 0 {
			filepath := fmt.Sprintf("%v%v", base, line)
			fmt.Printf("filepath: %v\n", filepath)
			fileBytes, err := ioutil.ReadFile(filepath)
			common.CheckError(err)

			newFileName := fmt.Sprintf("/root/tmp/files/%v", line)
			fmt.Printf("newFileName: %v\n", newFileName)

			var lastIndex int
			for i := len(newFileName) - 1; i >= 0; i-- {
				if newFileName[i] == '/' {
					lastIndex = i
					break
				}
			}

			dir := newFileName[:lastIndex]

			if _, ex := os.Stat(dir); os.IsNotExist(ex) {
				err = os.MkdirAll(dir, os.ModePerm)
				common.CheckError(err)
			}

			err = ioutil.WriteFile(newFileName, fileBytes, 0755)
			common.CheckError(err)
		}
	}
}
