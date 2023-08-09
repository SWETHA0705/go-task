package main

import(
    "fmt"
    "strconv"
)

func main(){
    str:= "4624748233249087"
    b,_:= strconv.Atoi(str)
    
     var a[16] int
     var d[16] int
     c:=1
     sum:=0
    for i:=15;i>=0;i--{
        a[i]= b%10
        b= b/10
    }
    i:= 0
    if a[0]==4 && len(a)==16{
        for j:= len(a)-2;j>=0;j-=2{
               prod := a[j]*2
            if prod>9{
                if c<3{
                  d[i]= prod%10
                  prod =prod/10
                  i++
                  c++
               }
               c=1
            }
               d[i]= prod
               i++
        
    }
    for i:=0;i<len(d);i++{
        sum= sum+d[i]
    }
    
    for j:=1;j<len(a);j+= 2{
       sum= sum+a[j]
    }
    if sum%10==0{
        fmt.Println("credit card is valid")
        
    }else{
        fmt.Println("credit card is not valid")
    }
} else{
    fmt.Println("credit card is not valid")
}
}
