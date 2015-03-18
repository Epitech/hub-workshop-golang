package main
import "fmt"

func condition() {
    i := 10
    if i == 10 {

    } else if i > 10 {

    } else {

    }
}

func main() {
    for i := 0; i < 100; i++ {
    }
    array := []string{"salut", "le", "monde"}
    for index, value := range array {
        fmt.Println(index, value)
    }
}