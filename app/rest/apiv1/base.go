package apiv1

import (
	"github/robotxt/propertyzag/app/rest"
	"net/http"
)

type baseResponse struct {
	Message string
}

// BaseAPI
func (api *ApiV1) BaseAPI(w http.ResponseWriter, r *http.Request) {
	rest.RespondJSON(w, http.StatusOK, &baseResponse{
		Message: "Hello World!",
	})
}
