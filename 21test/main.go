package main

import (
       "fmt"
       "strconv"
       "os"
       "gopl.io/ch2/21"
)

func main() {
     for _, arg := range os.Args[1:] {
     	 t, err := strconv.ParseFloat(arg, 64)
	 if err != nil {
	    fmt.Fprintf(os.Stderr, "main: %v^n", err)
	    os.Exit(1)
	 }

	 f:= tempconv.Fahrenheit(t)
	 c:= tempconv.Celsius(t)
	 k:= tempconv.Kelvin(t)
	 
	 fmt.Printf( "%s = %s, %s = %s\n",
	 	     f, tempconv.FToC(f), c, tempconv.CToF(c))
	 fmt.Printf( "%s = %s, %s = %s\n",
	 	     f, tempconv.FToK(f), c, tempconv.CToK(c))
	 fmt.Printf( "%s = %s, %s = %s\n",
	 	     k, tempconv.KToF(k), k, tempconv.KToC(k))		 
		     
     }
}