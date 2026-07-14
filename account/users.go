package account

import (
	"fmt"
	"time"
	u "zerago/utils"

	"github.com/dgrijalva/jwt-go"
)

type Token struct {
	UserId int
	jwt.StandardClaims
}
type User struct {
	ID                int       `json:"id"`
	Username          string    `json:"username"`
	Password          string    `json:"password"`
	Email             string    `json:"email"`
	Role              string    `json:"role"`
	Token             string    `json:"token"`
	Wallet            float64   `json:"wallet"`
	DeviceToken       string    `json:"device_token"`
	LastUpdate        time.Time `json:"last_update"`
	Created           time.Time `json:"created"`
	LastLogin         time.Time `json:"last_login"`
	Fullname          string    `json:"fullname"`
	Title             string    `json:"title"`
	PagePermissions   []string  `sql:",array" json:"page_permissions"`
	ActionPermissions []string  `sql:",array" json:"action_permissions"`
	ContactNo         string    `json:"contact_no"`
	Status            string    `json:"status"`
	Branch            int       `json:"branch"`
	FranchiseeID      int       `json:"franchisee_id"`
}
type UserListRes struct {
	ID           int       `json:"id"`
	Username     string    `json:"username"`
	Role         string    `json:"role"`
	Created      time.Time `json:"created"`
	LastLogin    time.Time `json:"last_login"`
	Fullname     string    `json:"fullname"`
	Title        string    `json:"title"`
	Status       string    `json:"status"`
	FranchiseeID int       `json:"franchisee_id"`
	Area         string    `json:"area"`
	Branch       string    `json:"branch"`
	Province     string    `json:"province"`
}
type UserRoles struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
}
type LoginTellerReq struct {
	Username string `json:"username"`
	Password string `json:"password"`
	BranchID int    `json:"branch_id"`
}

type Login struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

var CURRENTUSERID = 0

func RidersList() map[string]interface{} {
	var entries []UserListRes
	_, err := DBM.Query(&entries, `SELECT id, username, role, created, last_login, fullname, title, status from users where role=? order by id desc`, "Rider")
	if err != nil {
		panic(err)
		return u.Message(false, "Connection error. Please retry")
	}
	response := u.Message(true, "Entries")
	response["list"] = entries
	return response
}
func UserList() map[string]interface{} {
	var entries []UserListRes
	_, err := DBM.Query(&entries, `select u.id, username, role, created, last_login, fullname, title, u.status, f.area, b.name as branch ,p.name as province
	from users u left join franchisees f on f.id=u.franchisee_id 
	left join branches b on b.id=u.branch
	left join branches p on p.id=b.provincial order by id desc`)
	if err != nil {
		panic(err)
		return u.Message(false, "Connection error. Please retry")
	}
	response := u.Message(true, "Entries")
	response["list"] = entries
	return response
}

func (user *User) Validate() (map[string]interface{}, bool) {
	var p User
	_, err := DBM.Query(&p, `SELECT * FROM users where username=?`, user.Username)
	if err != nil {
		return u.Message(false, "Failed"), false
	}
	if p.ID != 0 {
		return u.Message(false, "Username already in use by another user."), false
	}
	_, err2 := DBM.Query(&p, `SELECT * FROM users where email=?`, user.Email)
	if err2 != nil {
		return u.Message(false, "Failed."), false
	}
	if p.ID != 0 {
		return u.Message(false, "Email already in use by another user."), false
	}

	return u.Message(false, "Requirement passed"), true
}

func GetRoles() map[string]interface{} {
	var roles []UserRoles

	_, err := DBM.Query(&roles, `SELECT * FROM user_roles order by id asc`)
	if err != nil {
		panic(err)
		return u.Message(false, "Connection error. Please retry")
	}

	response := u.Message(true, "Results")
	response["roles"] = roles
	return response
}
func (user *User) Create() map[string]interface{} {
	fmt.Println("create user!")
	if resp, ok := user.Validate(); !ok {
		return resp
	}
	user.Password = u.HashAndSalt([]byte(user.Password))
	user.Created = time.Now()
	user.LastUpdate = time.Now()
	//fmt.Println(user)
	errdb := DBM.Insert(user)
	if errdb != nil {
		panic(errdb)
		return u.Message(false, "Failed to create account, connection error")
	}
	user.Password = ""
	response := u.Message(true, "Account has been created")
	response["account"] = user
	return response
}

