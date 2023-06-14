package main

import "fmt"

//composing user with person
type Person struct {
	Name string
}

func (person *Person) introduce() {
	fmt.Printf("My name is %s", person.Name)
}

type User struct {
	*Person
	Age    int
	Father *User
}

func (user *User) increaseAge() {
	user.Age += 10
}

//factory function to create instance of user
func userFactory(name string, age int, father *User) *User {
	return &User{
		Person: &Person{name},
		Age:    age,
		Father: father,
	}
}

func main() {
	user1 := userFactory("Ashik", 27, nil)
	user1.increaseAge()
	user1.increaseAge()
	user1.introduce()
	user2 := userFactory("child", 13, user1)
	fmt.Printf("name and age of user %s,%d \n", user1.Name, user1.Age)
	fmt.Printf("name and age of user %s,%d \n", user2.Name, user2.Age)
	fmt.Printf("Father name %s and age %d", user2.Father.Name, user2.Father.Age)

}
