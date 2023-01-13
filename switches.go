package main

import (
	"math/rand"
	"time"
)

func calculateGrade(score int) string {
	var grade string
	switch {
	case score < 60:
		grade = "F"
	case score < 70:
		grade = "D"
	case score < 80:
		grade = "C"
	case score < 90:
		grade = "B"
	default:
		grade = "A"
	}
	return grade
}

func switch_example() {
	rand.Seed(time.Now().UnixNano())
	point := rand.Intn(100)
	grade := calculateGrade(point)
	println(point, grade)
}
