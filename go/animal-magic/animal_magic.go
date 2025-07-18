package chance

import "math/rand"

// RollADie returns a random int d with 1 <= d <= 20.
func RollADie() int {
	d := rand.Intn(20)
	d += 1
	return d
}

// GenerateWandEnergy returns a random float64 f with 0.0 <= f < 12.0.
func GenerateWandEnergy() float64 {
	r := rand.Float64()
	r *= 12
	return r
}

// ShuffleAnimals returns a slice with all eight animal strings in random order.
func ShuffleAnimals() []string {
	x := []string{"ant","beaver","cat","dog","elephant","fox","giraffe","hedgehog"}
	rand.Shuffle(len(x), func(i, j int) {
    	x[i], x[j] = x[j], x[i]
	})
		return x
}
