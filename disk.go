package main

type Disk int
const (
  Empty Disk = iota
  Black
  White
)

func (disk Disk) String() (str string) {
  switch disk {
  case Empty:
    str = "   "
  case Black:
    str = "(#)"
  case White:
    str = "( )"
  }
  return
}
