package main

import (
	"go_fiber/database"
	"go_fiber/models"
	"log"

	"github.com/gofiber/fiber/v2/middleware/cors"

	"github.com/gofiber/fiber/v2"
)

func main() {
	// Connexion à la base de données
	database.ConnectDB()
	database.DB.AutoMigrate(&models.Task{})

	app := fiber.New()
	app.Use(cors.New(cors.Config{
		AllowOrigins: "*", // Remplace par l'URL de ton frontend
		AllowMethods: "GET,POST,PUT,DELETE",
		AllowHeaders: "Content-Type, Authorization",
	}))
	// Routes CRUD
	app.Get("/tasks", getTasks)
	app.Post("/tasks", createTask)
	app.Get("/tasks/:id", getTask)
	app.Put("/tasks/:id", updateTask)
	app.Delete("/tasks/:id", deleteTask)

	log.Fatal(app.Listen(":9988"))
}

// 📌 Récupérer toutes les tâches
func getTasks(c *fiber.Ctx) error {
	var tasks []models.Task
	database.DB.Find(&tasks)
	return c.JSON(tasks)
}

// 📌 Créer une nouvelle tâche
func createTask(c *fiber.Ctx) error {
	task := new(models.Task)
	if err := c.BodyParser(task); err != nil {
		return c.Status(400).JSON(fiber.Map{"error": "Requête invalide"})
	}
	database.DB.Create(task)
	return c.JSON(task)
}

// 📌 Récupérer une tâche par ID
func getTask(c *fiber.Ctx) error {
	id := c.Params("id")
	var task models.Task
	result := database.DB.First(&task, id)
	if result.Error != nil {
		return c.Status(404).JSON(fiber.Map{"error": "Tâche non trouvée"})
	}
	return c.JSON(task)
}

// 📌 Mettre à jour une tâche
func updateTask(c *fiber.Ctx) error {
	id := c.Params("id")
	var task models.Task
	if database.DB.First(&task, id).Error != nil {
		return c.Status(404).JSON(fiber.Map{"error": "Tâche non trouvée"})
	}

	updateData := new(models.Task)
	if err := c.BodyParser(updateData); err != nil {
		return c.Status(400).JSON(fiber.Map{"error": "Requête invalide"})
	}

	database.DB.Model(&task).Updates(updateData)
	return c.JSON(task)
}

// 📌 Supprimer une tâche
func deleteTask(c *fiber.Ctx) error {
	id := c.Params("id")
	if database.DB.Delete(&models.Task{}, id).RowsAffected == 0 {
		return c.Status(404).JSON(fiber.Map{"error": "Tâche non trouvée"})
	}
	return c.SendStatus(204)
}
