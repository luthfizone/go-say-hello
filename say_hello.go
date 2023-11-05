package gosayhello

import "strconv"

func Hello(name string, age int) string {
	nameStr := name
	ageStr := strconv.Itoa(age)
	return "Halo" + " " + nameStr + ", your age is " + ageStr + " years old"
}
