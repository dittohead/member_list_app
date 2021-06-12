package main

import (
	"encoding/json"
	"log"
	"net/http"
)

type Member struct {
	UserNum  int    `json:"user_num"`
	UserName string `json:"user_name"`
	Email    string `json:"email"`
	RegDate  string `json:"reg_date"`
}

var members []Member
var userCount = 1

func getMemberHandler(w http.ResponseWriter, r *http.Request) {
	membersListBytes, err := json.Marshal(members)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		log.Fatal("Error: %v", err)
		return
	}
	w.Write(membersListBytes)
	log.Println("INFO: bytes sent")
}

func createMemberHandler(w http.ResponseWriter, r *http.Request) {
	member := Member{}
	err := r.ParseForm()
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		log.Fatal("Error: %v", err)
		return
	}

	email := r.Form.Get("email")
	userName := r.Form.Get("name")

	emailStatus := emailValidator(email)
	userNameStatus := nameValidator(userName)

	if emailStatus && userNameStatus {
		member.UserNum = userCount
		member.UserName = userName
		member.Email = email
		member.RegDate = getRegDate()

		members = append(members, member)
		userCount+=1
		log.Println("Users in storage:", userCount-1 )
	} else {
		if !emailStatus {
			log.Println("ERROR: bad email:", email)
		}

		if !userNameStatus {
			log.Println("ERROR: bad name:", userName)
		}
	}

	http.Redirect(w, r, "/assets/", http.StatusFound)
}
