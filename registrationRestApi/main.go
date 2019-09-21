package main 
import (
	"encoding/json"
	"fmt"
	"github.com/gorilla/mux"
	"log"
	"net/http"
)
type User struct{
	Email string `json:"email"`
	Password string `json:"password"`
	First_name string `json:"first_name"`
	Last_name string `json:"last_name"`
	City  string `json:"city"`
	Country string `json:"country"`
	Phoneno int64 `json:"phone_no"`
	Salary  int64 `json:"salary"`

}
var user []User

func GetUserDetailurl(w http.ResponseWriter , req *http.Request){
	params :=mux.Vars(req)
	for _,item := range user{
		if item.Email==params["email"]{
			json.NewEncoder(w).Encode(item)
			return
		}
	}
	json.NewEncoder(w).Encode(&User{})
}
func CreateUserDetailurl(w http.ResponseWriter , req *http.Request){
	//params :=mux.Vars(req)
	var users User
	_ = json.NewDecoder(req.Body).Decode(&users)
	//users.Email=params["email"]
	user = append(user,users)
    json.NewEncoder(w).Encode(user)
}
func DeleteUserDetailurl(w http.ResponseWriter , req *http.Request){
  params :=mux.Vars(req)
  for index,item :=range user{
	  if item.Email == params["email"]{
		  user = append(user[:index], user[index+1:]...)
		  break
	  }
	  json.NewEncoder(w).Encode(user)
  }	
}
func main() {
	 router:=mux.NewRouter()
	 user =append(user,User{Email:"sharad@gmail.com",Password:"Test",Phoneno:12344,Salary:3456,City:"Delhi",Country:"India",First_name:"Sharad",Last_name:"Maurya",})
	 router.HandleFunc("/getUser/{email}",GetUserDetailurl).Methods("GET")
	 router.HandleFunc("/createUser",CreateUserDetailurl).Methods("POST")
	 router.HandleFunc("/deleteUser/{email}",DeleteUserDetailurl).Methods("DELETE")
	 log.Fatal(http.ListenAndServe(":8002",router))
	 fmt.Print("TestRest")
}
