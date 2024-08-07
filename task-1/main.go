package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type GradeReport struct {
	studentName string
	grades      map[string]float64
}

func newGradeReport(name string) GradeReport {
	g := GradeReport{
		studentName: name,
		grades:      map[string]float64{},
	}
	return g
}

func (g *GradeReport) addGrade(course string, grade float64) {
	g.grades[course] = grade
}

func (g *GradeReport) formatReport() string {
	fs := fmt.Sprintf("%-25v%v\n", "Student Name:", g.studentName)

	for course, grade := range g.grades {
		fs += fmt.Sprintf("%-25v ... %v\n", course+":", grade)
	}
	fs += fmt.Sprintf("%-25v ... %v", "Average", g.calculateAverage())
	return fs
}

func (g *GradeReport) calculateAverage() float64 {
	var sum float64
	for _, grade := range g.grades {
		sum += grade
	}
	if len(g.grades) == 0 {
		fmt.Println("Grades are 0")
		return 0.0
	}
	average := sum / float64(len(g.grades))
	return average
}

func main() {
	reader := bufio.NewReader(os.Stdin)
	fmt.Print("Insert Your Name: ")
	name, err := reader.ReadString('\n')
	name = strings.TrimSpace(name)
	if err != nil {
		fmt.Println("Read Op Failed")
		return
	}
	gradeReport := newGradeReport(name)
	fmt.Print("Insert How many Subejects you have:")
	input, err := reader.ReadString('\n')
	if err != nil {
		fmt.Print("Error when reading")
		return
	}
	input = strings.TrimSpace(input)
	number, err := strconv.Atoi(input)
	if err != nil {
		fmt.Print("Error on your input")
		return
	}

	for i := 0; i < number; i++ {
		fmt.Printf("Input your %v's subject and grade separated by space: ", i+1)
		input, err := reader.ReadString('\n')
		if err != nil {
			fmt.Print("Error when reading")
			return
		}
		input = strings.TrimSpace(input)
		parts := strings.Split(input, " ")
		if len(parts) != 2 {
			fmt.Print("Wrong way to input")
			return
		}
		subj, numb := parts[0], parts[1]
		mark, err := strconv.Atoi(numb)
		if err != nil {
			fmt.Print("The number is not correct")
			return
		}
		if mark < 0 || mark > 100 {
			fmt.Print("Mark can't be > 100 or < 0")
			return
		}
		gradeReport.addGrade(subj, float64(mark))

	}
	fmt.Println(gradeReport.formatReport())
}
