package main

import "fmt"

func PrintAllOfNumberHousesInCity() {
	for cityName, city := range Cities {
		fmt.Printf("%s: одноэтажные: %d, двухэтажные: %d, трехэтажные: %d, четырехэтажные: %d, пятиэтажные: %d\n",
			cityName, city.NumberOfOneStoryHouse, city.NumberOfTwoStoryHouse, city.NumberOfThreeStoryHouse, city.NumberOfFourStoryHouse, city.NumberOfFiveStoryHouse)
	}
}

func PrintRepeatedRecordsWithNumberOfRepetitions() {
	for record, repetitionRate := range NumberRepeatedRecords {
		fmt.Printf("%s | повторений: %d\n", record.getAllInfo(), repetitionRate)
	}
}

func PrintNumberOfRecordsWithoutRepeated() {
	fmt.Println(NumberOfAllRecordWithoutRepeats)
}
