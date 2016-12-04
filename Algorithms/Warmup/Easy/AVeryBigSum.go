package main
import "fmt"

func main() {
    var num, i, res, a uint64
    res = 0    
    fmt.Scanf("%v", &num)
    for i=0; i<num; i++{
        fmt.Scanf("%v", &a)
        res += a
    }
    fmt.Println(res)
 //Enter your code here. Read input from STDIN. Print output to STDOUT
}
