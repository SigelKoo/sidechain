package main

import (
	"bufio"
	"encoding/csv"
	"fmt"
	"io"
	"log"
	"os"
	"strconv"
)

type etherTransfer struct {
	From  string  `json:"from"`
	To    string  `json:"to"`
	Value float64 `json:"value"`
}

type mapTransfer struct {
	From string
	To   string
}

func main() {
	csvFile, err := os.Open("E:\\awesomeProject\\sidechain\\bench\\ethereumAccount1.csv")
	if err != nil {
		log.Fatal(err)
	}
	reader := csv.NewReader(bufio.NewReader(csvFile))
	var transfers []etherTransfer
	result := make(map[mapTransfer]float64)

	for {
		line, error := reader.Read()
		if error == io.EOF {
			break
		} else if error != nil {
			log.Fatal(error)
		}
		valuein, _ := strconv.ParseFloat(line[7], 64)
		valueout, _ := strconv.ParseFloat(line[8], 64)
		transfers = append(transfers, etherTransfer{
			From:  line[4],
			To:    line[5],
			Value: valuein + valueout,
		})
	}
	for i := 0; i < len(transfers); i++ {
		transfer := mapTransfer{transfers[i].From, transfers[i].To}
		result[transfer] += transfers[i].Value
	}
	fmt.Println(result)
}
