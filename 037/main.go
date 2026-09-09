package main

import (
	"encoding/json"
	"log"
	"net/http"
	"sync"

	"github.com/go-webauthn/webauthn/protocol"
	"github.com/go-webauthn/webauthn/webauthn"
)

// User represents a user model implementing webauthn.User interface
type User struct {
	id          []byte
	name        string
	displayName string
	credentials []webauthn.Credential
}

func (u *User) WebAuthnID() []byte                         { return u.id }
func (u *User) WebAuthnName() string                       { return u.name }
func (u *User) WebAuthnDisplayName() string                { return u.displayName }
func (u *User) WebAuthnCredentials() []webauthn.Credential { return u.credentials }
func (u *User) WebAuthnIcon() string                       { return "" }

// In-memory user database & session store
var (
	userDB      = make(map[string]*User)
	sessionDB   = make(map[string]*webauthn.SessionData)
	dbMutex     sync.Mutex
	webAuthnObj *webauthn.WebAuthn
)

func main() {
	var err error

	// Initialize WebAuthn configuration
	webAuthnObj, err = webauthn.New(&webauthn.Config{
		RPDisplayName: "My Go App",
		RPID:          "localhost",                       // Domain without protocol/port
		RPOrigins:     []string{"http://localhost:8080"}, // Allowed origins
	})
	if err != nil {
		log.Fatalf("Failed to create WebAuthn instance: %v", err)
	}

	// Serve static UI (index.html) for testing
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		http.ServeFile(w, r, "index.html")
	})

	// WebAuthn API Handlers
	http.HandleFunc("/register/begin", handleBeginRegistration)
	http.HandleFunc("/register/finish", handleFinishRegistration)
	http.HandleFunc("/login/begin", handleBeginLogin)
	http.HandleFunc("/login/finish", handleFinishLogin)

	log.Println("Server running on http://localhost:8080")
	log.Fatal(http.ListenAndServe(":8080", nil))
}

// -----------------------------------------------------------------------------
// 1. REGISTRATION FLOW
// -----------------------------------------------------------------------------

func handleBeginRegistration(w http.ResponseWriter, r *http.Request) {
	username := r.URL.Query().Get("username")
	if username == "" {
		http.Error(w, "Username required", http.StatusBadRequest)
		return
	}

	dbMutex.Lock()
	user, exists := userDB[username]
	if !exists {
		user = &User{
			id:          []byte(username), // Simple ID generation for demo
			name:        username,
			displayName: username,
		}
		userDB[username] = user
	}
	dbMutex.Unlock()

	// Begin registration. go-webauthn automatically checks user.WebAuthnCredentials()
	// to build the exclusion list for existing passkeys.
	options, session, err := webAuthnObj.BeginRegistration(
		user,
		webauthn.WithAuthenticatorSelection(protocol.AuthenticatorSelection{
			UserVerification: protocol.VerificationPreferred,
		}),
	)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	// Save session data (In production, store in encrypted session cookie or Redis)
	dbMutex.Lock()
	sessionDB[username] = session
	dbMutex.Unlock()

	jsonResponse(w, options)
}

func handleFinishRegistration(w http.ResponseWriter, r *http.Request) {
	username := r.URL.Query().Get("username")

	dbMutex.Lock()
	user, userExists := userDB[username]
	session, sessionExists := sessionDB[username]
	dbMutex.Unlock()

	if !userExists || !sessionExists {
		http.Error(w, "Session not found", http.StatusBadRequest)
		return
	}

	// Parse webauthn client response
	credential, err := webAuthnObj.FinishRegistration(user, *session, r)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	// Save the new passkey credential to the user
	dbMutex.Lock()
	user.credentials = append(user.credentials, *credential)
	delete(sessionDB, username) // Clear session
	dbMutex.Unlock()

	jsonResponse(w, map[string]string{"status": "registration successful"})
}

// -----------------------------------------------------------------------------
// 2. AUTHENTICATION FLOW
// -----------------------------------------------------------------------------

func handleBeginLogin(w http.ResponseWriter, r *http.Request) {
	username := r.URL.Query().Get("username")

	dbMutex.Lock()
	user, exists := userDB[username]
	dbMutex.Unlock()

	if !exists {
		http.Error(w, "User not found", http.StatusBadRequest)
		return
	}

	options, session, err := webAuthnObj.BeginLogin(user)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	dbMutex.Lock()
	sessionDB[username] = session
	dbMutex.Unlock()

	jsonResponse(w, options)
}

func handleFinishLogin(w http.ResponseWriter, r *http.Request) {
	username := r.URL.Query().Get("username")

	dbMutex.Lock()
	user, userExists := userDB[username]
	session, sessionExists := sessionDB[username]
	dbMutex.Unlock()

	if !userExists || !sessionExists {
		http.Error(w, "Session not found", http.StatusBadRequest)
		return
	}

	// Verify the passkey response
	credential, err := webAuthnObj.FinishLogin(user, *session, r)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	// Check for credential clone warnings
	if credential.Authenticator.CloneWarning {
		log.Printf("Warning: Passkey credential clone suspected for user %s", username)
	}

	dbMutex.Lock()
	delete(sessionDB, username)
	dbMutex.Unlock()

	jsonResponse(w, map[string]string{"status": "login successful"})
}

func jsonResponse(w http.ResponseWriter, data interface{}) {
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(data)
}
