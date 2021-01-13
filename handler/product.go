package handler

import (
	"encoding/json"
	"github.com/gorilla/mux"
	"net/http"
)

func GetProduct(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)

	p := Product{}
	var err error
	if p, err = p.getProduct(vars["id"]); err != nil {
		respondWithError(w, http.StatusInternalServerError, err.Error())
		return
	}

	respondWithJSON(w, http.StatusOK, p)
}

func CreateProduct(w http.ResponseWriter, r *http.Request) {
	p := Product{}
	if err := json.NewDecoder(r.Body).Decode(&p); err != nil {
		respondWithError(w, http.StatusInternalServerError, err.Error())
		return
	}

	if err := p.createProduct(p); err != nil {
		respondWithError(w, http.StatusInternalServerError, err.Error())
		return
	}
	respondWithJSON(w, http.StatusOK, p)
}

func DeleteProduct(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	p := Product{}
	if err := p.deleteProduct(vars["id"]); err != nil {
		respondWithError(w, http.StatusInternalServerError, err.Error())
		return
	}
	w.WriteHeader(http.StatusOK)
}

func UpdateProduct(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	p := Product{}
	if err := json.NewDecoder(r.Body).Decode(&p); err != nil {
		respondWithError(w, http.StatusInternalServerError, err.Error())
		return
	}
	p.ID = stringToInt(vars["id"])
	if err := p.updateProduct(p, vars["id"]); err != nil {
		respondWithError(w, http.StatusInternalServerError, err.Error())
		return
	}
	respondWithJSON(w, http.StatusOK, p)
}


func respondWithError(w http.ResponseWriter, code int, message string) {
	respondWithJSON(w, code, map[string]string{"error": message})
}

func respondWithJSON(w http.ResponseWriter, code int, payload interface{}) {
	response, _ := json.Marshal(payload)

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(code)
	w.Write(response)
}
