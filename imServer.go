package main

import (
	"crypto/rand"
	"encoding/base64"
	"encoding/json"
	"errors"
	"fmt"
	"html/template"
	"log"
	"net/http"
	"os"
	"strconv"
	"strings"
)

type User struct {
	Id    int
	Name  string
	Token string
}

type AppConfig struct {
	msgc chan Message
}

type Message struct {
	content       string
	userIdSend    int
	userIdReceive int
}

type AppHandler struct {
	appConfig AppConfig
	h         func(app AppConfig, w http.ResponseWriter, r *http.Request) (int, error)
}

func (app AppConfig) findChan(uid int) chan Message {
	// todo this is a adapter sub
	return app.msgc
}

func addmsg(a Message, c chan Message) {
	// todo this is a adapter sub
	c <- a
}

func validateToken(token string) bool {
	return true
}

func (app AppHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	status, err := app.h(app.appConfig, w, r)
	if err != nil {
		log.Printf("HTTP err", status, err)
		switch status {
		case http.StatusNotFound:
			http.NotFound(w, r)
		case http.StatusInternalServerError:
			http.Error(w, http.StatusText(status), status)
		default:
			http.Error(w, http.StatusText(status), status)
		}
	}
}

func receiveMessage(app AppConfig, w http.ResponseWriter, r *http.Request) (int, error) {
	fmt.Fprintf(w, "receive Message")
	return http.StatusAccepted, nil
}

func sendMessage(app AppConfig, w http.ResponseWriter, r *http.Request) (int, error) {
	//get url param, for POST method, get the body
	r.ParseForm()

	//print out the form info
	if validateToken(strings.Join(r.Form["token"], "")) {
		fmt.Println("user loged in: ", r.Form["username"])

		// print the info to the server console
		fmt.Println(r.Form)
		fmt.Println("path", r.URL.Path)
		fmt.Println("scheme", r.URL.Scheme)
		fmt.Println(r.Form["url_long"])

		fmt.Println("username:", r.Form["username"])
		fmt.Println("user id:", r.Form["id"])
		fmt.Println("userid sendTo:", r.Form["sendToId"])
		fmt.Println("content:", r.Form["content"])

		for k, v := range r.Form {
			fmt.Println("key:", k)
			fmt.Println("val:", strings.Join(v, ""))
		}

		// construct a new message
		userId64, err := strconv.ParseInt(r.FormValue("id"), 0, 0)

		if err != nil {
			log.Println("parse user id err", err)
		}

		sendToId64, err := strconv.ParseInt(r.FormValue("sendToId"), 0, 0)
		if err != nil {
			log.Println("parse sendTo id err", err)
		}

		msg := Message{r.FormValue("content"), int(userId64), int(sendToId64)}

		// get the chan based on the sendToId
		// add the msg to the channel
		addmsg(msg, app.findChan(int(sendToId64)))

		fmt.Fprintf(w, "message sent")
		return http.StatusAccepted, nil
	} else {
		toJsonResponse("not valid token, please log in", w)
		return http.StatusUnauthorized, errors.New("not valid token")
	}
}

func login(app AppConfig, w http.ResponseWriter, r *http.Request) (int, error) {
	fmt.Println("method:", r.Method) // get the http method
	if r.Method == "GET" {
		t, _ := template.ParseFiles("login.gtpl")
		t.Execute(w, nil)
		return http.StatusAccepted, nil
	} else {
		r.ParseForm()
		//print out the form info
		fmt.Println("username:", r.Form["username"])

		// construct return json str
		token, ok := performLogin(strings.Join(r.Form["username"], ""), strings.Join(r.Form["password"], ""))
		if ok {
			var ur User
			ur.Id = 0
			ur.Name = strings.Join(r.Form["username"], "")
			ur.Token = token
			toJsonResponse(ur, w)
			return http.StatusAccepted, nil
		} else {
			toJsonResponse("wrong user name or password", w)
			return http.StatusUnauthorized, errors.New("not valid token")
		}

	}
}

func toJsonResponse(v interface{}, w http.ResponseWriter) {
	js, err := json.Marshal(v)
	if err != nil {
		log.Fatal("json parse err", err)
	}
	w.Header().Set("Content-Type", "application/json")
	w.Write(js)
}

func performLogin(uname string, pwd string) (string, bool) {
	if pwd == "password" {
		token, err := tokenGenerator()
		if err != nil {
			fmt.Println(err)
			return token, false
		} else {
			return token, true
		}
	} else {
		return "wrong password", false
	}

}

func tokenGenerator() (string, error) {
	// change the length of the generated random string here
	size := 32

	rb := make([]byte, size)
	_, err := rand.Read(rb)
	if err != nil {
		fmt.Println(err)
		return "error", errors.New("cannot generate token for user")
	}

	rs := base64.URLEncoding.EncodeToString(rb)

	return rs, nil
}

func serverSetup(port string) {
	fmt.Println("start setup server:")

	fmt.Println("setup send path (/send)")

	msgc := make(chan Message)
	appCofig := AppConfig{msgc}

	http.HandleFunc("/send", AppHandler{appCofig, sendMessage}.ServeHTTP)

	fmt.Println("setup receive path")

	http.HandleFunc("/receive", AppHandler{appCofig, receiveMessage}.ServeHTTP)

	fmt.Println("setup login path")

	http.HandleFunc("/", AppHandler{appCofig, login}.ServeHTTP)

	// setup the lisenting port
	err := http.ListenAndServe(":"+port, nil)
	if err != nil {
		log.Fatal("ListenAndServe err: ", err)
		os.Exit(1)
	}

}

func main() {
	serverSetup("9000")
}
