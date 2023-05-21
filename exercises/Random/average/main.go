package main

import "fmt"

//func main() {
//	args := os.Args[1:]
//	var sum float64 = 0
//	for _, num := range args {
//		num, err := strconv.ParseFloat(num, 64)
//		if err != nil {
//			log.Fatal(err)
//		}
//		sum += num
//	}
//
//	count := float64(len(args))
//	fmt.Printf("The average is: %.2f", sum/count)
//}

func main() {
	fmt.Printf("Average Number is: %.2f\n", Avg(3.4, 4.4, 5.5, 2.2))
	fmt.Printf("Average Number is: %.2f\n", Avg(100, 50))
	fmt.Printf("Average Number is: %.2f", Avg(99, 100, 98))
}

func Avg(f ...float64) float64 {
	var sum float64 = 0
	for _, v := range f {
		sum += v
	}
	return sum / float64(len(f))
}
