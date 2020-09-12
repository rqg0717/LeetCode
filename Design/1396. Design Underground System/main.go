package main

import "fmt"

// UndergroundSystem ...
type UndergroundSystem struct {
	station map[string]map[int]int
}

// Constructor ...
func Constructor() UndergroundSystem {
	obj := make(map[string]map[int]int)
	return UndergroundSystem{station: obj}
}

// CheckIn A customer with id card equal to id, gets in the station stationName at time t.
func (us *UndergroundSystem) CheckIn(id int, stationName string, t int) {
	if us.station[stationName] == nil {
		us.station[stationName] = make(map[int]int)
	}
	us.station[stationName][id] = t
}

// CheckOut A customer with id card equal to id, gets out from the station stationName at time t.
func (us *UndergroundSystem) CheckOut(id int, stationName string, t int) {
	if us.station[stationName] == nil {
		us.station[stationName] = make(map[int]int)
	}
	us.station[stationName][id] = t
}

// GetAverageTime Returns the average time to travel between the startStation and the endStation.
func (us *UndergroundSystem) GetAverageTime(startStation string, endStation string) float64 {
	count, sum := 0, 0
	start := us.station[startStation]
	end := us.station[endStation]
	for id, t := range start {
		if end[id] != 0 {
			sum = sum - t + end[id]
			count++
		}
	}
	return float64(sum) / float64(count)
}

func main() {
	undergroundSystem := Constructor()
	undergroundSystem.CheckIn(10, "Leyton", 3)
	undergroundSystem.CheckOut(10, "Paradise", 8)
	fmt.Println("Output: ", undergroundSystem.GetAverageTime("Leyton", "Paradise")) // return 5.00000
	undergroundSystem.CheckIn(5, "Leyton", 10)
	undergroundSystem.CheckOut(5, "Paradise", 16)
	fmt.Println("Output: ", undergroundSystem.GetAverageTime("Leyton", "Paradise")) // return 5.50000
	undergroundSystem.CheckIn(2, "Leyton", 21)
	undergroundSystem.CheckOut(2, "Paradise", 30)
	fmt.Println("Output: ", undergroundSystem.GetAverageTime("Leyton", "Paradise")) // return 6.66667
}
