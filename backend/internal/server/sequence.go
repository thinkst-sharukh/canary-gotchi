package server

import (
	"backend/internal/models"
	"backend/internal/utils"
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
)

func (s *FiberServer) RegenerateSequence(c *fiber.Ctx) error {
	id := c.Params("id")

	uid, err := uuid.Parse(id)

	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(map[string]interface{}{
			"error": "Gotchi ID is not valid",
		})
	}

	seq := models.Sequence{
		GotchiID: uid,
		Sequence: utils.GenerateRandomSequence(5),
	}

	seq, err = seq.Create(s.DB)

	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(map[string]interface{}{
			"error": err.Error(),
		})
	}

	return c.Status(fiber.StatusOK).JSON(map[string]interface{}{
		"result": "success",
		"data":   seq,
	})
}

func (s *FiberServer) ValidateSequence(c *fiber.Ctx) error {
	type requestBody struct {
		Sequence string `json:"sequence"`
	}
	rBody := new(requestBody)

	if err := c.BodyParser(rBody); err != err {
		return c.Status(fiber.StatusInternalServerError).JSON(map[string]interface{}{
			"err": err.Error(),
		})
	}

	if rBody.Sequence == "" {
		return c.Status(fiber.StatusBadRequest).JSON(map[string]interface{}{
			"error": "Sequence is missing",
		})
	}

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
			"error": "Failed to get Gotchi",
		})
	}

	if rBody.Sequence != gotchi.Sequence.Sequence {
		return c.Status(fiber.StatusBadRequest).JSON(map[string]interface{}{
			"error": "Please enter the correct Sequence",
		})
	}

	gotchi.Sequence.Expires = time.Now()
	gotchi.Verified = true

	err = gotchi.Save(s.DB)

	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(map[string]interface{}{
			"error": "Failed to save Gotchi",
		})
	}

	return c.Status(fiber.StatusOK).JSON(map[string]interface{}{
		"result": "success",
		"data":   gotchi,
	})
}
