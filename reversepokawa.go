package main

import (
	"crypto/hmac"
	"crypto/sha256"
	"encoding/base64"
	"encoding/json"
	"fmt"
	"net/http"
	"time"
)


const secretKey = "wQxTBLXYMXMbVjc7wnXkcsr2phZCRVnJNAa4CRJfUCqsQhc6qJ"


type SignatureRequest struct {
	CanonicalString string `json:"canonical_string_from_browser"`
}
type SignatureResponse struct {
	Signature string `json:"signature"`
}
type Credentials struct {
	Username string `json:"username"`
	Password string `json:"password"`
}
type CanonicalStringResponse struct {
	CanonicalString string `json:"canonical_string"`
}



type AuthPayload struct {
	Username string `json:"username"`
	Password string `json:"password"`
}


func generateCanonicalStringHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Méthode non autorisée", http.StatusMethodNotAllowed)
		return
	}

	var creds Credentials
	if err := json.NewDecoder(r.Body).Decode(&creds); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	payload := AuthPayload{
		Username: creds.Username,
		Password: creds.Password,
	}


	payloadBytes, err := json.Marshal(payload)
	if err != nil {
		http.Error(w, "Erreur lors de la création du JSON", http.StatusInternalServerError)
		return
	}


	base64Payload := base64.StdEncoding.EncodeToString(payloadBytes)

	currentTime := time.Now().UTC()
	adjustedTime := currentTime
	truncatedTime := adjustedTime.Truncate(time.Minute)
	formattedTime := truncatedTime.Format("2006-01-02T15:04:05.000Z")
	canonicalString := fmt.Sprintf("/auth/login_%s_%s", base64Payload, formattedTime)

	res := CanonicalStringResponse{CanonicalString: canonicalString}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(res)
}


func signatureHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Méthode non autorisée", http.StatusMethodNotAllowed)
		return
	}
	var req SignatureRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	secretKeyBytes := []byte(secretKey)
	canonicalStringBytes := []byte(req.CanonicalString)
	hasher := hmac.New(sha256.New, secretKeyBytes)
	hasher.Write(canonicalStringBytes)
	binarySignature := hasher.Sum(nil)
	signature := base64.StdEncoding.EncodeToString(binarySignature)
	res := SignatureResponse{Signature: signature}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(res)
}

func main() {
	http.HandleFunc("/generate-canonical-string", generateCanonicalStringHandler)
	http.HandleFunc("/signature", signatureHandler)
	fmt.Println("Le serveur démarre sur http://localhost:6767")
	fmt.Println("Points de terminaison disponibles :")
	fmt.Println("  POST /generate-canonical-string (prend username, password)")
	fmt.Println("  POST /signature (prend canonical_string_from_browser)")
	http.ListenAndServe(":6767", nil)
}
