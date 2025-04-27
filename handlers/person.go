package handlers

import (
	"net/http"
	"strconv"

	"t3_juniorGo/models"
	"t3_juniorGo/services"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type PersonHandler struct {
	db            *gorm.DB
	enrichmentSvc *services.EnrichmentService
}

func NewPersonHandler(db *gorm.DB) *PersonHandler {
	return &PersonHandler{
		db:            db,
		enrichmentSvc: services.NewEnrichmentService(),
	}
}

func (h *PersonHandler) CreatePerson(c *gin.Context) {
	var input models.PersonInput
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	person := models.Person{
		Name:       input.Name,
		Surname:    input.Surname,
		Patronymic: input.Patronymic,
	}

	// Обогащение данных
	age, err := h.enrichmentSvc.GetAge(input.Name)
	if err == nil {
		person.Age = age
	}

	gender, err := h.enrichmentSvc.GetGender(input.Name)
	if err == nil {
		person.Gender = gender
	}

	nationality, err := h.enrichmentSvc.GetNationality(input.Name)
	if err == nil {
		person.Nationality = nationality
	}

	if err := h.db.Create(&person).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, person)
}

func (h *PersonHandler) GetPeople(c *gin.Context) {
	var people []models.Person
	query := h.db.Model(&models.Person{})

	// Применяем фильтры
	if name := c.Query("name"); name != "" {
		query = query.Where("name ILIKE ?", "%"+name+"%")
	}
	if surname := c.Query("surname"); surname != "" {
		query = query.Where("surname ILIKE ?", "%"+surname+"%")
	}

	// Пагинация
	page, _ := strconv.Atoi(c.DefaultQuery("page", "1"))
	limit, _ := strconv.Atoi(c.DefaultQuery("limit", "10"))
	offset := (page - 1) * limit

	var total int64
	query.Count(&total)

	if err := query.Offset(offset).Limit(limit).Find(&people).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"data":  people,
		"total": total,
		"page":  page,
		"limit": limit,
	})
}

// UpdatePerson обновляет данные человека
// @Summary Обновить данные человека
// @Description Обновляет данные существующего человека
// @Tags people
// @Accept json
// @Produce json
// @Param id path int true "ID человека"
// @Param person body models.PersonInput true "Новые данные человека"
// @Success 200 {object} models.Person
// @Router /people/{id} [put]
func (h *PersonHandler) UpdatePerson(c *gin.Context) {
	id := c.Param("id")
	var input models.PersonInput
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	var person models.Person
	if err := h.db.First(&person, id).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Человек не найден"})
		return
	}

	person.Name = input.Name
	person.Surname = input.Surname
	person.Patronymic = input.Patronymic

	// Обновляем обогащенные данные
	age, err := h.enrichmentSvc.GetAge(input.Name)
	if err == nil {
		person.Age = age
	}

	gender, err := h.enrichmentSvc.GetGender(input.Name)
	if err == nil {
		person.Gender = gender
	}

	nationality, err := h.enrichmentSvc.GetNationality(input.Name)
	if err == nil {
		person.Nationality = nationality
	}

	if err := h.db.Save(&person).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, person)
}

// DeletePerson удаляет человека
// @Summary Удалить человека
// @Description Удаляет человека по ID
// @Tags people
// @Accept json
// @Produce json
// @Param id path int true "ID человека"
// @Success 204 "No Content"
// @Router /people/{id} [delete]
func (h *PersonHandler) DeletePerson(c *gin.Context) {
	id := c.Param("id")
	if err := h.db.Delete(&models.Person{}, id).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.Status(http.StatusNoContent)
}
