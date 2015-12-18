package main

import (
	"io/ioutil"
	"os"
	"regexp"
)

var (
	glRegexp    = regexp.MustCompile(`\sgl[A-Z][^\t\n\f\r (.,;/]+`)
	fileFuncMap = make(map[string]int)
	glFuncMap   = make(map[string]int)
)

func doDir(path string) {
	println("---------------")
	println(path)
	println("---------------")
	files, err := ioutil.ReadDir(path)
	if err != nil {
		println(err.Error())
		return
	}

	for _, f := range files {
		if f.IsDir() {
			doDir(path + "/" + f.Name())
		} else {
			println(f.Name())
			data, err := ioutil.ReadFile(path + "/" + f.Name())
			if err != nil {
				println(err.Error())
			} else {
				result := glRegexp.FindAllString(string(data), -1)
				for _, f := range result {
					val, has := fileFuncMap[f]
					if has == false {
						fileFuncMap[f] = 1
					} else {
						fileFuncMap[f] = val + 1
					}
				}
			}
		}
	}
}

func main() {
	dirs := os.Args[1:]
	for _, dir := range dirs {
		doDir(dir)
	}

	println("---------------")
	println("RESULT")
	println("---------------")
	for key, _ := range fileFuncMap {
		// println(key, " [", count, "]")
		println(key)
	}

	println("---------------")
	println("UNKNOWN")
	println("---------------")
	data, err := ioutil.ReadFile("gl3.h")
	if err == nil {
		result := glRegexp.FindAllString(string(data), -1)
		for _, f := range result {
			glFuncMap[f] = 1
		}

		for key, _ := range fileFuncMap {
			_, has := glFuncMap[key]
			if has == false {
				println(key)
			}
		}

	}
}
