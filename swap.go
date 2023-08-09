
package main
import "fmt"


func main() {
    string := "ABCDEFGHIJKL"
    s := []byte (string)
    
    var i int 
    for i= 2; i < len(s)-1; i += 4{
        
        s[i],s[i+1]=s[i+1],s[i]
    }
    fmt.Printf("%s \n",s) 

}
