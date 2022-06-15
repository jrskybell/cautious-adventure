package main

import "fmt"

func main() {
	const applePrice float64 = 5.99
	const pearPrice float64 = 7
	const ourBudget float64 = 23
	appleCount := 9
	pearCount := 8

	// Перший пункт
	fmt.Printf("1. Скільки грошей треба витратити, щоб купити %v яблук та %v груш?:\n", appleCount, pearCount)
	appleAndPearPrice := applePrice*float64(appleCount) + pearPrice*float64(pearCount)
	fmt.Println("— Потрібно витратити", appleAndPearPrice, "гривень.")
	println()

	// Другий пункт
	fmt.Println("2. Скільки груш ми можемо купити?:")
	pearsCanBuy := ourBudget / pearPrice
	fmt.Println("— Ми можемо купити", int(pearsCanBuy), "груш.")
	println()

	// Третій пункт
	fmt.Println("3. Скільки яблук ми можемо купити?:")
	applesCanBuy := ourBudget / applePrice
	fmt.Println("— Ми можемо купити", int(applesCanBuy), "яблук.")
	println()

	// Четвертий пункт
	appleCount = 2
	pearCount = 2
	fmt.Printf("4. Чи ми можемо купити %v груші та %v яблука?:\n", pearCount, appleCount)
	appleAndPearPrice = pearPrice*float64(pearCount) + applePrice*float64(appleCount)
	if appleAndPearPrice <= ourBudget {
		fmt.Printf("— Так! Ми можемо купити %v груші та %v яблука!", appleCount, pearCount)
	} else {
		fmt.Printf("— Ні. Нажаль, ми не можемо купити %v груші та %v яблука...", appleCount, pearCount)
	}
}
