package phonenumber

import (
	"errors"
	"fmt"
	"regexp"
	"unicode"
)

func Number(test string) (string, error){

	 var newNumber string

	for _, char := range test {
		c := fmt.Sprintf("%c", char)

		if !unicode.IsDigit(char) {

			validCharacter := is_string_valid(c)
			if validCharacter == false {
				return test, errors.New("not a valid char")
			}else {
				continue
			}
			// if no - return an error, if yes continue

		}		
		newNumber = newNumber + c
		fmt.Println("newNumber = ", newNumber)
	}
	return newNumber, nil

}

func is_string_valid (char string) bool {
	regex := regexp.MustCompile("^[.()\\+\\- ]")
	res := regex.MatchString(char)
	return res
}
