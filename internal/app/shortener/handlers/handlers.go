package handlers

import (
	"encoding/json"
	"fmt"
	"github.com/go-chi/chi/v5"
	"github.com/putalexey/go-practicum/internal/app/shortener/requests"
	"github.com/putalexey/go-practicum/internal/app/shortener/responses"
	"github.com/putalexey/go-practicum/internal/app/storage"
	"github.com/putalexey/go-practicum/internal/app/urlgenerator"
	"io"
	"net/http"
	"net/url"
)

func GetFullURLHandler(storage storage.Storager) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		id := chi.URLParam(r, "id")
		if id == "" {
			http.Error(w, "Bad request", http.StatusBadRequest)
			return
		}

		if fullURL, err := storage.Load(id); err == nil {
			http.Redirect(w, r, fullURL, http.StatusTemporaryRedirect)
			return
		}
		http.Error(w, "Not found", http.StatusNotFound)
	}
}

func CreateFullURLHandler(generator urlgenerator.URLGenerator, storage storage.Storager) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		body, err := io.ReadAll(r.Body)
		defer r.Body.Close()
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		if len(body) == 0 {
			http.Error(w, "Empty request", http.StatusBadRequest)
			return
		}

		fullURL := string(body)
		if _, err := url.ParseRequestURI(fullURL); err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}

		short := generator.GenerateShort(fullURL)
		if err := storage.Store(short, fullURL); err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
		}

		w.WriteHeader(http.StatusCreated)
		_, _ = fmt.Fprint(w, generator.GetURL(short))
	}
}

func JSONCreateShort(generator urlgenerator.URLGenerator, storage storage.Storager) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		body, err := io.ReadAll(r.Body)
		defer r.Body.Close()

		if err != nil {
			jsonError(w, err.Error(), http.StatusInternalServerError)
			return
		}
		if len(body) == 0 {
			jsonError(w, "Empty request", http.StatusBadRequest)
			return
		}

		createRequest := requests.CreateShortRequest{}
		if err = json.Unmarshal(body, &createRequest); err != nil {
			jsonError(w, "Request can't be parsed", http.StatusBadRequest)
			return
		}

		if _, err := url.ParseRequestURI(createRequest.URL); err != nil {
			jsonError(w, err.Error(), http.StatusBadRequest)
			return
		}

		short := generator.GenerateShort(createRequest.URL)
		if err := storage.Store(short, createRequest.URL); err != nil {
			jsonError(w, err.Error(), http.StatusInternalServerError)
			return
		}

		createResponse := responses.CreateShortResponse{Result: generator.GetURL(short)}
		data, err := json.Marshal(createResponse)
		if err != nil {
			jsonError(w, err.Error(), http.StatusInternalServerError)
			return
		}

		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusCreated)
		_, _ = w.Write(data)
	}
}

func BadRequestHandler() http.HandlerFunc {
	return func(w http.ResponseWriter, _ *http.Request) {
		http.Error(w, "Bad request", http.StatusBadRequest)
	}
}

func jsonError(w http.ResponseWriter, error string, code int) {
	_err := responses.ErrorResponse{Error: error}
	response, err := json.Marshal(_err)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(code)
	_, _ = w.Write(response)
}
