package lasagna

// TODO: define the 'PreparationTime()' function
func PreparationTime(layers []string, avgPrepTime int) int {
    time := avgPrepTime
    if time == 0 {
        time = 2
    }
    estimatedTotalPrepTime := len(layers) * time
    return estimatedTotalPrepTime
}


// TODO: define the 'Quantities()' function
func Quantities(layers []string) (int, float64) {
    noodleGrams := 50
    sauceLiters := 0.2

    noodle := 0
    sauce := 0

    for _, layer := range layers {
        switch layer {
        case "noodles":
            noodle++
        case "sauce":
            sauce++
        }
    }

    finalNoodleGrams := noodleGrams * noodle
    finalSauceLiters := sauceLiters * float64(sauce)
    return finalNoodleGrams, finalSauceLiters
}


// TODO: define the 'AddSecretIngredient()' function
func AddSecretIngredient(friendsList []string, myList []string) {
	secretItem := friendsList[len(friendsList)-1]
    myList[len(myList)-1] = secretItem
}

// TODO: define the 'ScaleRecipe()' function
func ScaleRecipe(quantities []float64, portions int) []float64 {
    scaledQuantities := make([]float64, len(quantities))
    scalingFactor := float64(portions) / 2.0
    for i, quantity := range quantities {
        scaledQuantities[i] = quantity * scalingFactor
    }
    return scaledQuantities
}


// Your first steps could be to read through the tasks, and create
// these functions with their correct parameter lists and return types.
// The function body only needs to contain `panic("")`.
// 
// This will make the tests compile, but they will fail.
// You can then implement the function logic one by one and see
// an increasing number of tests passing as you implement more 
// functionality.
