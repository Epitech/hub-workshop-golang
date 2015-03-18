package main

import "fmt"

type error interface{
    Error() string
}

type nbr interface {
    GetFloat() float64
}

type nbrI interface {
    GetInt() int64
}

type superNbr interface{
    nbr
    nbrI
}

type F struct {
}

func (f *F) Error() string {
    return "ERROR"
}

func (f *F) GetFloat() float64 {
    return 42.42
}

func returnError() error {
    f := &F{}
    return f
}

func Add(a, b nbr) float64 {
    return a.GetFloat() + b.GetFloat()
}

func main() {
    f1 := &F{}
    f2 := &F{}

    fmt.Println(returnError())
    fmt.Println(Add(f1, f2))
}
