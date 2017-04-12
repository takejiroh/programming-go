package main

import (
       "fmt"
       "os"
       "strconv"
)

func main() {
     var s, sep string

     for i := 0; i < len(os.Args); i++ {
     	 sep = strconv.Itoa(i)
     	 sep += ": "
     	 s = sep + os.Args[i]
	 
	 //fmt.Println( i, ": ", s)
	 fmt.Println(s)
     }

}