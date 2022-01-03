package urlshortner

import (
	"encoding/json"
	"net/http"

	"github.com/gorilla/mux"
)

type urlHandler struct {
	UrlHandlerService UrlShortnerService
}

type UrlEncoder struct {
	Url string `json:"url"`
}
type UrlDecoder struct {
	UrlCode string `json:"urlCode"`
}

func NewHandler(urlService UrlShortnerService, router *mux.Router) {
	handler := &urlHandler{UrlHandlerService: urlService}
	router.HandleFunc("/urlEncoder", handler.urlEncoder).Methods(http.MethodPost)
	router.HandleFunc("/{code}", handler.urlDecoder).Methods(http.MethodGet)
}

func (uh urlHandler) urlEncoder(w http.ResponseWriter, r *http.Request) {
	var urlInput UrlEncoder
	err := json.NewDecoder(r.Body).Decode(&urlInput)
	if err != nil {
		WriteErrorResponse(w, http.StatusOK, err.Error())
		return
	}
	urlStr, _ := uh.UrlHandlerService.encodeUrl(r.Context(), urlInput)

	WriteOKResponse(w, urlStr)
}

func (uh urlHandler) urlDecoder(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	code, ok := vars["code"]
	if !ok {
		WriteErrorResponse(w, http.StatusOK, "code is missing in parameters")
	}
	urlStr := uh.UrlHandlerService.decodeUrl(r.Context(), code)

	WriteOKResponse(w, urlStr)
}
