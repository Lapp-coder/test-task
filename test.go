package main

import "os"

const (
	testFile      = "address_test.xml"
	benchmarkFile = "address.xml"
)

func resetCities() {
	for cityName := range Cities {
		delete(Cities, cityName)
	}

	for record := range NumberRepeatedRecords {
		delete(NumberRepeatedRecords, record)
	}
}

func resetNumberRepeatedRecords() {
	for record := range NumberRepeatedRecords {
		delete(NumberRepeatedRecords, record)
	}
}

func resetNumberOfAllRecordWithoutRepeats() {
	NumberOfAllRecordWithoutRepeats = 0
}

func createFileForTest() (*os.File, func(f *os.File)) {
	f, _ := os.Create(testFile)
	return f, func(f *os.File) {
		f.Close()
		os.Remove(f.Name())
	}
}
