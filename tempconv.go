package tempconv

import (
	"fmt"
)

type Celsius float64
type Fahrenheit float64
type Kelvin float64

const (
	AbsoluteZeroC Celsius = -273.15
	FreezingC     Celsius = 0
	BoilingC      Celsius = 100
)

func (c Celsius) String() string { return fmt.Sprintf("%g°C", c) }

func (f Fahrenheit) String() string { return fmt.Sprintf("%g°F", f) }

//2.1 SOLUTION : converts a Kelvin temperature to Celsius.
func (k Kelvin) kelvinToCelsius() Celsius { return Celsius((k - 273.15)) }

// 2.2 SOLUTION : VARIOUS TYPE CONVERSIONS OF MEASURING ITEMS

type Measurement interface {
	String() string
}

type Distance struct {
	meters float64
}

func FromFeet(f float64) Distance {
	return Distance{f * 0.3048}
}

func FromMeters(m float64) Distance {
	return Distance{m}
}

func (d Distance) String() string {
	return fmt.Sprintf("%.3gm = %.3gft", d.meters, d.Feet())
}

func (d Distance) Meters() float64 {
	return d.meters
}

func (d Distance) Feet() float64 {
	return d.meters / 0.3048
}

// 2.3 SOLUTION: popcount table re-implementation

var pc [256]byte

func popCountImplementat(x uint64) int {

	for i := range pc {
		pc[i] = pc[i/2] + byte(i&1)

	}

	sum := 0
	for i := 0; i < 8; i++ {
		sum += int(pc[byte(x>>uint(i))])
	}
	return sum
}

// 2.4 SOLUTION: implementing shift value

func popcountShift(x uint64) int {
	count := 0
	mask := uint64(1)
	for i := 0; i < 64; i++ {
		if x&mask > 0 {
			count++
		}
		x >>= 1
	}
	return count
}

// 2.5 SOLUTION:

func popcount(x uint64) int {
	count := 0
	for x != 0 {
		x &= x - 1
		count++
	}
	return count
}
