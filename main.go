package main

import "fmt"

type Bread struct {
	yawDirection     float64
	topRollDirection float64
	adjacentSpread   singleSpread
}

type Spread struct {
	peanutButter PBSpread
	jelly        JellySpread
}

type singleSpread struct {
	peanutButter PBSpread
	jelly        JellySpread
	selector     int
}

type PBSpread struct {
	positionArray [][]PBParticle
}

type JellySpread struct {
	positionArray [][]JellyParticle
}

type PBParticle struct {
}

type JellyParticle struct {
}

type Sandwich struct {
	topBread    Bread
	spread      Spread
	bottomBread Bread
}

func main() {
	fmt.Println("hello!")
}
