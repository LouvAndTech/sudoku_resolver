package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
	"time"
)

func readFile(filename string) [9][9]int {
	//Init variables and open the file
	var output [9][9]int
	file, err := os.Open(filename)
	if err != nil {
		log.Fatalf("failed to open")
	} else {

		//Split the file into row
		scanner := bufio.NewScanner(file)
		scanner.Split(bufio.ScanLines)

		var row []string

		//For each line get the line
		ind := 0
		for scanner.Scan() {
			row = append(row, scanner.Text())
			pased := strings.Fields(row[ind])
			//for each caracter in the line convert it into a int
			for i, v := range pased {
				output[ind][i], err = strconv.Atoi(v)
				if err != nil {
					log.Fatalf("failed to convert to int")
				}
			}
			ind++
		}

		//Close the file
		file.Close()
	}

	return output
}

//The programm take as an argument the path of the sudoku file
func printTab(tab [9][9]int) string {
	var output string
	output = "╔═════════╦═════════╦═════════╗\n"
	for i := 0; i < 9; i++ {
		for j := 0; j < 9; j++ {
			if j%3 == 0 {
				output += "║"
			}
			if tab[i][j] == 0 {
				output += " _ "
			} else {
				output += fmt.Sprintf(" %d ", tab[i][j])
			}
		}
		if (i == 2) || (i == 5) {
			output += "║\n╠═════════╬═════════╬═════════╣\n"
		} else {
			output += "║\n"
		}
	}
	output += "╚═════════╩═════════╩═════════╝\n"
	return output
}

func main() {
	fmt.Println("Starting !\n")
	start := time.Now()

	//Get the sudoku from the arg or a user input
	var path string
	if len(os.Args) != 2 {
		fmt.Print("Enter the path of the sudoku file : ")
		fmt.Scanln(&path)
	} else {
		path = os.Args[1]
	}
	tab := readFile(path)

	//Print the sudoku
	fmt.Println("INTPUT SUDOKU :")
	fmt.Print(printTab(tab))

	//Solve the sudoku
	fmt.Println("\nWorking...")
	solved := solve(tab)
	fmt.Println("Solved !!!")

	//Print the sudoku solved
	fmt.Println("\nSOLUTION :")
	fmt.Print(printTab(solved))

	//Print the time and the end of the programm
	fmt.Printf("\nSolving time : %s!\n", time.Since(start))
	fmt.Println("Press the Enter Key to quit")
	fmt.Scanln()
	fmt.Println("Closing.\n")
}
