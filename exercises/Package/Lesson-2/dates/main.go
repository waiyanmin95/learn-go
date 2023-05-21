package dates

const DaysinWeek int = 7

func WeeksToDays(w int) int {
	wtd := w * DaysinWeek
	return wtd
}

func DaystoWeeks(d int) float64 {
	dtw := float64(d) / float64(DaysinWeek)
	return dtw
}
