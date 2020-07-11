// package main

// import (
// 	"database/sql"
// 	"fmt"

// 	_ "github.com/go-sql-driver/mysql"
// )

// var db *sql.DB

// func initDB() (err error) {

// 	dsn := "root:Ethan0909@tcp(127.0.0.1:3306)/sql_test"
// 	db, err = sql.Open("mysql", dsn)
// 	if err != nil {
// 		return err
// 	}
// 	err = db.Ping()
// 	if err != nil {
// 		return err
// 	}
// 	return nil
// }

// type User struct {
// 	id   int
// 	name string
// 	age  int
// }

// func queryRowDemo() {
// 	sqlStr := "select id, name, age from user where id=?"
// 	var u User
// 	// 非常重要：确保QueryRow之后调用Scan方法，否则持有的数据库链接不会被释放
// 	err := db.QueryRow(sqlStr, 1).Scan(&u.id, &u.name, &u.age)
// 	if err != nil {
// 		fmt.Printf("scan failed, err:%v\n", err)
// 		return
// 	}
// 	fmt.Printf("id:%d name:%s age:%d\n", u.id, u.name, u.age)
// }

// func queryMultipleRows() {
// 	sqlStr := "select * from user where id > ?"
// 	rows, err := db.Query(sqlStr, 0)
// 	if err != nil {
// 		fmt.Printf("query failed, err:%v\n", err)
// 	}
// 	defer rows.Close()

// 	for rows.Next() {
// 		var u User
// 		err := rows.Scan(&u.id, &u.name, &u.age)
// 		if err != nil {
// 			fmt.Printf("scan failed, err:%v\n", err)
// 		}
// 		fmt.Printf("id:%d name:%s age:%d\n", u.id, u.name, u.age)
// 	}
// }

// func insertRowDemo() {
// 	sqlStr := "insert into user (name, age) values (?, ?)"
// 	result, err := db.Exec(sqlStr, "PengPeng", 36)
// 	if err != nil {
// 		fmt.Println("insert failed", err)
// 	}
// 	insertID, err := result.LastInsertId()
// 	if err != nil {
// 		fmt.Println("get last insert id failed", err)
// 		return
// 	}
// 	fmt.Printf("insert id %d\n", insertID)
// }

// func updateRowDemo() {
// 	sqlStr := "update user set age = ? where id = ?"
// 	ret, err := db.Exec(sqlStr, 31, 2)
// 	if err != nil {
// 		fmt.Println("update failed", err)
// 		return
// 	}
// 	n, err := ret.RowsAffected()
// 	if err != nil {
// 		fmt.Println("get affected rows failed", err)
// 		return
// 	}
// 	fmt.Printf("update success, affected rows: %d\n", n)
// }

// func deleteDemo() {
// 	sqlStr := "delete from user where id = ?"
// 	ret, err := db.Exec(sqlStr, 3)
// 	if err != nil {
// 		fmt.Printf("delete failed, err:%v\n", err)
// 		return
// 	}
// 	n, err := ret.RowsAffected() // 操作影响的行数
// 	if err != nil {
// 		fmt.Printf("get RowsAffected failed, err:%v\n", err)
// 		return
// 	}
// 	fmt.Printf("delete success, affected rows:%d\n", n)
// }

// func main() {
// 	err := initDB()
// 	if err != nil {
// 		fmt.Printf("init db failed:%v\n", err)
// 		return
// 	}
// 	deleteDemo()
// }

package main

import (
	"fmt"
	"net/http"

	"GolangAPIPractice/routes"
)

func main() {
	router := routes.NewRouter()
	err := http.ListenAndServe(":3000", router)
	if err != nil {
		fmt.Println("start server error", err)
	}
	// http.HandleFunc("/createUser/", uploadHandle) // 上传
	// http.HandleFunc("/uploaded/", showPicHandle)  //显示图片
	// err := http.ListenAndServe(":3000", nil)
	// fmt.Println(err)
}

// func NewRouter() *mux.Router {
// 	router := mux.NewRouter()

// 	return router
// }

// 上传图像接口
// func uploadHandle(w http.ResponseWriter, req *http.Request) {
// 	w.Header().Set("Content-Type", "text/html")

// 	req.ParseForm()
// 	if req.Method != "POST" {
// 		w.Write([]byte(html))
// 	} else {
// 		// 接收图片
// 		uploadFile, handle, err := req.FormFile("image")
// 		errorHandle(err, w)

// 		// 检查图片后缀
// 		ext := strings.ToLower(path.Ext(handle.Filename))
// 		if ext != ".jpeg" && ext != ".png" {
// 			errorHandle(errors.New("只支持jpg/png图片上传"), w)
// 			return
// 			//defer os.Exit(2)
// 		}

// 		// 保存图片
// 		os.Mkdir("./uploaded/", 0777)
// 		saveFile, err := os.OpenFile("./uploaded/"+handle.Filename, os.O_WRONLY|os.O_CREATE, 0666)
// 		errorHandle(err, w)
// 		io.Copy(saveFile, uploadFile)

// 		defer uploadFile.Close()
// 		defer saveFile.Close()
// 		// 上传图片成功
// 		w.Write([]byte("upload link: <a target='_blank' href='/uploaded/" + handle.Filename + "'>" + handle.Filename + "</a>"))
// 	}
// }

// // 显示图片接口
// func showPicHandle(w http.ResponseWriter, req *http.Request) {
// 	file, err := os.Open("." + req.URL.Path)
// 	errorHandle(err, w)

// 	defer file.Close()
// 	buff, err := ioutil.ReadAll(file)
// 	errorHandle(err, w)
// 	w.Write(buff)
// }

// // 统一错误输出接口
// func errorHandle(err error, w http.ResponseWriter) {
// 	if err != nil {
// 		w.Write([]byte(err.Error()))
// 	}
// }

// const html = `<html>
//     <head></head>
//     <body>
//         <form method="post" enctype="multipart/form-data">
//             <input type="file" name="image" />
//             <input type="submit" />
//         </form>
//     </body>
// </html>`
