// Copyright © 2016 Alan A. A. Donovan & Brian W. Kernighan.
// License: https://creativecommons.org/licenses/by-nc-sa/4.0/

// See page 45.

// (Package doc comment intentionally malformed to demonstrate golint.)
//!+
package popcount

//import "fmt"

// pc[i] is the population count of i.
var pc [256]byte

func init() {
	for i := range pc {
		pc[i] = pc[i/2] + byte(i&1)
	}
}

// PopCount returns the population count (number of set bits) of x.
func PopCount(x uint64) int {
     return int(pc[byte(x>>(0*8))] +
		pc[byte(x>>(1*8))] +
		pc[byte(x>>(2*8))] +
		pc[byte(x>>(3*8))] +
		pc[byte(x>>(4*8))] +
		pc[byte(x>>(5*8))] +
		pc[byte(x>>(6*8))] +
		pc[byte(x>>(7*8))])
}

//書きなおせという課題ですが、比較のために新設させてください
func PopCountByLoop( x uint64 ) int{
     sum := 0
     
     for i:=0; i<8; i++ {
     	 sum += int(pc[byte(x>>(uint(i)*8))]);
     }
     
     return sum;
}

func PopCountByClearing( x uint64 ) int{
     ret := 0
     for i:=0; i<64; i++{
     	 x = x&(x-1)
	 if x == 0{
	    ret = int(i)
	    break
	 }
     }
     return ret
}

func PopCountByShifting( x uint64 ) int{
     
     sum := 0
     
     for i:=0; i<64; i++ {
     	 sum += int(x>>uint(i) & 1)
     }
     
     return sum;
}


//!-
