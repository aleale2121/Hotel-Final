package handler

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"github.com/aleale2121/Hotel-Final/entity"
	"github.com/aleale2121/Hotel-Final/form"
	"github.com/aleale2121/Hotel-Final/permission"
	"github.com/aleale2121/Hotel-Final/rtoken"
	"github.com/aleale2121/Hotel-Final/session"
	"golang.org/x/crypto/bcrypt"
	"html/template"
	"io/ioutil"
	"net/http"
	"net/url"
	"strconv"
	"strings"
)
type UserHandler struct {
	templ        *template.Template
	userSess       *entity.Session
	loggedInUser   *entity.User
	csrfSignKey    []byte
}
type contextKey string

var ctxUserSessionKey = contextKey("signed_in_user_session")

func NewUserHandler(T *template.Template,
	usrSess *entity.Session, csKey []byte) *UserHandler {
	return &UserHandler{templ:T,userSess:usrSess,csrfSignKey:csKey}
}
// Authenticated checks if a user is authenticated to access a given route
func (uh *UserHandler) Authenticated(next http.Handler) http.Handler {
	fn := func(w http.ResponseWriter, r *http.Request) {
		ok := uh.loggedIn(r)
		if !ok {
			http.Redirect(w, r, "/login", http.StatusSeeOther)
			return
		}
		ctx := context.WithValue(r.Context(), ctxUserSessionKey, uh.userSess)
		next.ServeHTTP(w, r.WithContext(ctx))
	}
	return http.HandlerFunc(fn)
}

// Authorized checks if a user has proper authority to access a give route
func (uh *UserHandler) Authorized(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if uh.loggedInUser == nil {
			fmt.Println("53")
			http.Error(w, http.StatusText(http.StatusUnauthorized), http.StatusUnauthorized)
			return
		}
		var roles []entity.Role
		client:=&http.Client{}
		output, err := json.MarshalIndent(uh.loggedInUser, "", "\t\t")
		if err!=nil{
			http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
			return
		}
		req, err := http.NewRequest(http.MethodGet,"http://localhost:9090/user/roles",bytes.NewBuffer(output))
		response,err:= client.Do(req)
		if (err!=nil)||(response.StatusCode==http.StatusNotFound){
			fmt.Println("62")
			http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
			return
		}
		responseData, err := ioutil.ReadAll(response.Body)
		if err != nil {
			fmt.Println("68")
			http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
			return
		}

		err = json.Unmarshal(responseData, &roles)
		if err!=nil {
			fmt.Println("75")
			http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
			return
		}
        fmt.Println(roles)
		for _, role := range roles {
			permitted := permission.HasPermission(r.URL.Path, role.Name, r.Method)
			if !permitted {
				fmt.Println("83")
				http.Error(w, http.StatusText(http.StatusUnauthorized), http.StatusUnauthorized)
				return
			}
		}
		if r.Method == http.MethodPost {
			ok, err := rtoken.ValidCSRF(r.FormValue("_csrf"), uh.csrfSignKey)
			if !ok || (err != nil) {
				fmt.Println("91")
				http.Error(w, http.StatusText(http.StatusUnauthorized), http.StatusUnauthorized)
				return
			}
		}
		next.ServeHTTP(w, r)
	})
}
// LogoutPost  handles the POST /login  requests

