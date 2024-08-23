package main

import (
	"fmt"
	"regexp"
	"strconv"
	"strings"
)

func main() {
	input := "-1/3-1/2+2/5"
	fmt.Println(fractionAddition(input))
}

// if first character inexpression is not minus (-), add (+)
// use regex to get array of fractions [-1/3, -1/2, +2/5]
// for each fraction, split by / and parse to integer
// calculate with previous numerator and denominator. If previous values not found, set previous values with current numerator and denominator
// at the end, find the irreducible fraction by dividing final numerator and denominator with GCD
// dont forget to handle (-) case
// Time complexity: O(n+m * log d) where n is the string lenght, m is number of fractions
// Space complexity: O(m) where m is number of fractions
func fractionAddition(expression string) string {
	if expression[0] != '-' {
		expression = "+" + expression
	}
	regexPattern := `[+-]?\d+/\d+`
	re := regexp.MustCompile(regexPattern)
	fractions := re.FindAllString(expression, -1)

	prevNumerator, prevDenominator := 0, 0
	for _, f := range fractions {
		s := strings.Split(f, "/")
		numerator, _ := strconv.Atoi(s[0])
		denominator, _ := strconv.Atoi(s[1])

		if prevNumerator != 0 && prevDenominator != 0 {
			lcm := lcm(prevDenominator, denominator)
			prevNumerator = prevNumerator * (lcm / prevDenominator)
			numerator = numerator * (lcm / denominator)
			prevNumerator = prevNumerator + numerator
			prevDenominator = lcm
		} else {
			prevNumerator, prevDenominator = numerator, denominator
		}
	}

	if prevNumerator == 0 {
		return "0/1"
	}
	if prevDenominator == 0 {
		return fmt.Sprintf("%d/1", prevNumerator)
	}

	prefix := ""
	if prevNumerator < 0 {
		prevNumerator *= -1
		prefix = "-"
	}

	gcd := gcd(prevNumerator, prevDenominator)
	return fmt.Sprintf("%s%d/%d", prefix, prevNumerator/gcd, prevDenominator/gcd)
}

func gcd(a, b int) int {
	if b > a {
		a, b = b, a
	}

	res := a % b
	if res == 0 {
		return b
	}

	return gcd(b, res)
}

func lcm(a, b int) int {
	return (a * b) / gcd(a, b)
}
