package main

import (
	"fmt"
	"math"
)

// default configuration values
const (
	DefaultMatrixHeight     = 5
	DefaultMatrixWidth      = 5
	DefaultYawDirection     = 0 // 0 is outwards, 180 is inwards
	DefaultTopRollDirection = 0 // 0 is upwards, 180 is downwards
)

type Bread struct {
	identity         string  // "A" or "B"
	yawDirection     float64 // where the upward is
	topRollDirection float64
	matrixWidth      int
	matrixHeight     int
}

type Spread struct {
	peanutButter PBSpread
	jelly        JellySpread
}

type PBSpread struct {
	positionArray [][]PBParticle
	adjacentBread Bread
}

type JellySpread struct {
	positionArray [][]JellyParticle
	adjacentBread Bread
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

var peanutButter PBSpread
var jelly JellySpread
var breadA Bread
var breadB Bread
var breads = []*Bread{&breadA, &breadB}
var sandwich Sandwich

func main() {
	prepare()
	pastePB(&peanutButter)
	pasteJelly(&jelly)
	sandwich = makeSandwich(breadA, breadB, peanutButter, jelly)
	fmt.Println(sandwich)
}

func prepare() {
	breadA.identity = "A"
	breadB.identity = "B"
	for _, bread := range breads {
		bread.matrixWidth = DefaultMatrixWidth
		bread.matrixHeight = DefaultMatrixHeight
		bread.yawDirection = DefaultYawDirection
		bread.topRollDirection = DefaultTopRollDirection
	}

	// set breadB as the top bread by facing down
	breadB.topRollDirection = math.Abs(180 - DefaultTopRollDirection)

	// note: breadA and breadB should not interchange

	// bind peanutButter to breadA and jelly to breadB
	peanutButter.adjacentBread = breadA
	jelly.adjacentBread = breadB
}

func pastePB(content *PBSpread) {
	// create a two-dimensional array signifying the position of particles

	var temp1DList []PBParticle
	var temp2DList [][]PBParticle
	for i := 0; i < content.adjacentBread.matrixWidth; i++ {
		temp1DList = append(temp1DList, PBParticle{}) // creates one-dimensional array with the length of adjacentBread's matrixWidth
	}
	for j := 0; j < content.adjacentBread.matrixHeight; j++ {
		temp2DList = append(temp2DList, temp1DList) // copies the 1D array by adjacentBread's matrixHeight
	}
	content.positionArray = temp2DList // copies the temporary 2D array to the content passed by
}

func pasteJelly(content *JellySpread) {
	// same function as pastePB, for JellySpread
	// refer to pastePB for explanations

	var temp1DList []JellyParticle
	var temp2DList [][]JellyParticle
	for i := 0; i < content.adjacentBread.matrixWidth; i++ {
		temp1DList = append(temp1DList, JellyParticle{})
	}
	for j := 0; j < content.adjacentBread.matrixHeight; j++ {
		temp2DList = append(temp2DList, temp1DList)
	}
	content.positionArray = temp2DList
}

func makeSandwich(topBread Bread, bottomBread Bread, peanutButter PBSpread, jelly JellySpread) Sandwich {
	// returns a Sandwich value consisting of the ingredients
	return Sandwich{
		topBread: topBread,
		spread: Spread{
			peanutButter: peanutButter,
			jelly:        jelly,
		},
		bottomBread: bottomBread,
	}
}
