package main

import (
       "fmt"
       "os"
       "time"
)

func main() {
     var s, sep string

     start := time.Now()

     for i := 0; i < len(os.Args); i++ {
     	 s += sep + os.Args[i]
	 sep = " "
     }
     fmt.Println(s)

     fmt.Printf("%.12fs elapsed", time.Since(start).Seconds())
}