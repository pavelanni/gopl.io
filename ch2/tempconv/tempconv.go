// Package tempconv performs Celsius and Fahrenheit conversions
package tempconv

import "fmt"

// Celsius is for Celsius temperature
type Celsius float64

// Fahrenheit is for Fahrenheit temperature
type Fahrenheit float64

// Kelvin is for Kelvin temperature
type Kelvin float64

const (
	// AbsoluteZeroC defines the absolute zero temperature in Celsius
	AbsoluteZeroC Celsius = -273.15
	// FreezingC defines the freezing temperature in Celsius
	FreezingC Celsius = 0
	// BoilingC defines the boiling temperature in Celsius
	BoilingC Celsius = 100
)

func (c Celsius) String() string    { return fmt.Sprintf("%g°C", c) }
func (f Fahrenheit) String() string { return fmt.Sprintf("%g°F", f) }
func (k Kelvin) String() string     { return fmt.Sprintf("%gK", k) }
