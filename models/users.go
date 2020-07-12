package models

import (
	"fmt"
)

type User struct {
	Id         int    `json:"id"`
	Name       string `json:"name"`
	Age        int    `json:"age"`
	AvatarPath string `json:"avatarPath"`
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

func ListAllUsers() []User {
	var users []User
	sql := "SELECT * FROM user"
	rows, err := db.Query(sql)
	if err != nil {
		fmt.Println("query error", err)
		return nil
	}
	defer rows.Close()

	for rows.Next() {
		var user User
		err := rows.Scan(&user.Id, &user.Name, &user.Age, &user.AvatarPath)
		if err != nil {
			fmt.Println("scan error", err)
		}
		users = append(users, user)
	}

	return users
}
