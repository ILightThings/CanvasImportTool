package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"
)

func main() {
	file, err := os.Open("users.csv")
	if err != nil {
		log.Fatal(err)
	}
	scanner := bufio.NewScanner(file)

	var rawline []string
	for scanner.Scan() {
		rawline = append(rawline, scanner.Text())
	}

	var courses []string
	var students []string
	var row = 0
	for _, eachln := range rawline {
		row++
		s := strings.SplitN(eachln, ",", 2)
		if row == 1 {
			continue
		}
		if s[0] != "" {
			courses = append(courses, s[0])
		}
		if s[1] != ",," {
			students = append(students, s[1])
		}

	}

	newfile, newfileError := os.Create("importme.csv")
	if newfileError != nil {
		log.Fatal(newfileError)
		return
	}

	_, _ = newfile.WriteString("course_id,user_id,role,status\n")

	for _, x := range courses {
		for _, y := range students {
			linetowrite := fmt.Sprintf("%s,%s \n", x, y)
			_, _ = newfile.WriteString(linetowrite)
		}
	}
	newfile.Close()
	fmt.Println("done")

}
