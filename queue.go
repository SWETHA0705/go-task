package main

import "fmt"

const maxsize = 100

type queue struct{
      front int
      rear int
      arr[100] int
  }
 
  func (q*queue) Isempty()bool{
      return q.front==-1
  }
 func (q*queue) Isfull()bool{
      return q.rear== maxsize-1
 }
 func (q*queue) enqueue( data int){
     if q.Isfull(){
         fmt.Printf("queue is full")
         return
     }
     
     if q.Isempty(){
         q.front =0
     }
     q.rear++
     q.arr[q.rear]= data
         
     }
 func (q*queue)dequeue() {
     if q.Isempty(){
         fmt.Println("queue is empty")
         return
     }
     if q.front==q.rear{
         q.front =-1
         q.rear =-1
         
     }else{
         q.front++
     }
 }
 func (q*queue)display() {
     if q.Isempty(){
         fmt.Println("queue is empty")
         return}
   
    for i:=q.front;i<=q.rear;i++{
        fmt.Printf("%d",q.arr[i])
    }
      fmt.Println()
     }
     
 
func main() {
 
 q:= queue{front :-1,rear:-1}
 q.enqueue(3)
 q.enqueue(4)
 q.enqueue(5)
 q.display()
 q.dequeue()
 q.display()
 
}
