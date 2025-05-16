//Emi Kostylev 10552236
//Example code for Assignment 2 Part B

/*
This is a simple program that asks the user for an integer input and then uses a switch case to turn that into a day of the week.
Then the program asks if you would like to add another amount of days, and it will calculate what day of the week that will be.
The program gives you a choice to ask if you want simply a day, or use time using a float. This is to demonstrate Go not allowing method overflowing, and similar methods require different names.
Also as a simple demonstration of concurrency, a goroutine gets the systems clock, and prints what day and time it actually is concurrently.
*/

package main

import (
	"fmt"
	"math"
	"time"
)

// Switch case for days of the week
func printDay(day int) {
	switch day {
	case 1:
		fmt.Println("Monday")
	case 2:
		fmt.Println("Tuesday")
	case 3:
		fmt.Println("Wednesday")
	case 4:
		fmt.Println("Thursday")
	case 5:
		fmt.Println("Friday")
	case 6:
		fmt.Println("Saturday")
	case 7:
		fmt.Println("Sunday")
	default:
		panic(fmt.Sprintf("User inputed an invalid day!", printDay)) //Simple demo of a panic, it stops execution and gives a predefined error message of what happened
	}
}

// Prints the day and time
func printDayAndTime(day float64) {
	dayInt := int(math.Floor(day))
	timeOfDay := day - float64(dayInt)

	//Wrap around the days so we can do more than just 7
	dayWrapped := ((dayInt-1)%7+7)%7 + 1
	printDay(dayWrapped)

	//Time of day (More like periods of the day)
	switch {
	case timeOfDay <= 0.25:

		fmt.Println("Morning")
	case timeOfDay <= 0.5:
		fmt.Println("Midday")
	case timeOfDay <= 0.75:

		fmt.Println("Afternoon")
	case timeOfDay <= 1:

		fmt.Println("Night")
	default:

		fmt.Println("-")
	}
}

// While this is an unessecary addition to the code, I wanted to demonstrate the Go does not have method overflowing, and if you want similar methods, different names are needed
func AddInt(dayOne, dayTwo int) int {
	return dayOne + dayTwo
}
func AddFloat(dayOne, dayTwo float64) float64 {
	return dayOne + dayTwo
}

// Goroutine to get the users current day from their system
func getUserDay(ch chan<- string) {
	weekday := time.Now().Weekday().String()
	ch <- fmt.Sprintf("Today is actually: %s", weekday)
}

// Goroutine to get the users current time from their system
func getUserTime(ch chan<- string) {
	currentTime := time.Now().Format("03:04 PM")
	ch <- fmt.Sprintf("The actual time is: %s", currentTime)
}

// Main
func main() {
	var todayInt int
	fmt.Print("What day is it today? (1-7): ")
	fmt.Scan(&todayInt)

	fmt.Print("Today is then: ")
	printDay(todayInt)

	//Run goroutines
	dayCh := make(chan string)
	timeCh := make(chan string)
	go getUserDay(dayCh)
	go getUserTime(timeCh)

	var choice string
	fmt.Print("Do you want to add more days or time? (Y/N): ")
	fmt.Scan(&choice)

	if choice == "Y" || choice == "y" { //Note that go uses || rather than or in Python which makes it less readable for users
		var calcType string
		fmt.Print("Add days only (D) or day and time (T)? (D/T): ")
		fmt.Scan(&calcType)

		if calcType == "D" || calcType == "d" {
			var addInt int
			fmt.Print("Enter number of days to add: ")
			fmt.Scan(&addInt)

			total := AddInt(todayInt, addInt)
			resultDay := ((total-1)%7+7)%7 + 1

			fmt.Print("Resulting day: ")
			printDay(resultDay)

		} else if calcType == "T" || calcType == "t" {
			var todayFloat = float64(todayInt)
			var addFloatVal float64
			fmt.Print("Enter number of days: ")
			fmt.Scan(&addFloatVal)

			sum := AddFloat(todayFloat, addFloatVal)
			fmt.Print("It will be: ")
			printDayAndTime(sum)

		} else {
			fmt.Println("Invalid input.")
		}
	} else {
		fmt.Println("\nYour actual system date and time:")
	}

	//Print system time and date from goroutines
	fmt.Println(<-dayCh)
	fmt.Println(<-timeCh)
}
