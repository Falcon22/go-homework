package main

import (
	"bufio"
	"bytes"
	"strings"
	"testing"
)

type TestCase struct {
	Name     string
	Input    string
	Expected string
	IsError  bool
}

func TestCalc(t *testing.T) {
	tests := []TestCase{
		TestCase{
			Name:     "TestAllOperations",
			Input:    "1 2 + =\n1 2 - =\n1 2 * =\n1 2 / =",
			Expected: "Result = 3\nResult = 1\nResult = 2\nResult = 2\n",
			IsError:  false,
		},
		TestCase{
			Name:     "TestFromExample1",
			Input:    "1 2 + =\n 3 4 *",
			Expected: "Result = 3\nResult = 12\n",
			IsError:  false,
		},
		TestCase{
			Name:     "TestFromExample2",
			Input:    "1 2 3 4 + * + =\n 1 2 + 3 4 + * =",
			Expected: "Result = 15\nResult = 21\n",
			IsError:  false,
		},
		TestCase{
			Name:     "BadInput",
			Input:    "bad input",
			Expected: "",
			IsError:  true,
		},
		TestCase{
			Name:     "EmptyStack+",
			Input:    "+",
			Expected: "",
			IsError:  true,
		},
		TestCase{
			Name:     "EmptyStack=",
			Input:    "=",
			Expected: "",
			IsError:  true,
		},
		TestCase{
			Name:     "EmptyStackWithResult",
			Input:    "1 2 + =\n 1 2 + + +",
			Expected: "Result = 3\n",
			IsError:  true,
		},
	}

	for _, test := range tests {
		t.Run(test.Name, func(t *testing.T) {
			in := bufio.NewReader(strings.NewReader(test.Input))
			out := new(bytes.Buffer)
			err := calc(out, in)
			if test.IsError {
				if err == nil {
					t.Fatalf("[%s] expected error", test.Name)
				}
			} else {
				if err != nil {
					t.Fatalf("[%s] error: %v", test.Name, err)
				}
			}
			if !strings.EqualFold(test.Expected, out.String()) {
				t.Fatalf("[%s] results not match\nGot : %#v\nWant: %#v", test.Name, out.String(), test.Expected)
			}
		})
	}
}
