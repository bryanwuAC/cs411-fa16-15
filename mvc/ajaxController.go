package mvc

import (
	"encoding/json"
	"log"
	"net/http"
	"io/ioutil"
	// "strconv"
	//"bufio"
    //"os"
    //"fmt"
    //"time"
	//"strconv"
	"strings"
    "github.com/ziutek/mymysql/mysql"
    _ "github.com/ziutek/mymysql/thrsafe"
)

type user struct {
	Name  string
	Password string
}
type username struct {
	Name  string
}

type like struct{
	Name string
	Favorite string
}
type search struct {
	Search_option string
	Keyword_option string
	Search_text string
}

type Paper struct{
	Title string 
	URL string  
}
type RecommendPaper struct{
	Title string 
	URL string  
	Author string
	Tag string
}
type Counter struct{
	Num string
}
type PaperSlice struct{
	Paper_array []Paper
}
type PaperSlices struct{
	Paper_array []RecommendPaper
	Paper_array_2 []Paper
}

type Result struct {
	Ret    int //hey
	Reason string
	Data   interface{}
}


type ajaxController struct {
}

func (this *ajaxController) ChangePasswordAction(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("content-type", "application/json")

	db := mysql.New("tcp", "", "localhost:3306", "root", "wwcl2016", "Administrator")
 	err := db.Connect()
	if err != nil {
		log.Println(err)
		OutputJson(w, 0, "failed to connect to", nil)
		return
	}

	defer db.Close()

	log.Println("body is",r.Body)
	var U user
	err = json.NewDecoder(r.Body).Decode(&U)	// body, err := ioutil.ReadAll(r.Body)

	if err != nil {
		log.Println("error:", err)
	}

	admin_name := U.Name
	admin_password := U.Password

	rows, _, err := db.Query("select * from Users where name = '%s'", admin_name)
	if rows == nil {
		OutputJson(w, 0, "user name does not exists", nil)
		return
	}

	_, _, err = db.Query("UPDATE Users SET password = '%s' where name = '%s'", admin_password, admin_name)

	OutputJson(w, 1, "Update successful!", nil)
	log.Println("out ajaxController")
	return
}


