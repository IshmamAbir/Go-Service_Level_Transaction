package response

import (
	"encoding/json"
	"net/http"
	"strings"

	CommonError "main.go/errors"
	util "main.go/utils"
)

type JsonResponse struct {
	ErrorMessage string      `json:"error_message"`
	Code         int         `json:"code"`
	Result       interface{} `json:"result"`
}

func Render(w http.ResponseWriter, err error, response interface{}) {
	jsonResponse := &JsonResponse{}
	if err != nil {
		if strings.Contains(err.Error(), CommonError.ErrNotFound.Error()) {
			jsonResponse.Code = http.StatusNotFound
			jsonResponse.ErrorMessage = err.Error()
			jsonResponse.Result = nil
			json.NewEncoder(w).Encode(jsonResponse)
		} else if strings.Contains(err.Error(), CommonError.ErrInternalServerError.Error()) {
			jsonResponse.Code = http.StatusInternalServerError
			jsonResponse.ErrorMessage = err.Error()
			jsonResponse.Result = nil
			json.NewEncoder(w).Encode(jsonResponse)
		} else if strings.Contains(err.Error(), CommonError.ErrConflict.Error()) {
			jsonResponse.Code = http.StatusConflict
			jsonResponse.ErrorMessage = err.Error()
			jsonResponse.Result = nil
			json.NewEncoder(w).Encode(jsonResponse)
		} else if strings.Contains(err.Error(), CommonError.ErrBadParamInput.Error()) {
			jsonResponse.Code = http.StatusBadRequest
			jsonResponse.ErrorMessage = err.Error()
			jsonResponse.Result = nil
			json.NewEncoder(w).Encode(jsonResponse)
		} else if strings.Contains(err.Error(), CommonError.ErrPasswordMismatched.Error()) {
			jsonResponse.Code = http.StatusUnauthorized
			jsonResponse.ErrorMessage = err.Error()
			jsonResponse.Result = nil
			json.NewEncoder(w).Encode(jsonResponse)
		} else if strings.Contains(err.Error(), CommonError.ErrBadReusePassword.Error()) {
			jsonResponse.Code = util.HTTP_REUSE_PASS
			jsonResponse.ErrorMessage = err.Error()
			jsonResponse.Result = nil
			json.NewEncoder(w).Encode(jsonResponse)
		} else if strings.Contains(err.Error(), CommonError.ErrFileSizeLimit.Error()) {
			jsonResponse.Code = util.HTTP_FILESIZE_LIMIT
			jsonResponse.ErrorMessage = err.Error()
			jsonResponse.Result = nil
			json.NewEncoder(w).Encode(jsonResponse)
		} else if strings.Contains(err.Error(), CommonError.ErrSingnin.Error()) ||
			strings.Contains(err.Error(), CommonError.ErrNotFoundCookie.Error()) {
			jsonResponse.Code = http.StatusUnauthorized
			jsonResponse.ErrorMessage = err.Error()
			jsonResponse.Result = nil
			json.NewEncoder(w).Encode(jsonResponse)
		} else if strings.Contains(err.Error(), CommonError.ErrAcountLocked.Error()) {
			jsonResponse.Code = http.StatusLocked
			jsonResponse.ErrorMessage = err.Error()
			jsonResponse.Result = nil
			json.NewEncoder(w).Encode(jsonResponse)
		}

	} else {
		jsonResponse.Code = http.StatusOK
		jsonResponse.ErrorMessage = ""
		jsonResponse.Result = response
		json.NewEncoder(w).Encode(jsonResponse)
	}
}
