package utils

import (
	"net/http"
)

type APIError struct {
	Status int
	Msg    string
	Code   string
}

func (e APIError) Error() string {
	return e.Msg
}

func NotFoundError() APIError {
	return APIError{
		Status: http.StatusNotFound,
		Code:   "NOT_FOUND_ERROR",
		Msg:    "Not found",
	}
}

func BadRequestError() APIError {
	return APIError{
		Status: http.StatusBadRequest,
		Code:   "INVALID_REQUEST_ERROR",
		Msg:    "Invalid request body",
	}
}

func PathParmMissingError() APIError {
	return APIError{
		Status: http.StatusBadRequest,
		Code:   "INVALID_REQUEST_ERROR",
		Msg:    "Path parameter missing",
	}
}

func DBCreateError() APIError {
	return APIError{
		Status: http.StatusInternalServerError,
		Code:   "DB_CREATE_ERROR",
		Msg:    "Unable to create a new TODO",
	}
}

func DBFetchError() APIError {
	return APIError{
		Status: http.StatusInternalServerError,
		Code:   "DB_FETCH_ERROR",
		Msg:    "Unable to fetch TODOs",
	}
}

func DBDeleteError() APIError {
	return APIError{
		Status: http.StatusInternalServerError,
		Code:   "DB_DELETE_ERROR",
		Msg:    "Unable to delete the TODO",
	}
}

func DBUpdateError() APIError {
	return APIError{
		Status: http.StatusInternalServerError,
		Code:   "DB_UPDATE_ERROR",
		Msg:    "Unable to update the TODO",
	}
}
