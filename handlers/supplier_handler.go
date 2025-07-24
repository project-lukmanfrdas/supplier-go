package handlers

import (
	"math"
	"net/http"
	"path/filepath"
	"strconv"
	"supplier-app/config"
	"supplier-app/models"

	"github.com/labstack/echo/v4"
)

// CreateSupplier godoc
// @Summary Buat supplier baru
// @Description Buat supplier baru dan upload logo
// @Tags Suppliers
// @Accept multipart/form-data
// @Produce json
// @Param supplier_name formData string true "Nama Supplier"
// @Param nick_name formData string false "Nama Panggilan"
// @Param address formData string false "Alamat"
// @Param status formData string false "Status"
// @Param logo formData file true "Logo (gambar)"
// @Success 201 {object} models.Supplier
// @Failure 400 {object} map[string]string
// @Failure 500 {object} map[string]string
// @Router /suppliers [post]
func CreateSupplier(c echo.Context) error {
	supplierName := c.FormValue("supplier_name")
	nickName := c.FormValue("nick_name")
	address := c.FormValue("address")
	status := c.FormValue("status")

	// File upload
	file, err := c.FormFile("logo")
	if err != nil {
		return c.JSON(http.StatusBadRequest, echo.Map{"error": "logo is required"})
	}
	src, err := file.Open()
	if err != nil {
		return c.JSON(http.StatusInternalServerError, echo.Map{"error": "failed to open uploaded file"})
	}
	defer src.Close()

	savePath := filepath.Join("uploads", file.Filename)
	if err := config.SaveFile(src, savePath); err != nil {
		return c.JSON(http.StatusInternalServerError, echo.Map{"error": "failed to save file"})
	}

	supplier := models.Supplier{
		SupplierName: supplierName,
		NickName:     nickName,
		Address:      address,
		Status:       status,
		Logo:         savePath,
	}

	if err := config.DB.Create(&supplier).Error; err != nil {
		return c.JSON(http.StatusInternalServerError, echo.Map{"error": "failed to create supplier"})
	}

	return c.JSON(http.StatusCreated, supplier)
}

// GetSuppliers godoc
// @Summary Ambil semua supplier
// @Description Menampilkan semua data supplier beserta kontak
// @Tags Suppliers
// @Produce json
// @Param page query int false "Halaman" default(1)
// @Param limit query int false "Jumlah data per halaman" default(10)
// @Success 200 {object} map[string]interface{}
// @Failure 500 {object} map[string]string
// @Router /suppliers [get]
func GetSuppliers(c echo.Context) error {
	// Ambil query params
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

	// Query supplier dengan kontak
	var suppliers []models.Supplier
	if err := config.DB.Preload("Contacts").Limit(limit).Offset(offset).Find(&suppliers).Error; err != nil {
		return c.JSON(http.StatusInternalServerError, echo.Map{"error": "failed to get suppliers"})
	}

	// Hitung total supplier
	var total int64
	config.DB.Model(&models.Supplier{}).Count(&total)

	// Return respons JSON
	return c.JSON(http.StatusOK, echo.Map{
		"data": suppliers,
		"meta": echo.Map{
			"page":       page,
			"limit":      limit,
			"total":      total,
			"total_page": int(math.Ceil(float64(total) / float64(limit))),
		},
	})
}
