package main

import (
	"bufio"
	"encoding/json"
	"net/http"
	"os"
)

type historyHandler struct {
	filename string
}

func (h historyHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case "GET":
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
		if len(entries) > 5 {
			index = len(entries) - 5
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
