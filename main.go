package main

import (
	"encoding/csv"
	"flag"
	"fmt"
	"os"
	"strconv"
	"time"
)

var (
	punkty    int = 0
	zadanie   int = 0
	odpowiedz int
	licznik   int = 1
	timeend   int = 0
)

func main() {
	timeLimit := flag.Int("limit", 5, "the time limit for the quiz in seconds")
	flag.Parse()
	timer := time.NewTimer(time.Duration(*timeLimit) * time.Second)
	go timerere(*timer)
	quiz()

}

func timerere(value time.Timer) {
	<-value.C
	timeend = 1
}
func quiz() {

	f, err := os.Open("problems.csv") // otwiera plik
	if err != nil {
		return
	}
	// reader := bufio.NewReader(os.Stdin)
	r := csv.NewReader(f)
	for {
		if timeend == 0 {
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
			continue
		} else {
			break
		}

	}
	fmt.Println()
	fmt.Println("Czas się skończył")
	fmt.Printf("Twój wynik to: %d / %d \n", punkty, zadanie)
}
