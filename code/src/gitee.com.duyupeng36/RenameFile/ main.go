package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"strings"
)

func main() {
	fd, _ := ioutil.ReadDir("./")
	for _, f := range fd {
		fileName := f.Name()
		nameSlice := strings.Split(fileName, ".")
		name := nameSlice[0]
		suffix := nameSlice[len(nameSlice)-1]
		nonPrefixName := strings.TrimLeft(name, " 0123456789.-")
		Prefix := strings.TrimRight(name, nonPrefixName+".-")
		if len(Prefix) == 1 {
			Prefix = "0"+Prefix
		}
		newFileName := Prefix + " "+nonPrefixName + "." +suffix
		fmt.Println(newFileName)
		os.Rename(fileName, newFileName)
	}
}

