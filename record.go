package main

import "fmt"

var UsedRecords = make(map[Record]struct{})
var NumberRepeatedRecords = make(map[Record]int)
var NumberOfAllRecordWithoutRepeats int

type Record struct {
	CityName string `xml:"city,attr"`
	Street   string `xml:"street,attr"`
	House    string `xml:"house,attr"`
	Floor    int    `xml:"floor,attr"`
}

func (r Record) getAllInfo() string {
	return fmt.Sprintf("%s, %s, дом %s, %d-этажный", r.CityName, r.Street, r.House, r.Floor)
}
