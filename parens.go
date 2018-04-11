package main

import "fmt"
import "errors"

func closingParen(sentence string, start int) (int, error) {
	if string(sentence[start]) != "(" {
		return -1, errors.New("not paren")
	}

	opens := 0
	for i := start + 1; i < len(sentence); i++ {
		if opens == 0 && string(sentence[i]) == ")" {
			return i, nil
		} else if string(sentence[i]) == "(" {
			opens = opens + 1
		} else if opens != 0 && string(sentence[i]) == ")" {
			opens = opens - 1
		}
	}

	return -1, errors.New("paren is unclosed")
}

func main() {
	end, err := closingParen("tr(a(())(()()))in)", 2)
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Printf("Closing paren at %d\n", end)
	}
}
