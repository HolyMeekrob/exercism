package lasagna

func PreparationTime(layers []string, avgPrepTime int) int {
	if avgPrepTime == 0 {
		avgPrepTime = 2
	}

	return len(layers) * avgPrepTime
}

func Quantities(layers []string) (int, float64) {
	noodles := 0
	sauce := 0.0

	for i := 0; i < len(layers); i++ {
		if layers[i] == "noodles" {
			noodles += 50
		} else if layers[i] == "sauce" {
			sauce += 0.2
		}
	}
	return noodles, sauce
}

func AddSecretIngredient(friendList, myList []string) {
	myList[len(myList)-1] = friendList[len(friendList)-1]
}

func ScaleRecipe(amounts []float64, portions int) []float64 {
	scaledAmounts := make([]float64, len(amounts))
	for i := 0; i < len(amounts); i++ {
		scaledAmounts[i] = amounts[i] * (float64(portions) / 2.0)
	}

	return scaledAmounts
}
