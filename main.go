package main

import (
	"crypto/rand"
	"errors"
	"fmt"
	"math/big"
)

//define a struct for determining password criteria
type PasswordCriteria struct{
	Upper bool
	Lower bool
	Numbers bool
	Special bool
	MinmumLength int
	MaximumLength int
}
/*
Generates a random password
@params criteria PasswordCriteria
*/
func GenerateRandomPassword(criteria PasswordCriteria)(string,error){
	var (
		password []byte
		characterSet string
	)
	//if upper is true, set the character set
	if criteria.Upper{
		characterSet += "ABCDEFGHIJKLMNOPQRSTUVWXYZ"
	}
	//if lower is true, set the character set
	if criteria.Lower{
		characterSet+="abcdefghijklmnopqrstuvwxyz" 
	}
	//if number is true, set the character set
	if criteria.Numbers{
		characterSet+="0123456789"
	}
	//if special character is true, set the character set
	if criteria.Special{
		characterSet += "!@#$%^&*()-_=+[]{}|;:,.<>?/"
	}
	//when character string is zero, return empty string
	if len(characterSet) == 0 {
		return "", errors.New("no character string")
	}
	passwordLength := criteria.MinmumLength+ RandomInt(criteria.MaximumLength-criteria.MinmumLength+1)
	fmt.Println("Random int:",RandomInt(criteria.MaximumLength-criteria.MinmumLength+1))
	fmt.Println("Password length:",passwordLength)
	for i:=0; i<passwordLength; i++{
		randomChar,err:=RandomChar(characterSet)
		if err != nil {
			return "",err
		}
		password =append(password,randomChar)
	}
	return string(password),nil
}
/*
Generates a random number between 0 and the (maximumLeght-minmumLenth)-1
@params defference between maximum and minmum legth
it return the generated int
*/
func RandomInt(maximum int) (int) {
	n,_:=rand.Int(rand.Reader, big.NewInt(int64(maximum)))
	
	return int(n.Int64())
}

/*
generates random char
@params characterSet string
*/
func RandomChar(characterSet string) (byte, error) {
	n,err := rand.Int(rand.Reader,big.NewInt(int64(len(characterSet))))
	if err != nil{
		return 0,err
	}
	return characterSet[n.Int64()],nil
}

//main function
func main() {
	fmt.Println("generating random password...")
	criteria :=PasswordCriteria{
		Upper: true,
		Lower: true,
		Numbers: true,
		Special: true,
		MinmumLength: 6,
		MaximumLength: 10,
	}
	
	password,err:=GenerateRandomPassword(criteria)
	if err != nil{
		fmt.Println("error generating random password:",err)
	}
	fmt.Println("generated random password:",password)
	fmt.Println("password length:",len(password))

}



