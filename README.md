# Sudoku_solver

## Sudoku solver coded in go
This program use Backtracking algorithm to brute force the solution for a given Sudoku.

# Usage :

To use it go to realease and download the last version for your OS.
Open a terminal in the folder where the file is located.

*You can either launch it and give the path to the input file when asked or give the path as an argument.*

### the input need to be a texte file formated like this :
```
0 0 0 0 0 0 0 0 0
0 0 0 0 0 3 0 8 5
0 0 1 0 2 0 0 0 0
0 0 0 5 0 7 0 0 0
0 0 4 0 0 0 1 0 0
0 9 0 0 0 0 0 0 0
5 0 0 0 0 0 0 7 3
0 0 2 0 1 0 0 0 0
0 0 0 0 4 0 0 0 9
```
For a sudoku like this this :
```
╔═════════╦═════════╦═════════╗
║ _  _  _ ║ _  _  _ ║ _  _  _ ║
║ _  _  _ ║ _  _  3 ║ _  8  5 ║
║ _  _  1 ║ _  2  _ ║ _  _  _ ║
╠═════════╬═════════╬═════════╣
║ _  _  _ ║ 5  _  7 ║ _  _  _ ║
║ _  _  4 ║ _  _  _ ║ 1  _  _ ║
║ _  9  _ ║ _  _  _ ║ _  _  _ ║
╠═════════╬═════════╬═════════╣
║ 5  _  _ ║ _  _  _ ║ _  7  3 ║
║ _  _  2 ║ _  1  _ ║ _  _  _ ║
║ _  _  _ ║ _  4  _ ║ _  _  9 ║
╚═════════╩═════════╩═════════╝
```

---
## Build the binaries :
**You need to have Go installed !**
```
go build .
```