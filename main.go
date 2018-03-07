package main

import (
	"flag"
	"fmt"

	"github.com/rbankston/goalpha/calc"
)

func main() {

	phasePtr := flag.Int("phase", 1, "Phase of the program you are on: Choose 1, 2, 3, or 4")
	weightPtr := flag.Int("weight", 250, "Your current weight in whole numbers")
	fatPtr := flag.Int("fat", 20, "Current Body Fat percentage in whole numbers")

	flag.Parse()

	var fat = *fatPtr
	var weight = *weightPtr
	var phase = *phasePtr

	leanBodyRate := calc.LbmRate(fat)
	leanBodyMass := calc.LbmNumber(weight, fat)
	maintenanceCalories := calc.MaintCalories(leanBodyRate, leanBodyMass)

	if phase == 1 {
		var woProtein = 0.8 * leanBodyMass
		var woCalories = maintenanceCalories - 300.00
		var nwoCalories = maintenanceCalories - 500.00
		var nwoProtein = 0.7 * leanBodyMass
		var woCarbs = 0.0
		var nwoCarbs = 0.0
		var woFat = (woCalories - (woProtein * 4.0) - (woCarbs * 4.0)) / 9.0
		var nwoFat = (nwoCalories - (nwoProtein * 4.0) - (nwoCarbs * 4.0)) / 9.0
		fmt.Printf("Calories on Workout days are %f. Protein on workout days are %f grams.\n", woCalories, woProtein)
		fmt.Printf("Carbs on workout days are %f grams. Fat on workout days are %f grams.\n", woCarbs, woFat)
		fmt.Printf("Calories on non workout days are %f. Protein on non workout days are %f grams.\n", nwoCalories, nwoProtein)
		fmt.Printf("Carbs on non workout days are %f grams. Fat on non workout days are %f.\n", nwoCarbs, nwoFat)
		fmt.Println("During Week 3 you start 30 grams of carbs on workout days.")
		fmt.Println("During Week 4 you get 100 grams of carbs on workout days and 70 on non-workout days")
	} else if phase == 2 {
		var woCalories = maintenanceCalories - 200.00
		var nwoCalories = maintenanceCalories - 600.00
		var woProtein = leanBodyMass
		var nwoProtein = 0.8 * leanBodyMass
		var woCarbs = 0.75 * leanBodyMass
		var nwoCarbs = 0.3 * leanBodyMass
		var woFat = (woCalories - (woProtein * 4.0) - (woCarbs * 4.0)) / 9.0
		var nwoFat = (nwoCalories - (nwoProtein * 4.0) - (nwoCarbs * 4.0)) / 9.0
		fmt.Printf("Calories on Workout days are %f. Protein on workout days are %f grams.\n", woCalories, woProtein)
		fmt.Printf("Carbs on workout days are %f grams. Fat on workout days are %f grams.\n", woCarbs, woFat)
		fmt.Printf("Calories on non workout days are %f. Protein on non workout days are %f grams.\n", nwoCalories, nwoProtein)
		fmt.Printf("Carbs on non workout days are %f grams. Fat on non workout days are %f.\n", nwoCarbs, nwoFat)
	} else if phase == 3 {
		var woCalories = maintenanceCalories + 400.00
		var nwoCalories = maintenanceCalories - 200.00
		var woProtein = 1.5 * leanBodyMass
		var nwoProtein = 1.25 * leanBodyMass
		var woCarbs = 1.00 * leanBodyMass
		var nwoCarbs = 0.5 * leanBodyMass
		var woFat = ((woCalories - (woProtein * 4.0) - (woCarbs * 4)) / 9)
		var nwoFat = (nwoCalories - (nwoProtein * 4) - (nwoCarbs * 4)) / 9
		fmt.Printf("Calories on Workout days are %f. Protein on workout days are %f grams.\n", woCalories, woProtein)
		fmt.Printf("Carbs on workout days are %f grams. Fat on workout days are %f grams.\n", woCarbs, woFat)
		fmt.Printf("Calories on non workout days are %f. Protein on non workout days are %f grams.\n", nwoCalories, nwoProtein)
		fmt.Printf("Carbs on non workout days are %f grams. Fat on non workout days are %f.\n", nwoCarbs, nwoFat)
	} else {
		var woCalories = maintenanceCalories + 300.00
		var nwoCalories = maintenanceCalories - 400.00
		var woProtein = 1.5 * leanBodyMass
		var nwoProtein = 1 * leanBodyMass
		var woCarbs = 1 * leanBodyMass
		var nwoCarbs = 0.25 * leanBodyMass
		var woFat = (woCalories - (woProtein * 4.0) - (woCarbs * 4.0)) / 9.0
		var nwoFat = (nwoCalories - (nwoProtein * 4.0) - (nwoCarbs * 4.0)) / 9.0
		fmt.Printf("Calories on Workout days are %f. Protein on workout days are %f grams.\n", woCalories, woProtein)
		fmt.Printf("Carbs on workout days are %f grams. Fat on workout days are %f grams.\n", woCarbs, woFat)
		fmt.Printf("Calories on non workout days are %f. Protein on non workout days are %f grams.\n", nwoCalories, nwoProtein)
		fmt.Printf("Carbs on non workout days are %f grams. Fat on non workout days are %f.\n", nwoCarbs, nwoFat)
	}

}
