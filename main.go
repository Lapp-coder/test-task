package main

import (
	"flag"
	"fmt"
	"log"
	"os"
)

const defaultFilename = "address.xml"

var filename string

func init() {
	flag.StringVar(&filename, "f", defaultFilename, "имя файла")
}

func main() {
	flag.Parse()

	file := NewFile(filename)

	content, err := file.GetContentInBytes()
	if err != nil {
		log.Fatalf("не удалось получить содержимое файла: %s", err.Error())
	}

	err = file.Unmarshal(content)
	if err != nil {
		log.Fatalf("не удалось распаковать содержимое из файла: %s", err.Error())
	}

	file.ParseRecords()

	var actionNumber int
	for {
		fmt.Print(
			"\t\t\tВыберите действие\n" +
				"1 - Вывести количество всех 1,2,3,4 и 5 этажных домов из всех городов.\n" +
				"2 - Вывести дублирующиеся записи с количеством повторений.\n" +
				"3 - Вывести общее количество записей без учёта дублей.\n" +
				"4 - Выход из программы\n>")
		fmt.Scan(&actionNumber)
		handleAction(actionNumber)
	}
}

func handleAction(number int) {
	switch number {
	case 1:
		PrintAllOfNumberHousesInCity()
	case 2:
		PrintRepeatedRecordsWithNumberOfRepetitions()
	case 3:
		PrintNumberOfRecordsWithoutRepeated()
	case 4:
		os.Exit(0)
	default:
		fmt.Println("Такого действия нету в списке.")
	}
}
