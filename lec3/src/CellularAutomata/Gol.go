package main

// GameBoard is a two-dimensional slice of booleans
type GameBoard []([]bool)

// here we will provide our functions for the Game of Life

//PlayGol takes an initial game board and a number of generations. It returns a slice of game boards corresponding to playing Game of Life numGens generations, starting with the initial board.
func PlayGol(initialBoard GameBoard, numGens int) []GameBoard {
    boards := make([]GameBoard, numGens + 1)
    boards[0] = initialBoard

    for i := 1; i <= numGens; i++ {
        boards[i] = UpdateBoard(boards[i-1])
    }

    return boards
}

//UpdateBoard takes a GameBoard and returns the board resulting from playing the Game of Life for one generation.
func UpdateBoard(currBoard GameBoard) GameBoard {
    // first, create a new board corresponding to the next generation.
    // let's have all cells dead to begin with.
    numRows := CountRows(currBoard)
    numCols := CountCols(currBoard)
    newBoard := InitializeBoard(numRows, numCols)

    //now, update values of newBoard
    //range through all cells of currBoard and update each one into newBoard.
    for r := range currBoard {
        // r will range over rows of board
        // current row is currBoard[r]
        // range over values in currBoard[r]
        for c := range currBoard[r] {
            //curr value is currBoard[r][c]
            newBoard[r][c] = UpdateCell(currBoard, r, c)
        }
    }
    return newBoard
}

//UpdateCell takes a GameBoard and row/col indices r and c, and it returns the state of the board at these row/col indices is in the next generation.
func UpdateCell(board GameBoard, r, c int) bool {
    numNeigbors := CountLiveNbrs(board, r, c)

    // now it's just a matter of consulting the rules.
    if board[r][c] == true { // i'm alive Now
        if numNeigbors == 2 || numNeigbors == 3 {
            return true // staying alive
        } else {
            return false // dying out
        }
    } else {// I'm dead Now
        if numNeigbors == 3 {
            // zombie!
            return true
        } else { // RIP
            return false
        }
    }
}

// CountLiveNbrs takes a GameBoard along with row and column indices and it counts the live neighbors of the cell at this position.
// it won't consider cells that are off the board.
func CountLiveNbrs(board GameBoard, r, c int) int {
    count := 0

    for i := r-1; i <= r+1; i++ {
        for j := c-1; j <= c+1; j++ {
            if (i != r || j != c) &&
            InField(board, i, j) {
              if board[i][j] == true {
                  count++
                }
            }
        }
    }
    return count
}

//InField takes a GameBoard board as well as row and col indices (i, j) and returns true if board[i][j] is in the board and false otherwise.
func InField(board GameBoard, i, j int) bool {
    if i < 0 || j < 0 {
        return false
    }
    if i >= CountRows(board) || j >= CountCols(board) {
        return false
    }
    // if we survive to here, then we are on the board
    return true
}

func CountRows(board GameBoard) int {
    return len(board)
}

func CountCols(board GameBoard) int {
    // assume that we have a rectangular board
    if CountRows(board) == 0 {
        panic("Error: empty board given to CountCols")
    }
    // give # elements in 0-th row
    return len(board[0])

}

//InitializeBoard takes a number of rows and columns as inputs and returns a gameboard with appropriate number of rows and columns.
func InitializeBoard(numRows, numCols int) GameBoard {
    // make a 2-D slice (default values = false)
    var board GameBoard // or immediately board := make(GameBoard, numRows)
    board = make(GameBoard, numRows)
    // now we need to make the rows too
    for r := range board {
        board[r] = make([]bool, numCols)
    }

    return board
}
