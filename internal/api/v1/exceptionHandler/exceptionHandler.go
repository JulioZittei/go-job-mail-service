package exceptionhandler

import (
	"net/http"

	internalerrors "github.com/JulioZittei/go-job-mail-service/internal/domain/internalErrors"
	"github.com/go-chi/render"
)


type ControllerFunc func(w http.ResponseWriter, r *http.Request) (interface{}, int, error )


func ExceptionHandler(controllerFunc ControllerFunc) http.HandlerFunc {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		obj, status, err := controllerFunc(w, r)

		if err != nil {
			errorResponse := catchError(r, err)
			render.Status(r, errorResponse.Code)
			render.JSON(w, r, errorResponse)
			return
		}
		render.Status(r, status)
		if obj != nil {
			render.JSON(w, r, obj)
		}
	})
}

func catchError(r *http.Request, err error) *internalerrors.ErrorResponse {
	switch err := err.(type) {
		case *internalerrors.ErrValidation:
			return buildErrValidationResponse(r, err)
		case *internalerrors.ErrInternal:
			return buildErrInternalResponse(r, err)
		case *internalerrors.ErrBadRequest: 
			return buildBadRequestResponse(r, err)
		default:
			return buildDefaultErrorResponse(r, err)
	}
}

func buildErrValidationResponse(r *http.Request, err *internalerrors.ErrValidation) *internalerrors.ErrorResponse {
	return &internalerrors.ErrorResponse{
		Code: err.StatusCode,
		Status: http.StatusText(err.StatusCode),
		Title: err.Title,
		Detail: err.Detail,
		Instance: r.RequestURI,
		InvalidParams: err.ErrorsParam,
	}
}

func buildErrInternalResponse(r *http.Request, err *internalerrors.ErrInternal) *internalerrors.ErrorResponse {
	return &internalerrors.ErrorResponse{
		Code: err.StatusCode,
		Status: http.StatusText(err.StatusCode),
		Title: err.Title,
		Detail: err.Detail,
		Instance: r.RequestURI,
		InvalidParams: []internalerrors.ErrorsParam{},
	}
}

func buildBadRequestResponse(r *http.Request, err *internalerrors.ErrBadRequest) *internalerrors.ErrorResponse {
	return &internalerrors.ErrorResponse{
		Code: err.StatusCode,
		Status: http.StatusText(err.StatusCode),
		Title: err.Title,
		Detail: err.Detail,
		Instance: r.RequestURI,
		InvalidParams: []internalerrors.ErrorsParam{},
	}
}

func buildDefaultErrorResponse(r *http.Request, err error) *internalerrors.ErrorResponse {
	return &internalerrors.ErrorResponse{
		Code: http.StatusInternalServerError,
		Status: http.StatusText(http.StatusInternalServerError),
		Title: "unmaped error",
		Detail: "contact support to report the problem",
		Instance: r.RequestURI,
		InvalidParams: []internalerrors.ErrorsParam{},
	}
}