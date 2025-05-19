package response

import (
	"encoding/json"
	"net/http"
)

// Response represents a standard JSON response structure.
type Response struct {
	Code    int         `json:"code"`
	Message string      `json:"message"`
	Errors  []Error     `json:"errors,omitempty"`
	Data    interface{} `json:"data,omitempty"`
}

// Error represents a single error detail.
type Error struct {
	Field   string `json:"field"`
	Message string `json:"message"`
}

// Write writes a Response to the http.ResponseWriter.
func Write(w http.ResponseWriter, statusCode int, message string, data interface{}, errs []Error) {
	if message == "" {
		message = defaultMessage(statusCode)
	}
	resp := Response{
		Code:    statusCode,
		Message: message,
		Errors:  errs,
		Data:    data,
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(statusCode)
	json.NewEncoder(w).Encode(resp)
}

// Success returns a 200 OK JSON response with data.
func Success(w http.ResponseWriter, data interface{}, message string) {
	if message == "" {
		message = "Success"
	}
	Write(w, http.StatusOK, message, data, nil)
}

// Created returns a 201 Created JSON response with data.
func Created(w http.ResponseWriter, data interface{}, message string) {
	if message == "" {
		message = "Created"
	}
	Write(w, http.StatusCreated, message, data, nil)
}

// BadRequest returns a 400 Bad Request response with error details.
func BadRequest(w http.ResponseWriter, message string, errs []Error) {
	if message == "" {
		message = "Bad Request"
	}
	Write(w, http.StatusBadRequest, message, nil, errs)
}

// InternalError returns a 500 Internal Server Error response.
func InternalError(w http.ResponseWriter, message string) {
	if message == "" {
		message = "Internal Server Error"
	}
	Write(w, http.StatusInternalServerError, message, nil, nil)
}

// Unauthorized returns a 401 Unauthorized response.
func Unauthorized(w http.ResponseWriter, message string) {
	if message == "" {
		message = "Unauthorized"
	}
	Write(w, http.StatusUnauthorized, message, nil, nil)
}

// Forbidden returns a 403 Forbidden response.
func Forbidden(w http.ResponseWriter, message string) {
	if message == "" {
		message = "Forbidden"
	}
	Write(w, http.StatusForbidden, message, nil, nil)
}

// NotFound returns a 404 Not Found response.
func NotFound(w http.ResponseWriter, message string) {
	if message == "" {
		message = "Not Found"
	}
	Write(w, http.StatusNotFound, message, nil, nil)
}

// defaultMessage maps HTTP status codes to default messages.
func defaultMessage(code int) string {
	switch code {
	case http.StatusOK:
		return "Request successful"
	case http.StatusCreated:
		return "Resource created"
	case http.StatusBadRequest:
		return "Invalid request"
	case http.StatusUnauthorized:
		return "Authentication required"
	case http.StatusForbidden:
		return "Access denied"
	case http.StatusNotFound:
		return "Resource not found"
	case http.StatusInternalServerError:
		return "Something went wrong on the server"
	default:
		return http.StatusText(code)
	}
}
