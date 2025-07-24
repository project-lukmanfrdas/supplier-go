package config

import (
	"log"
	"reflect"
	"supplier-app/models"
)

// ModelRegistry menyimpan semua model yang terdaftar
var ModelRegistry = []interface{}{
	&models.Contact{},
	&models.Supplier{},
	// Model baru akan ditambahkan di sini
}

// RegisterModel menambahkan model baru ke registry
func RegisterModel(model interface{}) {
	ModelRegistry = append(ModelRegistry, model)
}

// AutoMigrateAll melakukan migrate semua model yang terdaftar
func AutoMigrateAll() {
	log.Println("Starting auto migration for all models...")

	for _, model := range ModelRegistry {
		modelName := reflect.TypeOf(model).Elem().Name()
		if err := DB.AutoMigrate(model); err != nil {
			log.Printf("❌ Error migrating %s: %v", modelName, err)
		} else {
			log.Printf("✅ Successfully migrated %s", modelName)
		}
	}

	log.Println("Auto migration completed!")
}

// AutoMigrateModels melakukan migrate untuk model tertentu
func AutoMigrateModels(modelsToMigrate ...interface{}) {
	log.Println("Starting migration for specific models...")

	for _, model := range modelsToMigrate {
		modelName := reflect.TypeOf(model).Elem().Name()
		if err := DB.AutoMigrate(model); err != nil {
			log.Printf("❌ Error migrating %s: %v", modelName, err)
		} else {
			log.Printf("✅ Successfully migrated %s", modelName)
		}
	}

	log.Println("Specific models migration completed!")
}

// GetModelNames mengembalikan daftar nama model yang terdaftar
func GetModelNames() []string {
	var names []string
	for _, model := range ModelRegistry {
		names = append(names, reflect.TypeOf(model).Elem().Name())
	}
	return names
}

// PrintRegisteredModels menampilkan semua model yang terdaftar
func PrintRegisteredModels() {
	log.Println("📋 Registered models:")
	for i, model := range ModelRegistry {
		modelName := reflect.TypeOf(model).Elem().Name()
		log.Printf("  %d. %s", i+1, modelName)
	}
}
