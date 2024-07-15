package channel

func Pow(ref *int, val int) {
    isRef := *ref
    for i := 1; i < val; i++ {
        *ref = *ref * isRef
    }
}

func Pow2(ref *chan int, val, count int) {
    go func(val, count int) {
        isRef := val
        for i := 1; i < count; i++ {
            val = val * isRef
        }
        *ref <- val
    }(val, count)
}