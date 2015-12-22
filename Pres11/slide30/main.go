package main
import "fmt"

func main(){
        var a int =20
        var b *int =&a
        fmt.Println("a's value is ", a)
        fmt.Println("a's address is ", &a)
        fmt.Println("b's value is ", b)
        fmt.Println("b points to ", *b)
 }