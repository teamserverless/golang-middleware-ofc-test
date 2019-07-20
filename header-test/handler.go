package function

import (
	"fmt"
	"net/http"
)

func Handle(w http.ResponseWriter, r *http.Request) {

	if r.Body != nil {
		defer r.Body.Close()
	}

	headers := ""

	for k, v := range r.Header {
		headers = headers + fmt.Sprintf("%s=%s\n", k, v)
	}

	w.WriteHeader(http.StatusOK)
	w.Write([]byte(headers))
}