func (user *User) GetByID() map[string]interface{} {
	var account User

	_, err := DBM.Query(&account, `SELECT * FROM users where id=?`, user.ID)
	if err != nil {
		panic(err)
		return u.Message(false, "Connection error. Please retry")
	}
	if account.ID == 0 {
		return u.Message(false, "No account found!")
	}
	account.Password = ""
	response := u.Message(true, "Account Details")
	response["account"] = account
	return response
}
func (user *User) Update() map[string]interface{} {
	fmt.Println("Update user!")
	account := GetUserMeta(user.ID)
	user.Password = account.Password
	user.LastUpdate = time.Now()

	errdb := DBM.Update(user)
	if errdb != nil {
		panic(errdb)
		return u.Message(false, "Failed to create account, connection error")
	}
	response := u.Message(true, "Account has been updated")
	return response
}

func (user *User) ChangePassword() map[string]interface{} {
	fmt.Println("ChangePassword", user.ID)
	var account User
	_, err := DBM.Query(&account, `SELECT * FROM users where id=?`, user.ID)
	if err != nil {
		panic(err)
		return u.Message(false, "Connection error. Please retry")
	}
	if account.ID == 0 {
		return u.Message(false, "No account found!")
	}
	account.Password = u.HashAndSalt([]byte(user.Password))
	account.LastUpdate = time.Now()
	errdb := DBM.Update(&account)
	if errdb != nil {
		panic(errdb)
		return u.Message(false, "Failed to create account, connection error")
	}
	response := u.Message(true, "New password saved!")
	return response
}
func GetUserMeta(user int) *User {
	var account User
	_, err := DBM.Query(&account, `SELECT * FROM users where id=?`, user)
	if err != nil {
		//panic(err)
		return &account
	}
	return &account
}

func (user *LoginTellerReq) LoginTeller() map[string]interface{} {
	fmt.Println("LoginTeller: ", user.Username, user.Password)
	var account User
	_, err := DBM.Query(&account, `SELECT * FROM users where username=? and branch=?`, user.Username, user.BranchID)
	if err != nil {
		panic(err)
		return u.Message(false, "Connection error. Please retry")
	}
	if account.ID == 0 {
		return u.Message(false, "Teller not found")
	}
	fmt.Println("Login: ", user.Username, user.Password)
	if !u.ComparePasswords(account.Password, []byte(user.Password)) {
		return u.Message(false, "Invalid password. Please try again")
	}
	account.LastLogin = time.Now()
	errdb := DBM.Update(&account)
	if errdb != nil {
		panic(errdb)
		return u.Message(false, "Failed to create account, connection error")
	}
	account.Password = ""
	//Create JWT token
	tk := &Token{UserId: account.ID}
	token := jwt.NewWithClaims(jwt.GetSigningMethod("HS256"), tk)
	tokenString, _ := token.SignedString([]byte("token_mydog_app"))
	account.Token = tokenString //Store the token in the response
	resp := u.Message(true, "Logged-in successful")
	resp["account"] = account
	return resp
}
func (user *User) Login() map[string]interface{} {
	fmt.Println("Login: ", user.Username, user.Password)
	var account User
	_, err := DBM.Query(&account, `SELECT * FROM users where username=?`, user.Username)
	if err != nil {
		panic(err)
		return u.Message(false, "Connection error. Please retry")
	}
	if account.ID == 0 {
		return u.Message(false, "Email not found")
	}
	fmt.Println("Login: ", user.Username, user.Password)
	if !u.ComparePasswords(account.Password, []byte(user.Password)) {
		return u.Message(false, "Invalid password. Please try again")
	}
	account.LastLogin = time.Now()
	errdb := DBM.Update(&account)
	if errdb != nil {
		panic(errdb)
		return u.Message(false, "Failed to create account, connection error")
	}
	account.Password = ""
	//Create JWT token
	tk := &Token{UserId: account.ID}
	token := jwt.NewWithClaims(jwt.GetSigningMethod("HS256"), tk)
	tokenString, _ := token.SignedString([]byte("token_mydog_app"))
	account.Token = tokenString //Store the token in the response
	resp := u.Message(true, "Logged-in successful")
	resp["account"] = account
	return resp
}
func UpdateLogin(userID int) {
	_, err := DBM.Exec(`update users set last_login=? where id=?`, time.Now(), userID)
	if err != nil {
		panic(err)
	}
}
