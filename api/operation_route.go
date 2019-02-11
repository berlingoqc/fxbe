package api

import "net/http"

type DeplacementModel struct {
	Origin      string `json:"origin"`
	Destination string `json:"destination"`
	Recursive   bool   `json:"recursive"`
}

type RenameModel struct {
	Origin  string `json:"origin"`
	NewName string `json:"newname"`
}

type DeleteModel struct {
	FileName  string `json:"filename"`
	Recursive bool   `json:"recursive"`
}

func PostCopyHandler(w http.ResponseWriter, r *http.Request) {

}

func PostMoveHandler(w http.ResponseWriter, r *http.Request) {

}

func PostRenameHandler(w http.ResponseWriter, r *http.Request) {

}

func DeleteHandler(w http.ResponseWriter, r *http.Request) {

}
