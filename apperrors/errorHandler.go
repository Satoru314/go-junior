package apperrors

import (
	"encoding/json"
	"errors"
	"log"
	"myapi/common"
	"net/http"
)

func ErrorHandler(w http.ResponseWriter, req *http.Request, err error) {
	var appErr *MyAppError
	if !errors.As(err, &appErr) {
		appErr = &MyAppError{
			ErrCode: Unknown,
			Message: "internal process failed",
			Err:     err,
		}
	}

	traceID := common.GetTraceID(req.Context())
	log.Printf("[%d] %s", traceID, appErr)

	var statusCode int
	switch appErr.ErrCode {
	case NAData:
		statusCode = http.StatusNotFound
	case NoTargetData, ReqBodyDecodeFailed, BadParam, BadPathParam:
		statusCode = http.StatusBadRequest
	case RequiredAuthorizationHeader, Unauthorizated:
		statusCode = http.StatusUnauthorized
	default:
		statusCode = http.StatusInternalServerError
	}
	w.WriteHeader(statusCode)
	json.NewEncoder(w).Encode(appErr)
}
