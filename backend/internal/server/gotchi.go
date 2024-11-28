package server

import (
	"backend/internal/models"
	"backend/internal/utils"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"net/url"

	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
)

func (s *FiberServer) VerifyAuthKeyHandler(c *fiber.Ctx) error {
	type requestBody struct {
		Token string `json:"token"`
		Hash  string `json:"hash"`
		Id    string `json:"id"`
		Name  string `json:"name"`
	}
	rBody := new(requestBody)

	if err := c.BodyParser(rBody); err != err {
		return c.Status(fiber.StatusInternalServerError).JSON(map[string]interface{}{
			"err": err.Error(),
		})
	}

	if rBody.Hash == "" || rBody.Token == "" {
		return c.Status(fiber.StatusBadRequest).JSON(map[string]interface{}{
			"error": "Domain Hash or Token Missing",
		})
	}

	// Define the base URL
	baseURL := "https://" + rBody.Hash + ".canary.tools/api/v1/ping"

	// Define the payload (query parameters)
	params := url.Values{}
	params.Add("auth_token", rBody.Token)

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

	if pingResponse.Result != "success" {
		return c.Status(fiber.StatusBadRequest).JSON(map[string]interface{}{
			"error": "Failed to authorize the token",
		})
	}

	uid, err := uuid.Parse(rBody.Id)

	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(map[string]interface{}{
			"error": "ID is not valid",
		})
	}

	gotchi := models.Gotchi{
		ID:        uid,
		Name:      rBody.Name,
		Hash:      rBody.Hash,
		AuthToken: rBody.Token,
		Level:     1,
		Sequence: models.Sequence{
			Sequence: utils.GenerateRandomSequence(5),
		},
	}

	gotchi, err = gotchi.Create(s.DB)

	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(map[string]interface{}{
			"error":     err.Error(),
			"duplicate": gotchi.Hash == rBody.Hash,
		})
	}

	return c.Status(fiber.StatusOK).JSON(map[string]interface{}{
		"result": "success",
		"data":   gotchi,
	})
}

func (s *FiberServer) ExistingGotchiHandler(c *fiber.Ctx) error {
	id := c.Params("id")

	uid, err := uuid.Parse(id)

	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(map[string]interface{}{
			"error": "Gotchi ID is not valid",
		})
	}

	gotchi := models.Gotchi{
		ID: uid,
	}

	gotchi, err = gotchi.GetWithSequence(s.DB)

	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(map[string]interface{}{
			"error": "Gotchi not found",
		})
	}

	if gotchi.Sequence.ID == uuid.Nil {
		newSeq := models.Sequence{
			GotchiID: gotchi.ID,
			Sequence: utils.GenerateRandomSequence(5),
		}

		newSeq, err = newSeq.Create(s.DB)

		if err != nil {
			return c.Status(fiber.StatusInternalServerError).JSON(map[string]interface{}{
				"error": "Failed to create new sequence",
			})
		}

		gotchi.Sequence = newSeq
	}

	return c.Status(fiber.StatusOK).JSON(map[string]interface{}{
		"result": "success",
		"data":   gotchi,
	})
}
