package main
import "fmt"

func main() {
        var x,y int
        fmt.Println("Enter two numbers:")
        fmt.Scan(&x,&y)
        fmt.Println("The remainder is:",x%y)
}