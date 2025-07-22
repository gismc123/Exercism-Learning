package birdwatcher

// TotalBirdCount return the total bird count by summing
// the individual day's counts.
func TotalBirdCount(birdsPerDay []int) int {
	totalNumberOfBirds := 0
	for numberOfDays := 0; numberOfDays <= len(birdsPerDay)-1; numberOfDays++ {
	totalNumberOfBirds += birdsPerDay[numberOfDays]
	}
	return totalNumberOfBirds
}

// BirdsInWeek returns the total bird count by summing
// only the items belonging to the given week.
func BirdsInWeek(birdsPerDay []int, week int) int {
	startOfWeek := (week*7)-7
	endOfWeek := (week*7)-1
	totalNumberOfBirds := 0
	for numberOfDays := startOfWeek; numberOfDays <= endOfWeek; numberOfDays++ {
	totalNumberOfBirds += birdsPerDay[numberOfDays]
	}
	return totalNumberOfBirds
}

// FixBirdCountLog returns the bird counts after correcting
// the bird counts for alternate days.
func FixBirdCountLog(birdsPerDay []int) []int {
	for fixDayCount := 0; fixDayCount < len(birdsPerDay)-1; fixDayCount += 2 {
		birdsPerDay[fixDayCount] += 1
	}
	return birdsPerDay
}
