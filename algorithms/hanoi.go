package algorithms

import "fmt"

func TowersOfHanoi(disks uint) {
	towersOfHanoi(disks, "A", "C", "B")
}

func towersOfHanoi(n uint, from, to, aux string) {
	if n > 0 {
		towersOfHanoi(n - 1, from, aux, to)
		fmt.Printf("Moving from %s to %s\n", from, to)
		towersOfHanoi(n - 1, aux, to, from)
	}
}
