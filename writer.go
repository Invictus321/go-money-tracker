package main

import (
	"fmt"
	"os"
	"time"
)

type Writer interface {
	WriteAmount(amount int64) error
}

type writer struct {
	filename string
}

var dateFormat = "02/01/2006"

func (w writer) WriteAmount(amount int64) error {
	f, err := os.OpenFile(w.filename, os.O_APPEND|os.O_WRONLY, 0666)
	if err != nil {
		return err
	}

	floatAmount := float32(amount) / 100

	defer f.Close()

	if _, err = f.WriteString(fmt.Sprintf("%s,%.2f\n", time.Now().Format(dateFormat), floatAmount)); err != nil {
		return err
	}
	return nil
}

func WriterFromEnv() Writer {
	return writer{"moneytracker.csv"}
}
