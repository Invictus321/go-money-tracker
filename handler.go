package main

import (
	"net/http"
	"strconv"
)

type handler struct {
	writer Writer
}

func (h handler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	amount, err := strconv.ParseInt(r.FormValue("amount"), 10, 64)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	if amount > 10000 {
		w.WriteHeader(http.StatusForbidden)
		return
	}
	if err := h.writer.WriteAmount(amount); err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusOK)
}

func HandlerFromEnv() handler {
	return handler{
		WriterFromEnv(),
	}
}
