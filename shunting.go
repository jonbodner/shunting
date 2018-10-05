package sidecar

import (
	"errors"
	"fmt"
	"strconv"
	"strings"
)

func Calc(s string) float64 {
	var ops []string
	var nums []float64

	calcTop := func() error {
		if len(ops) == 0 || len(nums) < 2 {
			return errors.New("Syntax error")
		}
		op := ops[len(ops)-1]
		num1 := nums[len(nums)-1]
		num2 := nums[len(nums)-2]
		ops = ops[:len(ops)-1]
		nums = nums[:len(nums)-2]
		switch op {
		case "+":
			nums = append(nums, num2+num1)
		case "-":
			nums = append(nums, num2-num1)
		case "*":
			nums = append(nums, num2*num1)
		case "/":
			nums = append(nums, num2/num1)
		default:
			return fmt.Errorf("how did I get here? %s %f %f", op, num2, num1)
		}
		return nil
	}

	parts := strings.Split(s, " ")
	for _, v := range parts {
		switch v {
		case "+", "-":
			l := len(ops)
			for l > 0 && (ops[l-1] == "*" || ops[l-1] == "/") {
				err := calcTop()
				if err != nil {
					fmt.Println(err)
					return 0
				}
				ops = ops[:l-1]
				l = len(ops)
			}
			ops = append(ops, v)
		case "*", "/":
			l := len(ops)
			for l > 0 && (ops[l-1] == "*" || ops[l-1] == "/") {
				err := calcTop()
				if err != nil {
					fmt.Println(err)
					return 0
				}
				ops = ops[:l-1]
				l = len(ops)
			}
			ops = append(ops, v)
		case "(":
			ops = append(ops, v)
		case ")":
			found := false
			l := len(ops)
			for l > 0 {
				if ops[l-1] == "(" {
					ops = ops[:l-1]
					found = true
					break
				}
				err := calcTop()
				if err != nil {
					fmt.Println(err)
					return 0
				}
				l = len(ops)
			}
			if !found {
				return 0
			}
		default:
			//numbers
			n, err := strconv.ParseFloat(v, 64)
			if err != nil {
				fmt.Println(err)
				return 0
			}
			nums = append(nums, n)
		}
	}
	l := len(ops)
	for l > 0 {
		err := calcTop()
		if err != nil {
			fmt.Println(err)
			return 0
		}
		l = len(ops)
	}
	if len(nums) != 1 {
		return 0
	}
	return nums[0]
}
