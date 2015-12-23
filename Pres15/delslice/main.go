package main
        import "fmt"

        func main() {
        name := []string{"Harsh","Rohan","Sam","Mike"}
        fmt.Println(name)
        name = append(name[:2],name[3:]...)
        fmt.Println(name)
        }