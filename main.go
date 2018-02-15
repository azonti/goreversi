package main

import (
  "flag"
  "fmt"
  "os"
)

func main() {
  w := flag.Int("w", 8, "The width of the board.")
  h := flag.Int("h", 8, "The height of the board.")
  flag.Parse()

  board, err := NewBoard(*w, *h)
  if err != nil {
    fmt.Fprintf(os.Stderr, "Failed to create a board: %v\n", err)
    return
  }

  players := make([]Player, 2)
  for i, _ := range players {
    players[i] = NewLocal(Disk(i + 1))
  }

  for {
    i := 0
    for _, player := range players {
      if player.IsPassable(board) {
        i++
        continue
      }
      fmt.Fprintln(os.Stderr, board)
      player.Move(board)
    }
    if i == len(players) {
      break
    }
  }

  for _, player := range players {
    player.TellGameOver(board, players)
  }
}
