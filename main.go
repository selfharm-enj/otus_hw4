package main

import (
	"flag"
	"fmt"
	"slices"
	"strconv"
	"strings"
)

func main() {
	p1 := flag.String("p1", "player1", "")
	p2 := flag.String("p2", "player2", "")
	fs := flag.Int("fs", 8, "Field size `N` transform into field NxN")
	flag.Parse()

	bfigures := []rune{'♜', '♞', '♝', '♛', '♚', '♝', '♞', '♜'}
	bpawns := []rune{'♟', '♟', '♟', '♟', '♟', '♟', '♟', '♟'}
	wfigures := []rune{'♖', '♘', '♗', '♕', '♔', '♗', '♘', '♖'}
	wpawns := []rune{'♙', '♙', '♙', '♙', '♙', '♙', '♙', '♙'}

	board := strings.Builder{}
	fmt.Fprintln(&board, *p2)
	for r := range *fs {
		if slices.Contains([]int{0, 1, *fs - 1, *fs - 2}, r) {
			switch r {
			case 0:
				board.WriteString(strconv.Itoa(*fs))
				board.WriteString(DrawRow(bfigures, *fs))
			case 1:
				board.WriteString(strconv.Itoa(*fs - 1))
				board.WriteString(DrawRow(bpawns, *fs))
			case *fs - 2:
				board.WriteString(strconv.Itoa(*fs - r))
				board.WriteString(DrawRow(wpawns, *fs))
			case *fs - 1:
				board.WriteString(strconv.Itoa(*fs - r))
				board.WriteString(DrawRow(wfigures, *fs))
			}
		} else {
			board.WriteString(strconv.Itoa(*fs - r))
			board.WriteString(strings.Repeat(".", *fs))
			board.WriteString("\n")
		}
	}
	fmt.Fprintln(&board, *p1)
	fmt.Println(board.String())
}

func DrawRow(pieces []rune, boardSize int) string {
	out := strings.Builder{}
	piecesLen := len(pieces)

	for l := range boardSize {
		if l > (piecesLen - 1) {
			x := fmt.Sprintf("%s", string(pieces[(l%piecesLen)]))
			out.WriteString(x)
		} else {
			x := fmt.Sprintf("%s", string(pieces[l]))
			out.WriteString(x)
		}
	}
	out.WriteString("\n")
	return out.String()
}
