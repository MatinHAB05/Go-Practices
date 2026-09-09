package main

import (
	"fmt"
	"time"

	"aidanwoods.dev/go-paseto"
)

func main() {
	fmt.Println("======== PASETO V4 DEMO ========")
	// =========================================================================
	// 2. PASETO v4.public (Asymmetric Ed25519 Signatures)
	// =========================================================================
	fmt.Println("\n--- 2. Testing PASETO v4.public (Asymmetric) ---")

	// // Generate Ed25519 key pair
	// pubKey, privKey, err := ed25519.GenerateKey(rand.Reader)
	// if err != nil {
	// 	panic(err)
	// }

	pasetoSecretKey := paseto.NewV4AsymmetricSecretKey()
	pasetoPublicKey := pasetoSecretKey.Public()

	// Build token claims
	publicToken := paseto.NewToken()
	publicToken.SetIssuedAt(time.Now())
	publicToken.SetExpiration(time.Now().Add(30 * time.Minute))
	publicToken.SetSubject("user_1234")
	publicToken.SetString("scope", "read:write")

	// Sign token with Private Key
	signedTokenString := publicToken.V4Sign(pasetoSecretKey, nil)
	fmt.Println("Signed Public Token:")
	fmt.Println(signedTokenString)

	// Verify token with Public Key
	parsedPublic, err := paseto.NewParser().ParseV4Public(pasetoPublicKey, signedTokenString, nil)
	if err != nil {
		fmt.Printf("Verification Failed: %v\n", err)
	} else {
		sub, _ := parsedPublic.GetSubject()
		scope, _ := parsedPublic.GetString("scope")
		fmt.Printf("Success! Sub: %s, Scope: %s\n", sub, scope)
	}

}
