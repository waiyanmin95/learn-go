package main

import (
	"html/template"
	"log"
	"net/http"

	"github.com/google/uuid"
	"golang.org/x/crypto/bcrypt"
)

type user struct {
	UserName string
	Password []byte
	First    string
	Last     string
}

var tpl *template.Template

// Map of unique users using email as a key
var dbUsers = map[string]user{}

// Map of active sessions with session ID as the key
var dbSessions = map[string]string{}

func init() {
	tpl = template.Must(template.ParseGlob("templates/*"))

	// Define an initial user to test the registration and login process
	bs, _ := bcrypt.GenerateFromPassword([]byte("password"), bcrypt.MinCost)
	dbUsers["wai@amity.co"] = user{
		UserName: "wai@amity.co",
		Password: bs,
		First:    "Wai",
		Last:     "Min",
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
	tpl.ExecuteTemplate(w, "index.gohtml", u)
}

// Protected bar page handler accessible only if logged in
func bar(w http.ResponseWriter, req *http.Request) {
	u := getUser(w, req)
	if !alreadyLoggedIn(req) {
		http.Redirect(w, req, "/", http.StatusSeeOther)
		return
	}

	tpl.ExecuteTemplate(w, "bar.gohtml", u)
}

// Signup page handler
func signup(w http.ResponseWriter, req *http.Request) {
	// Redirect to index page if user is already logged in
	if alreadyLoggedIn(req) {
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

		http.SetCookie(w, c)

		// Associate session ID with the user's email
		dbSessions[c.Value] = un

		// Hash password and store user in the database
		bs, err := bcrypt.GenerateFromPassword([]byte(p), bcrypt.MinCost)
		if err != nil {
			http.Error(w, "Internal Server Error", http.StatusInternalServerError)
			return
		}

		u = user{un, bs, f, l}
		dbUsers[un] = u

		// Redirect to index page
		http.Redirect(w, req, "/", http.StatusSeeOther)
		return
	}
	tpl.ExecuteTemplate(w, "signup.gohtml", u)
}

// Login page handler
func login(w http.ResponseWriter, req *http.Request) {
	// Redirect to index page if user is already logged in
	if alreadyLoggedIn(req) {
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

		http.SetCookie(w, c)
		dbSessions[c.Value] = un
		http.Redirect(w, req, "/", http.StatusSeeOther)
		return
	}

	tpl.ExecuteTemplate(w, "login.gohtml", nil)
}

func logout(w http.ResponseWriter, req *http.Request) {
	if !alreadyLoggedIn(req) {
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

	http.Redirect(w, req, "/", http.StatusSeeOther)
}
