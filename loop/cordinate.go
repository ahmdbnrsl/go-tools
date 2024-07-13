package loop

import "fmt"
import "strings"
import "github.com/julien040/go-ternary"

func Cordinate(count int) {
    for i := 0; i < count * 2 + 1; i++ {
        if i == count {
            a := make([]any, count * 2 + 1, count * 2 + 1)
            for y := 0; y < count * 2; y++ {
                a[y] = y - count
                if y == count * 2 - 1 {
                    a[y + 1] = count
                }
            }
            fmt.Println(a...);
        } else {
            fmt.Println(
                strings.Repeat(" ", ternary.If(i <= count, count * 3 - 1, count * 3 - 2)),
                count - i,
            )
        }
    }
}
