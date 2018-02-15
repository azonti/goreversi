package main

type Player_ struct {
  disk Disk
}

func NewPlayer_(disk Disk) (player Player_) {
  player.disk = disk
  return
}

func (player Player_) Disk() Disk {
  return player.disk
}

func (player Player_) IsPassable(board Board) bool {
  for i := 0; i < board.Height(); i++ {
    for j := 0; j < board.Width(); j++ {
      if board.CheckDiskPlaceableOn(player.disk, i, j) {
        return false
      }
    }
  }
  return true
}

type Player interface {
  Disk() Disk
  IsPassable(board Board) bool
  Move(board Board)
  TellGameOver(board Board, players []Player)
}
