package server

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"net/url"

	"github.com/gofiber/fiber/v2"
)

func (s *FiberServer) RegisterFiberRoutes() {
	s.App.Get("/", s.HelloWorldHandler)

	s.App.Post("/verify-api-key", s.VerifyApiKeyHandler)

	s.App.Get("/health", s.healthHandler)
}

func (s *FiberServer) VerifyApiKeyHandler(c *fiber.Ctx) error {
	type requestBody struct {
		Token string `json:"token"`
		Hash  string `json:"hash"`
	}
	rbody := new(requestBody)

	if err := c.BodyParser(rbody); err != err {
		return c.Status(fiber.StatusInternalServerError).JSON(map[string]interface{}{
			"err": err.Error(),
		})
	}

	if rbody.Hash == "" || rbody.Token == "" {
		return c.Status(fiber.StatusBadRequest).JSON(map[string]interface{}{
			"error": "Domain Hash or Token Missing",
		})
	}

	// Define the base URL
	baseURL := "https://" + rbody.Hash + ".canary.tools/api/v1/ping"

	// Define the payload (query parameters)
	params := url.Values{}
	params.Add("auth_token", rbody.Token)

	// Construct the full URL with query parameters
	fullURL := fmt.Sprintf("%s?%s", baseURL, params.Encode())

	// Make the GET request
	resp, err := http.Get(fullURL)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(map[string]interface{}{
			"error": err.Error(),
		})
	}
	defer resp.Body.Close()

	// Check the response status code
	if resp.StatusCode != http.StatusOK {
		return c.Status(fiber.StatusInternalServerError).JSON(map[string]interface{}{
			"error": fmt.Sprintf("Error: received status code %d", resp.StatusCode),
		})
	}

	// Read and print the response body
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(map[string]interface{}{
			"error": err.Error(),
		})
	}

	// Define a struct for the expected JSON response
	type PingResponse struct {
		Result string `json:"result"`
	}

	// Unmarshal the JSON response into the struct
	var pingResponse PingResponse
	if err := json.Unmarshal(body, &pingResponse); err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(map[string]interface{}{
			"error": err.Error(),
		})
	}

	return c.Status(fiber.StatusOK).JSON(map[string]interface{}{
		"data": pingResponse,
		"test": string(body),
	})
}

func (s *FiberServer) HelloWorldHandler(c *fiber.Ctx) error {
	resp := fiber.Map{
		"message": "Hello World, from Fiber!",
	}

	return c.JSON(resp)
}

func (s *FiberServer) healthHandler(c *fiber.Ctx) error {
	return c.JSON(s.db.Health())
}
