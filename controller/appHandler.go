package controller

import (
	"encoding/json"
	"errors"
	config "github.com/config"
	model "github.com/model/DAO"
	util "github.com/util"
	"html/template"
	"log"
	"net/http"
	"os"
	"strconv"
	"strings"
)

type AppHandler struct {
	appConfig config.AppConfig
	h         func(app config.AppConfig, w http.ResponseWriter, r *http.Request) (int, error)
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

func ServerSetup(appConfig config.AppConfig, port string) {
	log.Println("start setup server:")

	log.Println("setup home page redirect (/)")

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		http.Redirect(w, r, "/login", http.StatusTemporaryRedirect)
	})

	log.Println("setup library handler (/js/)")

	http.HandleFunc("/js/", func(w http.ResponseWriter, r *http.Request) {
		http.ServeFile(w, r, "tmpl/"+r.URL.Path[1:])
	})

	log.Println("setup send path (/send)")

	http.HandleFunc("/send", AppHandler{appConfig, sendMessage}.ServeHTTP)

	log.Println("setup sync path (/sync)")

	http.HandleFunc("/sync", AppHandler{appConfig, syncMessages}.ServeHTTP)

	log.Println("setup login path (/login)")

	http.HandleFunc("/login", AppHandler{appConfig, login}.ServeHTTP)

	log.Println("setup get user by name path (/getuseridbyname)")

	http.HandleFunc("/getuseridbyname", AppHandler{appConfig, getUserIdbyName}.ServeHTTP)

	log.Println("setup register path (/register)")

	http.HandleFunc("/register", AppHandler{appConfig, register}.ServeHTTP)

	// setup the lisenting port
	err := http.ListenAndServe(":"+port, nil)
	if err != nil {
		log.Fatal("ListenAndServe err: ", err)
		os.Exit(1)
	}
}

/*
 * parameter needed:
 * userid (receive user)
 */
func syncMessages(app config.AppConfig, w http.ResponseWriter, r *http.Request) (int, error) {
	log.Println("receive Message")
	if r.Method == "GET" {
		t, _ := template.ParseFiles("tmpl/receive.gtpl")
		t.Execute(w, nil)
		return http.StatusAccepted, nil
	} else {
		r.ParseForm()

		if model.ValidateToken(r.FormValue("token")) {
			// print the info to the server console
			log.Println(r.Form)
			log.Println("path", r.URL.Path)
			log.Println("scheme", r.URL.Scheme)
			log.Println(r.Form["url_long"])

			log.Println("looking for message for userid: ", r.Form["id"])

			// accroding to the user id find his chan
			userId64, err := strconv.ParseInt(r.FormValue("id"), 0, 0)
			if err != nil {
				log.Println("parse user id err", err)
			} else {
				c := findChan(&app, int(userId64))
				log.Println("channel length", len(c))
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
func sendMessage(app config.AppConfig, w http.ResponseWriter, r *http.Request) (int, error) {
	//get url param, for POST method, get the body
	if r.Method == "GET" {
		t, _ := template.ParseFiles("tmpl/send.gtpl")
		t.Execute(w, nil)
		return http.StatusAccepted, nil
	} else {
		log.Println("method send:", r.Method) // get the http method
		r.ParseForm()

		// everytime need to ensure the user is an available user by checking
		// the token of the user
		if model.ValidateToken(strings.Join(r.Form["token"], "")) {
			log.Println("user loged in: ", r.Form["username"])

			// print the info to the server console
			log.Println(r.Form)
			log.Println("path", r.URL.Path)
			log.Println("scheme", r.URL.Scheme)
			log.Println(r.Form["url_long"])

			log.Println("username:", r.Form["username"])
			log.Println("user id:", r.Form["id"])
			log.Println("userid sendTo:", r.Form["sendToId"])
			log.Println("content:", r.Form["content"])

			for k, v := range r.Form {
				log.Println("key:", k)
				log.Println("val:", strings.Join(v, ""))
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
			c := findChan(&app, int(sendToId64))
			go model.Addmsg(msg, c)

			log.Println("message sent", len(findChan(&app, 0)))
			toJsonResponse("message received", w)
			return http.StatusAccepted, nil
		} else {
			toJsonResponse(util.ErrMessage{"not valid token, please log in"}, w)
			return http.StatusUnauthorized, errors.New("not valid token")
		}
	}
}

func login(app config.AppConfig, w http.ResponseWriter, r *http.Request) (int, error) {
	log.Println("method login:", r.Method) // get the http method
	if r.Method == "GET" {
		t, _ := template.ParseFiles("tmpl/login.gtpl")
		t.Execute(w, nil)
		return http.StatusAccepted, nil
	} else {
		r.ParseForm()
		//print out the form info
		log.Println("username:", r.Form["username"])
		// construct return json str
		sessionid, ok := model.PerformLogin(w, r)
		if ok {
			var ur model.User
			ur.Name = strings.Join(r.Form["username"], "")
			uid, err := model.GetUserIdByName(ur.Name)
			if err == nil {
				ur.Id = uid
				ur.SessionID = sessionid
				toJsonResponse(ur, w)
				return http.StatusAccepted, nil
			} else {
				toJsonResponse("please register", w)
				return http.StatusBadRequest, nil
			}
		} else {
			toJsonResponse(util.ErrMessage{"wrong user name or password"}, w)
			return http.StatusUnauthorized, errors.New("not valid token")
		}
	}
}

func getUserIdbyName(app config.AppConfig, w http.ResponseWriter, r *http.Request) (int, error) {
	log.Println("method get user id by name:", r.Method) // get the http method
	r.ParseForm()
	//print out the form info
	log.Println("username:", r.Form["username"])
	// construct return json str
	uname := strings.Join(r.Form["username"], "")
	userDao := model.NewUserDAO()
	user := userDao.GetUserByName(uname)
	toJsonResponse(user, w)
	return http.StatusAccepted, nil
}

func register(app config.AppConfig, w http.ResponseWriter, r *http.Request) (int, error) {
	log.Println("method get user id by name:", r.Method) // get the http method
	if r.Method == "GET" {
		t, _ := template.ParseFiles("tmpl/register.gtpl")
		t.Execute(w, nil)
		return http.StatusAccepted, nil
	} else {
		r.ParseForm()
		//print out the form info
		log.Println("username:", r.Form["username"])
		// construct return json str
		uname := strings.Join(r.Form["username"], "")
		pwd := strings.Join(r.Form["password"], "")
		userDao := model.NewUserDAO()
		user := model.User{0, uname, pwd, ""} // use a default value of 0 as id
		userDao.Save(user)
		toJsonResponse(user, w)
		return http.StatusAccepted, nil
	}
}

// find the chanel for the user by the userid, if the chanel is not
// exist, will create one and add to the map
func findChan(app *config.AppConfig, uid int) chan model.Message {
	if app.GetMsgcs()[uid] != nil {
		return app.GetMsgcs()[uid]
	}
	// create a new channel for the user
	msgc := make(chan model.Message, config.MAX_CHAN)
	app.SetMsgcs(uid, msgc)
	return app.GetMsgcs()[uid]
}

func toJsonResponse(v interface{}, w http.ResponseWriter) {
	js, err := json.Marshal(v)
	if err != nil {
		log.Fatal("json parse err", err)
	}
	w.Header().Set("Content-Type", "application/json")
	w.Write(js)
}
