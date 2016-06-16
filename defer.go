package main

import "log"

// DeferExample demos defer
type DeferExample struct {
  variable int32
}

// NewDeferExample returns a pointer to a DeferExample
func NewDeferExample() *DeferExample {
  return &DeferExample{
    variable: 1,
  }
}

type intTaker func(int32)
func (d *DeferExample) showIt(f intTaker) {
  d.changeVar()
  defer d.resetVar()

  f(d.variable)
}

func (d *DeferExample) changeVar() {
  d.variable = 2
}

func (d *DeferExample) resetVar() {
  d.variable = 1
}

func (d *DeferExample) printVar() {
  log.Printf("%d", d.variable)
}

func main() {
  de := NewDeferExample()
  de.printVar()
  de.showIt(func (num int32) {
    log.Printf("%d", num)
    de.printVar()
  })
  de.printVar()
}
