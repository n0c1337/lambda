section .text
global convert
convert:
    push rbp
    mov rbp, rsp
    sub rsp, 8
    mov qword[rsp-8], 2147483647
    cmp qword[rsp-8], 2147483647
    je .labelTrue
    mov rax, 42
    jmp .labelEnd
.labelTrue:
    mov rax, 1
    add rax, [rsp-8]
.labelEnd:
    add rsp, 8
    leave
    ret
