package gosayhello

import "strconv"

func Hello(name string, age int) string {
	return "Halo" + " " + name + " " + "your age" + " " + strconv.Itoa(age) + " years old"
}
