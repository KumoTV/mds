package controllers

import (
	"encoding/json"
	"github.com/KumoTV/mds/api/v1/models"
	"github.com/google/uuid"
	"github.com/gorilla/mux"
	"net/http"
)

type Env struct {
	db models.MetadataStore
}

func (env *Env) getMetadataById(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	metadata_id, err := uuid.Parse(params["id"])
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		panic(err)
	}
	m := env.db.GetMetadataById(metadata_id)
	json.NewEncoder(w).Encode(m)
}

func (env *Env) getMetadata(w http.ResponseWriter, r *http.Request) {}

func (env *Env) importMetadata(w http.ResponseWriter, r *http.Request) {
	m := models.VodMetadata{}
	_ = json.NewDecoder(r.Body).Decode(&m.Metadata)
	metadata_id := env.db.ImportMetadata(m)
	json.NewEncoder(w).Encode(metadata_id)
}

func StartHttpServer(ds models.MetadataStore) {
	env := &Env{db: ds}

	router := mux.NewRouter()
	s := router.PathPrefix("/api/v1").Subrouter()

	//GET Methods
	s.HandleFunc("/vod", env.getMetadata).Methods("GET")
	s.HandleFunc("/vod/{id}", env.getMetadataById).Methods("GET")

	//POST Methods
	s.HandleFunc("/vod", env.importMetadata).Methods("POST")

	http.ListenAndServe(":8080", router)
}
