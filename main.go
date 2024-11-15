package main

import (
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
)

type AutoPart struct {
	ID           string    `json:"id"`
	Category     string    `json:"category"`
	Name         string    `json:"name"`
	Manufacturer string    `json:"manufacturer"`
	Stock        int       `json:"stock"`
	Price        float64   `json:"price"`
	DateAdded    time.Time `json:"date_added"`
	Description  string    `json:"description"`
}

var autoParts = []AutoPart{
	{ID: "1", Category: "Двигатель", Name: "Масляный фильтр", Manufacturer: "Bosch", Stock: 50, Price: 500.00, DateAdded: time.Now(), Description: "Подходит для большинства легковых автомобилей."},
	{ID: "2", Category: "Тормоза", Name: "Тормозные колодки", Manufacturer: "Brembo", Stock: 20, Price: 1500.00, DateAdded: time.Now(), Description: "Высококачественные тормозные колодки."},
	{ID: "3", Category: "Электроника", Name: "Аккумулятор", Manufacturer: "Varta", Stock: 15, Price: 8000.00, DateAdded: time.Now(), Description: "Надежный автомобильный аккумулятор."},
	{ID: "4", Category: "Подвеска", Name: "Амортизатор", Manufacturer: "KYB", Stock: 10, Price: 4000.00, DateAdded: time.Now(), Description: "Амортизатор для комфортной езды."},
	{ID: "5", Category: "Освещение", Name: "Фара", Manufacturer: "Philips", Stock: 25, Price: 3000.00, DateAdded: time.Now(), Description: "Светодиодная фара высокого качества."},
}

func main() {
	router := gin.Default()

	// Получение списка запчастей
	router.GET("/parts", getAutoParts)

	// Получение запчасти по её ID
	router.GET("/parts/:id", getAutoPartByID)

	// Добавление новой запчасти
	router.POST("/parts", createAutoPart)

	// Обновление параметров существующей запчасти
	router.PUT("/parts/:id", updateAutoPart)

	// Удаление запчасти
	router.DELETE("/parts/:id", deleteAutoPart)

	// Получение количества запчастей
	router.GET("/parts/stats", getAutoPartStatistics)

	// Получение запчастей по категории
	router.GET("/parts/category", getAutoPartsByCategory)

	router.Run(":8080")
}

func getAutoParts(c *gin.Context) {
	c.JSON(http.StatusOK, autoParts)
}

func getAutoPartByID(c *gin.Context) {
	id := c.Param("id")

	for _, part := range autoParts {
		if part.ID == id {
			c.JSON(http.StatusOK, part)
			return
		}
	}

	c.JSON(http.StatusNotFound, gin.H{"message": "Запчасть не найдена"})
}

func createAutoPart(c *gin.Context) {
	var newPart AutoPart

	if err := c.BindJSON(&newPart); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": "Некорректный запрос"})
		return
	}

	autoParts = append(autoParts, newPart)
	c.JSON(http.StatusCreated, newPart)
}

func updateAutoPart(c *gin.Context) {
	id := c.Param("id")
	var updatedPart AutoPart

	if err := c.BindJSON(&updatedPart); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": "Некорректный запрос"})
		return
	}

	for i, part := range autoParts {
		if part.ID == id {
			autoParts[i] = updatedPart
			c.JSON(http.StatusOK, updatedPart)
			return
		}
	}

	c.JSON(http.StatusNotFound, gin.H{"message": "Запчасть не найдена"})
}

func deleteAutoPart(c *gin.Context) {
	id := c.Param("id")

	for i, part := range autoParts {
		if part.ID == id {
			autoParts = append(autoParts[:i], autoParts[i+1:]...)
			c.JSON(http.StatusOK, gin.H{"message": "Запчасть удалена"})
			return
		}
	}

	c.JSON(http.StatusNotFound, gin.H{"message": "Запчасть не найдена"})
}

func getAutoPartStatistics(c *gin.Context) {
	totalParts := len(autoParts)
	c.JSON(http.StatusOK, gin.H{"Всего запчастей": totalParts})
}

func getAutoPartsByCategory(c *gin.Context) {
	category := c.Query("category")
	var filteredParts []AutoPart

	for _, part := range autoParts {
		if part.Category == category {
			filteredParts = append(filteredParts, part)
		}
	}

	if len(filteredParts) > 0 {
		c.JSON(http.StatusOK, filteredParts)
	} else {
		c.JSON(http.StatusNotFound, gin.H{"message": "Запчасти не найдены"})
	}
}
