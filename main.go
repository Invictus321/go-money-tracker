package main

import (
	"net/http"
)

func main() {
	handler := HandlerFromEnv()
	historyHandler := HistoryHandlerFromEnv()

	http.Handle("/history", historyHandler)
	http.Handle("/", handler)

	if err := http.ListenAndServe(":5000", nil); err != nil {
		panic(err)
	}
}
