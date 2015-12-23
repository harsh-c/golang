package main
        import "fmt"

        type student struct{
        name string
        age int
        }

        func main(){
        h := student{name: "Harsh", age:21}
        ro := student{name:"Rohan"}

        fmt.Println(h.name, h.age)
        fmt.Println(ro.name)

        ro.age=19
        fmt.Println(ro.age)
        }