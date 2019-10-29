// Convert converts the input into various differnt units of length, temperature, and weight etc.
package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"

	"gobook/ch2/2.2/convert"
)

/*
 * Write a general-purpose unit-conversion program analogous to cf
 * that reads numbers from its command-line arguments or from the
 * standard input if there are no arguments, and converts each
 * number into units like temperature in Celsius and Fahrenheit,
 * length in feet and meters, weight in pounds and kilograms, and
 * the like.
 */

func main() {
	if len(os.Args) > 1 {
		for _, arg := range os.Args[1:] {
			displayConversions(arg)
		}
	} else {
		fmt.Println("Enter numbers to be converted one at a time")
		scanner := bufio.NewScanner(os.Stdin)
		for scanner.Scan() {
			displayConversions(scanner.Text())
		}
	}
}

func displayConversions(arg string) {
	num, err := strconv.Atoi(arg)
	if err != nil {
		fmt.Fprintf(os.Stderr, "%v\n", err)
		os.Exit(1)
	}
	fmt.Printf("%s\n", temperatures(num))
	fmt.Printf("%s\n", weights(num))
	fmt.Printf("%s\n\n", distance(num))

}

func temperatures(num int) string {
	c := convert.Celsius(num)
	f := convert.Fahrenheit(num)
	return fmt.Sprintf("%s = %s, %s = %s", c, convert.CToF(c), f, convert.FToC(f))
}

func weights(num int) string {
	p := convert.Pounds(num)
	k := convert.Kilograms(num)
	return fmt.Sprintf("%s = %s, %s = %s", p, convert.PToK(p), k, convert.KToP(k))
}

func distance(num int) string {
	m := convert.Meters(num)
	f := convert.Feet(num)
	return fmt.Sprintf("%s = %s, %s = %s", m, convert.MToF(m), f, convert.FToM(f))
}
