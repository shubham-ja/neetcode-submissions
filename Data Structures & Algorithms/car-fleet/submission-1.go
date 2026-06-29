func carFleet(target int, position []int, speed []int) int {
	if len(speed) < 1{
		return 0
	}
	type car struct {
		position int
		speed int
	}
	carList := make([]car,len(position))
	for i, val := range position {
		carList[i] = car {
			position: val,
			speed: speed[i],
		}
	}

	sort.Slice(carList, func(i,j int) bool {
		return carList[i].position < carList[j].position
	})
	
	time := 0.0
	count := 0
	for i := len(carList)-1; i >= 0; i-- {
		cc := carList[i]
		computedTime := float64(target - cc.position)/float64(cc.speed)
		
		if computedTime > time {
			count +=1
			time = computedTime
		}
	}	
	return count
}
