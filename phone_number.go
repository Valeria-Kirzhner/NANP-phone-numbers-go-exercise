package phonenumber

import (
	"errors"
	"fmt"
	"regexp"
	"unicode"
	"unicode/utf8"
)

func Number(test string) (string, error){
	var  err error
	var newNumber string

	for _, char := range test {
		c := fmt.Sprintf("%c", char)

		if !unicode.IsDigit(char) {

			validCharacter := is_string_valid(c)
			if validCharacter == false {
				return test, errors.New("not a valid char")
			} else {
				continue
			}
		}
		newNumber = newNumber + c
	}
		
	len := len(newNumber)

	if (len == 11 ){
		newNumber, err = removeCountryCode(newNumber)
		if err != nil {
			return newNumber, err
		}	
	} else if (len < 10 || len > 10) {
		return newNumber, errors.New("number length must not be less or bigger then 10")
	} 
	fmt.Println("newNumber = ", newNumber)
	return newNumber, nil

}

func is_string_valid(char string) bool {
	regex := regexp.MustCompile("^[.()\\+\\- ]")
	res := regex.MatchString(char)
	return res
}
func removeCountryCode(newNumber string) (string,error) {
	firstCharacter:= newNumber[0:1]
	if ( firstCharacter == "1" ){
		_, i := utf8.DecodeRuneInString(newNumber)
		return newNumber[i:], nil
	} else {
		return newNumber, errors.New("Country code can be '1' only")
	}
}
// func check_code_area_validity() {

// }