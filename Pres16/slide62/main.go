package main
        import "fmt"

        func printFriends(names ...string) {
        for _,val := range names{
        fmt.Println(val)
        }
        }

        func main() {
        printFriends("Harsh","Chauhan","Hitansh","Sam","Mike")
        }