package router

import (
	"fmt"
	foodninja "goTest/FoodNinja"
	initializers "goTest/Initializers"
	models "goTest/Models"

	"github.com/gofiber/fiber/v2"
	"go.mongodb.org/mongo-driver/bson"
)

func AddNutritionGroup(app *fiber.App) {
	nutritionGroup := app.Group("/nutrition")

	nutritionGroup.Get("/all/:username", GetAllFacts)
	nutritionGroup.Post("/add/:username/:food", AddFood)
	nutritionGroup.Put("/update/:username/:name", UpdateSavedFood)
	nutritionGroup.Delete("/delete/:username/:name", DeleteSavedFood)
}

func GetAllFacts(c *fiber.Ctx) error {
	collection := initializers.GetDBCollection("SavedFoods")
	username := c.Params("username")

	nutritionFacts := make([]models.NutritionFacts, 0)
	cursor, err := collection.Find(c.Context(), bson.M{"username": username})
	if err != nil {
		return c.Status(500).JSON(fiber.Map{
			"error": err.Error(),
		})
	}

	for cursor.Next(c.Context()) {
		item := models.NutritionFacts{}
		err := cursor.Decode(&item)
		if err != nil {
			return c.Status(500).JSON(fiber.Map{
				"error": err.Error(),
			})
		}
		nutritionFacts = append(nutritionFacts, item)
	}

	return c.Status(200).JSON(fiber.Map{
		"data": nutritionFacts,
	})
}

func CheckFacts(c *fiber.Ctx) error {
	food := c.Params("food")

	url := fmt.Sprintf("https://nutrition-by-api-ninjas.p.rapidapi.com/v1/nutrition?query=%s", food)

	if food == "" {
		return c.Status(400).JSON(fiber.Map{
			"error": "food required",
		})
	}
	facts := foodninja.GetNutritionFacts(url)

	return c.Status(200).JSON(fiber.Map{
		"data": facts,
	})

}
func AddFood(c *fiber.Ctx) error {
	food := c.Params("food")
	username := c.Params("username")
	url := fmt.Sprintf("https://nutrition-by-api-ninjas.p.rapidapi.com/v1/nutrition?query=%s", food)

	if food == "" {
		return c.Status(400).JSON(fiber.Map{
			"error": "please enter a food",
		})
	}

	facts := foodninja.GetNutritionFacts(url)
	facts.Username = username

	collection := initializers.GetDBCollection("SavedFoods")

	result, err := collection.InsertOne(c.Context(), facts)
	if err != nil {
		return c.Status(500).JSON(fiber.Map{
			"error":   "failed to add food",
			"message": err.Error(),
		})
	}

	return c.Status(200).JSON(fiber.Map{
		"data": result,
	})

}

func UpdateSavedFood(c *fiber.Ctx) error {
	username := c.Params("username")
	updated := models.UpdatedNutritionFacts{}
	if err := c.BodyParser(updated); err != nil {
		return c.Status(400).JSON(fiber.Map{
			"error": "invalid body",
		})
	}

	name := c.Params("name")
	if name == "" {
		return c.Status(400).JSON(fiber.Map{
			"error": "please enter a name",
		})
	}

	collection := initializers.GetDBCollection("SavedFoods")
	result, err := collection.UpdateOne(c.Context(), bson.M{"name": name, "username": username}, bson.M{"$set": updated})
	if err != nil {
		return c.Status(400).JSON(fiber.Map{
			"error": err.Error(),
		})
	}

	return c.Status(200).JSON(fiber.Map{
		"result": result,
	})
}

func DeleteSavedFood(c *fiber.Ctx) error {
	username := c.Params("username")
	collection := initializers.GetDBCollection("SavedFoods")
	name := c.Params("name")
	if name == "" {
		return c.Status(400).JSON(fiber.Map{
			"error": "name was not given",
		})
	}

	result, err := collection.DeleteOne(c.Context(), bson.M{"name": name, "username": username})
	if err != nil {
		return c.Status(400).JSON(fiber.Map{
			"error": "name does not exist in db",
		})
	}

	return c.Status(200).JSON(fiber.Map{
		"result": result,
	})
}
