package main

import (
    "fmt"
    "bufio"
    "os"
    "strings"
    "strconv"
)

func main() {
    gameBody := make(map[int]map[int]string)
    for i := range 3 {
        gameBody[i] = make(map[int]string)
    }
    scanner := bufio.NewScanner(os.Stdin)
    fmt.Println(
`
\\\\\\\\\\\\\\\\\\\\\\\\\\\\
\\\\\\\\|@@@@@@@@@@@|\\\\\\\
\\\\\\\\|TackTackGo!|\\\\\\\
\\\\\\\\|@@@@@@@@@@@|\\\\\\\
\\\\\\\\\\\\\\\\\\\\\\\\\\\\
`)
    isX := true
    count := 0
    currentChar := "X"
    for {
        if isX {
            fmt.Printf("X's turn (row, col)> ")
        } else {
            fmt.Printf("O's turn (row, col)> ")
        }
        scanner.Scan()
        text := strings.Fields(scanner.Text())

        if len(text) != 2 {
            fmt.Printf("\ninvalid format, please enter again\n")
            continue
        }

        row, _ := strconv.Atoi(text[0])
        col, _ := strconv.Atoi(text[1])

        if row > 2 || row < 0 || col > 2 || col < 0 {
            fmt.Printf("\ninvalid index, please enter again\n")
            continue
        }

        for {
            if _, ok := gameBody[row][col]; ok {
                fmt.Println("invalid position, please enter a different place")
                if isX {
                    fmt.Printf("X's turn (row, col)> ")
                } else {
                    fmt.Printf("O's turn (row, col)> ")
                }
                scanner.Scan()
                text = strings.Fields(scanner.Text())
                row, _ = strconv.Atoi(text[0])
                col, _ = strconv.Atoi(text[1])
            } else {
                gameBody[row][col] = currentChar
                break
            }
        }
        err := drawFrame(gameBody)
        if err != nil {
            fmt.Println("frames failed to generate...Closing")
            os.Exit(1)
        }
        fmt.Printf("\nEntered:\nrow: %d\ncol: %d\n", row, col)
        if currentChar == "X" {
            currentChar = "O"
        } else { currentChar = "X" }
        isX = !isX
        status := gameStatus(gameBody, &count)
        if status != 0 {
            if status == 1 {
                fmt.Println(`
                    \\\\\\\\\\\\\\\\\\\\\\\\\\\\
                    \\\\\\\\|@@@@@@@@@@@|\\\\\\\
                    \\\\\\\\|   X WON!  |\\\\\\\
                    \\\\\\\\|@@@@@@@@@@@|\\\\\\\
                    \\\\\\\\\\\\\\\\\\\\\\\\\\\\
                    `)
            } else if status == -1 {
                fmt.Println(
                    `
                    \\\\\\\\\\\\\\\\\\\\\\\\\\\\
                    \\\\\\\\|@@@@@@@@@@@|\\\\\\\
                    \\\\\\\\|   O WON!  |\\\\\\\
                    \\\\\\\\|@@@@@@@@@@@|\\\\\\\
                    \\\\\\\\\\\\\\\\\\\\\\\\\\\\
                    `)
            } else if status == 2 {
                fmt.Println(
                    `
                    \\\\\\\\\\\\\\\\\\\\\\\\\\\\
                    \\\\\\\\|@@@@@@@@@@@|\\\\\\\
                    \\\\\\\\|   DRAW!   |\\\\\\\
                    \\\\\\\\|@@@@@@@@@@@|\\\\\\\
                    \\\\\\\\\\\\\\\\\\\\\\\\\\\\
                    `)
            }
            os.Exit(1)
        }
    }
}
