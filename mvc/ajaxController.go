package mvc

import (
	"encoding/json"
	"log"
	"net/http"
	// "io/ioutil"
	//"bufio"
    //"os"
    //"fmt"
    //"time"
	//"strconv"
	//"strings"
    "github.com/ziutek/mymysql/mysql"
    _ "github.com/ziutek/mymysql/thrsafe"
)

type user struct {
	Name  string
	Password string
}

type Admin struct{
	Name string   `json:"username"`
}

type AdminSlice struct{
	Admin_array []Admin
}

type Result struct {
	Ret    int //hey
	Reason string
	Data   interface{}
}


type ajaxController struct {
}

func (this *ajaxController) SignupAction(w http.ResponseWriter, r *http.Request) {

}
func (this *ajaxController) LoginAction(w http.ResponseWriter, r *http.Request) {
log.Println("In ajaxController getting logging")
	w.Header().Set("content-type", "application/json")

	db := mysql.New("tcp", "", "localhost:3306", "root", "wwcl2016", "Administrator")
 	err := db.Connect()
	if err != nil {
		log.Println(err)
		OutputJson(w, 0, "数据库连接失败", nil)
		return
	}
	defer db.Close()


	log.Println("body is",r.Body)
	var U user
	err = json.NewDecoder(r.Body).Decode(&U)	// body, err := ioutil.ReadAll(r.Body)

	if err != nil {
		log.Println("error:", err)
	}

	log.Println(U)
	var admin_name string
	var admin_password string

	admin_name = U.Name
	admin_password = U.Password

	log.Println("admin_name is:", admin_name, "admin_password is:",admin_password)
	
	if admin_name == "" || admin_password == "" {
		OutputJson(w, 0, "帐号或密码不能回空", nil)
		return
	}

	message := ""
	rows, res, err := db.Query("select * from Users where name = '%s'", admin_name)
	if err != nil {
		log.Println(err)
		message = "Query failed"
	}
	if rows == nil {
		message = "找不到用户："+admin_name
	}

	name := res.Map("password")	//returns the index of column :"admin_password"
	admin_password_db := rows[0].Str(name)

	Flag := true
	if admin_password_db != admin_password {
		Flag = false
	}

	if Flag == false{
		OutputJson(w, 0, message, nil)
	}else{
		OutputJson(w, 1, message, nil)
	}

}

func (this *ajaxController) GetAdminsAction(w http.ResponseWriter, r *http.Request) {
	log.Println("In ajaxController getting admins")
	w.Header().Set("content-type", "application/json")
	err := r.ParseForm()
	if err != nil {
		OutputJson(w, 0, "参数错误", nil)
		return
	}	

	db := mysql.New("tcp", "", "localhost:3306", "root", "wwcl2016", "Administrator")
 	err = db.Connect()
	if err != nil {
		log.Println(err)
		OutputJson(w, 0, "数据库连接失败", nil)
		return
	}
	defer db.Close()

	rows, _, err := db.Query("select * from Admins")
	if err != nil {
		log.Println(err)
		OutputJson(w, 0, "数据库操作失败", nil)
		return
	}else{
		log.Println("连接成功")
	}

	var Slice AdminSlice

	for _, row := range rows {
		Admin := Admin{}
		Admin.Name = row.Str(0)	
		Slice.Admin_array = append(Slice.Admin_array, Admin)
   	}
   	body, err := json.Marshal(Slice)
	if err != nil {
	    panic(err.Error())
	    return
	}
	w.Write(body)
	return

}

func OutputJson(w http.ResponseWriter, ret int, reason string, i interface{}) {
	out := &Result{ret, reason, i}
	b, err := json.Marshal(out)
	if err != nil {
		return
	}
	w.Write(b)
}
