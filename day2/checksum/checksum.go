package checksum

import (
)

func Checksum(s string) (twos, threes bool){
    listed := make(map[rune]int)
    for _, rune := range s {
        listed[rune]+=1
    }
    for _,v := range listed{
        if v==2 {twos =true}
        if v==3 {threes = true}
    }
return
}
