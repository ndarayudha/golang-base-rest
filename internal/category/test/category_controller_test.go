// package test

// import (
// 	"context"
// 	"database/sql"
// 	"encoding/json"
// 	"fmt"
// 	"golang_restfull_api/config"
// 	"golang_restfull_api/internal/category/controller"
// 	"golang_restfull_api/internal/category/delivery"
// 	"golang_restfull_api/internal/category/repository"
// 	"golang_restfull_api/internal/category/service"
// 	"golang_restfull_api/internal/model"
// 	"golang_restfull_api/pkg/logger"
// 	"golang_restfull_api/pkg/utils"
// 	"log"
// 	"os"

// 	"io"
// 	"net/http"
// 	"net/http/httptest"
// 	"strconv"
// 	"strings"
// 	"testing"
// 	"time"

// 	"github.com/go-playground/validator"
// 	_ "github.com/go-sql-driver/mysql"
// 	"github.com/stretchr/testify/assert"
// )

// func setupRouter(db *sql.DB) http.Handler {
// 	configPath := utils.GetConfigPath(os.Getenv("config"))
// 	configFile, err := config.LoadConfig(configPath)
// 	if err != nil {
// 		log.Fatalf("LoadConfig: %v", err)
// 	}
// 	config, err := config.ParseConfig(configFile)
// 	if err != nil {
// 		log.Fatalf("ParseConfig: %v", err)
// 	}

// 	appLogger := logger.NewApiLogger(config)
	
// 	validate := validator.New()
// 	categoryRepository := repository.NewCategoryRepository()
// 	categoryService := service.NewCategoryService(categoryRepository, db, validate)
// 	categoryController := controller.NewCategoryController(categoryService)
// 	router := delivery.NewCategoryRoutes(categoryController, appLogger)

// 	return router
// }

// func setupTestDB() *sql.DB {
// 	db, err := sql.Open("mysql", "root@tcp(localhost:3306)/golang-restfull")
// 	utils.PanicIfError(err)

// 	db.SetMaxIdleConns(5)
// 	db.SetMaxOpenConns(20)
// 	db.SetConnMaxLifetime(60 * time.Minute)
// 	db.SetConnMaxIdleTime(10 * time.Minute)

// 	return db
// }

// func truncateCategory(db *sql.DB) {
// 	db.Exec("TRUNCATE category")
// }

// func TestCreateCategorySuccess(t *testing.T) {
// 	db := setupTestDB()
// 	truncateCategory(db)
// 	router := setupRouter(db)

// 	requestBody := strings.NewReader(`{"name" : "Gadget"}`)
// 	request := httptest.NewRequest(http.MethodPost, "http://localhost:3000/api/categories", requestBody)
// 	request.Header.Add("Content-Type", "application/json")
// 	request.Header.Add("X-API-Key", "SECRET")

// 	recorder := httptest.NewRecorder()

// 	router.ServeHTTP(recorder, request)

// 	response := recorder.Result()
// 	assert.Equal(t, 200, response.StatusCode)

// 	body, _ := io.ReadAll(response.Body)
// 	var responseBody map[string]interface{}
// 	json.Unmarshal(body, &responseBody)

// 	assert.Equal(t, 200, int(responseBody["code"].(float64)))
// 	assert.Equal(t, "OK", responseBody["status"])
// 	assert.Equal(t, "Gadget", responseBody["data"].(map[string]interface{})["name"])
// }

// func TestCreateCategoryFailed(t *testing.T) {
// 	db := setupTestDB()
// 	truncateCategory(db)
// 	router := setupRouter(db)

// 	requestBody := strings.NewReader(`{"name" : ""}`)
// 	request := httptest.NewRequest(http.MethodPost, "http://localhost:3000/api/categories", requestBody)
// 	request.Header.Add("Content-Type", "application/json")
// 	request.Header.Add("X-API-Key", "SECRET")

// 	recorder := httptest.NewRecorder()

// 	router.ServeHTTP(recorder, request)

// 	response := recorder.Result()
// 	assert.Equal(t, 400, response.StatusCode)

// 	body, _ := io.ReadAll(response.Body)
// 	var responseBody map[string]interface{}
// 	json.Unmarshal(body, &responseBody)

// 	assert.Equal(t, 400, int(responseBody["code"].(float64)))
// 	assert.Equal(t, "BAD_REQUEST", responseBody["status"])
// }

// func TestUpdateCategorySuccess(t *testing.T) {
// 	db := setupTestDB()
// 	truncateCategory(db)

// 	tx, _ := db.Begin()
// 	categoryRepository := repository.NewCategoryRepository()
// 	category := categoryRepository.Save(context.Background(), tx, model.Category{
// 		Name: "Gadget",
// 	})
// 	tx.Commit()

