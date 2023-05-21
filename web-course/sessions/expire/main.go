package main

import (
	"html/template"
	"log"
	"net/http"
	"time"

	"github.com/google/uuid"
	"golang.org/x/crypto/bcrypt"
)

type user struct {
	UserName string
	Password []byte
	First    string
	Last     string
	Role     string
}

type session struct {
	un           string
	lastActivity time.Time
}

var tpl *template.Template

// Map of unique users using email as a key
var dbUsers = map[string]user{}

// Map of active sessions with session ID as the key
var dbSessions = map[string]session{}

var dbSessionsCleaned time.Time

const sessionLength int = 30

func init() {
	tpl = template.Must(template.ParseGlob("templates/*"))

	// Define an initial user to test the registration and login process
	bs, _ := bcrypt.GenerateFromPassword([]byte("password"), bcrypt.MinCost)
	dbUsers["wai@amity.co"] = user{
		UserName: "wai@amity.co",
		Password: bs,
		First:    "Wai",
		Last:     "Min",
		Role:     "007",
	}
}

func main() {
	http.HandleFunc("/", index)
	http.HandleFunc("/bar", bar)
	http.HandleFunc("/signup", signup)
	http.HandleFunc("/login", login)
	http.HandleFunc("/logout", logout)
	http.Handle("/favicon.ico", http.NotFoundHandler())
	log.Fatal(http.ListenAndServe("localhost:8080", nil))
}

// Index page handler
func index(w http.ResponseWriter, req *http.Request) {
	u := getUser(w, req)
	showSessions() // for demonstration purposes
	tpl.ExecuteTemplate(w, "index.gohtml", u)
}

// Protected bar page handler accessible only if logged in
func bar(w http.ResponseWriter, req *http.Request) {
	u := getUser(w, req)
	if !alreadyLoggedIn(w, req) {
		http.Redirect(w, req, "/", http.StatusSeeOther)
		return
	}

	if u.Role != "007" {
		http.Error(w, "You must be 007 to enter the bar.", http.StatusForbidden)
		return
	}

	showSessions() // for demonstration purposes
	tpl.ExecuteTemplate(w, "bar.gohtml", u)
}

// Signup page handler
func signup(w http.ResponseWriter, req *http.Request) {
	// Redirect to index page if user is already logged in
	if alreadyLoggedIn(w, req) {
		http.Redirect(w, req, "/", http.StatusSeeOther)
		return
	}

	var u user
	// process form submission
	if req.Method == http.MethodPost {
		// Read form values
		un := req.FormValue("username")
		p := req.FormValue("password")
		f := req.FormValue("firstname")
		l := req.FormValue("lastname")
		r := req.FormValue("role")

		// Check if username is already taken
		if _, ok := dbUsers[un]; ok {
			http.Error(w, "Username Already Taken", http.StatusForbidden)
			return
		}

		// Generate a session ID
		sID, err := uuid.NewUUID()
		if err != nil {
			http.Error(w, "Internal Server Error", http.StatusInternalServerError)
			return
		}

		// Set session cookie
		c := &http.Cookie{
			Name:  "session",
			Value: sID.String(),
		}

		c.MaxAge = sessionLength

		http.SetCookie(w, c)

		// Associate session ID with the user's email
		dbSessions[c.Value] = session{un, time.Now()}

		// Hash password and store user in the database
		bs, err := bcrypt.GenerateFromPassword([]byte(p), bcrypt.MinCost)
		if err != nil {
			http.Error(w, "Internal Server Error", http.StatusInternalServerError)
			return
		}

		u = user{un, bs, f, l, r}
		dbUsers[un] = u

		// Redirect to index page
		http.Redirect(w, req, "/", http.StatusSeeOther)
		return
	}
	showSessions() // for demonstration purposes
	tpl.ExecuteTemplate(w, "signup.gohtml", u)
}

// Login page handler
func login(w http.ResponseWriter, req *http.Request) {
	// Redirect to index page if user is already logged in
	if alreadyLoggedIn(w, req) {
		http.Redirect(w, req, "/", http.StatusSeeOther)
		return
	}

	// Process form submission
	if req.Method == http.MethodPost {
		// Read form values
		un := req.FormValue("username")
		p := req.FormValue("password")

		// Retrieve user from the database
		u, ok := dbUsers[un]
		if !ok {
			http.Error(w, "Username and/or Password do not match", http.StatusForbidden)
			return
		}

		// Compare password hash with the one stored in the database
		err := bcrypt.CompareHashAndPassword(u.Password, []byte(p))
		if err != nil {
			http.Error(w, "Username and/or Password do not match", http.StatusForbidden)
			return
		}

		// Generate a session ID
		sID, err := uuid.NewUUID()
		if err != nil {
			http.Error(w, "Internal Server Error", http.StatusInternalServerError)
			return
		}

		// Set session cookie
		c := &http.Cookie{
			Name:  "session",
			Value: sID.String(),
		}

		c.MaxAge = sessionLength

		http.SetCookie(w, c)
		dbSessions[c.Value] = session{un, time.Now()}
		http.Redirect(w, req, "/", http.StatusSeeOther)
		return
	}
	showSessions() // for demonstration purposes
	tpl.ExecuteTemplate(w, "login.gohtml", nil)
}

func logout(w http.ResponseWriter, req *http.Request) {
	if !alreadyLoggedIn(w, req) {
		http.Redirect(w, req, "/", http.StatusSeeOther)
		return
	}

	c, err := req.Cookie("session")
	if err != nil {
		http.Error(w, "Cookie Not Found", http.StatusNotFound)
		return
	}

	// delete session
	delete(dbSessions, c.Value)
	// Expire the cookie
	c = &http.Cookie{
		Name:   "session",
		Value:  "",
		MaxAge: -1,
	}
	http.SetCookie(w, c)

	// clean up dbSessions
	if time.Now().Sub(dbSessionsCleaned) > (time.Second * 30) {
		go cleanSessions()
	}

	http.Redirect(w, req, "/", http.StatusSeeOther)
}
