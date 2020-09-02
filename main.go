package main

import (
	"encoding/csv"
	"fmt"
	"os"
	"strconv"
)

var (
	punkty    int = 0
	zadanie   int = 0
	odpowiedz int
	licznik   int = 1
)

func main() {

	f, err := os.Open("problems.csv") // otwiera plik
	if err != nil {
		return
	}
	// reader := bufio.NewReader(os.Stdin)
	r := csv.NewReader(f)
	for {
		line, err := r.Read() // czyta linie
		if err != nil {
			fmt.Printf("Twój wynik to: %d / %d \n", punkty, zadanie)
			break
		}
		zadanie++
		fmt.Printf("Problem #%d: %s = ", licznik, line[0])
		licznik++
		fmt.Scan(&odpowiedz)
		i, _ := strconv.Atoi(line[1])
		if odpowiedz == i {
			punkty++
		}
	}
}