// 	router := setupRouter(db)

// 	requestBody := strings.NewReader(`{"name" : "Gadget"}`)
// 	request := httptest.NewRequest(http.MethodPut, "http://localhost:3000/api/categories/"+strconv.Itoa(category.Id), requestBody)
// 	request.Header.Add("Content-Type", "application/json")
// 	request.Header.Add("X-API-Key", "SECRET")

// 	recorder := httptest.NewRecorder()

// 	router.ServeHTTP(recorder, request)

// 	response := recorder.Result()
// 	assert.Equal(t, 200, response.StatusCode)

// 	body, _ := io.ReadAll(response.Body)
// 	var responseBody map[string]interface{}
// 	json.Unmarshal(body, &responseBody)

// 	assert.Equal(t, 200, int(responseBody["code"].(float64)))
// 	assert.Equal(t, "OK", responseBody["status"])
// 	assert.Equal(t, category.Id, int(responseBody["data"].(map[string]interface{})["id"].(float64)))
// 	assert.Equal(t, "Gadget", responseBody["data"].(map[string]interface{})["name"])
// }

// func TestUpdateCategoryFailed(t *testing.T) {
// 	db := setupTestDB()
// 	truncateCategory(db)

// 	tx, _ := db.Begin()
// 	categoryRepository := repository.NewCategoryRepository()
// 	category := categoryRepository.Save(context.Background(), tx, model.Category{
// 		Name: "Gadget",
// 	})
// 	tx.Commit()

// 	router := setupRouter(db)

// 	requestBody := strings.NewReader(`{"name" : ""}`)
// 	request := httptest.NewRequest(http.MethodPut, "http://localhost:3000/api/categories/"+strconv.Itoa(category.Id), requestBody)
// 	request.Header.Add("Content-Type", "application/json")
// 	request.Header.Add("X-API-Key", "SECRET")

// 	recorder := httptest.NewRecorder()

// 	router.ServeHTTP(recorder, request)

// 	response := recorder.Result()
// 	assert.Equal(t, 400, response.StatusCode)

// 	body, _ := io.ReadAll(response.Body)
// 	var responseBody map[string]interface{}
// 	json.Unmarshal(body, &responseBody)

// 	assert.Equal(t, 400, int(responseBody["code"].(float64)))
// 	assert.Equal(t, "BAD_REQUEST", responseBody["status"])
// }

// func TestGetCategorySuccess(t *testing.T) {
// 	db := setupTestDB()
// 	truncateCategory(db)

// 	tx, _ := db.Begin()
// 	categoryRepository := repository.NewCategoryRepository()
// 	category := categoryRepository.Save(context.Background(), tx, model.Category{
// 		Name: "Gadget",
// 	})
// 	tx.Commit()

// 	router := setupRouter(db)

// 	request := httptest.NewRequest(http.MethodGet, "http://localhost:3000/api/categories/"+strconv.Itoa(category.Id), nil)
// 	request.Header.Add("X-API-Key", "SECRET")

// 	recorder := httptest.NewRecorder()

// 	router.ServeHTTP(recorder, request)

// 	response := recorder.Result()
// 	assert.Equal(t, 200, response.StatusCode)

// 	body, _ := io.ReadAll(response.Body)
// 	var responseBody map[string]interface{}
// 	json.Unmarshal(body, &responseBody)

// 	assert.Equal(t, 200, int(responseBody["code"].(float64)))
// 	assert.Equal(t, "OK", responseBody["status"])
// 	assert.Equal(t, category.Id, int(responseBody["data"].(map[string]interface{})["id"].(float64)))
// 	assert.Equal(t, category.Name, responseBody["data"].(map[string]interface{})["name"])
// }

// func TestGetCategoryFailed(t *testing.T) {
// 	db := setupTestDB()
// 	truncateCategory(db)
// 	router := setupRouter(db)

// 	request := httptest.NewRequest(http.MethodGet, "http://localhost:3000/api/categories/404", nil)
// 	request.Header.Add("X-API-Key", "SECRET")

// 	recorder := httptest.NewRecorder()

// 	router.ServeHTTP(recorder, request)

// 	response := recorder.Result()
// 	assert.Equal(t, 404, response.StatusCode)

// 	body, _ := io.ReadAll(response.Body)
// 	var responseBody map[string]interface{}
// 	json.Unmarshal(body, &responseBody)

// 	assert.Equal(t, 404, int(responseBody["code"].(float64)))
// 	assert.Equal(t, "NOT_FOUND", responseBody["status"])
// }