func(uh *UserHandler) LoginPostHandler(w http.ResponseWriter, r *http.Request) {
	err := r.ParseForm()
	if err != nil {
		http.Error(w, http.StatusText(http.StatusBadRequest), http.StatusBadRequest)
		return
	}
	loginForm := form.Input{Values: r.PostForm, VErrors: form.ValidationErrors{}}
	client:=http.Client{}
	user:=entity.User{}
	user.Password = r.FormValue("password")
	user.Id=0
	user.Orders=nil
	//user.Roles=nil
	user.Email= r.FormValue("email")
	user.Phone=""
	output, err := json.MarshalIndent(user, "", "\t\t")
	req, err := http.NewRequest("POST","http://localhost:9090/user/check",bytes.NewBuffer(output))
	response, err:=client.Do(req)
	if http.StatusNotFound == response.StatusCode {
		fmt.Println("login54")
		loginForm.VErrors.Add("generic", "Your email address or password is wrong")
		uh.templ.ExecuteTemplate(w, "login.layout", loginForm)
		return
	}
	if http.StatusUnprocessableEntity == response.StatusCode {
		fmt.Println("login60")
		loginForm.VErrors.Add("generic", "Cannot Process Email")
		uh.templ.ExecuteTemplate(w, "login.layout", loginForm)
		return

	}
	responseData, err := ioutil.ReadAll(response.Body)
	if err != nil {
		fmt.Println("login68")
		loginForm.VErrors.Add("generic", "Your email address or password is wrong")
		uh.templ.ExecuteTemplate(w, "login.layout", loginForm)
		return
	}
	var userReturned entity.User
	_ = json.Unmarshal(responseData, &userReturned)
	err = bcrypt.CompareHashAndPassword([]byte(userReturned.Password), []byte(r.FormValue("password")))
	if err == bcrypt.ErrMismatchedHashAndPassword {
		fmt.Println("login77")
		loginForm.VErrors.Add("generic", "Your email address or password is wrong")
		uh.templ.ExecuteTemplate(w, "login.layout", loginForm)
		return
	}
	uh.loggedInUser = &userReturned
	claims := rtoken.Claims(userReturned.Email, uh.userSess.Expires)
	session.Create(claims, uh.userSess.UUID, uh.userSess.SigningKey, w)
	output, err = json.MarshalIndent(uh.userSess, "", "\t\t")
	req, err = http.NewRequest("POST","http://localhost:9090/user/session/new",bytes.NewBuffer(output))
	response, err=client.Do(req)
	if err != nil {
		fmt.Println("login90")
		loginForm.VErrors.Add("generic", "Failed to store session ")
		uh.templ.ExecuteTemplate(w, "login.layout", loginForm)
		return
	}
	responseData, err = ioutil.ReadAll(response.Body)
	if err != nil {
		fmt.Println("login97")
		loginForm.VErrors.Add("generic", "Failed to store session ")
		uh.templ.ExecuteTemplate(w, "login.layout", loginForm)
		return
	}
	var sessionReturned entity.Session
	_ = json.Unmarshal(responseData, &sessionReturned)
	uh.userSess = &sessionReturned
	output, err = json.MarshalIndent(uh.loggedInUser, "", "\t\t")
	req, err = http.NewRequest("GET","http://localhost:9090/user/roles",bytes.NewBuffer(output))
	response, err=client.Do(req)
	if err != nil {
		loginForm.VErrors.Add("generic", "Failed to Authenticate ")
		uh.templ.ExecuteTemplate(w, "login.layout", loginForm)
		return
	}
	responseData, err = ioutil.ReadAll(response.Body)
	if err != nil {
		loginForm.VErrors.Add("generic", "Failed to  Authorize ")
		uh.templ.ExecuteTemplate(w, "login.layout", loginForm)
		return
	}
	var roleReturned []entity.Role
	_ = json.Unmarshal(responseData, &roleReturned)
    fmt.Println(roleReturned)
	if uh.checkAdmin(roleReturned) {
		http.Redirect(w, r, "/admin/rooms", http.StatusSeeOther)
		return
	}
	http.Redirect(w, r, "/", http.StatusSeeOther)

}
func(uh *UserHandler) LoginGetHandler(w http.ResponseWriter, r *http.Request) {
	token, err := rtoken.CSRFToken(uh.csrfSignKey)
	if err != nil {
		http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
	}

	if r.Method==http.MethodGet {
		loginForm := struct {
			Values  url.Values
			VErrors form.ValidationErrors
			CSRF    string
		}{
			Values:  nil,
			VErrors: nil,
			CSRF:    token,
		}
		_ = uh.templ.ExecuteTemplate(w, "login.layout", loginForm)
	}else {
		uh.LoginPostHandler(w,r)
	}

}
func(uh *UserHandler) Logout(w http.ResponseWriter, r *http.Request) {
	userSess, _ := r.Context().Value(ctxUserSessionKey).(*entity.Session)
	session.Remove(userSess.UUID, w)
	cliet:=&http.Client{}
	url:=fmt.Sprintf("http://localhost:9090/user/session/delete/%d",uh.loggedInUser.Id)
	req, _ := http.NewRequest(http.MethodDelete,url, nil)
	_,err:=cliet.Do(req)
	if err!=nil{
		http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
		return
	}
	http.Redirect(w, r, "/login", http.StatusSeeOther)
}
func(uh *UserHandler) SignupGetHandler(w http.ResponseWriter, r *http.Request) {
	token, err := rtoken.CSRFToken(uh.csrfSignKey)
	if err != nil {
		http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
	}
	if r.Method==http.MethodGet{
		signUpForm := struct {
			Values  url.Values
			VErrors form.ValidationErrors
			CSRF    string
		}{
			Values:  nil,
			VErrors: nil,
			CSRF:    token,
		}
		uh.templ.ExecuteTemplate(w, "signup.layout", signUpForm)
		return
	}else {
		uh.SignupPostHandler(w,r)
	}
}
func(uh *UserHandler) SignupPostHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Post Login page")
	token, err := rtoken.CSRFToken(uh.csrfSignKey)
	if err != nil {
		http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
		return
	}
	// Parse the form data
	err = r.ParseForm()
	if err != nil {
		http.Error(w, http.StatusText(http.StatusBadRequest), http.StatusBadRequest)
		return
	}
	//
	singnUpForm := form.Input{Values: r.PostForm, VErrors: form.ValidationErrors{}}
	singnUpForm.Required("full_name", "email", "password", "password_confirmed")
	singnUpForm.MatchesPattern("email", form.EmailRX)
	singnUpForm.MatchesPattern("phone", form.PhoneRX)
	singnUpForm.MinLength("password", 8)
	singnUpForm.PasswordMatches("password", "password_confirmed")
	singnUpForm.CSRF = token

	// If there are any errors, redisplay the signup form.
	if !singnUpForm.Valid() {
		uh.templ.ExecuteTemplate(w, "signup.layout", singnUpForm)
		return
	}

	client:=&http.Client{}
	urlEmail :=fmt.Sprintf("http://localhost:9090/user/email/%s",r.FormValue("email"))
	req, err := http.NewRequest(http.MethodGet,urlEmail,nil)
	response,err:= client.Do(req)
	if (err!=nil)||(response.StatusCode==http.StatusNotFound){
		fmt.Println("Signup2111")
		http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
		return
	}
	responseData, err := ioutil.ReadAll(response.Body)
	if err != nil {
		fmt.Println("Signup217")
		fmt.Println(" Marshal Error")
		http.Redirect(w, r, "/signup", 302)
		return
	}
	var ok bool
	_ = json.Unmarshal(responseData, &ok)
	if ok {
		fmt.Println("Signup225")
		singnUpForm.VErrors.Add("email", "Email Already Exists")
		uh.templ.ExecuteTemplate(w, "signup.layout", singnUpForm)
		return
	}
	///
	urlPhone :=fmt.Sprintf("http://localhost:9090/user/phone/%s",r.FormValue("phone"))
	req, err = http.NewRequest(http.MethodGet,urlPhone,nil)
	response,err= client.Do(req)
	if (err!=nil)||(response.StatusCode==http.StatusNotFound){
		fmt.Println("Signup235")
		http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
		return
	}
	responseData, err = ioutil.ReadAll(response.Body)
	if err != nil {
		fmt.Println("Signup241")
		fmt.Println(" Marshal Error")
		http.Redirect(w, r, "/signup", 302)
		return
	}
	var ok2 bool
	_ = json.Unmarshal(responseData, &ok2)
	if ok2 {
		fmt.Println("Signup249")
		singnUpForm.VErrors.Add("phone", "Phone Already Exists")
		uh.templ.ExecuteTemplate(w, "signup.layout", singnUpForm)
		return
	}
	///
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(r.FormValue("password")), 12)
	if err != nil {
		fmt.Println("Signup257")
		singnUpForm.VErrors.Add("password", "Password Could not be stored")
		uh.templ.ExecuteTemplate(w, "signup.layout", singnUpForm)
		return
	}
	req, err = http.NewRequest(http.MethodGet,"http://localhost:9090/roles/USER",nil)
	response,err=client.Do(req)
	if err !=nil{
		fmt.Println("Signup265")
		http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
		return
	}
	role := entity.Role{}
	responseData, err = ioutil.ReadAll(response.Body)
	if err != nil {
		fmt.Println("Signup272")
		http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
		return
	}

	_ = json.Unmarshal(responseData, &role)
	roleId :=int(role.ID)
	var user = &entity.User{
		FullName: r.FormValue("full_name"),
		Email:    r.FormValue("email"),
		Phone:    r.FormValue("phone"),
		Password: string(hashedPassword),
		RoleID:   roleId,
	}
	output, err := json.MarshalIndent(user, "", "\t\t")
	if err != nil {
		http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
		return
	}
	response, err = http.Post("http://localhost:9090/user/new","application/json", bytes.NewBuffer(output))
	if (response==nil)||(err!=nil) {
		fmt.Println("Signup293")
		http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
		return
	}
	http.Redirect(w, r, "/login", 302)
	defer response.Body.Close()

}
func(uh *UserHandler) loggedIn(r *http.Request) bool {
	if uh.userSess == nil {
		return false
	}
	userSess := uh.userSess
	//retrieve the cookie from the request:
	c, err := r.Cookie(userSess.UUID)
	if err != nil {
		fmt.Println("390")
		return false
	}
	ok, err := session.Valid(c.Value, userSess.SigningKey)
	if !ok || (err != nil) {
		fmt.Println("395")
		return false
	}

	return true
}
func(uh *UserHandler) checkAdmin(rs []entity.Role) bool {
	for _, r := range rs {
		if strings.ToUpper(r.Name) == strings.ToUpper("Admin") {
			return true
		}
	}
	return false
}
// AdminUsers handles Get /admin/users request
func (uh *UserHandler) AdminUsers(w http.ResponseWriter, r *http.Request) {
	token, err := rtoken.CSRFToken(uh.csrfSignKey)
	if err != nil {
		http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
	}
	var users []entity.User

	client:=&http.Client{}
	req, err := http.NewRequest(http.MethodGet,"http://localhost:9090/user/users",nil)
	response,err:= client.Do(req)
	if (err!=nil)||(response.StatusCode==http.StatusNotFound){
		http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
		return
	}
	responseData, err := ioutil.ReadAll(response.Body)
	if err != nil {
		http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
		return
	}

	err = json.Unmarshal(responseData, &users)
	if err!=nil {
		http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
		return
	}
	fmt.Println(users)
	tmplData := struct {
		Values  url.Values
		VErrors form.ValidationErrors
		Users   []entity.User
		CSRF    string
	}{
		Values:  nil,
		VErrors: nil,
		Users:   users,
		CSRF:    token,
	}
	uh.templ.ExecuteTemplate(w, "admin.users.layout", tmplData)
}
// AdminUsersNew handles GET/POST /admin/users/new request
func (uh *UserHandler) AdminUsersNew(w http.ResponseWriter, r *http.Request) {
	token, err := rtoken.CSRFToken(uh.csrfSignKey)
	if err != nil {
		http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
		return
	}
	if r.Method == http.MethodGet {
		var roles []entity.Role
		client:=&http.Client{}
		req, err := http.NewRequest(http.MethodGet,"http://localhost:9090/roles",nil)
		response,err:= client.Do(req)
		if (err!=nil)||(response.StatusCode==http.StatusNotFound){
			http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
			return
		}
		responseData, err := ioutil.ReadAll(response.Body)
		if err != nil {
			http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
			return
		}

		err = json.Unmarshal(responseData, &roles)
		if err!=nil {
			http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
			return
		}
		formInput:=form.Input{
			Values:  nil,
			VErrors: nil,
			CSRF:    token,
		}
		accountForm := struct {

			Roles   []entity.Role
			FormInput form.Input
		}{

			Roles:   roles,
			FormInput:formInput,
		}
		uh.templ.ExecuteTemplate(w, "admin.user.new.layout", accountForm)
		return
	}

	if r.Method == http.MethodPost {
		// Parse the form data
		err := r.ParseForm()
		if err != nil {
			w.WriteHeader(http.StatusBadRequest)
			return
		}
		vl:=url.Values{}
		vl.Add("fullname",r.FormValue("fullname"))
		vl.Add("email",r.FormValue("email"))
		vl.Add("password",r.FormValue("password"))
		vl.Add("confirmpassword",r.FormValue("confirmpassword"))
		vl.Add("phone",r.FormValue("phone"))
		vl.Add("role",r.FormValue("role"))
        fmt.Println(r.FormValue("role"))
		// Validate the form contents
		roles:=uh.GetRoles()
		formInput := form.Input{Values: vl, VErrors: form.ValidationErrors{}}
		formInput.Required("fullname", "email", "password", "confirmpassword")
		formInput.MatchesPattern("email", form.EmailRX)
		formInput.MatchesPattern("phone", form.PhoneRX)
		formInput.MinLength("password", 8)
		formInput.PasswordMatches("password", "confirmpassword")
		formInput.CSRF = token
		accountForms := struct {
			Roles   []entity.Role
			FormInput    form.Input
		}{
			Roles:     roles,
			FormInput: formInput,
		}
		fmt.Println("not passed")
		// If there are any errors, redisplay the signup form.
		if !formInput.Valid() {
			uh.templ.ExecuteTemplate(w, "admin.user.new.layout", accountForms)
			return
		}
		fmt.Println("passed")

		client:=&http.Client{}
		urlEmail :=fmt.Sprintf("http://localhost:9090/user/email/%s",r.FormValue("email"))
		req, err := http.NewRequest(http.MethodGet,urlEmail,nil)
		response,err:= client.Do(req)
		if (err!=nil)||(response.StatusCode==http.StatusNotFound){
			fmt.Println("466")
			http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
			return
		}
		responseData, err := ioutil.ReadAll(response.Body)
		if err != nil {

			fmt.Println("473")
			http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
			return
		}

		var ok bool
		_ = json.Unmarshal(responseData, &ok)
		if ok {
			fmt.Println("Signup481")
			formInput.VErrors.Add("email", "Email Already Exists")
			accountForms = struct {
				Roles   []entity.Role
				FormInput    form.Input
			}{
				Roles:     roles,
				FormInput: formInput,
			}
			uh.templ.ExecuteTemplate(w, "admin.user.new.layout", accountForms)
			return
		}
		///
		urlPhone :=fmt.Sprintf("http://localhost:9090/user/phone/%s",r.FormValue("phone"))
		req, err = http.NewRequest(http.MethodGet,urlPhone,nil)
		response,err= client.Do(req)
		if (err!=nil)||(response.StatusCode==http.StatusNotFound){
			fmt.Println("Signup498")
			http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
			return
		}
		responseData, err = ioutil.ReadAll(response.Body)
		if err != nil {
			fmt.Println("Signup504")
			http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
			return

		}

		var ok2 bool
		_ = json.Unmarshal(responseData, &ok2)
		if ok2 {
			fmt.Println("Signup513")
			formInput.VErrors.Add("phone", "Phone Already Exists")
			accountForms = struct {
				Roles   []entity.Role
				FormInput    form.Input
			}{
				Roles:     roles,
				FormInput: formInput,
			}
			uh.templ.ExecuteTemplate(w, "admin.user.new.layout", accountForms)
			return
		}

		hashedPassword, err := bcrypt.GenerateFromPassword([]byte(r.FormValue("password")), 12)
		if err != nil {
			formInput.VErrors.Add("password", "Password Could not be stored")
			accountForms = struct {
				Roles   []entity.Role
				FormInput    form.Input
			}{
				Roles:     roles,
				FormInput: formInput,
			}
			uh.templ.ExecuteTemplate(w, "admin.user.new.layout", accountForms)
			return
		}

		roleID, err := strconv.Atoi(r.FormValue("role"))
		if err != nil {
			formInput.VErrors.Add("role", "could not retrieve role id")
			accountForms = struct {
				Roles   []entity.Role
				FormInput    form.Input
			}{
				Roles:     roles,
				FormInput: formInput,
			}
			uh.templ.ExecuteTemplate(w, "admin.user.new.layout", accountForms)
			return
		}
		user := &entity.User{
			FullName: r.FormValue("fullname"),
			Email:    r.FormValue("email"),
			Phone:    r.FormValue("phone"),
			Password: string(hashedPassword),
			RoleID:   roleID,
		}
		output, err := json.MarshalIndent(user, "", "\t\t")
		if err != nil {
			http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
			return
		}
		response, err = http.Post("http://localhost:9090/user/new","application/json", bytes.NewBuffer(output))
		if (response==nil)||(err!=nil) {
			fmt.Println("Signup567")
			http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
			return
		}
		http.Redirect(w, r, "/admin/users", http.StatusSeeOther)
		defer response.Body.Close()

	}

}

