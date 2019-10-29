// Package tempconv performs Celsius and Fahrenheit conversions.
package tempconv

import "fmt"

/*
 * Exercise 2.1
 * Add types, constants, and functions to tempconv for processing
 * temperatures in the Kelvin scale, where zero Kelvin is -273.15ºC
 * and a difference of 1K has the same magnitude as 1ºc
 */

type Celsius float64
type Fahrenheit float64
type Kelvin float64

const (
	AbsoluteZeroC Celsius = -273.15
	FreezingC     Celsius = 0
	BoilingC      Celsius = 100
	FreezingK     Kelvin  = 273.15
	BoilingK      Kelvin  = 373.15
	AbsoluteZeroK Kelvin  = 0
)

func (c Celsius) String() string    { return fmt.Sprintf("%g°C", c) }
func (f Fahrenheit) String() string { return fmt.Sprintf("%g°F", f) }
func (k Kelvin) String() string     { return fmt.Sprintf("%gK", k) }
