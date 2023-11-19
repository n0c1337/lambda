convert :: i64 -> i64
convert n =
    let x = 2147483647 + n
    if x == 2147483648
        x = x + 1
    else
        x = 42

    asm! {
        mov rax, 81
    }

    return x
