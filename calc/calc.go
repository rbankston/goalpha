package calc

//LbmRate is the Lean Body Mass Rate for calucations
func LbmRate(f int) int {
	if f < 12 {
		return 17
	} else if f < 15 {
		return 16
	} else if f < 19 {
		return 15
	} else if f < 22 {
		return 14
	} else {
		return 13
	}
}

//LbmNumber is the Lean Body Mass Number for calucations
func LbmNumber(w, f int) float64 {
	wf := float64(w)
	ff := float64(f)
	var ffm = (100.00 - ff) / 100.00
	var number = ffm * wf
	return number
}

//MaintCalories is the Maintenance Calories needed for calucations
func MaintCalories(r int, n float64) float64 {
	var rf = float64(r)
	var calories = rf * n
	return calories
}
