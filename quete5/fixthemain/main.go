package main

import "github.com/01-edu/z01"

type Door struct {
	state string
}

func PrintStr(s string) {
	for _, r := range s {
		z01.PrintRune(r)
	}
}

func CloseDoor(Door *Door) {
	PrintStr("Door Closing...\n")
	Door.state = "CLOSE"
}

func OpenDoor(Door *Door) {
	PrintStr("Door Opening...\n")
	Door.state = "OPEN"
}

func IsDoorOpen(Door *Door) bool {
	PrintStr("is the Door opened ?\n")
	if Door.state == "OPEN" {
		return true
	} else {
		return false
	}
}

func IsDoorClose(Door *Door) bool {
	PrintStr("is the Door closed ?\n")
	if Door.state == "CLOSE" {
		return true
	} else {
		return false
	}
}

func main() {
	door := &Door{}
	OpenDoor(door)
	if IsDoorClose(door) {
		OpenDoor(door)
	}
	if IsDoorOpen(door) {
		CloseDoor(door)
	}
	if door.state == "OPEN" {
		CloseDoor(door)
	}
}
