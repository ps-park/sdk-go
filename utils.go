package sdk_go

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"math/rand"
	"net/http"
	"strings"
	"time"

	"github.com/golang-jwt/jwt/v4"
)

type SignedToken struct {
	Token  string
	Claims map[string]interface{}
}

// SignToken generates a JWT token with the given secret and custom body
func signToken(secret string, body interface{}) (SignedToken, error) {

	now := time.Now()
	nonce := now.Unix() + rand.Int63n(1000)

	claims := jwt.MapClaims{
		"nonce": nonce,
		"iat":   now.Unix(),
		"exp":   now.Add(time.Second * TokenExpirationTime).Unix(),
	}

	bodyJson, err := json.Marshal(body)

	if err != nil {
		return SignedToken{}, err
	}

	var bodyMap map[string]interface{}

	err = json.Unmarshal(bodyJson, &bodyMap)

	if err != nil {
		return SignedToken{}, err
	}

	for key, value := range bodyMap {
		claims[key] = value
	}

	signedToken := SignedToken{
		Claims: claims,
	}

	key, err := jwt.ParseRSAPrivateKeyFromPEM([]byte(secret))

	if err != nil {
		return signedToken, err
	}

	token, err := jwt.NewWithClaims(jwt.SigningMethodRS256, claims).SignedString(key)

	signedToken.Token = token

	if err != nil {
		return signedToken, err
	}

	return signedToken, nil
}

func makeAuthenticatedRequest[T any](url string, requestBody interface{}, client *PSPark) (T, error) {
	var response T
	fullURL := fmt.Sprintf("%s/%s/%s", client.BaseURL, APIVersion, url)

	token, err := signToken(client.Secret, requestBody)
	if err != nil {
		fmt.Println("Error signing token")
		return response, err
	}

	jsonBody, err := json.Marshal(token.Claims)
	if err != nil {
		return response, err
	}

	req, err := http.NewRequest("POST", fullURL, bytes.NewBuffer(jsonBody))
	if err != nil {
		return response, err
	}

	req.Header.Set("X-API-Key", client.APIKey)
	req.Header.Set("Authorization", "Bearer "+token.Token)
	req.Header.Set("Content-Type", "application/json")

	httpClient := &http.Client{}
	resp, err := httpClient.Do(req)

	if err != nil {
		return response, err
	}
	defer resp.Body.Close()

	bodyBytes, err := io.ReadAll(resp.Body)
	if err != nil {
		return response, err
	}

	var genericResponse ResponseDTO[interface{}]

	if err = json.Unmarshal(bodyBytes, &genericResponse); err != nil {
		return response, err
	}

	if genericResponse.Code > 0 || strings.ToLower(genericResponse.Message) != "ok" {
		return response, &ResponseValidationError{
			Message: genericResponse.Message,
			Code:    genericResponse.Code,
			Data:    genericResponse.Data,
		}
	}

	var data ResponseDTO[T]
	if err = json.Unmarshal(bodyBytes, &data); err != nil {
		return response, err
	}

	response = data.Data

	return response, nil
}
