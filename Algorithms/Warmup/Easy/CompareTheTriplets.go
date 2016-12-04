package main
import "fmt"

func main() {
    var a, b [3] uint32
    var pointA, pointB uint32
    pointA = 0
    pointB = 0
    fmt.Scanf("%v %v %v", &a[0], &a[1], &a[2])
    fmt.Scanf("%v %v %v", &b[0], &b[1], &b[2])
    for i:=0; i<3; i++{
        switch {
            case a[i]>b[i]:
                pointA++
            case a[i]<b[i]:
                pointB++
        }
    }
    fmt.Printf("%v %v", pointA, pointB)
 //Enter your code here. Read input from STDIN. Print output to STDOUT
}
