package server

import (
	"backend/internal/models"
	"backend/internal/utils"

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

	_, err := utils.ValidateAuthToken(rBody.Hash, rBody.Token)

	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(map[string]interface{}{
			"error": err.Error(),
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

func (s *FiberServer) GetGotchiHandler(c *fiber.Ctx) error {
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
		return c.Status(fiber.StatusNotFound).JSON(map[string]interface{}{
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

func (s *FiberServer) UpdateGotchiHandler(c *fiber.Ctx) error {
	type requestBody struct {
		Token string `json:"token"`
		Hash  string `json:"hash"`
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

	_, err := utils.ValidateAuthToken(rBody.Hash, rBody.Token)

	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(map[string]interface{}{
			"error": err.Error(),
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
		return c.Status(fiber.StatusNotFound).JSON(map[string]interface{}{
			"error": "Gotchi not found",
		})
	}

	if gotchi.Sequence.ID == uuid.Nil {
		gotchi.Sequence.Sequence = utils.GenerateRandomSequence(5)
	}

	gotchi.Hash = rBody.Hash
	gotchi.AuthToken = rBody.Token

	err = gotchi.Save(s.DB)

	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(map[string]interface{}{
			"error": "Failed to update gotchi",
		})
	}

	return c.Status(fiber.StatusOK).JSON(map[string]interface{}{
		"result": "success",
		"data":   gotchi,
	})
}

func (s *FiberServer) GetAllHandler(c *fiber.Ctx) error {
	gotchis, err := models.GetAllGotchis(s.DB, "id", "name", "level", "verified")

	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(map[string]interface{}{
			"error": "Failed to get all gotchis",
		})
	}

	return c.Status(fiber.StatusOK).JSON(map[string]interface{}{
		"result": "success",
		"data":   gotchis,
	})
}
