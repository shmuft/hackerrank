package main
import "fmt"

func main() {
    var v, n, val int
    fmt.Scanf("%v", &v)
    fmt.Scanf("%v", &n)
    for i:=0; i < n; i++{
        fmt.Scanf("%v", &val)
        if val == v {
            fmt.Println(i)
        }
    }
 //Enter your code here. Read input from STDIN. Print output to STDOUT
}
