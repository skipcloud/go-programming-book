package convert

import "fmt"

type Celsius float64
type Fahrenheit float64

type Meters float64
type Feet float64

type Pounds float64
type Kilograms float64

func (c Celsius) String() string    { return fmt.Sprintf("%.3gºC", c) }
func (f Fahrenheit) String() string { return fmt.Sprintf("%.3gºF", f) }

func (m Meters) String() string { return fmt.Sprintf("%.3gm", m) }
func (f Feet) String() string   { return fmt.Sprintf("%.3gft", f) }

func (p Pounds) String() string    { return fmt.Sprintf("%.3glb", p) }
func (k Kilograms) String() string { return fmt.Sprintf("%.3gkg", k) }

// CToF converts a Celsius temperature to Fahrenheit.
func CToF(c Celsius) Fahrenheit { return Fahrenheit(c*9/5 + 32) }

// FToC converts a Fahrenheit temperature to Celsius.
func FToC(f Fahrenheit) Celsius { return Celsius((f - 32) * 5 / 9) }

// MToF converts a distance in Meters to Feet.
func MToF(m Meters) Feet { return Feet(m * 3.28084) }

// FToM converts a distance in Feet to Meters.
func FToM(f Feet) Meters { return Meters(f / 3.28084) }

// PToK converts a weight in Pounds to KIlograms.
func PToK(p Pounds) Kilograms { return Kilograms(p / 2.20462) }

// KToP converts a weight in Kilograms to Pounds.
func KToP(k Kilograms) Pounds { return Pounds(k * 2.20462) }
