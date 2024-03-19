package main

import "fmt"

func main() {
	var g Greeter
	g = Helloer{}
	g.Greet()
}

type Greeter interface{ Greet() }
type Helloer struct{}
type Goodbyer struct{}

var _ Greeter = Helloer{}
var _ Greeter = Goodbyer{}

func (Helloer) Greet()  { hello() }
func (Goodbyer) Greet() { goodbye() }
func hello()            { fmt.Println("你好，111！") }
func goodbye()          { fmt.Println("再见，222！") }

//deadcode .

//deadcode -whylive=github.com/xwen/gofeature/tools/deadcode.hello .
