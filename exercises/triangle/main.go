package main

import "fmt"

func main() {
	for i := 0; i < 5; i++ {
		for j := 0; j < 5-i; j++ {
			fmt.Print("")
		}
		for j := 0; j <= i; j++ {
			fmt.Print("*")
		}
		fmt.Println("")
	}
}

//func main() {
//	for i := 5; i > 0; i-- {
//		for g := 5; g > i; g-- {
//			fmt.Printf(" ")
//			//fmt.Printf("Inner 1, G = %d, I = %d", g, i)
//		}
//		for j := 0; j < i; j++ {
//			fmt.Printf("* ")
//			//fmt.Printf("Inner 2, J = %d, I = %d", j, i)
//		}
//
//		fmt.Println(" ")
//	}
//}

//func main() {
//	for i := 1; i < 6; i++ {
//		for j := 0; j < i; j++ {
//			fmt.Print("* ")
//		}
//		fmt.Println(" ")
//	}
//}

//func main() {
//	for i := 0; i < 5; i++ {
//		for j := 0; j < 5-i; j++ {
//			fmt.Printf("* ")
//			//fmt.Printf("Inner 2, J = %d, I = %d", j, i)
//		}
//
//		fmt.Println(" ")
//	}
//}
