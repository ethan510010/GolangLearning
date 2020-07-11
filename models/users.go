package models

import "fmt"

type User struct {
	Id         int
	Name       string
	Age        int
	AvatarPath string
}

func CreateUserProfile(u *User) int64 {
	sql := "INSERT INTO user (name, age, avatarPath) values (?, ?, ?)"
	stmt, err := db.Prepare(sql)

	if err != nil {
		fmt.Println("prepare failed", err)
	}
	defer stmt.Close()
	result, err := stmt.Exec(u.Name, u.Age, u.AvatarPath)
	if err != nil {
		fmt.Println("insert failed", err)
		return 0
	}
	insertID, _ := result.LastInsertId()
	fmt.Println("insert id ", insertID)
	return insertID
}
