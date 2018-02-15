package main

import (
  "fmt"
  "os"
)

type Local struct {
  Player_
}

func NewLocal(disk Disk) (local Local) {
  local.Player_ = NewPlayer_(disk)
  return
}

func (local Local) Move(board Board) {
  for {
    fmt.Printf("You are the %v player. Type the row and column number of the grid to place your disk on.\n", local.disk)
    var i, j int
    if n, _ := fmt.Scanf("%d %d", &i, &j); n != 2 {
      continue
    }
    if err := board.PlaceDiskOn(local.disk, i, j); err != nil {
      fmt.Fprintf(os.Stderr, "Failed to place your disk on the board: %v\n", err)
      continue
    }
    break
  }
}

func (local Local) TellGameOver(board Board, players []Player) {
  i := 0;
  for _, player := range players {
    if board.CountDisk(local.disk) < board.CountDisk(player.Disk()) {
      i++
    }
  }
  if i == 0 {
    fmt.Printf("The %v player win!\n", local.disk)
  } else {
    fmt.Printf("The %v player lose...\n")
  }
}