// func TestDeleteCategorySuccess(t *testing.T) {
// 	db := setupTestDB()
// 	truncateCategory(db)

// 	tx, _ := db.Begin()
// 	categoryRepository := repository.NewCategoryRepository()
// 	category := categoryRepository.Save(context.Background(), tx, model.Category{
// 		Name: "Gadget",
// 	})
// 	tx.Commit()

// 	router := setupRouter(db)

// 	request := httptest.NewRequest(http.MethodDelete, "http://localhost:3000/api/categories/"+strconv.Itoa(category.Id), nil)
// 	request.Header.Add("Content-Type", "application/json")
// 	request.Header.Add("X-API-Key", "SECRET")

// 	recorder := httptest.NewRecorder()

// 	router.ServeHTTP(recorder, request)

// 	response := recorder.Result()
// 	assert.Equal(t, 200, response.StatusCode)

// 	body, _ := io.ReadAll(response.Body)
// 	var responseBody map[string]interface{}
// 	json.Unmarshal(body, &responseBody)

// 	assert.Equal(t, 200, int(responseBody["code"].(float64)))
// 	assert.Equal(t, "OK", responseBody["status"])
// }

// func TestDeleteCategoryFailed(t *testing.T) {
// 	db := setupTestDB()
// 	truncateCategory(db)
// 	router := setupRouter(db)

// 	request := httptest.NewRequest(http.MethodDelete, "http://localhost:3000/api/categories/404", nil)
// 	request.Header.Add("Content-Type", "application/json")
// 	request.Header.Add("X-API-Key", "SECRET")

// 	recorder := httptest.NewRecorder()

// 	router.ServeHTTP(recorder, request)

// 	response := recorder.Result()
// 	assert.Equal(t, 404, response.StatusCode)

// 	body, _ := io.ReadAll(response.Body)
// 	var responseBody map[string]interface{}
// 	json.Unmarshal(body, &responseBody)

// 	assert.Equal(t, 404, int(responseBody["code"].(float64)))
// 	assert.Equal(t, "NOT_FOUND", responseBody["status"])
// }

// func TestListCategoriesSuccess(t *testing.T) {
// 	db := setupTestDB()
// 	truncateCategory(db)

// 	tx, _ := db.Begin()
// 	categoryRepository := repository.NewCategoryRepository()
// 	category1 := categoryRepository.Save(context.Background(), tx, model.Category{
// 		Name: "Gadget",
// 	})
// 	category2 := categoryRepository.Save(context.Background(), tx, model.Category{
// 		Name: "Computer",
// 	})
// 	tx.Commit()

// 	router := setupRouter(db)

// 	request := httptest.NewRequest(http.MethodGet, "http://localhost:3000/api/categories", nil)
// 	request.Header.Add("X-API-Key", "SECRET")

// 	recorder := httptest.NewRecorder()

// 	router.ServeHTTP(recorder, request)

// 	response := recorder.Result()
// 	assert.Equal(t, 200, response.StatusCode)

// 	body, _ := io.ReadAll(response.Body)
// 	var responseBody map[string]interface{}
// 	json.Unmarshal(body, &responseBody)

// 	assert.Equal(t, 200, int(responseBody["code"].(float64)))
// 	assert.Equal(t, "OK", responseBody["status"])

// 	fmt.Println(responseBody)

// 	var categories = responseBody["data"].([]interface{})

// 	categoryResponse1 := categories[0].(map[string]interface{})
// 	categoryResponse2 := categories[1].(map[string]interface{})

// 	assert.Equal(t, category1.Id, int(categoryResponse1["id"].(float64)))
// 	assert.Equal(t, category1.Name, categoryResponse1["name"])

// 	assert.Equal(t, category2.Id, int(categoryResponse2["id"].(float64)))
// 	assert.Equal(t, category2.Name, categoryResponse2["name"])
// }

// func TestUnauthorized(t *testing.T) {
// 	db := setupTestDB()
// 	truncateCategory(db)
// 	router := setupRouter(db)

// 	request := httptest.NewRequest(http.MethodGet, "http://localhost:3000/api/categories", nil)
// 	request.Header.Add("X-API-Key", "SALAH")

// 	recorder := httptest.NewRecorder()

// 	router.ServeHTTP(recorder, request)

// 	response := recorder.Result()
// 	assert.Equal(t, 401, response.StatusCode)

// 	body, _ := io.ReadAll(response.Body)
// 	var responseBody map[string]interface{}
// 	json.Unmarshal(body, &responseBody)

// 	assert.Equal(t, 401, int(responseBody["code"].(float64)))
// 	assert.Equal(t, "UNAUTHORIZED", responseBody["status"])
// }