package main

import (
	"fmt"
)
type Planet string

const (
	Earth = 31557600
	Mercury = 0.2408467 * Earth
	Venus   = 0.61519726 * Earth
	Mars    = 1.8808158 * Earth
	Jupiter = 11.862615 * Earth
	Saturn  = 29.447498 * Earth
	Uranus  = 84.016846 * Earth
	Neptune = 164.79132 * Earth
)

func Age(sec float64, planet Planet) float64 {
	var result float64
	switch planet {
	case "Earth":
		result = sec/Earth
	case "Mercury":
		result = sec/Mercury
	case "Venus":
		result = sec/Venus
	case "Mars":
		result = sec/Mars
	case "Jupiter":
		result = sec/Jupiter
	case "Saturn":
		result = sec/Saturn
	case "Uranus":
		result = sec/Uranus
	case "Neptune":
		result = sec/Neptune
	default:
		fmt.Println("Not know this planet")
	}
	return result
}