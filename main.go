package main

import (
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
)

type Animal struct {
	ID           string    `json:"id"`
	Type         string    `json:"type"`
	Name         string    `json:"name"`
	Location     string    `json:"location"`
	Age          int       `json:"age"`
	Gender       string    `json:"gender"`
	Weight       float64   `json:"weight"`
	DateAdded    time.Time `json:"date_added"`
	HealthStatus string    `json:"health_status"`
	Description  string    `json:"description"`
}

var animals = []Animal{
	{ID: "1", Type: "Обезьяна", Name: "Чичичи", Location: "Вьетнам", Age: 8, Gender: "ж",
		Weight: 20, DateAdded: time.Now(), HealthStatus: "намана", Description: "Ворует кирпичи"},
	{ID: "2", Type: "Мадагаскарский яйценосный дрозд", Name: "BigBalls", Location: "Мадагаскар", Age: 4, Gender: "ж",
		Weight: 0.1, DateAdded: time.Now(), HealthStatus: "намана", Description: `Единственная известная науке птица, 
	не высиживающая яйца и не вьющая гнезд. После завершения беременности, 
	самка откладывает в специальную кожаную сумку два, в редких случаях - три яйца, практически идеальной круглой формы.`},
	{ID: "3", Type: "Собака", Name: "Мухтар", Location: "Россия", Age: 10, Gender: "м",
		Weight: 10, DateAdded: time.Now(), HealthStatus: "намана", Description: "Снимается в кино"},
	{ID: "4", Type: "Кирпич", Name: "Кирпич", Location: "Россия", Age: 100, Gender: "-",
		Weight: 3, DateAdded: time.Now(), HealthStatus: "намана", Description: "Как он сюда попал??"},
	{ID: "5", Type: "Обезьяна", Name: "Чачача", Location: "Израиль", Age: 8, Gender: "ж",
		Weight: 20, DateAdded: time.Now(), HealthStatus: "намана", Description: "Не ворует кирпичи"},
}

func main() {
	router := gin.Default()

	// Получение списка животных
	router.GET("/animals", getAnimals)

	// Получение животного по его ID
	router.GET("/animals/:id", getAnimalByID)

	// Создание нового животного
	router.POST("/animals", createAnimal)

	// Обновление параметров существующего животного
	router.PUT("/animals/:id", updateAnimal)

	// Удаление животного(
	router.DELETE("/animals/:id", deleteAnimal)

	// Получение количества животных
	router.GET("/animals/stats", getAnimalStatistics)

	// Получение животного только по типу
	router.GET("/animals/type", getAnimalsByType)

	router.Run(":8080")
}

func getAnimals(c *gin.Context) {
	c.JSON(http.StatusOK, animals)
}

func getAnimalByID(c *gin.Context) {
	id := c.Param("id")

	for _, animal := range animals {
		if animal.ID == id {
			c.JSON(http.StatusOK, animal)
			return
		}
	}

	c.JSON(http.StatusNotFound, gin.H{"message": "нет такого животного"})
}

func createAnimal(c *gin.Context) {
	var newAnimal Animal

	if err := c.BindJSON(&newAnimal); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": "invalid request"})
		return
	}

	animals = append(animals, newAnimal)
	c.JSON(http.StatusCreated, newAnimal)
}

func updateAnimal(c *gin.Context) {
	id := c.Param("id")
	var updatedAnimal Animal

	if err := c.BindJSON(&updatedAnimal); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": "invalid request"})
		return
	}

	for i, animal := range animals {
		if animal.ID == id {
			animals[i] = updatedAnimal
			c.JSON(http.StatusOK, updatedAnimal)
			return
		}
	}

	c.JSON(http.StatusNotFound, gin.H{"message": "нет такого животного"})
}

func deleteAnimal(c *gin.Context) {
	id := c.Param("id")

	for i, animal := range animals {
		if animal.ID == id {
			animals = append(animals[:i], animals[i+1:]...)
			c.JSON(http.StatusOK, gin.H{"message": "животное удалено("})
			return
		}
	}

	c.JSON(http.StatusNotFound, gin.H{"message": "нет такого животного"})
}

func getAnimalStatistics(c *gin.Context) {
	totalAnimals := len(animals)
	c.JSON(http.StatusOK, gin.H{"всего животных": totalAnimals})
	return
}

func getAnimalsByType(c *gin.Context) {
	animalType := c.Query("type")
	var filteredAnimals []Animal

	for _, animal := range animals {
		if animal.Type == animalType {
			filteredAnimals = append(filteredAnimals, animal)
		}
	}

	if len(filteredAnimals) > 0 {
		c.JSON(http.StatusOK, filteredAnimals)
	} else {
		c.JSON(http.StatusNotFound, gin.H{"message": "животные не найдены"})
	}
	return
}
