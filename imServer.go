package main

import (
	"crypto/rand"
	"encoding/base64"
	"encoding/json"
	"errors"
	"fmt"
	model "github.com/model"
	"html/template"
	"log"
	"net/http"
	"os"
	"strconv"
	"strings"
)

const (
	MAX_CHAN = 10000
)

type AppConfig struct {
	msgcs map[int]chan model.Message
}

type ErrMessage struct {
	Err string
}

type AppHandler struct {
	appConfig AppConfig
	h         func(app AppConfig, w http.ResponseWriter, r *http.Request) (int, error)
}

func (app *AppConfig) findChan(uid int) chan model.Message {
	// todo this is a adapter sub
	if app.msgcs[uid] != nil {
		return app.msgcs[uid]
	}
	// create a new channel for the user
	msgc := make(chan model.Message, MAX_CHAN)
	app.msgcs[uid] = msgc
	return app.msgcs[uid]
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

/*
 * parameter needed:
 * userid (receive user)
 */
func sync(app AppConfig, w http.ResponseWriter, r *http.Request) (int, error) {
	fmt.Println("receive Message")
	if r.Method == "GET" {
		t, _ := template.ParseFiles("tmpl/receive.gtpl")
		t.Execute(w, nil)
		return http.StatusAccepted, nil
	} else {
		r.ParseForm()

		if validateToken(r.FormValue("token")) {
			// print the info to the server console
			fmt.Println(r.Form)
			fmt.Println("path", r.URL.Path)
			fmt.Println("scheme", r.URL.Scheme)
			fmt.Println(r.Form["url_long"])

			fmt.Println("looking for message for userid: ", r.Form["id"])

			// accroding to the user id find his chan
			userId64, err := strconv.ParseInt(r.FormValue("id"), 0, 0)
			if err != nil {
				log.Println("parse user id err", err)
			} else {
				c := app.findChan(int(userId64))
				fmt.Println("channel length", len(c))
				if c != nil && len(c) > 0 {
					var msgs []model.Message
					count := len(c)
					for i := 0; i < count; i++ {
						msgs = append(msgs, <-c)
					}

					toJsonResponse(msgs, w)
				} else {
					toJsonResponse("no message", w)
				}
			}
		}
		return http.StatusAccepted, nil
	}
}

/* parameters need:
 * username(send user)
 * id (send user)
 * sendTo userId
 * content (message content)
 */
func sendMessage(app AppConfig, w http.ResponseWriter, r *http.Request) (int, error) {
	//get url param, for POST method, get the body
	if r.Method == "GET" {
		t, _ := template.ParseFiles("tmpl/send.gtpl")
		t.Execute(w, nil)
		return http.StatusAccepted, nil
	} else {
		fmt.Println("method send:", r.Method) // get the http method

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

			msg := model.Message{r.FormValue("content"), int(userId64), int(sendToId64)}

			// get the chan based on the sendToId
			// add the msg to the channel
			c := app.findChan(int(sendToId64))
			go model.Addmsg(msg, c)

			fmt.Println("message sent", len(app.findChan(0)))
			toJsonResponse("message received", w)
			return http.StatusAccepted, nil
		} else {
			toJsonResponse(ErrMessage{"not valid token, please log in"}, w)
			return http.StatusUnauthorized, errors.New("not valid token")
		}
	}
}

func login(app AppConfig, w http.ResponseWriter, r *http.Request) (int, error) {
	fmt.Println("method login:", r.Method) // get the http method
	if r.Method == "GET" {
		t, _ := template.ParseFiles("tmpl/login.gtpl")
		t.Execute(w, nil)
		return http.StatusAccepted, nil
	} else {
		r.ParseForm()
		//print out the form info
		fmt.Println("username:", r.Form["username"])
		// construct return json str
		token, ok := performLogin(strings.Join(r.Form["username"], ""), strings.Join(r.Form["password"], ""))
		if ok {
			var ur model.User
			ur.Name = strings.Join(r.Form["username"], "")
			ur.Token = token
			uid, err := model.GetUserIdByName(ur.Name)
			if err == nil {
				ur.Id = uid
				toJsonResponse(ur, w)
				return http.StatusAccepted, nil
			} else {
				toJsonResponse("please register", w)
				return http.StatusBadRequest, nil
			}
		} else {
			toJsonResponse(ErrMessage{"wrong user name or password"}, w)
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

func serverSetup(appConfig AppConfig, port string) {
	fmt.Println("start setup server:")
	http.HandleFunc("/js/", func(w http.ResponseWriter, r *http.Request) {
		http.ServeFile(w, r, "tmpl/"+r.URL.Path[1:])
	})

	fmt.Println("setup send path (/send)")

	http.HandleFunc("/send", AppHandler{appConfig, sendMessage}.ServeHTTP)

	fmt.Println("setup sync path")

	http.HandleFunc("/sync", AppHandler{appConfig, sync}.ServeHTTP)

	fmt.Println("setup login path")

	http.HandleFunc("/login", AppHandler{appConfig, login}.ServeHTTP)

	// setup the lisenting port
	err := http.ListenAndServe(":"+port, nil)
	if err != nil {
		log.Fatal("ListenAndServe err: ", err)
		os.Exit(1)
	}

}

func main() {
	msgcs := make(map[int]chan model.Message)
	appConfig := AppConfig{msgcs}
	serverSetup(appConfig, "9000")
}
