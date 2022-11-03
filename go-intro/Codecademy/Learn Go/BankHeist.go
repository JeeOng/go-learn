package main

import (
	"fmt"       //Format
	"math/rand" //Random Number Generator
	"time"      //Time calculation
)

func main() {
	rand.Seed(time.Now().UnixNano())

	isHeistOn := true
	eludedGuards := rand.Intn(100)

	//First Condition - Sneak past guards
	if eludedGuards >= 50 {
		fmt.Println("Looks like you've managed to make it past the guards. Good job, but remember, this is the first step.")
	} else {
		isHeistOn = false
		fmt.Println("Plan a better disguise next time?")
	}
	//Second Condition - Open the Vault
	openedVault := rand.Intn(75)
	if isHeistOn && openedVault >= 50 {
		fmt.Println("Grab and GO!")
	} else if isHeistOn == true || openedVault <= 70 {
		isHeistOn = false
		fmt.Println("The Vault can't be opened!")
	}
	//Third Condition - Leaving
	leftSafely := rand.Intn(5000)
	if isHeistOn {
		switch leftSafely {
		case 1:
			isHeistOn = false
			fmt.Println("Plan a better disguise next time?")
		case 2:
			isHeistOn = false
			fmt.Println("Turns out vault doors don't open from the inside...")
		case 3:
			isHeistOn = false
			fmt.Println("Maybe next time don't trip the alarm as a practical joke")
		case 4:
			isHeistOn = false
			fmt.Println("Who didn't watch the hostages? How did they call the cops?!")
		default:
			fmt.Println("Start the getaway car!")
		}
	}
	//Wrapping Up
	if isHeistOn {
		amtStolen := 10000 + rand.Intn(1000000)
		fmt.Println("amtStolen", amtStolen)
	}

	//Check if the Heist is On
	fmt.Println("isHeistOn is current:", isHeistOn)
}
