package main

import (
    "fmt"
    "math/rand"
    "sync"
)

const (
    Reponse  = 42
    Question = 21
)

type User struct {
    Name string
}

func (u *User) Mange(nourriture int) {
    switch nourriture {
        case Reponse:
        fmt.Println("Je mange la reponse")
        case Question:
        fmt.Println("Hein?")
        default:
        fmt.Println("C'est quoi ca?", nourriture)
    }
}

type Factory struct {
}

func (f *Factory) Produce() int {
    random := rand.Intn(1)
    switch random {
        case 0:
        return Question
        default:
        return Reponse
    }
}

func produce(channel chan int) {
    for i := 0; i < 100; i++ {
        channel <- i
    }
    close(channel)
    fmt.Println("Channel closed!")
}

func consume(channel chan int, wg *sync.WaitGroup) {
    u := &User{
        Name: "Vincent",
    }
    for nourriture := range channel {
        u.Mange(nourriture)
    }
    wg.Done()
}

func main() {
    wg := &sync.WaitGroup{}
    channel := make(chan int)
    go produce(channel)
    wg.Add(1)
    go consume(channel, wg)
    wg.Wait()
}
