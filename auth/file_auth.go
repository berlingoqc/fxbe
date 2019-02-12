package auth

import (
	"net/http"

	"github.com/berlingoqc/fxbe/utility"
)

type LoginModel struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

func GetLoginHandler(w http.ResponseWriter, r *http.Request) {
	iau := GetIAuth()
	defer iau.Close()
	l := &LoginModel{}
	if err := utility.QueryToStruct(r.URL.Query(), l); err != nil {
		utility.RespondWithError(w, http.StatusBadRequest, err)
		return
	}
	user, err := iau.LoginUser(l.Username, l.Password)
	if err != nil {
		utility.RespondWithError(w, http.StatusBadRequest, err)
		return
	}
	err = SetCookieForUser(w, user)
	if err != nil {
		utility.RespondWithError(w, http.StatusBadRequest, err)
	}

}
