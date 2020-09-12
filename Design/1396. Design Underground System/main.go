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
	undergroundSystem := Constructor1()
	undergroundSystem.CheckIn1(10, "Leyton", 3)
	undergroundSystem.CheckOut1(10, "Paradise", 8)
	fmt.Println("Output: ", undergroundSystem.GetAverageTime1("Leyton", "Paradise")) // return 5.00000
	undergroundSystem.CheckIn1(5, "Leyton", 10)
	undergroundSystem.CheckOut1(5, "Paradise", 16)
	fmt.Println("Output: ", undergroundSystem.GetAverageTime1("Leyton", "Paradise")) // return 5.50000
	undergroundSystem.CheckIn1(2, "Leyton", 21)
	undergroundSystem.CheckOut1(2, "Paradise", 30)
	fmt.Println("Output: ", undergroundSystem.GetAverageTime1("Leyton", "Paradise")) // return 6.66667
}

// Station ...
type Station struct {
	stationName string
	time        int
}

// UndergroundSystem1 ...
type UndergroundSystem1 struct {
	in  map[int]*Station
	out map[string][]int
}

// Constructor1 ...
func Constructor1() UndergroundSystem1 {
	us := new(UndergroundSystem1)
	us.in = map[int]*Station{}
	us.out = map[string][]int{}
	return *us
}

// CheckIn1 A customer with id card equal to id, gets in the station stationName at time t.
func (us *UndergroundSystem1) CheckIn1(id int, stationName string, t int) {
	s := &Station{stationName, t}
	us.in[id] = s
}

// CheckOut1 A customer with id card equal to id, gets out from the station stationName at time t.
func (us *UndergroundSystem1) CheckOut1(id int, stationName string, t int) {
	s := us.in[id]
	name := s.stationName + stationName
	time := t - s.time
	out := us.out[name]
	out = append(out, time)
	us.out[name] = out
}

// GetAverageTime1 Returns the average time to travel between the startStation and the endStation.
func (us *UndergroundSystem1) GetAverageTime1(startStation string, endStation string) float64 {
	name := startStation + endStation
	out := us.out[name]
	sum := 0
	for _, t := range out {
		sum = sum + t
	}
	return float64(sum) / float64(len(out))
}
