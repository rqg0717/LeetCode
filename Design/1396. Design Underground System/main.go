package main

import "fmt"

// UndergroundSystem ...
type UndergroundSystem struct {
	station map[string]map[int]int
}

// Constructor ...
func Constructor() UndergroundSystem {
	us := make(map[string]map[int]int)
	return UndergroundSystem{station: us}
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
	for id, time := range start {
		if end[id] > 0 {
			sum += end[id] - time
			count++
		}
	}
	return float64(sum) / float64(count)
}

func main() {
	undergroundSystem := Constructor()
	undergroundSystem.CheckIn(45, "Leyton", 3)
	undergroundSystem.CheckIn(32, "Paradise", 8)
	undergroundSystem.CheckIn(27, "Leyton", 10)
	undergroundSystem.CheckOut(45, "Waterloo", 15)
	undergroundSystem.CheckOut(27, "Waterloo", 20)
	undergroundSystem.CheckOut(32, "Cambridge", 22)
	fmt.Println("Output: ", undergroundSystem.GetAverageTime("Paradise", "Cambridge")) // return 14.00000. There was only one travel from "Paradise" (at time 8) to "Cambridge" (at time 22)
	fmt.Println("Output: ", undergroundSystem.GetAverageTime("Leyton", "Waterloo"))    // return 11.00000. There were two travels from "Leyton" to "Waterloo", a customer with id=45 from time=3 to time=15 and a customer with id=27 from time=10 to time=20. So the average time is ( (15-3) + (20-10) ) / 2 = 11.00000
	undergroundSystem.CheckIn(10, "Leyton", 24)
	fmt.Println("Output: ", undergroundSystem.GetAverageTime("Leyton", "Waterloo")) // return 11.00000
	undergroundSystem.CheckOut(10, "Waterloo", 38)
	fmt.Println("Output: ", undergroundSystem.GetAverageTime("Leyton", "Waterloo")) // return 12.00000

	undergroundSystem1 := Constructor1()
	undergroundSystem1.CheckIn1(10, "Leyton", 3)
	undergroundSystem1.CheckOut1(10, "Paradise", 8)
	fmt.Println("Output: ", undergroundSystem1.GetAverageTime1("Leyton", "Paradise")) // return 5.00000
	undergroundSystem1.CheckIn1(5, "Leyton", 10)
	undergroundSystem1.CheckOut1(5, "Paradise", 16)
	fmt.Println("Output: ", undergroundSystem1.GetAverageTime1("Leyton", "Paradise")) // return 5.50000
	undergroundSystem1.CheckIn1(2, "Leyton", 21)
	undergroundSystem1.CheckOut1(2, "Paradise", 30)
	fmt.Println("Output: ", undergroundSystem1.GetAverageTime1("Leyton", "Paradise")) // return 6.66667
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
	in := map[int]*Station{}
	out := map[string][]int{}
	return UndergroundSystem1{in: in, out: out}
}

// CheckIn1 A customer with id card equal to id, gets in the station stationName at time t.
func (us *UndergroundSystem1) CheckIn1(id int, stationName string, t int) {
	station := &Station{stationName, t}
	us.in[id] = station
}

// CheckOut1 A customer with id card equal to id, gets out from the station stationName at time t.
func (us *UndergroundSystem1) CheckOut1(id int, stationName string, t int) {
	station := us.in[id]
	name := station.stationName + stationName
	time := t - station.time
	out := us.out[name]
	out = append(out, time)
	us.out[name] = out
}

// GetAverageTime1 Returns the average time to travel between the startStation and the endStation.
func (us *UndergroundSystem1) GetAverageTime1(startStation string, endStation string) float64 {
	name := startStation + endStation
	out := us.out[name]
	sum := 0
	for _, time := range out {
		sum = sum + time
	}
	return float64(sum) / float64(len(out))
}
