func carFleet(target int, position []int, speed []int) int {
	type carData struct {
		position 	int
		time 		float64
	}
	data := make([]carData, len(position))
	for i := range position {
		data[i] = carData{
			position: position[i],
			time: float64(target - position[i]) / float64(speed[i]),
		}
	}
	sort.Slice(data, func(i int, j int) bool {
		return data[i].position > data[j].position
	})

	fleets := 0
	fleetTime := -1.0
	for i := range data {
		if data[i].time > fleetTime {
			fleets++
			fleetTime = data[i].time
		}
	}

	return fleets
}
