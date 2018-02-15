package main

import (
  "fmt"
  "strings"
)

type Board [][]Disk

func NewBoard(width, height int) (Board, error) {
  if width < 2 {
    return nil, fmt.Errorf("The width must be bigger than 2.")
  }
  if height < 2 {
    return nil, fmt.Errorf("The height must be bigger than 2.")
  }
  if width % 2 != 0 {
    return nil, fmt.Errorf("The width must be a multiple of 2.")
  }
  if height % 2 != 0 {
    return nil, fmt.Errorf("The height must be a multiple of 2.")
  }

  board := make([][]Disk, height)
  for i := range board {
    board[i] = make([]Disk, width)
  }
  board[height / 2 - 1][width / 2 - 1] = Black
  board[height / 2 - 1][width / 2] = White
  board[height / 2][width / 2 - 1] = White
  board[height / 2][width / 2] = Black

  return board, nil
}

func (board Board) Height() int {
  return len(board)
}

func (board Board) Width() int {
  return len(board[0])
}

func (board Board) CheckDiskPlaceableOn(disk Disk, i, j int) bool {
  for h := 0; h < 8; h++ {
    var m, n int
    switch h {
    case 0:
      m, n = -1, -1
    case 1:
      m, n = -1, 0
    case 2:
      m, n = -1, 1
    case 3:
      m, n = 0, -1
    case 4:
      m, n = 0, 1
    case 5:
      m, n = 1, -1
    case 6:
      m, n = 1, 0
    case 7:
      m, n = 1, 1
    }

    if i + m < 0 || i + m >= board.Height() {
      continue
    }
    if j + n < 0 || j + n >= board.Width() {
      continue
    }
    if board[i + m][j + n] == Empty || board[i + m][j + n] == disk {
      continue
    }
    for k := 2; 0 <= i + m * k && i + m * k < board.Height() && 0 <= j + n * k && j + n * k < board.Width(); k++ {
      if board[i + m * k][j + n * k] == Empty {
        break
      }
      if board[i + m * k][j + n * k] == disk {
        return true
      }
    }
  }
  return false
}

func (board Board) PlaceDiskOn(disk Disk, i, j int) error {
  if i < 0 || i >= board.Height() {
    return fmt.Errorf("The row number is invalid.")
  }
  if j < 0 || j >= board.Width() {
    return fmt.Errorf("The column number is invalid.")
  }

  placeable := false
  for h := 0; h < 8; h++ {
    var m, n int
    switch h {
    case 0:
      m, n = -1, -1
    case 1:
      m, n = -1, 0
    case 2:
      m, n = -1, 1
    case 3:
      m, n = 0, -1
    case 4:
      m, n = 0, 1
    case 5:
      m, n = 1, -1
    case 6:
      m, n = 1, 0
    case 7:
      m, n = 1, 1
    }

    if i + m < 0 || i + m >= board.Height() {
      continue
    }
    if j + n < 0 || j + n >= board.Width() {
      continue
    }
    if board[i + m][j + n] == Empty || board[i + m][j + n] == disk {
      continue
    }
    for k := 2; 0 <= i + m * k && i + m * k < board.Height() && 0 <= j + n * k && j + n * k < board.Width(); k++ {
      if board[i + m * k][j + n * k] == Empty {
        break
      }
      if board[i + m * k][j + n * k] == disk {
        placeable = true
        for l := 0; l < k; l++ {
          board[i + m * l][j + n * l] = disk
        }
      }
    }
  }

  if !placeable {
    return fmt.Errorf("The disk can not be placed there.")
  }

  return nil
}

func (board Board) CountDisk(disk Disk) (count int) {
  for i := 0; i < board.Height(); i++ {
    for j := 0; j < board.Width(); j++ {
      if board[i][j] == disk {
        count++;
      }
    }
  }
  return count
}

func (board Board) String() (str string) {
  str += "+" + strings.Repeat("---+", board.Width())
  for i := 0; i < board.Height(); i++ {
    str += "\n|"
    for j := 0; j < board.Width(); j++ {
      substr := fmt.Sprint(board[i][j])
      str += substr + "+"
    }
    str += "\n+" + strings.Repeat("---+", board.Width())
  }
  return
}
