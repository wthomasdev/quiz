package main

import (
	"bufio"
	"encoding/csv"
	"fmt"
	"io"
	"log"
	"os"
	"strings"
	"time"
)

type Question struct {
	Question string `json:"question"`
	Answer   string `json:"answer"`
}

const timeLimit = 20

func main() {
	questions := readFileAndParse("problems.csv")
	reader := bufio.NewReader(os.Stdin)
	answersCorrect := 0

	timer := time.NewTimer(time.Second * timeLimit)

	for _, v := range questions {
		select {
		case <-timer.C:
			fmt.Printf("Quiz finished, you got %v/%v right. \n", answersCorrect, len(questions))
			return
		default:
			var answer string
			fmt.Printf("What is: %s?    ", v.Question)
			answer, err := reader.ReadString('\n')
			if err != nil {
				fmt.Println("There was an error, reading your answer")
			}
			sanitizedAnswer := strings.TrimSpace(answer)
			if strings.Compare(v.Answer, sanitizedAnswer) == 0 {
				fmt.Println("You are correct.")
				answersCorrect = answersCorrect + 1
			} else {
				fmt.Println("Incorrect.")
			}
		}

	}
	fmt.Printf("Quiz finished, you got %v/%v right. \n", answersCorrect, len(questions))

}

func readFileAndParse(fileName string) []Question {
	csvFile, err := os.Open(fileName)
	if err != nil {
		log.Fatal(err)
	}
	reader := csv.NewReader(bufio.NewReader(csvFile))
	var questions []Question
	for {
		line, error := reader.Read()
		if error == io.EOF {
			break
		} else if error != nil {
			log.Fatal(error)
		}
		q := strings.Join(line[:len(line)-1], ",")
		a := line[len(line)-1]

		questions = append(questions, Question{q, a})
	}
	return questions
}
