package main

import (
	"log"

	"github.com/schollz/progressbar/v3"
)

func isIn(a int, list [9]int) bool {
	for _, b := range list {
		if b == a {
			return true
		}
	}
	return false
}

func checkLine(tab [9][9]int, line int, val int) bool {
	return !isIn(val, tab[line])
}
func checkCol(tab [9][9]int, col int, val int) bool {
	//format the column into a line
	var lineFormat [9]int
	for i := 0; i < 9; i++ {
		lineFormat[i] = tab[i][col]
	}
	return !isIn(val, lineFormat)
}
func checkSquare(tab [9][9]int, line int, col int, val int) bool {
	//find the square where we are
	realCol := col - (col % 3)
	realLine := line - (line % 3)
	//format the square into a line
	var lineFormat [9]int
	for i := 0; i < 3; i++ {
		for j := 0; j < 3; j++ {
			lineFormat[i*3+j] = tab[realLine+i][realCol+j]
		}
	}
	return !isIn(val, lineFormat)
}

func checkCase(tab [9][9]int, i int, j int) ([]int, bool) {
	var output []int
	possible := false
	for x := 1; x <= 9; x++ {
		if checkLine(tab, i, x) && checkCol(tab, j, x) && checkSquare(tab, i, j, x) {
			output = append(output, x)
			possible = true
		}
	}
	return output, possible
}

type Pos struct {
	x int
	y int
}

func initPos(tab [9][9]int) []Pos {
	var output []Pos
	index := 0
	for i := 0; i < 9; i++ {
		for j := 0; j < 9; j++ {
			if tab[i][j] == 0 {
				output = append(output, Pos{i, j})
				index++
			}
		}
	}
	return output
}

//solve a sudoku
func solve(tab [9][9]int) [9][9]int {
	output := tab
	var possible [9][9][]int
	//save the coord for each case we can change
	pos := initPos(tab)
	log.Printf("%d cases to solve :", len(pos))
	var available bool

	//declare the progress bar to the number of cases to solve
	bar := progressbar.Default(int64(len(pos)))

	cycle := 0
	i := 0
	//For each case of the sudoku that we can change
	for i < (len(pos)) {
		//get the possibilities
		possible[pos[i].x][pos[i].y], available = checkCase(output, pos[i].x, pos[i].y)

		if available {
			//if there is any, we choose the first one and pass to the next case
			output[pos[i].x][pos[i].y] = possible[pos[i].x][pos[i].y][0]
			i++
		} else {

			//if there is no possibility we work on the case before
			for len(possible[pos[i-1].x][pos[i-1].y]) == 1 {
				//If there was no other possibility, we go back one more case
				output[pos[i-1].x][pos[i-1].y] = 0
				i--
				if i == 0 {
					//If we are at the first case, we can't go back anymore
					log.Fatalf("No solution found")
				}
			}
			//Once we are back to a case with multiple possibilities, we remove the last possibility
			possible[pos[i-1].x][pos[i-1].y][0] = possible[pos[i-1].x][pos[i-1].y][len(possible[pos[i-1].x][pos[i-1].y])-1]
			possible[pos[i-1].x][pos[i-1].y] = possible[pos[i-1].x][pos[i-1].y][:len(possible[pos[i-1].x][pos[i-1].y])-1]
			//Them use the next available and go back in the loop
			output[pos[i-1].x][pos[i-1].y] = possible[pos[i-1].x][pos[i-1].y][0]
		}
		cycle++

		//update the progress bar to the case we are solving
		bar.Set(i)
	}
	log.Printf("Cycles required : %d ", cycle)
	return output
}
