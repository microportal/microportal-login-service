package controller

import (
	"encoding/json"
	"github.com/Nerzal/gocloak"
	"microportal-resource-service/model"
	"net/http"
	"os"
)

type LoginController struct{}

var (
	clientId     string // = "microportal"
	clientSecret string // = "b56f5388-f2b6-4f9c-be59-48fed7294f43"
	realm        string // = "MicroportalRealm"
	keycloakUrl  string // = "http://localhost:7000/"
)

func (c LoginController) Init() {
	clientId = os.Getenv("KEYCLOAK_CLIENT_ID")
	clientSecret = os.Getenv("KEYCLOAK_CLIENT_SECRET")
	realm = os.Getenv("KEYCLOAK_REALM")
	keycloakUrl = os.Getenv("KEYCLOAK_URL")
}

func (c LoginController) Login(w http.ResponseWriter, r *http.Request) {
	client := gocloak.NewClient(keycloakUrl)

	var formLogin model.FormLogin
	_ = json.NewDecoder(r.Body).Decode(&formLogin)

	token, err := client.Login(clientId, clientSecret, realm, formLogin.Username, formLogin.Password)

	if err != nil {
		w.WriteHeader(http.StatusUnauthorized)
	}

	w.Header().Set("Content-Type", "application/json")
	_ = json.NewEncoder(w).Encode(token)
}

func (c LoginController) ValidateToken(w http.ResponseWriter, r *http.Request) {

	client := gocloak.NewClient(keycloakUrl)

	accessToken := r.Header.Get("Authorization")

	introspect, err := client.RetrospectToken(accessToken[7:], clientId, clientSecret, realm)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		_ = json.NewEncoder(w).Encode(err.Error())
		return
	}

	if !introspect.Active {
		w.WriteHeader(http.StatusUnauthorized)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	_ = json.NewEncoder(w).Encode(introspect.Active)
}
