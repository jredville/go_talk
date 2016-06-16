package main

import (
  "fmt"
  "log"
 )

// Foo is an example struct
type Foo struct {
  Foo string
  baz int32
}

// NewFoo returns an initilized Foo
func NewFoo(foo string) *Foo {
  return &Foo{
    Foo: foo, //capitalized export
    baz: 12,  //lower case not exported
  }
}

// Bar prints a method
func (f *Foo) Bar() string { //Exported with receiver
  return fmt.Sprintf("I haz %v", f.Foo)   // . access to pointer members
}

func (f *Foo) quux() int32 { // no receiver
  return f.baz + 3
}

type stringer func(string) string
func iTakeAFunc(blk stringer) string {
  return blk("world")
}

func iCallWithAFunc() {
  res := iTakeAFunc(func(f string) string {
    return "hello" + f
  })
  log.Println(res)
}

func enumeration(arr []interface{}) {
  for v,i := range arr {
    log.Printf("%v is at %v", v, i)
  }
}

func main() {
  f := NewFoo("abc")
  log.Println(f.Bar())
  iCallWithAFunc()
  arr := []interface{}{1, 2, 3, 4}
  arr2 := []interface{}{"1", "2", "3", "4"}

  enumeration(arr)
  enumeration(arr2)
}
