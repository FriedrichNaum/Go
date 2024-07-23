package main

import (
    "fmt"
)

func main() {
    var x int = 10 // declaração de variável
    y := 20        // declaração de variável com inferência de tipo
    z := x + y     // operação aritmética
    if z > 15 {    // estrutura condicional
        fmt.Println("z é maior que 15") // chamada de função
    } else {
        fmt.Println("z não é maior que 15")
    }

    // loop for
    for i := 0; i < 5; i++ {
        fmt.Println(i)
    }

    // comentário de linha única
    /* comentário de 
       múltiplas linhas */
    str := "Hello, World!" // string literal
    fmt.Println(str)

    // operador lógico e bit a bit
    a := true && false
    b := 1 | 2

    fmt.Println(a)
    fmt.Println(b)
}
