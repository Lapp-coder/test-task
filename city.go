package main

var Cities = make(map[string]*City)

type City struct {
	NumberOfOneStoryHouse   int
	NumberOfTwoStoryHouse   int
	NumberOfThreeStoryHouse int
	NumberOfFourStoryHouse  int
	NumberOfFiveStoryHouse  int
}

func (c *City) incrementNumberOfHousesByFloor(floor int) {
	switch floor {
	case 1:
		c.NumberOfOneStoryHouse++
	case 2:
		c.NumberOfTwoStoryHouse++
	case 3:
		c.NumberOfThreeStoryHouse++
	case 4:
		c.NumberOfFourStoryHouse++
	case 5:
		c.NumberOfFiveStoryHouse++
	default:
	}
}
