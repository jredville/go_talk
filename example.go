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

func (f *Foo) baz() int32 { // no receiver
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

func showBoo() {
 var b Booer
 callBoo(b)
}

type Inventory struct {
	Material string
	Count    uint
}
sweaters := Inventory{"wool", 17}
tmpl, err := template.New("test").Parse("{{.Count}} items are made of {{.Material}}")
if err != nil { panic(err) }
err = tmpl.Execute(os.Stdout, sweaters)
if err != nil { panic(err) }
