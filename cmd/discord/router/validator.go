package router

import (
	"bytes"
	"crypto/ed25519"
	"encoding/hex"
	"io"
	"net/http"

	"github.com/pkg/errors"
)

type RequestValidator interface {
	Validate(r *http.Request) (bool, error)
}

type discordPublicKeyValidator struct {
	key ed25519.PublicKey
}

func (v discordPublicKeyValidator) Validate(r *http.Request) (bool, error) {
	valid := verifyPublicKey(r, v.key)
	if !valid {
		return false, errors.New("public key verification failed")
	}
	return true, nil
}

func NewPublicKeyValidator(key ed25519.PublicKey) RequestValidator {
	return discordPublicKeyValidator{key}
}

// https://github.com/bsdlp/discord-interactions-go/blob/main/interactions/verify.go
func verifyPublicKey(r *http.Request, key ed25519.PublicKey) bool {
	var msg bytes.Buffer

	signature := r.Header.Get("X-Signature-Ed25519")
	if signature == "" {
		return false
	}

	sig, err := hex.DecodeString(signature)
	if err != nil {
		return false
	}

	if len(sig) != ed25519.SignatureSize || sig[63]&224 != 0 {
		return false
	}

	timestamp := r.Header.Get("X-Signature-Timestamp")
	if timestamp == "" {
		return false
	}

	msg.WriteString(timestamp)

	defer r.Body.Close()
	var body bytes.Buffer

	// at the end of the function, copy the original body back into the request
	defer func() {
		r.Body = io.NopCloser(&body)
	}()

	// copy body into buffers
	_, err = io.Copy(&msg, io.TeeReader(r.Body, &body))
	if err != nil {
		return false
	}

	return ed25519.Verify(key, msg.Bytes(), sig)
}
