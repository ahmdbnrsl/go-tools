package pattern

import "go-gabut/console"
import "strings"
import "github.com/julien040/go-ternary"

func Diamond(count int) {
    for i := 0; i < count * 2 - 1; i++ {
        if i == count - 1 {
            console.Log("", strings.Repeat("* ", count))
        } else if i < count - 1 {
            console.Log(strings.Repeat(" ", count - (i + 1)), strings.Repeat("* ", i + 1))
        } else {
            console.Log(strings.Repeat(" ", i + 1 - count), strings.Repeat("* ", i - (((i - count) * 2) + 1)))
        }
    }
}

func Diamond2(count int) {
    for i := 0; i < count * 2 - 1; i++ {
        switch {
        case i == count - 1:
            console.Log(
                "", "*", strings.Repeat(" ",(i - 1) + (i - 2)), "*",
            )
            break
        case i < 2:
            console.Log(
                strings.Repeat(" ", count - (i + 1)),
                "*",
                ternary.If(i < 1, "", "*"),
            )
            break
        case i < count - 1:
            console.Log(
                strings.Repeat(" ", count - (i + 1)),
                "*",
                strings.Repeat(" ", (i - 1) + (i - 2)),
                "*",
            )
            break
        case i < count * 2 - 3:
            console.Log(
                strings.Repeat(" ", i + 1 - count),
                "*",
                strings.Repeat(" ", (i - (((i - count) * 2) + 3)) * 2 - 1),
                "*",
            )
            break
        default:
            console.Log(
                strings.Repeat(" ", i + 1 - count),
                "*",
                ternary.If(i < count * 2 - 2, "*", ""),
            )
            break
        }
    }
}
