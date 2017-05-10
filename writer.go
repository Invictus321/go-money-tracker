package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
	"time"
)

type Writer interface {
	AddThisMonth(month, year int64) (float32, error)
	WriteAmount(amount int64) error
}

type writer struct {
	filename string
}

var dateFormat = "02/01/2006"

func (w writer) AddThisMonth(month, year int64) (float32, error) {
	monthString := fmt.Sprintf("%d/%d", month, year)
	f, err := os.Open(w.filename)
	if err != nil {
		return 0, err
	}
	defer f.Close()

	amount := float32(0.0)

	scanner := bufio.NewScanner(f)
	for scanner.Scan() {
		fields := strings.Split(scanner.Text(), ",")
		if strings.Contains(fields[0], monthString) {
			amountToAdd, err := strconv.ParseFloat(fields[1], 32)
			if err != nil {
				return 0, err
			}
			amount = amount + float32(amountToAdd)
		}
	}
	return amount, nil
}

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
