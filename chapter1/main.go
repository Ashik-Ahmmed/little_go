package main

import "fmt"

func getPower() int {
	return 9000
}
func nameAndPower() (string, int) {
	return "Ashik", 9000
}
func main() {
	//var power int = 9000
	name, _ := nameAndPower()
	_, age := nameAndPower()
	fmt.Printf("It is %s over %d \n", name, age)
}
