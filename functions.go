package main
import "fmt"

func drawFrame(body map[int]map[int]string) error {
    template := []rune(`

         |       |
         |       |    
  _______|_______|_______
         |       |
         |       |    
  _______|_______|_______
         |       |
         |       |    
         |       |
`)

	placeHolders := [][]int{[]int{26, 34, 42},
                            []int{94, 102, 110},
                            []int{162, 170, 178}}

    for keyRow, rows := range body {
        for keyCol, val := range rows {
            if val == "X" {
                template[placeHolders[keyRow][keyCol]] = 'X'
            } else if val == "O" {
                template[placeHolders[keyRow][keyCol]] = 'O'
            } else { return fmt.Errorf("unknown error") }
        }
    }
	fmt.Println(string(template))
    return nil
}

func gameStatus(body map[int]map[int]string, count *int) int {

    *count++

    switch true {
    case body[0][0] == "X" && body[0][1] == "X" && body[0][2] == "X":
        return 1
    case body[1][0] == "X" && body[1][1] == "X" && body[1][2] == "X":
        return 1
    case body[2][0] == "X" && body[2][1] == "X" && body[2][2] == "X":
        return 1
    case body[0][0] == "X" && body[1][0] == "X" && body[2][0] == "X":
        return 1
    case body[0][1] == "X" && body[1][1] == "X" && body[2][1] == "X":
        return 1
    case body[0][2] == "X" && body[1][2] == "X" && body[2][2] == "X":
        return 1
    case body[0][0] == "X" && body[1][1] == "X" && body[2][2] == "X":
        return 1
    case body[0][2] == "X" && body[1][1] == "X" && body[2][0] == "X":
        return 1


    case body[0][0] == "O" && body[0][1] == "O" && body[0][2] == "O":
        return -1
    case body[1][0] == "O" && body[1][1] == "O" && body[1][2] == "O":
        return -1
    case body[2][0] == "O" && body[2][1] == "O" && body[2][2] == "O":
        return -1
    case body[0][0] == "O" && body[1][0] == "O" && body[2][0] == "O":
        return -1
    case body[0][1] == "O" && body[1][1] == "O" && body[2][1] == "O":
        return -1
    case body[0][2] == "O" && body[1][2] == "O" && body[2][2] == "O":
        return -1
    case body[0][0] == "O" && body[1][1] == "O" && body[2][2] == "O":
        return -1
    case body[0][2] == "O" && body[1][1] == "O" && body[2][0] == "O":
        return -1
    case *count == 9:
        return 2
    default:
        return 0
    }
}
