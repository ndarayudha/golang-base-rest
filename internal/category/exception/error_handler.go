package exception

import (
	response "rest_base/internal/category/web/response"
	"rest_base/pkg/utils"

	"net/http"

	"github.com/go-playground/validator"
)

func ErrorHandler(writer http.ResponseWriter, request *http.Request, err interface{}) {
	if notFoundError(writer, request, err) {
		return
	}

	if alreadyExistError(writer, request, err) {
		return
	}

	if validationErrors(writer, request, err) {
		return
	}

	internalServerError(writer, request, err)
}

func validationErrors(writer http.ResponseWriter, request *http.Request, err interface{}) bool {
	exception, ok := err.(validator.ValidationErrors)
	if ok {
		writer.Header().Set("Content-Type", "application/json")
		writer.WriteHeader(http.StatusBadRequest)

		webResponse := response.WebResponse{
			Code:   http.StatusBadRequest,
			Status: "BAD_REQUEST",
			Data:   exception.Error(),
		}

		utils.WriteToResponseBody(writer, webResponse)
		return true
	} else {
		return false
	}
}

func notFoundError(writer http.ResponseWriter, request *http.Request, err interface{}) bool {
	exception, ok := err.(NotFoundError)

	if ok {
		writer.Header().Set("Content-Type", "application/json")
		writer.WriteHeader(http.StatusNotFound)

		webResponse := response.WebResponse{
			Code:   http.StatusNotFound,
			Status: "NOT_FOUND",
			Data:   exception.Error,
		}

		utils.WriteToResponseBody(writer, webResponse)
		return true
	} else {
		return false
	}
}

func alreadyExistError(writer http.ResponseWriter, request *http.Request, err interface{}) bool {
	exception, ok := err.(AlredyExistError)
	if ok {
		writer.Header().Set("Content-Type", "application/json")
		writer.WriteHeader(http.StatusUnprocessableEntity)

		webResponse := response.WebResponse{
			Code:   http.StatusUnprocessableEntity,
			Status: "ALREADY_EXIST",
			Data:   exception.Error,
		}

		utils.WriteToResponseBody(writer, webResponse)
		return true
	} else {
		return false
	}
}

func internalServerError(writer http.ResponseWriter, request *http.Request, err interface{}) {
	writer.Header().Set("Content-Type", "application/json")
	writer.WriteHeader(http.StatusInternalServerError)

	webResponse := response.WebResponse{
		Code:   http.StatusInternalServerError,
		Status: "INTERNAL_SERVER_ERROR",
		Data:   err,
	}

	utils.WriteToResponseBody(writer, webResponse)
}
