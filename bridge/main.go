package main

import (
    "fmt"
)

type Implementor interface{
    OperationImp()
}

type ConcreteImplementorA struct{}

func (this *ConcreteImplementorA) OperationImp() {
    fmt.Println("ConcreteImplementor A")
}

type ConcreteImplementorB struct{}

func (this *ConcreteImplementorB) OperationImp() {
    fmt.Println("ConcreteImplementor B")
}

type Abstraction struct{
    imp Implementor
}

func (this *Abstraction) SetImplementor(i Implementor){
    this.imp = i
}

func (this *Abstraction) Operation(){
    fmt.Print("Abstraction: ")
    this.imp.OperationImp()
}

type RedefinedAbsraction struct{
    Abstraction
}

func (this *RedefinedAbsraction) Operation(){
    fmt.Print("RedefinedAbsraction: ")
    this.imp.OperationImp()
}

func main(){
    ab := new(RedefinedAbsraction)
    ab.SetImplementor(new(ConcreteImplementorA))
    ab.Operation()
    ab.SetImplementor(new(ConcreteImplementorB))
    ab.Operation()
}
