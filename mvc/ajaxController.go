package mvc

import (
	"encoding/json"
	"log"
	"net/http"
	//"io/ioutil"
	//"bufio"
    //"os"
    //"fmt"
    //"time"
	//"strconv"
	//"strings"
    "github.com/ziutek/mymysql/mysql"
    _ "github.com/ziutek/mymysql/thrsafe"
)

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

func (this *ajaxController) LoginAction(w http.ResponseWriter, r *http.Request) {
log.Println("In ajaxController getting logging")
	w.Header().Set("content-type", "application/json")
	// err := r.ParseForm()
	// if err != nil {
	// 	OutputJson(w, 0, "参数错误", nil)
	// 	return
	// }	

	db := mysql.New("tcp", "", "localhost:3306", "root", "wwcl2016", "Administrator")
 	err := db.Connect()
	if err != nil {
		log.Println(err)
		OutputJson(w, 0, "数据库连接失败", nil)
		return
	}
	defer db.Close()

	var admin_name string
	var admin_password string

	log.Println("57")
	//if r.Method == "POST"{ 
		admin_name = r.FormValue("name")
		admin_password = r.FormValue("password")
	log.Println("61")
	//}else{
	//	admin_name = r.Form["user.name"][0]
	//	admin_password = r.Form["user.password"][0]
	//}
	
	if admin_name == "" || admin_password == "" {
		OutputJson(w, 0, "帐号或密码不能回空", nil)
		return
	}

	log.Println("admin_name is:", admin_name, "admin_password is:",admin_password)

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
