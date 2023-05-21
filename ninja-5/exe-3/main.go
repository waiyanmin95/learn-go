package main

import "fmt"

type vehicle struct {
	door  string
	color string
}
type truck struct {
	vehicle
	fourWheel bool
}

type sedan struct {
	vehicle
	luxury bool
}

func main() {
	v1 := vehicle{
		door:  "four",
		color: "white",
	}
	fmt.Println(v1)

	t1 := truck{
		vehicle: vehicle{
			door:  "four",
			color: "yellow",
		},
		fourWheel: true,
	}
	fmt.Println(t1)

	s1 := sedan{
		vehicle: vehicle{
			door:  "two",
			color: "black",
		},
		luxury: true,
	}

	fmt.Println(s1)

	fmt.Println(v1.door, v1.color)
	fmt.Println(t1.door, t1.fourWheel, t1.color)
	fmt.Println(s1.luxury, s1.door, s1.color)
}
