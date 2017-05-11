package main

import (
	"bufio"
	"encoding/json"
	"net/http"
	"os"
	"strconv"
)

type historyHandler struct {
	filename string
}

func (h historyHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case "GET":
		size := 5
		if newSize, err := strconv.Atoi(r.FormValue("size")); err == nil && newSize > 0 {
			size = newSize
		}
		f, err := os.Open(h.filename)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			return
		}
		defer f.Close()
		scanner := bufio.NewScanner(f)
		var entries []string
		for scanner.Scan() {
			entries = append(entries, scanner.Text())
		}
		index := 0
		if len(entries) > size {
			index = len(entries) - size
		}
		if err := json.NewEncoder(w).Encode(entries[index:]); err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			return
		}
	default:
		w.WriteHeader(http.StatusMethodNotAllowed)
	}
}

func HistoryHandlerFromEnv() historyHandler {
	return historyHandler{"moneytracker.csv"}
}
