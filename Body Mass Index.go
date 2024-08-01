package main 

import "fmt"

func main() {
	var weightInMeter, heightInKg float64 

	fmt.Println("Enter body weight: ")
	fmt.Scanln(&weightInMeter)

	fmt.Println("Enter body height: ")
	fmt.Scanln(&heightInKg)

	bodyMassIndex := weightInMeter / (heightInKg * heightInKg)

	switch {
	case bodyMassIndex < 18.5:
		fmt.Println("Underweight")

	case bodyMassIndex >= 18.5 && bodyMassIndex < 24.9:
		fmt.Println("Normal weight")

	case bodyMassIndex >= 25.0 && bodyMassIndex < 30:
		fmt.Println("Overweight")

	default:
		fmt.Println("Obese")

	}
	fmt.Println("Your BMI is: ", bodyMassIndex)
}




