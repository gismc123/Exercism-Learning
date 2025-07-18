package speed

// TODO: define the 'Car' type struct
type Car struct {
    speed       int
    battery     int
    batteryDrain int
    distance int
}

// NewCar creates a new remote controlled car with full battery and given specifications.
func NewCar(speed, batteryDrain int) Car {
	return Car{
        speed:       speed,
        battery:     100, // Full battery
        batteryDrain: batteryDrain,
        distance: 0,
    }
}

// TODO: define the 'Track' type struct
type Track struct {
    distance int
}

// NewTrack creates a new track
func NewTrack(distance int) Track {
	return Track{distance}
}

// Drive drives the car one time. If there is not enough battery to drive one more time,
// the car will not move.
func Drive(car Car) Car {
    if car.battery < car.batteryDrain {
        return car // Not enough battery to drive
    }
    car.battery -= car.batteryDrain
    car.distance += car.speed
    return car
}

// CanFinish checks if a car is able to finish a certain track.
func CanFinish(car Car, track Track) bool {
    carTimesDrive := car.battery/car.batteryDrain
    carCapacity := carTimesDrive*car.speed
    if carCapacity >= track.distance{
        return true
    }
    return false
    //return car.battery >= track.distance*car.batteryDrain

}
