package example

import "fmt"
import "net/http"
import "strconv"
import "encoding/json"

func (model *ApiError) Check(r *http.Request) error {
	return nil
}

func (model *MyApi) Check(r *http.Request) error {
	return nil
}

func (model *ProfileParams) Check(r *http.Request) error {
	model.Login = r.FormValue("login")
	if model.Login == "" {
		return fmt.Errorf("login must me not empty")
	}
	return nil
}

func (model *CreateParams) Check(r *http.Request) error {
	model.Login = r.FormValue("login")
	if model.Login == "" {
		return fmt.Errorf("login must me not empty")
	}	
	if len(model.Login) < 10 {
		return fmt.Errorf("loginlen must be >= 10")
	}

	model.Name = r.FormValue("full_name")
	model.Status = r.FormValue("status")
	if model.Status == "" {
		model.Status = "user"
	}
	switch model.Status {
	case "user":
	case "moderator":
	case "admin":
	case "":
	default:
		return fmt.Errorf("class must be one of [user moderator admin]")
	}

	val, err := strconv.Atoi(r.FormValue("age"))
	if err != nil {
		return fmt.Errorf("age must be int")
	}
	model.Age = val	
	if model.Age < 0 {
		return fmt.Errorf("age must be >= 0")
	}
	if model.Age > 128 {
		return fmt.Errorf("age must be <= 128")
}

	return nil
}

func (model *User) Check(r *http.Request) error {
	return nil
}

func (model *NewUser) Check(r *http.Request) error {
	return nil
}

func (srv *MyApi) ProfileHandler(w http.ResponseWriter, r *http.Request) { 
	params := ProfileParams{}
	err := params.Check(r)
	w.Header()["Content-Type"] = []string{"application/json"}
	if err != nil {
		w.WriteHeader(400)
		body := Body{
			Error: err.Error(),
		}
		bodyBytes, _ := json.Marshal(body)
		w.Write(bodyBytes)
		return
	}
	result, err := srv.Profile(r.Context(), params)
	if err != nil {
		if apierror, ok := err.(ApiError); ok {
			w.WriteHeader(apierror.HTTPStatus)
		} else {
			w.WriteHeader(500)
		}
		body := Body{
			Error: err.Error(),
		}
		bodyBytes, _ := json.Marshal(body)
		w.Write(bodyBytes)
		return
	}
	body := Body{
		Response: result,
	}
	bodyBytes, _ := json.Marshal(body)
	w.Write(bodyBytes)
}

func (srv *MyApi) CreateHandler(w http.ResponseWriter, r *http.Request) { 
		if r.Header.Get("X-Auth") != "100500" {
		w.WriteHeader(403)
		body := Body{
			Error: "unauthorized",
		}
		bodyBytes, _ := json.Marshal(body)
		w.Write(bodyBytes)
		return
	}

	if r.Method != "POST" {
		w.WriteHeader(406)
		body := Body{
			Error: "bad method",
		}
		bodyBytes, _ := json.Marshal(body)
		w.Write(bodyBytes)
		return
	}

	params := CreateParams{}
	err := params.Check(r)
	w.Header()["Content-Type"] = []string{"application/json"}
	if err != nil {
		w.WriteHeader(400)
		body := Body{
			Error: err.Error(),
		}
		bodyBytes, _ := json.Marshal(body)
		w.Write(bodyBytes)
		return
	}
	result, err := srv.Create(r.Context(), params)
	if err != nil {
		if apierror, ok := err.(ApiError); ok {
			w.WriteHeader(apierror.HTTPStatus)
		} else {
			w.WriteHeader(500)
		}
		body := Body{
			Error: err.Error(),
		}
		bodyBytes, _ := json.Marshal(body)
		w.Write(bodyBytes)
		return
	}
	body := Body{
		Response: result,
	}
	bodyBytes, _ := json.Marshal(body)
	w.Write(bodyBytes)
}

func (model *OtherApi) Check(r *http.Request) error {
	return nil
}

func (model *OtherCreateParams) Check(r *http.Request) error {
	model.Username = r.FormValue("username")
	if model.Username == "" {
		return fmt.Errorf("username must me not empty")
	}	
	if len(model.Username) < 3 {
		return fmt.Errorf("usernamelen must be >= 3")
	}

	model.Name = r.FormValue("account_name")
	model.Class = r.FormValue("class")
	if model.Class == "" {
		model.Class = "warrior"
	}
	switch model.Class {
	case "warrior":
	case "sorcerer":
	case "rouge":
	case "":
	default:
		return fmt.Errorf("class must be one of [warrior sorcerer rouge]")
	}

	val, err := strconv.Atoi(r.FormValue("level"))
	if err != nil {
		return fmt.Errorf("level must be int")
	}
	model.Level = val	
	if model.Level < 1 {
		return fmt.Errorf("level must be >= 1")
	}
	if model.Level > 50 {
		return fmt.Errorf("level must be <= 50")
}

	return nil
}

func (model *OtherUser) Check(r *http.Request) error {
	return nil
}

func (srv *OtherApi) CreateHandler(w http.ResponseWriter, r *http.Request) { 
		if r.Header.Get("X-Auth") != "100500" {
		w.WriteHeader(403)
		body := Body{
			Error: "unauthorized",
		}
		bodyBytes, _ := json.Marshal(body)
		w.Write(bodyBytes)
		return
	}

	if r.Method != "POST" {
		w.WriteHeader(406)
		body := Body{
			Error: "bad method",
		}
		bodyBytes, _ := json.Marshal(body)
		w.Write(bodyBytes)
		return
	}

	params := OtherCreateParams{}
	err := params.Check(r)
	w.Header()["Content-Type"] = []string{"application/json"}
	if err != nil {
		w.WriteHeader(400)
		body := Body{
			Error: err.Error(),
		}
		bodyBytes, _ := json.Marshal(body)
		w.Write(bodyBytes)
		return
	}
	result, err := srv.Create(r.Context(), params)
	if err != nil {
		if apierror, ok := err.(ApiError); ok {
			w.WriteHeader(apierror.HTTPStatus)
		} else {
			w.WriteHeader(500)
		}
		body := Body{
			Error: err.Error(),
		}
		bodyBytes, _ := json.Marshal(body)
		w.Write(bodyBytes)
		return
	}
	body := Body{
		Response: result,
	}
	bodyBytes, _ := json.Marshal(body)
	w.Write(bodyBytes)
}


func (h *MyApi) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	switch r.URL.Path {
	case "/user/profile":
		h.ProfileHandler(w, r)
	case "/user/create":
		h.CreateHandler(w, r)
	default:
		w.WriteHeader(404)
		body := Body{
			Error: "unknow method",
		}
		bodyBytes, _ := json.Marshal(body)
		w.Write(bodyBytes)
		return
	}
}

func (h *OtherApi) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	switch r.URL.Path {
	case "/user/create":
		h.CreateHandler(w, r)
	default:
		w.WriteHeader(404)
		body := Body{
			Error: "unknow method",
		}
		bodyBytes, _ := json.Marshal(body)
		w.Write(bodyBytes)
		return
	}
}