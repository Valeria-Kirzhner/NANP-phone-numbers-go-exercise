package phonenumber

import (
	"errors"
	"fmt"
	"regexp"
	"strconv"
	"unicode/utf8"
)
var  err error

func Number(input string) (string, error){

	cleanedNumber, err := cleanInputNumber(input)
	if err != nil {
		return cleanedNumber, err
	}
	cleanedNumber, err = removeCountryCode(cleanedNumber)
	if err != nil {
		return cleanedNumber, err
	}
	err = checkNNumber(cleanedNumber)
	if err != nil { 
		return cleanedNumber, err
	}	

	return cleanedNumber, nil
}

func cleanInputNumber(input string) (string, error) {
	var char rune
	var digitChar bool
	var newNumber string = ""
	
	for _, char = range input {
		c := fmt.Sprintf("%c", char)

		digitChar, err = isCharisADigit(c, input)
		if err != nil {
			return input, err
		} 
		if !digitChar {
			continue
		} 
		newNumber = newNumber + c
	}
	fmt.Println("newNumber ", newNumber)
	return newNumber, nil
}

func isCharisADigit(char string, input string) (bool, error) {

	if _, err := strconv.Atoi(char); err != nil {

		characterValid := isStringiAllowed(char)

		if characterValid == false {
			return false, errors.New("not a valid char")
		} else if characterValid == true {
			return false, nil
		}
	} 
	return true, nil
}

func removeCountryCode(cleanedNumber string) (string, error) {
	fmt.Println("cleanedNumber is " ,cleanedNumber)
	cleanedNumLen := len(cleanedNumber)
	if (cleanedNumLen == 11 ){
		cleanedNumber, err = removeFirstChar(cleanedNumber)
		if err != nil {
			return cleanedNumber, err
		}	
	} else if (cleanedNumLen < 10 || cleanedNumLen > 10) {
		return cleanedNumber, errors.New("number length must not be less or bigger then 10")
	} 
	return cleanedNumber, nil
}

func isStringiAllowed(char string) bool {
	regex := regexp.MustCompile("^[.()\\+\\- ]")
	res := regex.MatchString(char)
	return res
}
func removeFirstChar(newNumber string) (string,error) {
	firstCharacter:= newNumber[0:1]
	if ( firstCharacter == "1" ){
		_, i := utf8.DecodeRuneInString(newNumber)
		return newNumber[i:], nil
	} else {
		return newNumber, errors.New("Country code can be '1' only")
	}
}
func checkNNumber(newNumber string) (error) {
	firstN:= newNumber[0:1]
	firthCharacter:= newNumber[3:4]
	if ( firstN == "1" || firstN == "0"  || firthCharacter == "1" || firthCharacter == "0"  ){
		return errors.New("N number can be 2-9 only")
	} 
	return nil
}
func AreaCode(test string) (string, error) {
	var codeArea string
	phoneNumber, err := Number(test)
	if err != nil { 
	 return test, errors.New("This input not pass Number function")
	}
	codeArea = codeArea + phoneNumber[0:3]
	return codeArea, nil
}

func Format(test string) (string, error) {

	phoneNumber, err := Number(test)
	if err != nil { 
	 return test, errors.New("This input not pass Number function")
	}
    formatedNumber := fmt.Sprintf("(%s) %s-%s", phoneNumber[0:3], phoneNumber[3:6], phoneNumber[6:10])
	return formatedNumber, nil
}