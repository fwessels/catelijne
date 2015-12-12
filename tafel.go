package main

import (
	"fmt"
	"strconv"
)

func main() {
	fmt.Println("tafel spel")
	fmt.Println("we beginnen met de tafel van 1 veel succes")
	fmt.Println("bij tien goede antwoorden kan je naar de volgende tafel")
	for tafel := 1; tafel < 10; tafel++ {
		for keer := 1; keer < 10; keer++ {
			fmt.Printf(" hoeveel is %d keer %d?\n", keer, tafel)
			answer := ""
			fmt.Scanln(&answer)
			result, _ := strconv.Atoi(answer)
			if result == tafel*keer {
				fmt.Println("goed zo!")
			} else {
				fmt.Println("jammer dat is niet goed")
			}
		}

	}

}
