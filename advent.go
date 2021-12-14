package main

import (
	"io"
	"log"
	"net/http"
	"strconv"
	"strings"
)

func main() {
	client := &http.Client{}
	req, err := http.NewRequest("GET", "https://adventofcode.com/2021/day/3/input", nil)
	resp, err := client.Do(req)
	if err != nil {
		// handle error
	}
	defer resp.Body.Close()
	body, err := io.ReadAll(resp.Body)
	temp := strings.Split(string(body), "\n")

	var size = len(temp[0]);
	var gammaRate string
	var epsilonRate string
	for i := 0; i < size; i++ {
		var positionBits []string
		for _, v := range temp {
			if v != "" {
				positionBits = append(positionBits, string(v[i]));
			}
		}
		dict := make(map[string]int)

		for _, e := range positionBits {
			dict[e] = dict[e] + 1
		}
		log.Print("Dict", dict)
		if dict["0"] < dict["1"] {
			gammaRate = gammaRate + "1"
			epsilonRate = epsilonRate + "0"
		} else {
			gammaRate = gammaRate + "0"
			epsilonRate = epsilonRate + "1"
		}

	}
	log.Print("Gamma ", gammaRate)
	log.Print("Epsilon ", epsilonRate)

	gammaRateInt, err := strconv.ParseInt(gammaRate, 2, 64)
	epsilonRateInt, err := strconv.ParseInt(epsilonRate, 2, 64)

	log.Println("Power consumption", epsilonRateInt * gammaRateInt)
	
}