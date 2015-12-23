package main
import "fmt"

func main(){
      var num int
      fmt.Print("Enter a number: ")
      fmt.Scanln (&num)
      newNum := num * 10
      fmt.Println(num, "x 10 =", newNum)
}
