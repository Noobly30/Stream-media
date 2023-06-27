package dbops

import (
	"strconv"
	"testing"
	"time"
	"fmt"
	"golang.org/x/tools/go/analysis/passes/nilfunc"
)

var tempvid string

func clearTables(){
	dbConn.Exec("truncate users")
	dbConn.Exec("truncate video_info")
	dbConn.Exec("truncate comments")
	dbConn.Exec("truncate sessions")
}

func TestMain(m *testing.M) {
	clearTables()
	m.Run()
	clearTables()
}

func TestUserWorkFlow(t *testing.T){
	t.Run("Add", testAddUser)
	t.Run("Get",testGetUser)
	t.Run("Del",testDeleteUser)
	t.Run("Reget",testRegetUser)
}

func testAddUser(t *testing.T){
	err := AddUserCredential("avenssi","123")
	if err != nil{
		t.Errorf("Error of Adduser: %v",err)
	}
}

func testGetUser(t *testing.T){
	pwd, err := GetUserCredential("avenssi")
	if err != nil{
		t.Error("Error of GetUser: %v", err)
	}
}

func testDeleteUser(t *testing.T){
	err := DeleteUser("avenssi","123")
	if err != nil{
		t.Errorf("Error of DeleteUser: %v", err)
	}
}

func testRegetUser(t *testing.T){
	pwd, err := GetUserCredential("avenssi")
	if err != nil{
		t.Errorf("Error of RegetUser: %v",err)
	}
	if pwd != ""{
		t.Errorf("Deleting user test failed")
	}
}

func TestComments(t *testing.T){
	clearTables()
	t.Run("AddUser",test)
	t.Run("AddComments",testAddComments)
	t.Run("ListComments",testListComments)
}

func testAddComments(t *testing.T){
	vid := "12345"
	aid := 1
	content := "I like this video"

	err := AddNewComments(vid,aid,content)

	if err != nil{
		t.Errorf("Error of AddComments: %v",err)
	}
}

func testListComments(t *testing.T){
	vid = "12345"
	from := 1514764800
	to, _ := strconv.Atoi(strconv.FormatInt(time.Now().UnixNano()/1000000000,10))

	res, err := ListComments(vid,from,to)
	if err != nil{
		t.Errorf("Error of LisiComments:%v",err)
	}
	for i, ele := range res {
		fmt.Printf("comment: %d, %v\n",i,ele)
	}
}