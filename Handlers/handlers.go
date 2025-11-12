package Handlers

import (
	"backend/DBAdapter"
	"encoding/json"
	"fmt"
	"strconv"

	"github.com/gofiber/fiber/v2"
	"gorm.io/datatypes"
)

func GetRoadmap(adapter *DBAdapter.Adapter) fiber.Handler {
	return func(c *fiber.Ctx) error {
		id, err := strconv.ParseUint(c.Params("id"), 10, 64)
		if err != nil {
			return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Can't parse id. It must be an integer"})
		}

		var nodes_json []datatypes.JSON
		var roadmap_id uint64
		roadmap_id, nodes_json, err = adapter.GetRoadmapNodesJSON(id)
		if err != nil {
			return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": err.Error()})
		}
		return c.Status(200).JSON(fiber.Map{"roadmap_id": roadmap_id, "parts": nodes_json})
	}
}

func InsertRoadmap(adapter *DBAdapter.Adapter) fiber.Handler {
	return func(c *fiber.Ctx) error {
		id, err := strconv.ParseUint(c.Params("id"), 10, 64)
		if err != nil {
			return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Can't parse id. It must be an integer"})
		}

		var body []datatypes.JSON
		if err := c.BodyParser(&body); err != nil {
			return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
				"error": "Cannot parse JSON array",
			})
		}

		if DEBUG_OUTPUT {
			fmt.Println("=== COMPACT DEBUG ===")
			fmt.Printf("ID: %d | Items: %d\n", id, len(body))
			for i, item := range body {
				var data map[string]interface{}
				if err := json.Unmarshal(item, &data); err == nil {
					fmt.Printf("  [%d]: %+v\n", i, data)
				} else {
					fmt.Printf("  [%d]: %s\n", i, string(item))
				}
			}
		}

		err = adapter.CreateRoadmap(id, body)

		if err != nil {
			return c.Status(400).JSON(fiber.Map{"error": "Something went wrong creating roadmap. Try later"})
		}

		return c.SendStatus(200)
	}
}
