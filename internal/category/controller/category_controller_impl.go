package controller

import (
	"net/http"
	"rest_base/internal/category"

	dto "rest_base/internal/category/web/dto"
	response "rest_base/internal/category/web/response"
	"rest_base/pkg/utils"
	"strconv"

	"github.com/julienschmidt/httprouter"
)

type CategoryControllerImpl struct {
	CategoryService category.CategoryService
}

func NewCategoryController(categoryService category.CategoryService) *CategoryControllerImpl {
	return &CategoryControllerImpl{
		CategoryService: categoryService,
	}
}

func (controller *CategoryControllerImpl) Create(writer http.ResponseWriter, request *http.Request, params httprouter.Params) {
	categoryCreateRequest := dto.CategoryCreateRequest{}

	utils.ReadFromRequestBody(request, &categoryCreateRequest)

	categoryResponse := controller.CategoryService.Create(request.Context(), categoryCreateRequest)
	webResponse := response.WebResponse{
		Code:   201,
		Status: "OK",
		Data:   categoryResponse,
	}

	utils.WriteToResponseBody(writer, webResponse)
}

func (controller *CategoryControllerImpl) Update(writer http.ResponseWriter, request *http.Request, params httprouter.Params) {
	categoryUpdateRequest := dto.CategoryUpdateRequest{}
	utils.ReadFromRequestBody(request, &categoryUpdateRequest)

	categoryId := params.ByName("id")
	id, err := strconv.Atoi(categoryId)
	utils.PanicIfError(err)

	categoryUpdateRequest.Id = id

	categoryResponse := controller.CategoryService.Update(request.Context(), categoryUpdateRequest)
	webResponse := response.WebResponse{
		Code:   200,
		Status: "OK",
		Data:   categoryResponse,
	}

	utils.WriteToResponseBody(writer, webResponse)
}

func (controller *CategoryControllerImpl) Delete(writer http.ResponseWriter, request *http.Request, params httprouter.Params) {
	categoryId := params.ByName("id")
	id, err := strconv.Atoi(categoryId)
	utils.PanicIfError(err)

	controller.CategoryService.Delete(request.Context(), id)
	webResponse := response.WebResponse{
		Code:   200,
		Status: "OK",
	}

	utils.WriteToResponseBody(writer, webResponse)
}

func (controller *CategoryControllerImpl) FindById(writer http.ResponseWriter, request *http.Request, params httprouter.Params) {
	categoryId := params.ByName("id")
	id, err := strconv.Atoi(categoryId)
	utils.PanicIfError(err)

	categoryResponse := controller.CategoryService.FindById(request.Context(), id)
	webResponse := response.WebResponse{
		Code:   200,
		Status: "OK",
		Data:   categoryResponse,
	}

	utils.WriteToResponseBody(writer, webResponse)
}

func (controller *CategoryControllerImpl) FindByName(writer http.ResponseWriter, request *http.Request, params httprouter.Params) {
	name := params.ByName("name")

	categoryResponse := controller.CategoryService.FindByName(request.Context(), name)
	webResponse := response.WebResponse{
		Code:   200,
		Status: "OK",
		Data:   categoryResponse,
	}

	utils.WriteToResponseBody(writer, webResponse)
}

func (controller *CategoryControllerImpl) FindByName2(writer http.ResponseWriter, request *http.Request, params httprouter.Params) {
	name := params.ByName("yok")

	categoryResponse := controller.CategoryService.FindByName(request.Context(), name)
	webResponse := response.WebResponse{
		Code:   200,
		Status: "OK",
		Data:   categoryResponse,
	}

	utils.WriteToResponseBody(writer, webResponse)
}

func (controller *CategoryControllerImpl) FindAll(writer http.ResponseWriter, request *http.Request, params httprouter.Params) {
	categoryResponses := controller.CategoryService.FindAll(request.Context())
	webResponse := response.WebResponse{
		Code:   200,
		Status: "OK",
		Data:   categoryResponses,
	}

	utils.WriteToResponseBody(writer, webResponse)
}
