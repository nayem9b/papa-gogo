package student

import (
	"encoding/json"
	"errors"
	"hello/internal/utils/response"
	"io"

	// "go/types" removed as it's not needed
	"log/slog"
	"net/http"
)

func New() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {

		type Student struct {
			ID   int    `json:"id"`
			Name string `json:"name"`
			Email string `json:"email"`
		}

		var students Student
		err := json.NewDecoder(r.Body).Decode(&students)
		if errors.Is(err, io.EOF){
			response.Writejson(w, http.StatusBadRequest, err.Error())
			return
		}
		// if err != nil {
        //     http.Error(w, err.Error(), http.StatusBadRequest)
        //     return
        // }
		slog.Info("Request received")
		// w.Write([]byte("Welcome to the student API"))
		response.Writejson(w, http.StatusOK, students)
	}

}