package main

import (
	"fmt"
	"math"
	"strconv"
	"strings"
)

type OpeandErroer rune

func (oe OpeandErroer) Error() string {
	return fmt.Sprintf("operan %s not valid", string(oe))
}

type Process struct {
	Number1 float64
	Number2 float64
	Operand rune
}

func main() {
	fmt.Println("hellow")

	for {
		//get input----------
		input := getinput()
		//pars input--------
		process, err := parse(input)
		if err != nil {
			fmt.Println(err)
			continue
		}
		//calculate
		result, err := process.calculate()
		if err != nil {
			fmt.Println(err)
			continue
		}
		//print----------
		fmt.Println(result)
	}

}

func getinput() string {
	var input string
	fmt.Print(">")
	fmt.Scanln(&input)
	return input
}

func parse(input string) (*Process, error) {
	for _, operand := range "+-/*^" {
		subs := strings.Split(input, string(operand))
		if len(subs) != 1 {
			process := &Process{
				Operand: operand,
			}
			var err error
			process.Number1, err = strconv.ParseFloat(subs[0], 64)
			if err != nil {
				return nil, err
			}
			process.Number2, err = strconv.ParseFloat(subs[1], 64)
			if err != nil {
				return nil, err
			}

			return process, nil

		}
	}
	return nil, OpeandErroer(' ')
}

func (p *Process) calculate() (float64, error) {
	switch p.Operand {
	case '+':
		return p.Number1 + p.Number2, nil
	case '-':
		return p.Number1 - p.Number2, nil
	case '*':
		return p.Number1 * p.Number2, nil
	case '%':
		return math.Mod(p.Number1, p.Number2), nil
	case '/':
		return p.Number1 / p.Number2, nil
	case '^':
		return math.Pow(p.Number1, p.Number2), nil
	}
	return 0, OpeandErroer(p.Operand)
}
