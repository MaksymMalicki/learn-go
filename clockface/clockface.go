package clockface

import (
	"encoding/xml"
	"math"
	"time"
)

type Point struct {
	X float64
	Y float64
}

type SVG struct {
	XMLName xml.Name `xml:"svg"`
	Xmlns   string   `xml:"xmlns,attr"`
	Width   string   `xml:"width,attr"`
	Height  string   `xml:"height,attr"`
	ViewBox string   `xml:"viewBox,attr"`
	Version string   `xml:"version,attr"`
	Circle  Circle   `xml:"circle"`
	Line    []Line   `xml:"line"`
}

type Circle struct {
	Cx float64 `xml:"cx,attr"`
	Cy float64 `xml:"cy,attr"`
	R  float64 `xml:"r,attr"`
}

type Line struct {
	X1 float64 `xml:"x1,attr"`
	Y1 float64 `xml:"y1,attr"`
	X2 float64 `xml:"x2,attr"`
	Y2 float64 `xml:"y2,attr"`
}

const secondHandLength = 90
const minuteHandLength = 80
const hourHandLength = 50
const clockCentreX = 150
const clockCentreY = 150

const (
	secondsInHalfClock = 30
	secondsInClock     = 2 * secondsInHalfClock
	minutesInHalfClock = 30
	minutesInClock     = 2 * minutesInHalfClock
	hoursInHalfClock   = 6
	hoursInClock       = 2 * hoursInHalfClock
)

func SecondHand(t time.Time) Point {
	p := secondHandPoint(t)
	p = Point{p.X * secondHandLength, p.Y * secondHandLength} // scale
	p = Point{p.X, -p.Y}                                      // flip
	p = Point{p.X + clockCentreX, p.Y + clockCentreY}         // translate
	return p
}

func secondsInRadians(tm time.Time) float64 {
	return (math.Pi / (secondsInHalfClock / (float64(tm.Second()))))
}

func secondHandPoint(tm time.Time) Point {
	return angleToPoint(secondsInRadians(tm))

}

func minutesInRadians(tm time.Time) float64 {
	return (secondsInRadians(tm) / secondsInClock) +
		(math.Pi / (minutesInHalfClock / float64(tm.Minute())))
}

func minuteHandPoint(tm time.Time) Point {
	return angleToPoint(minutesInRadians(tm))

}

func hoursInRadians(tm time.Time) float64 {
	return (minutesInRadians(tm) / hoursInClock) + (math.Pi / (hoursInHalfClock / float64(tm.Hour()%hoursInClock)))
}

func hourHandPoint(tm time.Time) Point {
	return angleToPoint(hoursInRadians(tm))

}

func angleToPoint(angle float64) Point {
	x := math.Sin(angle)
	y := math.Cos(angle)

	return Point{x, y}
}
