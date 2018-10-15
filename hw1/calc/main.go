package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strconv"
	"strings"

	"github.com/pkg/errors"
)

type stack struct {
	data []int
}

func newStack() *stack {
	return &stack{
		data: make([]int, 0),
	}
}

func (s *stack) push(a int) {
	s.data = append(s.data, a)
}

func (s *stack) pop() (int, error) {
	len := len(s.data)
	if len == 0 {
		return 0, errors.New("empty stack")
	}
	res := s.data[len-1]
	s.data = s.data[:len-1]
	return res, nil
}

func calcStr(out io.Writer, s *stack, strings []string) error {
	for _, str := range strings {
		switch str {
		case " ", "\n", "":
		case "+", "-", "*", "/":
			num1, err := s.pop()
			if err != nil {
				return errors.Wrap(err, "can't calculate line")
			}
			num2, err := s.pop()
			if err != nil {
				return errors.Wrap(err, "can't calculate line")
			}
			var num int
			switch str {
			case "+":
				num = num1 + num2
			case "-":
				num = num1 - num2
			case "*":
				num = num1 * num2
			case "/":
				num = num1 / num2
			}
			s.push(num)
		case "=":
			res, err := s.pop()
			if err != nil {
				return errors.Wrap(err, "can't get res")
			}
			fmt.Fprintf(out, "Result = %d\n", res)
		default:
			num, err := strconv.Atoi(str)
			if err != nil {
				return errors.Wrap(err, "can't read str")
			}
			s.push(num)
		}
	}
	return nil
}

func calc(out io.Writer, in io.Reader) error {
	input := bufio.NewScanner(in)
	s := newStack()
	for input.Scan() {
		str := strings.Split(input.Text(), " ")
		if err := calcStr(out, s, str); err != nil {
			return errors.Wrap(err, "can't read integer")
		}
	}
	if res, err := s.pop(); err == nil {
		fmt.Fprintf(out, "Result = %d\n", res)
	}
	return nil
}

func main() {
	err := calc(os.Stdout, os.Stdin)
	if err != nil {
		panic(err)
	}
}
