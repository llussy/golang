package main

import "fmt"

type User struct {
	ID   int64
	Name string `gorm:"default:'galeone'"`
	Age  int64
}

func main() {
	DBInit()
	defer DB.Close()

	if !DB.HasTable(&User{}) {
		if err := DB.CreateTable(&User{}).Error; err != nil {
			fmt.Println("create table failed!", err)
			return
		}
		fmt.Println("create table success")
		return
	}

	//user := User{ID: 1, Name: "test", Age: 27}
	//DB.Create(&user)

	user := []User{}
	DB.Find(&user)
	for _, v := range user {
		fmt.Println("Name: ", v.Name, "Age: ", v.Age)
	}

}
