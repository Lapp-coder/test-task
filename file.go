package main

import (
	"encoding/xml"
	"io/ioutil"
)

type File struct {
	Filename string
	Records  []Record `xml:"item"`
}

func NewFile(filename string) *File {
	return &File{
		Filename: filename,
	}
}

func (f File) GetContentInBytes() ([]byte, error) {
	return ioutil.ReadFile(f.Filename)
}

func (f *File) Unmarshal(content []byte) error {
	return xml.Unmarshal(content, f)
}

func (f *File) ParseRecords() {
	for _, record := range f.Records {
		if isRepeatRecord(record) {
			incrementNumberOfRepeatedRecord(record)
			continue
		}

		if isCityNotRecorded(record.CityName) {
			city := &City{}
			Cities[record.CityName] = city
			city.incrementNumberOfHousesByFloor(record.Floor)
		} else {
			city := Cities[record.CityName]
			city.incrementNumberOfHousesByFloor(record.Floor)
		}

		NumberOfAllRecordWithoutRepeats++
		UsedRecords[record] = struct{}{}
	}
}

func isRepeatRecord(record Record) bool {
	if _, isUsed := UsedRecords[record]; isUsed {
		return true
	}
	return false
}

func isCityNotRecorded(cityName string) bool {
	return Cities[cityName] == nil
}

func incrementNumberOfRepeatedRecord(record Record) {
	NumberRepeatedRecords[record]++
}
