package main

import (
	"fmt"
	"github.com/nurik-nuxt/homeWork2/atoi"
	"github.com/nurik-nuxt/homeWork2/fibonacci"
	"github.com/nurik-nuxt/homeWork2/reverse"
	"github.com/nurik-nuxt/homeWork2/itoa"
)

func main()  {
	fmt.Println("Startig HomeWork2")
	fmt.Println(atoi.Atoi(""))
	fmt.Println(reverse.Reverse("Құдалар"))
	fmt.Println(reverse.Reverse("Адата"))
	fmt.Println(reverse.Reverse("Adata"))
	generator:=fibonacci.Fibonacci()
	for i:=0; i < 10; i++ {
		fmt.Print(generator(),"")
	}
	fmt.Println()
	fmt.Println(itoa.Itoa(123))
}