func (this *ajaxController) DeleteAccountAction(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("content-type", "application/json")

	db := mysql.New("tcp", "", "localhost:3306", "root", "wwcl2016", "Administrator")
 	err := db.Connect()
	if err != nil {
		log.Println(err)
		OutputJson(w, 0, "failed to connect to", nil)
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
	admin_name := U.Name

	rows, _, err := db.Query("select * from Users where name = '%s'", admin_name)
	if rows == nil {
		OutputJson(w, 0, "user name does not exists", nil)
		return
	}
	_, _, err = db.Query("DELETE FROM Users where name = '%s'", admin_name)


	OutputJson(w, 1, "Delete successful!", nil)
	log.Println("out ajaxController")
	return

}
func (this *ajaxController) SignupAction(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("content-type", "application/json")

	db := mysql.New("tcp", "", "localhost:3306", "root", "wwcl2016", "Administrator")
 	err := db.Connect()
	if err != nil {
		log.Println(err)
		OutputJson(w, 0, "failed to connect to", nil)
		return
	}
	defer db.Close()

	log.Println("body is",r.Body)
	var U user
	err = json.NewDecoder(r.Body).Decode(&U)	
	// body, err := ioutil.ReadAll(r.Body)

	if err != nil {
		log.Println("error:", err)
	}

	log.Println(U)

	admin_name := U.Name
	admin_password := U.Password

	_, _, err = db.Query("INSERT INTO Users VALUES ('%s','%s')", admin_name, admin_password)
	if err != nil {
		log.Println(err)
		OutputJson(w, 0, "User name exists", nil)
		return
	}

	OutputJson(w, 1, "Sign up successful!", nil)
	log.Println("out ajaxController")
	return
}
func (this *ajaxController) Like_paperAction(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("content-type", "application/json")

	db := mysql.New("tcp", "", "localhost:3306", "root", "wwcl2016", "dblp_csv")
 	err := db.Connect()
	if err != nil {
		log.Println(err)
		OutputJson(w, 0, "failed to connect to", nil)
		return
	}
	defer db.Close()

	// log.Println("body is",r.Body)
	var L like
	err = json.NewDecoder(r.Body).Decode(&L)	
	if err != nil {
		log.Println("error:", err)
	}

	log.Println(L)

	_, _, err = db.Query("INSERT INTO favorite VALUES ('%s','%s')", L.Name, L.Favorite)
	if err != nil {
		log.Println(err)
		OutputJson(w, 0, "favorite failed", nil)
		return
	}

	OutputJson(w, 1, "favorite successful!", nil)
	log.Println("out ajaxController")
	return

}
func (this *ajaxController) LoginAction(w http.ResponseWriter, r *http.Request) {

	log.Println("In ajaxController getting logging")
	w.Header().Set("content-type", "application/json")

	db := mysql.New("tcp", "", "localhost:3306", "root", "wwcl2016", "Administrator")
 	err := db.Connect()
	if err != nil {
		log.Println(err)
		OutputJson(w, 0, "db connection failed", nil)
		return
	}
	defer db.Close()


	log.Println("body is",r.Body)
	var U user
	err = json.NewDecoder(r.Body).Decode(&U)	

	if err != nil {
		log.Println("error:", err)
	}

	log.Println(U)

	admin_name := U.Name
	admin_password := U.Password

	log.Println("admin_name is:", admin_name, "admin_password is:",admin_password)
	
	if admin_name == "" || admin_password == "" {
		OutputJson(w, 0, "Username or Password could not be NULL", nil)
		return
	}

	rows, res, err := db.Query("select * from Users where name = '%s'", admin_name)

	if rows == nil {
		OutputJson(w, 0, "Sorry, could not find user:"+admin_name, nil)
		return
	}

	name := res.Map("password")	//returns the index of column :"admin_password"
	admin_password_db := rows[0].Str(name)

	if admin_password_db != admin_password {
		OutputJson(w, 0, "Wrong password!", nil)
		return
	}else{
		OutputJson(w, 1, "Success!", nil)
	}
	cookie := http.Cookie{Name: "admin_name", Value: admin_name, Path: "/"}
	http.SetCookie(w, &cookie)

}
func (this *ajaxController) GetFavAndRecAction(w http.ResponseWriter, r *http.Request) {
	log.Println("In ajaxController GetFavAndRecAction")
	w.Header().Set("content-type", "application/json")
	//get Recommendation First 
	db := mysql.New("tcp", "", "localhost:3306", "root", "wwcl2016", "dblp_csv")
 	err := db.Connect()
	if err != nil {
		log.Println(err)
		OutputJson(w, 0, "failed to connect to db", nil)
		return
	}
	defer db.Close()

	// body, _ := ioutil.ReadAll(r.Body)
	// log.Println(string(body))

	var U username
	err = json.NewDecoder(r.Body).Decode(&U)	
	if err != nil {
		log.Println("276 error:", err)
	}
	admin_name := U.Name
	log.Println(admin_name)
	rows, _, err := db.Query("select DISTINCT paper.title, paper.URL, people.name, paper.tag from paper, writtenby, favorite, people where paper.Title not in (select Title from favorite where User = '%s') and paper.id = writtenby.paper and writtenby.PERSON = people.ID and people.name in ( select name from people, paper, writtenby, favorite where favorite.User = '%s' and favorite.Title = paper.Title and paper.id = writtenby.paper and writtenby.PERSON = people.ID ) and paper.tag in ( select tag from people, paper, writtenby, favorite where favorite.User = '%s' and favorite.Title = paper.Title and paper.id = writtenby.paper and writtenby.PERSON = people.ID)",admin_name, admin_name, admin_name)
	print("rec rows length:",len(rows))
	var Slices PaperSlices

	for _, row := range rows {
		Paper := RecommendPaper{}
		Paper.Title = row.Str(0)
		Paper.URL = row.Str(1)	
		Paper.Author = row.Str(2)	
		Paper.Tag = row.Str(3)	
		Slices.Paper_array = append(Slices.Paper_array, Paper)
   	}	

   	rows, _, err = db.Query("select paper.Title, paper.URL from favorite, paper where favorite.User = '%s' and favorite.Title = paper.title", admin_name)
   	print("\nfav rows length:",len(rows))
   	for _, row := range rows {
		Paper := Paper{}
		Paper.Title = row.Str(0)
		Paper.URL = row.Str(1)
		Slices.Paper_array_2 = append(Slices.Paper_array_2, Paper)
   	}	

   	log.Println(len(Slices.Paper_array),len(Slices.Paper_array_2))

	OutputJson(w, 1, "success", Slices)
	return

}
func (this *ajaxController) SearchAction(w http.ResponseWriter, r *http.Request) {
	log.Println("In ajaxController Searching")
	w.Header().Set("content-type", "application/json")

	db := mysql.New("tcp", "", "localhost:3306", "root", "wwcl2016", "dblp_csv")
 	err := db.Connect()
	if err != nil {
		log.Println(err)
		OutputJson(w, 0, "failed to connect to db", nil)
		return
	}
	defer db.Close()

	body, _ := ioutil.ReadAll(r.Body)
	log.Println(string(body))

	var S search
	err = json.Unmarshal(body, &S)	// body, err := ioutil.ReadAll(r.Body)

	if err != nil {
		log.Println("error:", err)
	}

	log.Println(S)

	search_text := S.Search_text
	search_option := S.Search_option
	keyword_option := S.Keyword_option

	log.Println(search_text, search_option, keyword_option)

	var Slice PaperSlice

	if search_option == "1" && keyword_option == "1"{
		var original string
		if strings.Contains(search_text, "'") {
			original = search_text
			search_text = strings.Replace(search_text, "'", "%"+"\\"+"'%", -1)
		}
		rows, _, err := db.Query("select TITLE, URL from paper, writtenby where paper.ID = writtenby.paper and writtenby.PERSON = (select ID from people where Name = '%s')", search_text)

		if err != nil {
			log.Println(err)
			OutputJson(w, 0, "Query execution failed", nil)
			return
		}else{
			log.Println("Query execution succeeded.")
		}	
		if rows == nil{
			OutputJson(w, 0, "No paper of " + original +"has been found!", Slice)
		}
		for _, row := range rows {
			Paper := Paper{}
			Paper.Title = row.Str(0)
			Paper.URL = row.Str(1)		
			Slice.Paper_array = append(Slice.Paper_array, Paper)
   		}	

	}else if search_option == "2" && keyword_option == "1"{
		text := "%"+search_text+ "%"
		rows, _, err := db.Query("select TITLE, URL from paper where TITLE like '%s'",text)
		

		if err != nil {
			log.Println(err)
			OutputJson(w, 0, "Query execution failed", nil)
			return
		}else{
			log.Println("Query execution succeeded.")
		}	

		for _, row := range rows {
			Paper := Paper{}
			Paper.Title = row.Str(0)
			Paper.URL = row.Str(1)	
			Slice.Paper_array = append(Slice.Paper_array, Paper)
   		}	
   		log.Println("array length: ",len(Slice.Paper_array))

   		if len(Slice.Paper_array) > 50{
   			OutputJson(w, 0, "Too many items! Please be more specific!", nil)
   			return
   		}
	}else if search_option == "2" && keyword_option == "2"{
		rows, _, err := db.Query("select TITLE, URL from paper where tag = '%s'",search_text)
		if err != nil {
			log.Println(err)
			OutputJson(w, 0, "Query execution failed", nil)
			return
		}else{
			log.Println("Query execution succeeded.")
		}	
		count := 0
		for _, row := range rows {
			Paper := Paper{}
			Paper.Title = row.Str(0)
			Paper.URL = row.Str(1)	
			Slice.Paper_array = append(Slice.Paper_array, Paper)
			count += 1
			if count >= 100{
				break
			}
   		}	
   		log.Println("array length: ",len(Slice.Paper_array))

	}
	OutputJson(w, 1, "success", Slice)

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
