package binarySearch

// 二分法，类比 875
func shipWithinDays(weights []int, days int) int {
	sum := 0
	for _, weight := range weights {
		sum += weight
	}
	light, heavy := maxOfInts(weights), sum
	for light <= heavy {
		mid := light + (heavy-light)>>1
		if totalTimeForShip(weights, mid) <= days {
			heavy = mid - 1
		} else {
			light = mid + 1
		}
	}
	return light
}

func totalTimeForShip(weights []int, capacity int) int {
	time := 1 // 初始化为第一天
	c := capacity
	i := 0
	for i < len(weights) {
		if c >= weights[i] {
			c -= weights[i] // 这一天还可以装
			i++
		} else {
			time++ // 这一天已经装满了，在下一天再装
			c = capacity
		}
	}
	return time
}
