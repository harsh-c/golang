package main
        import "fmt"

        func main() {
        name1 := []string{"Harsh","Rohan"}
        name2 := []string{"Sam","Mike"}
        name := append(name1,name2...)
        fmt.Println(name)
        }