package main 

import "fmt"

func main(){

	// favNums3 := [5]float64 {1,2,3,4,5}

	// for _, value := range favNums3 {
	// 	fmt.Println(value);
	// }

	// presAge := make(map[string] int)

	// presAge["key"] = 42

	// fmt.Println(presAge["key"])
	// fmt.Println(len(presAge))

	// listNums := []float64{1,2,3,4,5}

	// for _, value := range listNums {
	// 	fmt.Println(value);
	// }

	// num1, num2 := next2Values(5)
	// fmt.Println(num1, num2)

	fmt.Println(subtractThem(1,2,3,4,5))

}

// func next2Values(number int) (int, int){

// 	return number+1, number+2

// }

func subtractThem(args ...int) int {
	finalValue := 0

	for _, value := range args{
		finalValue -= value
	}

	return finalValue
}