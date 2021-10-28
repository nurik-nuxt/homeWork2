package sort

import (
	"fmt"
	"io/ioutil"
	"regexp"
	"strings"
)

func Sort(path string)  {
	content, _ :=ioutil.ReadFile(path)
	reg:=regexp.MustCompile(`\((.*?)\)`)
	fmt.Printf("Pattern: %v\n", reg.String())
	submatchall:=reg.FindAllString(string(content),-1)
	for _,element:=range submatchall{
		element = strings.Trim(element,"(")
		element = strings.Trim(element,")")
		fmt.Println(element)
	}
	//matches:=reg.FindAllStringSubmatch(string(content),-1)
	//for _,v:=range matches{
	//	imports:=strings.Split(v[1],"\n")
	//	fmt.Println(imports)
	//
	//}
}
