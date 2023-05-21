package main

import (
	"fmt"
	"github.com/ivan-amity/learn-go/exercises/Package/Lesson-2/dates"
)

func main() {
	days := 3
	fmt.Println("You appointment is in", days, "Days")
	fmt.Println("With follow-up in", days+dates.DaysinWeek, "Days")

	fmt.Println("Weeks To Days:", dates.WeeksToDays(8), "Days")
	fmt.Printf("Days to Weeks: %.2f %v \n", dates.DaystoWeeks(14), "Weeks")
	fmt.Printf("Days to Weeks: %.2f %v", dates.DaystoWeeks(77), "Weeks")
}