// AdminUsersUpdate handles GET/POST /admin/users/update?id={id} request
func (uh *UserHandler) AdminUsersUpdate(w http.ResponseWriter, r *http.Request) {
	token, err := rtoken.CSRFToken(uh.csrfSignKey)
	if err != nil {
		http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
	}
	if r.Method == http.MethodGet {
		idRaw := r.URL.Query().Get("id")
		id, err := strconv.Atoi(idRaw)
		if err != nil {
			http.Error(w, http.StatusText(http.StatusBadRequest), http.StatusBadRequest)
			return
		}
		var user entity.User
		client:=&http.Client{}
		urlRequested :=fmt.Sprintf("http://localhost:9090/user/user/%d",id)
		req, err := http.NewRequest(http.MethodGet, urlRequested,nil)
		response,err:= client.Do(req)
		if (err!=nil)||(response.StatusCode==http.StatusNotFound){
			http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
			return
		}
		responseData, err := ioutil.ReadAll(response.Body)
		if err != nil {
			http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
			return
		}

		err = json.Unmarshal(responseData, &user)
		if err!=nil {
			http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
			return
		}
		var roles []entity.Role
		req, err = http.NewRequest(http.MethodGet,"http://localhost:9090/roles",nil)
		response,err= client.Do(req)
		if (err!=nil)||(response.StatusCode==http.StatusNotFound){
			http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
			return
		}
		responseData, err = ioutil.ReadAll(response.Body)
		if err != nil {
			http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
			return
		}

		err = json.Unmarshal(responseData, &roles)
		if err!=nil {
			http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
			return
		}
		output, err := json.MarshalIndent(user, "", "\t\t")
		if err != nil {
			http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
			return
		}
		var role entity.Role

		urlRequested =fmt.Sprintf("http://localhost:9090/role/%d",user.RoleID)
		req, err = http.NewRequest(http.MethodGet, urlRequested,bytes.NewBuffer(output))
		response,err= client.Do(req)
		if (err!=nil)||(response.StatusCode==http.StatusNotFound){
			http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
			return
		}

		responseData, err = ioutil.ReadAll(response.Body)
		if err != nil {
			http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
			return
		}

		err = json.Unmarshal(responseData, &role)
		if err!=nil {
			http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
			return
		}
		values :=url.Values{}
		values.Add("userid", idRaw)
		values.Add("fullname", user.FullName)
		values.Add("email", user.Email)
		values.Add("role", string(user.RoleID))
		values.Add("phone", user.Phone)
		values.Add("rolename", role.Name)
		formInputs:=form.Input{
			Values:  values,
			VErrors: form.ValidationErrors{},
			CSRF:    token,
		}
		upAccForm := struct {

			Roles   []entity.Role
			User    *entity.User
			FormInputs form.Input
		}{

			Roles:   roles,
			User:    &user,
			FormInputs:formInputs,
		}
		uh.templ.ExecuteTemplate(w, "admin.user.update.layout", upAccForm)
		return
	}

	if r.Method == http.MethodPost {
		// Parse the form data

		v:=url.Values{}
		v.Add("fullname",r.FormValue("fullname"))
		v.Add("email",r.FormValue("email"))
		v.Add("phone",r.FormValue("phone"))
		v.Add("userid",r.FormValue("userid"))
		v.Add("role",r.FormValue("role"))
		// Validate the form contents
		upAccForm := form.Input{Values: v, VErrors: form.ValidationErrors{}}
		upAccForm.Required("fullname", "email", "phone")
		upAccForm.MatchesPattern("email", form.EmailRX)
		upAccForm.MatchesPattern("phone", form.PhoneRX)
		upAccForm.CSRF = token
		roles:=uh.GetRoles()
		uid,_:=strconv.Atoi(r.FormValue("userid"))
		rid,_:=strconv.Atoi(r.FormValue("role"))
		user:=entity.User{
			Id:       uint32(uid),
			FullName: r.FormValue("fullname"),
			Email:    r.FormValue("email"),
			Phone:    r.FormValue("phone"),
			RoleID:   rid,
		}
		upAccForm2 := struct {

			Roles   []entity.Role
			User    *entity.User
			FormInputs form.Input
		}{

			Roles:   roles,
			User:    &user,
			FormInputs:upAccForm,
		}
		if !upAccForm.Valid() {
			uh.templ.ExecuteTemplate(w, "admin.user.update.layout", upAccForm2)
			return
		}
		client:=&http.Client{}
		urlRequested :=fmt.Sprintf("http://localhost:9090/user/user/%d",uid)
		req, err := http.NewRequest(http.MethodGet, urlRequested,nil)
		response,err:= client.Do(req)
		if (err!=nil)||(response.StatusCode==http.StatusNotFound){
			http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
			return
		}
		responseData, err2 := ioutil.ReadAll(response.Body)
		if err2 != nil {
			http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
			return
		}
		err = json.Unmarshal(responseData, &user)
		if err!=nil {
			http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
			return
		}

		urlEmail :=fmt.Sprintf("http://localhost:9090/user/email/%s",r.FormValue("email"))
		req, err = http.NewRequest(http.MethodGet,urlEmail,nil)
		response,err= client.Do(req)
		if (err!=nil)||(response.StatusCode==http.StatusNotFound){
			http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
			return
		}
		responseData, err = ioutil.ReadAll(response.Body)
		if err != nil {
			http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
			return
		}

		var ok bool
		_ = json.Unmarshal(responseData, &ok)
		if ok {
			upAccForm.VErrors.Add("email", "Email Already Exists")
			upAccForm2 := struct {

				Roles   []entity.Role
				User    *entity.User
				FormInputs form.Input
			}{

				Roles:   roles,
				User:    &user,
				FormInputs:upAccForm,
			}
			uh.templ.ExecuteTemplate(w, "admin.user.update.layout", upAccForm2)
			return
		}

		urlPhone :=fmt.Sprintf("http://localhost:9090/user/phone/%s",r.FormValue("phone"))
		req, err = http.NewRequest(http.MethodGet,urlPhone,nil)
		response,err= client.Do(req)
		if (err!=nil)||(response.StatusCode==http.StatusNotFound){
			http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
			return
		}
		responseData, err = ioutil.ReadAll(response.Body)
		if err != nil {
			http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
		}

		var ok2 bool
		_ = json.Unmarshal(responseData, &ok2)
		if ok2 {
			upAccForm.VErrors.Add("phone", "Phone Already Exists")
			upAccForm2 := struct {

				Roles   []entity.Role
				User    *entity.User
				FormInputs form.Input
			}{

				Roles:   roles,
				User:    &user,
				FormInputs:upAccForm,
			}
			uh.templ.ExecuteTemplate(w, "admin.user.update.layout", upAccForm2)
			return
		}

		roleID, err := strconv.Atoi(r.FormValue("role"))
		if err != nil {
			upAccForm.VErrors.Add("role", "could not retrieve role id")
			upAccForm2 := struct {

				Roles   []entity.Role
				User    *entity.User
				FormInputs form.Input
			}{

				Roles:   roles,
				User:    &user,
				FormInputs:upAccForm,
			}
			uh.templ.ExecuteTemplate(w, "admin.user.update.layout", upAccForm2)
			return
		}

		usr := &entity.User{
			Id:       user.Id,
			FullName: r.FormValue("fullname"),
			Email:    r.FormValue("email"),
			Phone:    r.FormValue("phone"),
			Password: user.Password,
			RoleID:   roleID,
		}
		output, err := json.MarshalIndent(usr, "", "\t\t")
		if err != nil {
			fmt.Println("902")
			http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
			return
		}
		urlUpdate:=fmt.Sprintf("http://localhost:9090/user/update/%d",usr.Id)
		req, err = http.NewRequest(http.MethodPut, urlUpdate,bytes.NewBuffer(output))
		response,err= client.Do(req)
		if (err!=nil)||(response.StatusCode==http.StatusNotFound){
			fmt.Println("918")
			http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
			return
		}
		responseData, err = ioutil.ReadAll(response.Body)
		if err != nil {
			fmt.Println("924")
			http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
			return
		}

		err = json.Unmarshal(responseData, &usr)
		if err!=nil {
			fmt.Println("931")
			http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
			return
		}

		http.Redirect(w, r, "/admin/users", http.StatusSeeOther)
	}
}

