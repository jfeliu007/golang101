package main

import "fmt"

func main() {
	fmt.Println("Cycles -> Result")
	fmt.Println("0 -> " + fmt.Sprint(getYearSize(0)))
	fmt.Println("1 -> " + fmt.Sprint(getYearSize(1)))
	fmt.Println("4 -> " + fmt.Sprint(getYearSize(4)))
}

func getYearSize(cycleNumber int) int {
	result := 1
	for i := 1; i <= cycleNumber; i++ {
		if i%2 != 0 {
			result = springCycle(result)
		} else {
			result = summerCycle(result)
		}
	}
	return result
}

func springCycle(size int) int {
	return size * 2
}
func summerCycle(size int) int {
	return size + 1
}
