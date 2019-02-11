package api

import (
	"net/http"
	"os"
	"path"
	"strings"
	"time"

	"git.wquintal.ca/berlingoqc/fxbe/auth"

	"git.wquintal.ca/berlingoqc/fxbe/files"
	"git.wquintal.ca/berlingoqc/fxbe/utility"

	"github.com/gorilla/mux"
)

var Context = make(files.Ctx)

func getFileReq(r *http.Request) (string, string) {
	fileReq := strings.TrimPrefix(r.RequestURI, "/files")
	fullPath := path.Join(Context.GetRoot(), fileReq)
	return fileReq, fullPath
}

// PostFilesHandler is the handler for the post request for uploading files
func PostFilesHandler(w http.ResponseWriter, r *http.Request) {
	// Get le fichier qu'on veut upload pis path complet
	_, fullPath := getFileReq(r)
	if _, err := os.Stat(fullPath); err == nil {
		utility.RespondWithError(w, http.StatusBadRequest, err)
		return
	}
	file, _, err := r.FormFile("file")
	if err != nil {
		utility.RespondWithError(w, http.StatusBadRequest, err)
		return
	}

	err = files.UploadFile(fullPath, file)
	if err != nil {
		utility.RespondWithError(w, http.StatusInternalServerError, err)
		return
	}

	w.Write([]byte("OK"))

}

// GetFilesHandler is the handler that return the get request for a file path
func GetFilesHandler(w http.ResponseWriter, r *http.Request) {
	fileReq := strings.TrimPrefix(r.RequestURI, "/files")
	fullPath := path.Join(Context.GetRoot(), fileReq)
	// Regarde s'il s'agit du dossier ou d'un fichier
	if s, err := os.Stat(fullPath); err == nil {
		if s.IsDir() {
			// Retourn le contenu du repertoire
			items, err := files.ListingDirectory(Context, fileReq)
			if err != nil {
				utility.RespondWithError(w, http.StatusBadRequest, err)
				return
			}
			utility.RespondWithJSON(w, http.StatusOK, items)
			return
		}
		// Retourne le fichier
		f, err := os.Open(fullPath)
		if err != nil {
			utility.RespondWithError(w, http.StatusBadRequest, err)
			return
		}

		http.ServeContent(w, r, s.Name(), time.Now(), f)

	} else {
		utility.RespondWithError(w, http.StatusBadRequest, err)
		return
	}
}

// AddFileRoute add the files serving api route
func AddFileRoute(r *mux.Router) {
	srf := r.PathPrefix("/files/").Subrouter()
	srf.Methods("GET").HandlerFunc(GetFilesHandler)
	srf.Methods("POST").HandlerFunc(PostFilesHandler)
	srf.Use(auth.MiddlewareFile)
}

// AddAccountRoute add the route for login and account management
func AddAccountRoute(r *mux.Router) {
	ar := r.PathPrefix("/auth").Subrouter()
	ar.HandleFunc("/", auth.GetLoginHandler)
}

// StartWebServer start the web server
func StartWebServer() {
	r := mux.NewRouter()
	AddFileRoute(r)
	AddAccountRoute(r)

	srv := &http.Server{
		Addr:    "0.0.0.0:8080",
		Handler: r,
	}
	if err := srv.ListenAndServe(); err != nil {
		panic(err)
	}
}
