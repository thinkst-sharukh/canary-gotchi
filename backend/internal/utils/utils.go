package utils

import (
	"backend/internal/types"
	"encoding/json"
	"fmt"
	"io"
	"math/rand"
	"net/http"
	"net/url"
	"os"
	"strings"
	"time"
)

var directions = []string{"up", "right", "left", "down"}

func IsDev() bool {
	return os.Getenv("ENVIRONMENT") == "development"
}

func GenerateRandomSequence(length int) string {
	// Create a new random generator with a unique seed
	rng := rand.New(rand.NewSource(time.Now().UnixNano()))
	sequence := make([]string, length)
	for i := 0; i < length; i++ {
		sequence[i] = directions[rng.Intn(len(directions))]
	}
	return strings.Join(sequence, ",")
}

func ValidateAuthToken(hash, authToken string) (types.PingResponse, error) {
	pingResponse := types.PingResponse{}
	// Define the base URL
	baseURL := "https://" + hash + ".canary.tools/api/v1/ping"

	// Define the payload (query parameters)
	params := url.Values{}
	params.Add("auth_token", authToken)

	// Construct the full URL with query parameters
	fullURL := fmt.Sprintf("%s?%s", baseURL, params.Encode())

	// Make the GET request
	resp, err := http.Get(fullURL)
	if err != nil {
		return pingResponse, fmt.Errorf("Failed to make request")
	}
	defer resp.Body.Close()

	// Check the response status code
	if resp.StatusCode != http.StatusOK {
		return pingResponse, fmt.Errorf("Auth token or hash is invalid")
	}

	// Read and print the response body
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return pingResponse, fmt.Errorf("Failed to read response body")
	}

	if err := json.Unmarshal(body, &pingResponse); err != nil {
		return pingResponse, fmt.Errorf("Failed to unmarshal response body")
	}

	if pingResponse.Result != "success" {
		return pingResponse, fmt.Errorf("Error: %s", pingResponse.Result)
	}

	return pingResponse, nil
}
