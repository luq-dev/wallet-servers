package routing

import (
	"encoding/json"
	"fmt"
	"net/http"
	"servers/data"
)

// Hello World

func hello_world(w http.ResponseWriter, req *http.Request) {
	fmt.Fprint(w, "Hello World\n")
}

// Adding Users

func addUser(w http.ResponseWriter, req *http.Request){
	var user data.UserDAO

	if err := json.NewDecoder(req.Body).Decode(&user); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	
	if (user.HasMinDetails()) {
		stmt, err := data.Db.Prepare("INSERT INTO users(fullname, email, password, phone_number) values (?,?,?,?)")
		if err != nil {
			fmt.Println(err.Error())
		}
		res, err := stmt.Exec(user.Name, user.Email, user.Password, user.Phone_number)

		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
		}else {
			fmt.Println(res)
		}
	} else {
		http.Error(w, "Missing Important data", http.StatusBadRequest)
		return
	}

	fmt.Fprint(w, "User Added")
}


