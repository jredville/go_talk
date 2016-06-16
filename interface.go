package main

import "log"

// Booable is an interface
type Booable interface {
  Boo() string
}

func callBoo(booer Booable) {
  log.Println(booer.Boo())
}

// Booer is a Booable
type Booer struct {
}

// Boo satisfies Booable
func (b *Booer) Boo() string {
  return "boo"
}

func main() {
 var b Booer
 callBoo(&b)
}
