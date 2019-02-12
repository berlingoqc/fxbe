package api

import (
	"net/http"
	"os"
	"path/filepath"

	"github.com/berlingoqc/fxbe/auth"

	"github.com/berlingoqc/fxbe/files"

	"github.com/berlingoqc/fxbe/utility"

	"github.com/gorilla/mux"
)

// DeplacementModel is the model for request for deplacment (copy or move)
type DeplacementModel struct {
	Origin      string `json:"origin"`
	Destination string `json:"destination"`
	Recursive   bool   `json:"recursive"`
}

// RenameModel is the model for the rename request
type RenameModel struct {
	Origin  string `json:"origin"`
	NewName string `json:"newname"`
}

// DeleteModel ...
type DeleteModel struct {
	FileName  string `json:"filename"`
	Recursive bool   `json:"recursive"`
}

// PostCopyHandler ...
func PostCopyHandler(w http.ResponseWriter, r *http.Request) {
	model := &DeplacementModel{}
	if err := utility.ParsePostBody(r, model); err != nil {
		utility.RespondWithError(w, http.StatusBadRequest, err)
		return
	}
	src := filepath.Join(Context.GetRoot(), model.Origin)
	dst := filepath.Join(Context.GetRoot(), model.Destination)
	if err := files.Copy(src, dst); err != nil {
		utility.RespondWithError(w, http.StatusBadRequest, err)
		return
	}
}

// PostMoveHandler ...
func PostMoveHandler(w http.ResponseWriter, r *http.Request) {
	model := &DeplacementModel{}
	if err := utility.ParsePostBody(r, model); err != nil {
		utility.RespondWithError(w, http.StatusBadRequest, err)
		return
	}
	src := filepath.Join(Context.GetRoot(), model.Origin)
	dst := filepath.Join(Context.GetRoot(), model.Destination)
	if err := os.Rename(src, dst); err != nil {
		utility.RespondWithError(w, http.StatusBadRequest, err)
		return
	}
}

// DeleteHandler ...
func DeleteHandler(w http.ResponseWriter, r *http.Request) {
	model := &DeleteModel{}
	var err error
	if err = utility.ParsePostBody(r, model); err != nil {
		utility.RespondWithError(w, http.StatusBadRequest, err)
		return
	}
	file := filepath.Join(Context.GetRoot(), model.FileName)

	if model.Recursive {
		err = os.RemoveAll(file)
	} else {
		err = os.Remove(file)
	}
	if err != nil {
		utility.RespondWithError(w, http.StatusBadRequest, err)
	}
}

// AddOperationRoute ...
func AddOperationRoute(r *mux.Router) {
	subrouter := r.PathPrefix("/op").Subrouter()
	subrouter.Path("/cp").Methods("POST").HandlerFunc(PostCopyHandler)
	subrouter.Path("/mv").Methods("POST").HandlerFunc(PostMoveHandler)
	subrouter.Path("/rm").Methods("DELETE").HandlerFunc(DeleteHandler)
	subrouter.Use(auth.MiddlewareFile)
}
