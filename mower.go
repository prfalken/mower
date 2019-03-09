package main

import "container/ring"

type mower struct {
	orientation  string
	xPos, yPos   int
	instructions []string
}

func (mower *mower) moveForward(lawn lawn) {
	switch mower.orientation {
	case "N":
		if mower.yPos+1 > lawn.size {
			break
		}
		mower.yPos++
	case "S":
		if mower.yPos-1 < 0 {
			break
		}
		mower.yPos--
	case "W":
		if mower.xPos-1 < 0 {
			break
		}
		mower.xPos--
	case "E":
		if mower.xPos+1 > lawn.size {
			break
		}
		mower.xPos++

	}
}

func (mower *mower) turn(instruction string) {
	var compass = ring.New(4)
	for _, o := range []string{"N", "E", "S", "W"} {
		compass.Value = o
		compass = compass.Next()
	}
	for {
		if mower.orientation == compass.Value {
			break
		}
		compass = compass.Next()
	}

	if instruction == "L" {
		compass = compass.Prev()
	}
	if instruction == "R" {
		compass = compass.Next()
	}
	mower.orientation = string(compass.Value.(string))

}
