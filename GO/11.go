package main

import "fmt"

type Model interface{
	save()
	update()
	delete()
	read()
}

type User struct {
	name string
	age int
}
type Produt struct {
	name string
	Price int
}

func (user User) save(){
	fmt.Println("the user is save in db")
}
func (user Produt) save(){
	fmt.Println("the user is save in db")
}
func (user User) update(){
	fmt.Println("the user is updated in db")
}
func (user Produt) update(){
	fmt.Println("the user is updated in db")
}
func (user User) delete(){
	fmt.Println("the user is deleted from the db")
}
func (user Produt) delete(){
	fmt.Println("the user is deleted from the db")
}
func (user User) read(){
	fmt.Println("the user is readed from the db")
}
func (user Produt) read(){
	fmt.Println("the user is readed from the db")
}

func save(m Model){
	m.save()
}
func update(m Model){
	m.update()
}
func delete(m Model){
	m.delete()
}
func read(m Model){
	m.read()
}

func main(){
	user:=User{"Sanket",21}
	save(user)
	Produt:=Produt{"Sanket",21}
	update(Produt)
	// user:=User{"Sanket",21}
	// save(user)
	// user:=User{"Sanket",21}
	// save(user)
}

