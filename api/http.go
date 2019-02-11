package api

import (
	"net/http"
	"os"
	"path"
	"strings"

	"git.wquintal.ca/berlingoqc/fxbe/files"

	"github.com/gorilla/mux"
)

var Context = make(files.Ctx)

// GetFilesHandler is the handler that return the get request for a file path
func GetFilesHandler(w http.ResponseWriter, r *http.Request) {
	fileReq := strings.TrimPrefix(r.RequestURI, "/files")
	println(fileReq)
	// Regarde s'il s'agit du dossier ou d'un fichier
	if s, err := os.Stat(path.Join(Context.GetRoot(), fileReq)); err == nil {
		if s.IsDir() {
			// Retourn le contenu du repertoire
			items, err := files.ListingDirectory(Context, fileReq)
			if err != nil {
				RespondWithError(w, http.StatusBadRequest, err)
			}
			RespondWithJSON(w, http.StatusOK, items)
			return
		}
		// Retourne le fichier

	} else {
		RespondWithError(w, http.StatusBadRequest, err)
		return
	}
}

// AddFileRoute add the files serving api route
func AddFileRoute(r *mux.Router) {
	r.PathPrefix("/files/").Methods("GET").HandlerFunc(GetFilesHandler)
}

// StartWebServer start the web server
func StartWebServer() {
	r := mux.NewRouter()
	AddFileRoute(r)

	srv := &http.Server{
		Addr:    "0.0.0.0:8080",
		Handler: r,
	}
	if err := srv.ListenAndServe(); err != nil {
		panic(err)
	}
}