// AdminUsersDelete handles Delete /admin/users/delete?id={id} request
func (uh *UserHandler) AdminUsersDelete(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodGet {
		idRaw := r.URL.Query().Get("id")
		id, err := strconv.Atoi(idRaw)
		if err != nil {
			http.Error(w, http.StatusText(http.StatusBadRequest), http.StatusBadRequest)
			return
		}
		cliet:=&http.Client{}
		url:=fmt.Sprintf("http://localhost:9090/user/delete/%d",id)
		req, err := http.NewRequest(http.MethodDelete,url, nil)
		response,err:=cliet.Do(req)
		if http.StatusNotFound == response.StatusCode {
			http.Error(w, http.StatusText(http.StatusBadRequest), http.StatusBadRequest)
			return
		}
		if http.StatusUnprocessableEntity == response.StatusCode{
			http.Error(w, http.StatusText(http.StatusBadRequest), http.StatusBadRequest)
			return
		}
		if err!=nil{
			http.Error(w, http.StatusText(http.StatusBadRequest), http.StatusBadRequest)
			return
		}
		http.Redirect(w, r, "/admin/users", http.StatusSeeOther)

	}
	http.Redirect(w, r, "/admin/users", http.StatusSeeOther)
}
func (uh *UserHandler) GetRoles()[]entity.Role {
	var roles []entity.Role
	client:=&http.Client{}
	req, err := http.NewRequest(http.MethodGet,"http://localhost:9090/roles",nil)
	response,err:= client.Do(req)
	if (err!=nil)||(response.StatusCode==http.StatusNotFound){
		return nil
	}
	responseData, err := ioutil.ReadAll(response.Body)
	if err != nil {
		return nil
	}

	err = json.Unmarshal(responseData, &roles)
	if err!=nil {
		return nil
	}
	return roles
}
