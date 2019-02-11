package auth

import (
	"errors"
	"net/http"

	"github.com/gorilla/schema"
	"github.com/gorilla/securecookie"
)

const (
	authCookieName = "idc"
)

var hashKey = securecookie.GenerateRandomKey(64)
var blockKey = securecookie.GenerateRandomKey(32)

var secureCookie = securecookie.New(hashKey, blockKey)

var decoder = schema.NewDecoder()
var encoder = schema.NewEncoder()

// SetCookieForUser crée une cookie de securiter avec les informations d'un account
func SetCookieForUser(w http.ResponseWriter, user *User) error {
	var err error
	if encoded, err := secureCookie.Encode(authCookieName, user); err == nil {
		cookie := &http.Cookie{
			Name:  authCookieName,
			Value: encoded,
			Path:  "/",
		}
		http.SetCookie(w, cookie)
	}
	return err
}

// DecodeCookieForUser decode le cookie
func DecodeCookieForUser(r *http.Request) (*User, error) {
	var e error
	if cookie, e := r.Cookie(authCookieName); e == nil {
		print("lol")
		u := &User{}
		return u, secureCookie.Decode(authCookieName, cookie.Value, &u)
	}
	return nil, e
}

// ValidUserCookie valide que le cookie de la request contient
// l'information d'un user valide
func ValidUserCookie(r *http.Request) (*User, error) {
	user, err := DecodeCookieForUser(r)
	if err != nil {
		return nil, err
	}
	if user == nil {
		return nil, errors.New("No cookie")
	}

	irepo := GetIAuth()
	defer irepo.Close()
	if err = irepo.IsValidUser(user.ID, user.SaltedPW); err != nil {
		return nil, err
	}
	return user, nil
}
