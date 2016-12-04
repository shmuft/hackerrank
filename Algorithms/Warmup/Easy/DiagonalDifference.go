package main
import "fmt"

func main() {
    var res, i, j, num, data, mainDiag, secondDiag int32
    fmt.Scanf("%v", &num)
    mainDiag = 0
    secondDiag = 0
    for i=0; i<num; i++{
        for j=0; j<num; j++{
            fmt.Scanf("%v", &data)
            if (i == j) {
                mainDiag += data
            }
            if (i == num-j-1) {
                secondDiag += data
            }
        }
    }
    res = mainDiag - secondDiag
    if res < 0 {
        res = res * (-1)
    }
    fmt.Println(res)
}
