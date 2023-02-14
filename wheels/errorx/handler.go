package errorx

import (
	"encoding/json"
	"log"
	"net/http"

	"google.golang.org/grpc/status"
)

func ErrorHandler(r *http.Request, w http.ResponseWriter, resp interface{}, err error) {
	if err == nil {
		writeJSON(w, http.StatusOK, resp)

		return
	}

	var respErr error
	var c int

	sts, ok := status.FromError(err)
	if ok {
		c = int(HTTPStatusFromCode(sts.Code()))
		respErr = &Err{
			CodeF:    c,
			StatusF:  sts.Code().String(),
			MessageF: sts.Message(),
		}

		writeJSON(w, c, struct {
			Error any `json:"error"`
		}{
			Error: respErr,
		})

		return
	}

	if e, ok := err.(*Err); ok {
		c = e.Code()
		respErr = &Err{CodeF: e.Code(), StatusF: e.Status(), MessageF: e.Message()}
	} else {
		respErr = &Err{CodeF: http.StatusInternalServerError, StatusF: http.StatusText(http.StatusInternalServerError), MessageF: err.Error()}
	}

	writeJSON(w, c, struct {
		Error any `json:"error"`
	}{
		Error: respErr,
	})

	return
}

func writeJSON(w http.ResponseWriter, c int, v any) {
	bytes, err := json.Marshal(v)
	if err != nil {
		http.Error(w, err.Error(), c)

		return
	}

	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	w.WriteHeader(http.StatusOK)

	if n, err := w.Write(bytes); err != nil {
		if err != http.ErrHandlerTimeout {
			log.Printf("write response failed, error: %s", err)
		}
	} else if n < len(bytes) {
		log.Printf("actual bytes: %d, written bytes: %d", len(bytes), n)
	}

}
