package main

import "fmt"

func main() {
	scheme, host, err := ParseURL("https://akozadaev.ru")
	fmt.Println(scheme, host, err)
}
