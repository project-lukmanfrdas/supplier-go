package handlers

import (
	"math"
	"net/http"
	"strconv"
	"supplier-app/config"
	"supplier-app/dto"
	"supplier-app/models"

	"github.com/labstack/echo/v4"
)

// CreateContact godoc
// @Summary Tambah contact
// @Description Tambah contact baru untuk supplier tertentu
// @Tags Contacts
// @Accept json
// @Produce json
// @Param contact body dto.ContactRequest true "Data Contact"
// @Success 201 {object} models.Contact
// @Failure 400 {object} map[string]string
// @Failure 500 {object} map[string]string
// @Router /contacts [post]
func CreateContact(c echo.Context) error {
	var req dto.ContactRequest
	if err := c.Bind(&req); err != nil {
		return c.JSON(http.StatusBadRequest, echo.Map{"error": "invalid request"})
	}

	contact := models.Contact{
		Name:        req.Name,
		Phone:       req.Phone,
		Email:       req.Email,
		JobPosition: req.JobPosition,
		Mobile:      req.Mobile,
		SupplierID:  req.SupplierID,
	}

	if err := config.DB.Create(&contact).Error; err != nil {
		return c.JSON(http.StatusInternalServerError, echo.Map{"error": "failed to create contact"})
	}

	return c.JSON(http.StatusCreated, contact)
}

// GetContactsBySupplier godoc
// @Summary List kontak berdasarkan supplier
// @Description Ambil semua contact berdasarkan supplier ID dengan opsi pagination
// @Tags Contacts
// @Produce json
// @Param supplier_id path int true "Supplier ID"
// @Param page query int false "Halaman" default(1)
// @Param limit query int false "Jumlah data per halaman" default(10)
// @Success 200 {object} map[string]interface{}
// @Failure 400 {object} map[string]string
// @Failure 500 {object} map[string]string
// @Router /contacts/supplier/{supplier_id} [get]
func GetContactsBySupplier(c echo.Context) error {
	supplierIDStr := c.Param("supplier_id")
	supplierID, err := strconv.Atoi(supplierIDStr)
	if err != nil {
		return c.JSON(http.StatusBadRequest, echo.Map{"error": "invalid supplier_id"})
	}

	// Pagination
	pageStr := c.QueryParam("page")
	limitStr := c.QueryParam("limit")

	page := 1
	limit := 10
	if p, err := strconv.Atoi(pageStr); err == nil && p > 0 {
		page = p
	}
	if l, err := strconv.Atoi(limitStr); err == nil && l > 0 {
		limit = l
	}
	offset := (page - 1) * limit

	// Ambil data contact berdasarkan supplier_id
	var contacts []models.Contact
	if err := config.DB.Where("supplier_id = ?", supplierID).Limit(limit).Offset(offset).Find(&contacts).Error; err != nil {
		return c.JSON(http.StatusInternalServerError, echo.Map{"error": "failed to get contacts"})
	}

	// Hitung total
	var total int64
	if err := config.DB.Model(&models.Contact{}).Where("supplier_id = ?", supplierID).Count(&total).Error; err != nil {
		return c.JSON(http.StatusInternalServerError, echo.Map{"error": "failed to count contacts"})
	}

	return c.JSON(http.StatusOK, echo.Map{
		"data": contacts,
		"meta": echo.Map{
			"page":       page,
			"limit":      limit,
			"total":      total,
			"total_page": int(math.Ceil(float64(total) / float64(limit))),
		},
	})
}

// DeleteContact godoc
// @Summary Hapus contact
// @Description Hapus contact berdasarkan ID
// @Tags Contacts
// @Param id path int true "ID contact"
// @Success 200 {object} map[string]string
// @Failure 404 {object} map[string]string
// @Failure 500 {object} map[string]string
// @Router /contacts/{id} [delete]
func DeleteContact(c echo.Context) error {
	id := c.Param("id")

	var contact models.Contact
	if err := config.DB.First(&contact, id).Error; err != nil {
		return c.JSON(http.StatusNotFound, echo.Map{"error": "contact not found"})
	}

	if err := config.DB.Delete(&contact).Error; err != nil {
		return c.JSON(http.StatusInternalServerError, echo.Map{"error": "failed to delete contact"})
	}

	return c.JSON(http.StatusOK, echo.Map{"message": "contact deleted"})
}
