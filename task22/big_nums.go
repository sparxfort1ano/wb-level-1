package main

import (
	"errors"
	"fmt"
	"log"
	"math/big"
)

func main() {
	fmt.Println("Enter 1 number:")
	num1, err := InputNum()
	if err != nil {
		log.Println(err.Error())
		return
	}

	fmt.Println("Enter 2 number:")
	num2, err := InputNum()
	if err != nil {
		log.Println(err.Error())
		return
	}

	// Sum.
	sumRes := new(big.Int)
	sumRes.Add(num1, num2)
	fmt.Println("Sum result:", sumRes.String())

	// Subtraction.
	subRes := new(big.Int)
	subRes.Sub(num1, num2)
	fmt.Println("Subtraction result:", subRes.String())

	// Multiplication.
	mulRes := new(big.Int)
	mulRes.Mul(num1, num2)
	fmt.Println("Multiplication result:", mulRes.String())

	// Division.
	fmt.Print("Division result: ")
	if err := IsZero(num2); err != nil {
		log.Println(err.Error())
	} else {
		divRes := new(big.Int)
		divRes.Div(num1, num2)
		fmt.Println(divRes.String())
	}
}

// InputNum returns *big.Int value after os.Stdin input.
func InputNum() (*big.Int, error) {
	var numStr string
	fmt.Scan(&numStr)

	numBig := new(big.Int)
	if _, ok := numBig.SetString(numStr, 10); !ok {
		return nil, errors.New("failed to convert str to big.Int")
	}

	return numBig, nil
}

// IsZero return error if *big.Int contains zero value.
func IsZero(num *big.Int) error {
	zero := big.NewInt(0)
	if num.Cmp(zero) == 0 {
		return errors.New("failed to divide by zero")
	}
	return nil
}
