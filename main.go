package main

import (
	"fmt"
	readfile "json_spreader/internal/readFile"
	readjson "json_spreader/internal/readJson"
	writecsv "json_spreader/internal/writeCSV"
	"log"
	"os"
	"path/filepath"
	"strings"
)

func processAll() {
	files, err := os.ReadDir(".")
	if err != nil {
		log.Fatal(err)
	}

	for _, file := range files {
		if strings.ToLower(filepath.Ext(file.Name())) == ".json" {

			processOne(file.Name(), "")
		}
	}
}

func processOne(pth, des string) {

	data, err := readfile.ReadOneFile(pth)
	if err != nil {
		fmt.Println("Problem trying to read file")
	}
	//fmt.Println("File read:\n", data)
	matrix, err := readjson.Process(data)
	if err != nil {
		fmt.Println("Problem trying to read file")
	}
	if des == "" {
		des, _, _ = strings.Cut(pth, ".")
	}
	msg, err := writecsv.WriteCsv(matrix, des+".csv")

	if err != nil {
		log.Fatal(err)

	}
	fmt.Println(msg)
}

func listPront() {
	files, err := os.ReadDir(".")
	if err != nil {
		log.Fatal(err)
	}
	list := []string{}
	for _, f := range files {
		if f.IsDir() {
			continue
		}
		list = append(list, f.Name())
	}
	fmt.Println(" --- Chose file to proces via index number --- ")

	for i, file := range list {
		fmt.Println(i, ": ", file)
	}

	fmt.Print("->> ")
	var input int
	_, err = fmt.Scan(&input)
	if err != nil {
		log.Fatal(err)
	}
	processOne(list[input], "")

}

func helpPront() {
	println(
		`-help -h open help(this)
-list -l index all files in directory. After you can selet file via index.
-all  -a create .cvs from each .json in directory
1 atribute, atribute is source, destination is curern directory
2 atributes, first is source, second is destination
	
json-spreader is custom json to csv tool (creates copy).
It only works with json shape of List of SingularKey:List of [structs of key:value pairs]
Keys and values are mapped to csv that opes as spreadsheet, keys are atributes and all values are recors`)

	var input string
	_, err := fmt.Scan(&input)
	if err != nil {
		log.Fatal(err)
	}
	pront(strings.Fields(input))
}

func pront(input []string) {
	fmt.Printf("-- input <%v>\n", input)

	switch len(input) {
	case 0:
		helpPront()
	case 1:
		switch input[0][:2] {
		case "-h":
			helpPront()

		case "-l":
			listPront()

		case "-a":
			processAll()

		default:
			processOne(input[1], "")
		}

	case 2:
		processOne(input[0], input[1])
	default:
		fmt.Println(" --- Too many arguments --- ")
		helpPront()
	}
}

func main() {
	runArguments := os.Args
	pront(runArguments[1:])

}
