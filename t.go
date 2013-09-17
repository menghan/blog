package main

import "fmt"

type Handle interface {
	ServeHTTP(c *int, req *int)
}

func NotFound(c *int, req *int) {
}

type HandleFunc func(c *int, req *int)

func (f HandleFunc) ServeHTTP(c *int, req *int) {
	f(c, req)
}

func (f HandleFunc) String() string {
	return "hahah"
}

type HandleImpl struct {
}

func (hi HandleImpl) ServeHTTP(c *int, req *int) {
}

func main() {
	var Handle404 = HandleFunc(NotFound)
	var hi HandleImpl
	var hi2 = HandleImpl(hi)
	fmt.Printf("%v %v %v\n", Handle404, hi, hi2)
}
