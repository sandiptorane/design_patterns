package proxy

import (
	"math/rand"
	"testing"
)

func TestUserProxyList_UserFinder(t *testing.T) {
	someDatabase := UserList{}

	rand.Seed(2445648)
	for i:=0;i<1000000;i++{
		num := rand.Int()
		someDatabase = append(someDatabase,User{Id:num})
	}

	proxy := UserProxyList{
		SomeDatabase: someDatabase,
		StackCapacity: 2,
		StackCache: UserList{},
	}

	knownIDs := [3]int{someDatabase[0].Id,someDatabase[1].Id,someDatabase[2].Id}

	t.Run("Find User - Empty Cache",func(t *testing.T){
		user,err := proxy.UserFinder(knownIDs[0])

		if err!=nil{
			t.Error(err)
		}

		if user.Id != knownIDs[0]{
			t.Error("Returned user Id doesn't match with expected")
		}

		if len(proxy.StackCache) != 1 {
			t.Error("After one successful search in an empty cache, the size of it must be one")
		}
		if proxy.DidLastSearchUsedCache {
			t.Error("No user can be returned from an empty cache")
		}
	})

	t.Run("Find User - Empty Cache for the same user",func(t *testing.T){
		user,err := proxy.UserFinder(knownIDs[0])

		if err!=nil{
			t.Error(err)
		}

		if user.Id != knownIDs[0]{
			t.Error("Returned user Id doesn't match with expected")
		}

		if len(proxy.StackCache) != 1 {
			t.Error("Cache must not grow if we asked for an object that is stored on it")
		}
		if !proxy.DidLastSearchUsedCache {
			t.Error("The user should have been returned from the cache")
		}
	})

	t.Run("Overflowing the stack", func(t *testing.T) {
		user1 ,err := proxy.UserFinder(knownIDs[0])
		if err !=nil{
			t.Error(err)
		}

		user2 ,err := proxy.UserFinder(knownIDs[1])
		if proxy.DidLastSearchUsedCache{
			t.Error("the user wasn't stored on the proxy cache yet")
		}

		user3, err := proxy.UserFinder(knownIDs[2])
		if proxy.DidLastSearchUsedCache{
			t.Error("the user wasn't stored on the proxy cache yet")
		}

		for i:=0; i< len(proxy.StackCache);i++{
			if proxy.StackCache[i].Id == user1.Id{
				t.Error("the user1 should be gone was found")
			}
		}

		if len(proxy.StackCache) !=2 {
			t.Error("After inserting 3 users the cache should not be grown more than 2")
		}

		for _,v := range proxy.StackCache{
			if v != user2 && v != user3 {
				t.Error("The non expected users found on the cache")
			}
		}
	})

}

