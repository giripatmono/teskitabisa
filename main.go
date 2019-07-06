package main

import (
	"fmt"
	"math"
	"strconv"
	"strings"
)

type Command interface {
	Execute(nums ...int) string
}

// Simple hello command
// Return: 'Hello World!'
type HelloCommand struct{}

func (p *HelloCommand) Execute(nums ...int) string {
	return "Hello World!"
}

// Calculate the sum from input argument, and print the result
type SumXYCommand struct{}

func (p *SumXYCommand) Execute(nums ...int) string {
	retval := 0
	for _, num := range nums {
		retval += num
	}
	return fmt.Sprintf("%v", retval)
}

// Calculate the total multiplication from input argument, and print the result
type MultiplyXYCommand struct{}

func (p *MultiplyXYCommand) Execute(nums ...int) string {
	retval := 1
	for _, num := range nums {
		retval *= num
	}
	return fmt.Sprintf("%v", retval)
}

// Find first N prime number, and print the result
type FindFirstNPrimeCommand struct{}

func (p *FindFirstNPrimeCommand) Execute(nums ...int) string {
	var retval []string
	N := nums[0]
	number := 2
	primeCount := 0
	for primeCount < N {
		if isPrime(number) {
			retval = append(retval, strconv.Itoa(number))
			primeCount += 1
		}
		number += 1
	}
	return fmt.Sprintf("%v", strings.Join(retval, ", "))
}

// Check whether a number is a prime number. Return boolean
func isPrime(num int) bool {
	for i := 2; i <= int(math.Floor(math.Sqrt(float64(num)))); i++ {
		if num%i == 0 {
			return false
		}
	}
	return true
}

//Find first N Fibonacci sequence, and print the result
type FindFirstNFibonacciCommand struct{}

func (p *FindFirstNFibonacciCommand) Execute(nums ...int) string {
	var retval []string
	N := nums[0]
	one := 0
	two := 1
	var temp int
	for i := 0; i < N; i++ {
		retval = append(retval, strconv.Itoa(one))
		temp = one + two
		one = two
		two = temp
	}
	return fmt.Sprintf("%v", strings.Join(retval, ", "))
}

func ExecFuncByName(name string, nums ...int) string {
	// Register commands
	commands := map[string]Command{
		"Hello":               &HelloCommand{},
		"SumXY":               &SumXYCommand{},
		"MultiplyXY":          &MultiplyXYCommand{},
		"FindFirstNPrime":     &FindFirstNPrimeCommand{},
		"FindFirstNFibonacci": &FindFirstNFibonacciCommand{},
	}

	if command := commands[name]; command == nil {
		return "No such command found!"
	} else {
		return command.Execute(nums...)
	}
}

func main() {

	fmt.Println(ExecFuncByName("Hello"))

	fmt.Println("\n==========\nCalculate Sum X & Y \n==========")
	fmt.Printf("Input : 1, 2 Output : %v", ExecFuncByName("SumXY", 1, 2))

	fmt.Println("\n\n==========\nCalculate Multiply X & Y \n==========")
	fmt.Printf("Input : 1, 2 Output : %v", ExecFuncByName("MultiplyXY", 1, 2))

	fmt.Println("\n\n==========\nFind First N Prime Number \n==========")
	fmt.Printf("Input : 4 Output : %v", ExecFuncByName("FindFirstNPrime", 4))

	fmt.Println("\n\n==========\nFind First N Fibonacci Sequence \n==========")
	fmt.Printf("Input : 4 Output : %v", ExecFuncByName("FindFirstNFibonacci", 4))

	fmt.Println()

}
