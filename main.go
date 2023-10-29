package main

import (
	"encoding/csv"
	"fmt"
	"os"
	"strconv"
)

type Grade string

const (
	A Grade = "A"
	B Grade = "B"
	C Grade = "C"
	F Grade = "F"
)

type student struct {
	firstName, lastName, university                string
	test1Score, test2Score, test3Score, test4Score int
}

type studentStat struct {
	student
	finalScore float32
	grade      Grade
}

func parseCSV(filePath string) []student {

	osFile, err := os.Open(filePath)
	if err != nil {
		fmt.Println("error in open the file")
	}
	defer osFile.Close()

	csvReader := csv.NewReader(osFile)

	var students []student
	var singleRow []string

	index := 0
	for {
		singleRow, err = csvReader.Read()
		if err != nil {
			fmt.Println("error in reading the row")
			fmt.Println(err)
			return students
		}
		if index == 0 {
			index += 1
			continue
		}
		singleStudent := student{}
		singleStudent.firstName = singleRow[0]
		singleStudent.lastName = singleRow[1]
		singleStudent.university = singleRow[2]
		singleStudent.test1Score, _ = strconv.Atoi(singleRow[3])
		singleStudent.test2Score, _ = strconv.Atoi(singleRow[4])
		singleStudent.test3Score, _ = strconv.Atoi(singleRow[5])
		singleStudent.test4Score, _ = strconv.Atoi(singleRow[6])

		students = append(students, singleStudent)
	}

}

func calculateGrade(students []student) []studentStat {

	var studentStats []studentStat

	for _, singleStudent := range students {
		singleStudentStat := studentStat{}
		singleStudentStat.student = singleStudent
		finalScore := float32(singleStudent.test1Score+singleStudent.test2Score+singleStudent.test3Score+singleStudent.test4Score) / 4

		singleStudentStat.grade = A
		if finalScore < 35 {
			singleStudentStat.grade = F
		} else if finalScore < 50 {
			singleStudentStat.grade = C
		} else if finalScore < 70 {
			singleStudentStat.grade = B
		}
		singleStudentStat.finalScore = finalScore
		studentStats = append(studentStats, singleStudentStat)
	}

	return studentStats
}

func findOverallTopper(gradedStudents []studentStat) studentStat {
	maxScoreStudentStat := studentStat{}
	maxScore := -1

	for _, gradedStudent := range gradedStudents {
		if gradedStudent.finalScore > float32(maxScore) {
			maxScore = int(gradedStudent.finalScore)
			maxScoreStudentStat = gradedStudent
		}
	}

	return maxScoreStudentStat
}

func findTopperPerUniversity(gs []studentStat) map[string]studentStat {

	var solution = make(map[string]studentStat, len(gs))

	for _, singleStudentStat := range gs {
		elem, ok := solution[singleStudentStat.university]
		if !ok || elem.finalScore < singleStudentStat.finalScore {
			solution[singleStudentStat.university] = singleStudentStat
		}
	}

	return solution
}
