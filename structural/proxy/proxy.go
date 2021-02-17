package proxy

import (
	"errors"
	"fmt"
)

type UserFinder interface {
	FindUser(id int) (User,error)
}

type User struct{
	Id int
}

type UserList []User

type UserProxyList struct{
	SomeDatabase UserList
	StackCache UserList
	StackCapacity int
	DidLastSearchUsedCache bool
}

func (t *UserList) UserFinder(id int) (User,error){
	for i:=0; i<len(*t); i++{
		if (*t)[i].Id==id{
			return (*t)[i],nil
		}
	}
	return User{},errors.New(fmt.Sprintf("user %d could not be found",id))
}

func (u *UserProxyList) UserFinder(id int) (User,error){
	user, err := u.StackCache.UserFinder(id)
	if err == nil {
		fmt.Println("returning user from cache")
		u.DidLastSearchUsedCache = true
		return user, nil
	}

	//if the error was not nil, it means that it couldn't find the user in the stack. So,search in the database:
	 user, err = u.SomeDatabase.UserFinder(id)
	 if err == nil {
		 fmt.Println("Returning user from database")
		 u.addToStack(user)
		 u.DidLastSearchUsedCache = false
		 return user, nil
	 }
	 return User{}, err
}


//When the user is found , then add the user to the stack.
func (u *UserProxyList)addToStack(user User){
	if len(u.StackCache) >= u.StackCapacity{    //if stackCache is full or exceed capacity then remove first user and append new user
		u.StackCache = append(u.StackCache[1:],user)
	} else{
		u.StackCache.addUser(user)
	}
}

func (t *UserList)addUser(newUser User){
	 *t = append(*t,newUser)
}

