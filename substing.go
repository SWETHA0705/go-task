package main

import(
    "fmt"
)

func main(){
    count :=0
    var s string 
    fmt.Scanln(&s)
    for i:=0;i<len(s)-1;i++{
        
            if s[i]=='g'&&s[i+1] =='o'{
                count++
            }
        }
    
    fmt.Printf("%d",count)
}
