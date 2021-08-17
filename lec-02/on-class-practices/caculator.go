package main

import "fmt"

func main() {

}
func calculator(operator string, a, b int) {
	switch operator {
	case "mul":
		return a * b
	case "sub":
		return a - b
	case "add":
		return a + b
	case "div":
		if b == 0 {
			err := fmt.Errorf("Can not divide for 0")
			return 0, err
		}
		else{
			return a / b
		}
	}
}
