package main

import (
	"fmt"
	"math/rand"
	"net/http"
	"strconv"
	"text/template"
	"time"
	uuid "github.com/satori/go.uuid"
)

func init() {
	tpl = template.Must(template.ParseGlob("templates/*html"))

}

func main() {
	http.HandleFunc("/",index)
	http.HandleFunc("/gamePlay/",gamePlay)
	http.Handle("/favicon.ico",http.NotFoundHandler())
	http.ListenAndServe(":8080",nil)
}

func GetSet(num1 int, num2 int) set {
	guess := num1
	bulls := 0
	cows := 0
	var a, b [4]int
	for i :=0; i < 4; i++ {
		a[3-i] = num1 % 10
		b[3-i] = num2 % 10
		num1 = num1 / 10
		num2 = num2 / 10
	}
	if a[0] == b[0]{
		bulls++	
	}else if a[0] == b[1] || a[0] == b[2] || a[0] == b[3] {
		cows++
	}
	if a[1] == b[1]{
		bulls++	
	}else if a[1] == b[0] || a[1] == b[2] || a[1] == b[3] {
		cows++
	}
	if a[2] == b[2]{
		bulls++	
	}else if a[2] == b[1] || a[2] == b[0] || a[2] == b[3] {
		cows++
	}
	if a[3] == b[3]{
		bulls++	
	}else if a[3] == b[1] || a[3] == b[2] || a[3] == b[0] {
		cows++
	}
	return set{guess, bulls, cows}

}

func generateSecretNumber() int {
	rand.Seed(time.Now().UTC().UnixNano())
	return numbers[rand.Intn(4536)]
}

func index(res http.ResponseWriter, req *http.Request) {
	tpl.ExecuteTemplate(res, "index.html", nil)
}

func gamePlay(res http.ResponseWriter, req *http.Request) {
	c, err := req.Cookie("gamePlay")
	if err != nil {
		sID, _ := uuid.NewV4()
		c = &http.Cookie{
			Name: "gamePlay",
			Value: sID.String(),
		}
	}
	http.SetCookie(res, c)

	if _,ok := data[c.Value]; !ok {
		data[c.Value] = game{generateSecretNumber(), []set{}, false}
	}

	if req.Method == "POST" {
		var newgame game
		if req.FormValue("play_again") == "Play Again" || req.FormValue("restart") == "Restart"{
			newgame= game{generateSecretNumber(), []set{}, false}
		}else {
			temp := data[c.Value]
			guess,_ := strconv.Atoi(req.FormValue("guess"))
			newSet := GetSet(guess,temp.secretNumber)
			if newSet.Bulls == 4 {
				newgame = game{temp.secretNumber, append(temp.Sets, newSet), true}
			} else {
				newgame = game{temp.secretNumber, append(temp.Sets, newSet), false}
			}
		}
		data[c.Value] = newgame
	}
	fmt.Println(data[c.Value])
	tpl.ExecuteTemplate(res,"gamePlay.html", data[c.Value])
}