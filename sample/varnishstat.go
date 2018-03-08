package main

import "C"

import(
    "fmt"
    "../head"
)

func main(){
  varnishapi.StatInit()
  for name,v :=range varnishapi.StatGet(){
    fmt.Printf("%50s %20d %s\n",name, v.Val, v.Sdesc)
  }
  varnishapi.StatFini()
  
}