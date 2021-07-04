package main

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestFile_GetContentInBytes(t *testing.T) {
	f, rm := createFileForTest()
	defer rm(f)

	testTable := []struct {
		name            string
		expectedContent []byte
		wantGetErr      bool
	}{
		{
			name: "OK",
			expectedContent: []byte(`
					<?xml version="1.0" encoding="utf-8"?>
					<root>
					<item city="Тест" street="Улица Тестовая" house="666" floor="5" />
					<item city="Тест" street="Улица Тестовая" house="555" floor="4" />
					<item city="Тест" street="Улица Тестовая" house="444" floor="3" />
					</root>`),
			wantGetErr: false,
		},
		{
			name:            "Empty content file",
			expectedContent: []byte{},
			wantGetErr:      false,
		},
	}

	for _, tc := range testTable {
		t.Run(tc.name, func(t *testing.T) {
			f.Truncate(0)

			file := NewFile(f.Name())
			f.Write(tc.expectedContent)

			content, err := file.GetContentInBytes()
			if tc.wantGetErr {
				assert.Error(t, err)
			} else {
				assert.NoError(t, err)
				assert.Equal(t, tc.expectedContent, content)
			}
		})
	}
}

func TestFile_Unmarshal(t *testing.T) {
	testTable := []struct {
		name            string
		contentFile     string
		expectedRecords []Record
		wantGetErr      bool
	}{
		{
			name: "OK",
			contentFile: `<?xml version="1.0" encoding="utf-8"?>
	 				 	  	<root>
							<item city="Тест" street="Улица Тестовая" house="666" floor="5" />
							<item city="Тест" street="Улица Тестовая" house="555" floor="4" />
							<item city="Тест" street="Улица Тестовая" house="444" floor="3" />
							<item city="Тест" street="Улица Тестовая" house="333" floor="2" />
							<item city="Тест" street="Улица Тестовая" house="222" floor="1" />
							</root>`,
			expectedRecords: []Record{
				{CityName: "Тест", Street: "Улица Тестовая", House: "666", Floor: 5},
				{CityName: "Тест", Street: "Улица Тестовая", House: "555", Floor: 4},
				{CityName: "Тест", Street: "Улица Тестовая", House: "444", Floor: 3},
				{CityName: "Тест", Street: "Улица Тестовая", House: "333", Floor: 2},
				{CityName: "Тест", Street: "Улица Тестовая", House: "222", Floor: 1},
			},
			wantGetErr: false,
		},
		{
			name: "Invalid content",
			contentFile: `<?xml version="1.0" encoding="utf-8"?>
							<root>
							<item ="Тест" ="Улица Тестовая" ="666" ="5" />
							<item ="Тест" ="Улица Тестовая" ="555" ="4" />
							<item ="Тест" ="Улица Тестовая" ="444" ="3" />
							<item ="Тест" ="Улица Тестовая" ="333" ="2" />
							<item ="Тест" ="Улица Тестовая" ="222" ="1" />
							</root>`,
			wantGetErr: true,
		},
		{
			name:       "Empty content file",
			wantGetErr: true,
		},
	}

	for _, tc := range testTable {
		t.Run(tc.name, func(t *testing.T) {
			var file File

			err := file.Unmarshal([]byte(tc.contentFile))
			if tc.wantGetErr {
				assert.Error(t, err)
			} else {
				assert.NoError(t, err)
				assert.Equal(t, tc.expectedRecords, file.Records)
			}
		})
	}
}

func TestFile_ParseRecords(t *testing.T) {
	testTable := []struct {
		name                                           string
		records                                        []Record
		expectedAllOfNumberHouseInCity                 map[string]*City
		expectedRepeatedRecordsWithNumberOfRepetitions map[Record]int
		expectedNumberOfRecordsWithoutRepeated         int
	}{
		{
			name: "OK",
			records: []Record{
				{CityName: "Тест", Street: "Улица Тестовая", House: "666", Floor: 5},
				{CityName: "Тест", Street: "Улица Тестовая", House: "666", Floor: 5},
				{CityName: "Тест2", Street: "Улица Тестовая2", House: "56", Floor: 4},
				{CityName: "Тест2", Street: "Улица Тестовая2", House: "56", Floor: 4},
				{CityName: "Тест3", Street: "Улица Тестовая3", House: "127", Floor: 1},
				{CityName: "Тест3", Street: "Улица Тестовая3", House: "127", Floor: 1},
			},
			expectedAllOfNumberHouseInCity: map[string]*City{
				"Тест":  {NumberOfFiveStoryHouse: 1},
				"Тест2": {NumberOfFourStoryHouse: 1},
				"Тест3": {NumberOfOneStoryHouse: 1},
			},
			expectedRepeatedRecordsWithNumberOfRepetitions: map[Record]int{
				{CityName: "Тест", Street: "Улица Тестовая", House: "666", Floor: 5}:   1,
				{CityName: "Тест2", Street: "Улица Тестовая2", House: "56", Floor: 4}:  1,
				{CityName: "Тест3", Street: "Улица Тестовая3", House: "127", Floor: 1}: 1,
			},
			expectedNumberOfRecordsWithoutRepeated: 3,
		},
		{
			name:                           "Empty records",
			records:                        []Record{},
			expectedAllOfNumberHouseInCity: map[string]*City{},
			expectedRepeatedRecordsWithNumberOfRepetitions: map[Record]int{},
			expectedNumberOfRecordsWithoutRepeated:         0,
		},
	}

	for _, tc := range testTable {
		t.Run(tc.name, func(t *testing.T) {
			resetCities()
			resetNumberRepeatedRecords()
			resetNumberOfAllRecordWithoutRepeats()

			var file File
			file.Records = tc.records

			file.ParseRecords()

			assert.Equal(t, tc.expectedAllOfNumberHouseInCity, Cities)
			assert.Equal(t, tc.expectedRepeatedRecordsWithNumberOfRepetitions, NumberRepeatedRecords)
			assert.Equal(t, tc.expectedNumberOfRecordsWithoutRepeated, NumberOfAllRecordWithoutRepeats)
		})
	}
}

func BenchmarkFile_GetContentInBytes(b *testing.B) {
	file := NewFile(benchmarkFile)

	b.ResetTimer()

	if _, err := file.GetContentInBytes(); err != nil {
		b.Fatal(err)
	}
}

func BenchmarkFile_Unmarshal(b *testing.B) {
	file := NewFile(benchmarkFile)

	content, err := file.GetContentInBytes()
	if err != nil {
		b.Fatal(err)
	}

	b.ResetTimer()

	if err = file.Unmarshal(content); err != nil {
		b.Fatal(err)
	}
}

func BenchmarkFile_ParseRecords(b *testing.B) {
	file := NewFile(benchmarkFile)

	content, err := file.GetContentInBytes()
	if err != nil {
		b.Fatal(err)
	}

	if err = file.Unmarshal(content); err != nil {
		b.Fatal(err)
	}

	b.ResetTimer()

	file.ParseRecords()
}
