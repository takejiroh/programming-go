// Copyright © 2016 Alan A. A. Donovan & Brian W. Kernighan.
// License: https://creativecommons.org/licenses/by-nc-sa/4.0/

// See page 43.
//!+

// Cf converts its numeric argument to Celsius and Fahrenheit.
package main

import (
	"fmt"
	"os"
	"strconv"

//	"gopl.io/ch2/tempconv"
//	"gopl.io/ch2/ex2-1"
)


type Celsius float64
type Fahrenheit float64

type Feat float64
type Metor float64

type Pond float64
type Kilo float64


func (c Celsius) String() string	{return fmt.Sprintf("%gC",c)}
func (f Fahrenheit) String() string	{return fmt.Sprintf("%gF",f)}

func (f Feat) String() string	{return fmt.Sprintf("%gfeat",f)}
func (m Metor) String() string	{return fmt.Sprintf("%gmetor",m)}

func (p Pond) String() string	{return fmt.Sprintf("%gPond",p)}
func (k Kilo) String() string	{return fmt.Sprintf("%gKilo-gram",k)}


func CToF(c Celsius) Fahrenheit		{return Fahrenheit(c*9/5 + 32)}
func FToC(f Fahrenheit) Celsius		{return Celsius((f-32) * 5/9)}


func FToM(f Feat) Metor			{return Metor(f/3.2808)}
func MToF(m Metor) Feat			{return Feat(m*3.2808)}

func PToK(p Pond) Kilo			{return Kilo(p*2.2046)}
func KToP(k Kilo) Pond			{return Pond(k/2.2046)}


func Convert(tgt float64) {
     t := tgt
     f :=Fahrenheit(t)
     c := Celsius(t)
     feat := Feat(t)
     m := Metor(t)
     p := Pond(t)
     k := Kilo(t)

     fmt.Printf("%s = %s, %s = %s\n",
     		    f, FToC(f), c, CToF(c) )
	fmt.Printf("%s = %s, %s = %s\n",
		feat, FToM(feat), m, MToF(m) )
	fmt.Printf("%s = %s, %s = %s\n",
		p, PToK(p), k, KToP(k) )    
}

func main() {
	for _, arg := range os.Args[1:] {
		t, err := strconv.ParseFloat(arg, 64)
		if err != nil {
			fmt.Fprintf(os.Stderr, "cf: %v\n", err)
			os.Exit(1)
		}
		Convert(t)
	}

	if len(os.Args) < 2 {
		var a float64
		fmt.Scan(&a)
		Convert(a)
	}
}

//!-
