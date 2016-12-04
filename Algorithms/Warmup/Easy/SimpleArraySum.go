package main
import "fmt"

func main() {
 //Enter your code here. Read input from STDIN. Print output to STDOUT
    var num, i, item, res uint32
    fmt.Scanf("%v", &num)
    res = 0
    for i=0;i<num;i++{
        fmt.Scanf("%v", &item)
        res += item
    }
    fmt.Println(res)
}
