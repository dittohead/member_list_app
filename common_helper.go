package main

import (
	"net/mail"
	"regexp"
	"time"
)

func getRegDate() string{
currentTime := time.Now()
return currentTime.Format("02.01.2006")
}

func emailValidator(email string) bool{
	_, err := mail.ParseAddress(email)
	count := 0
	for _, item := range members{
		if email==item.Email{
			count+=1
		}
	}
	if count != 0 || err != nil{
		return false
	}else{
		return true}
}

func nameValidator(name string) bool{
	isAlpha := regexp.MustCompile(`^[A-Za-z\s(.)]+$`).MatchString

	if !isAlpha(name){
		return false
	} else {
		return true}
}