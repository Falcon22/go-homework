// Сгенерированный код 
package example

import "net/http"
import "strconv"
import "fmt"
import "encoding/json"

    type Body struct {
      Error    string      "json:\"error\""
      Response interface{} "json:\"response,omitempty\""
    }
    
// Генерация функции Bind для структуры ApiError
func (model *ApiError) Bind(r *http.Request) error {
return nil
}
// Генерация функции Bind для структуры MyApi
func (model *MyApi) Bind(r *http.Request) error {
return nil
}
// Генерация функции Bind для структуры ProfileParams
func (model *ProfileParams) Bind(r *http.Request) error {
model.Login = r.FormValue("login") 

      if model.Login == "" {
        return fmt.Errorf("login must me not empty")
      }
      return nil
}
// Генерация функции Bind для структуры CreateParams
func (model *CreateParams) Bind(r *http.Request) error {
model.Login = r.FormValue("login") 

      if model.Login == "" {
        return fmt.Errorf("login must me not empty")
      }
      
      if len(model.Login) < 10 {
        return fmt.Errorf("login len must be >= 10")
      }
      model.Name = r.FormValue("full_name") 
model.Status = r.FormValue("status") 
switch model.Status {
case "user":
case "moderator":
case "admin":
case "":
default:
        return fmt.Errorf("status must be one of [user, moderator, admin]")
      }
      if model.Status == "" {
                model.Status = "user"
              }
              
            val, err := strconv.Atoi(r.FormValue("age"))
            if err != nil {
              return fmt.Errorf("age must be int")
            }
            model.Age = val
            
      if model.Age > 128 {
        return fmt.Errorf("age must be <= 128")
      }
      
      if model.Age < 0 {
        return fmt.Errorf("age must be >= 0")
      }
      return nil
}
// Генерация функции Bind для структуры User
func (model *User) Bind(r *http.Request) error {
return nil
}
// Генерация функции Bind для структуры NewUser
func (model *NewUser) Bind(r *http.Request) error {
return nil
}
//Сгенерированный хэндлер для функции Profile

          func (srv *MyApi) ProfileHandler(w http.ResponseWriter, r *http.Request) {
            
          params := ProfileParams{}
          err := params.Bind(r)
          w.Header()["Content-Type"] = []string{"application/json"}
          if err != nil {
            w.WriteHeader(400)
            body := Body {
            Error: err.Error(),
          }
          bodyBytes, _:= json.Marshal(body)
          w.Write(bodyBytes)
          return
        }
        
        result, err := srv.Profile(r.Context(), params)
        if err!=nil {
          if apierror, ok := err.(ApiError); ok {
            w.WriteHeader(apierror.HTTPStatus)
            } else {
              w.WriteHeader(500)
            }
            body := Body{
              Error: err.Error(),
            }
            bodyBytes, _:= json.Marshal(body)
            w.Write(bodyBytes)
            return
          }
          body := Body{
            Response: result,
          }
          bodyBytes, _ := json.Marshal(body)
          w.Write(bodyBytes)
          }
//Сгенерированный хэндлер для функции Create

          func (srv *MyApi) CreateHandler(w http.ResponseWriter, r *http.Request) {
            
              if r.Header.Get("X-Auth") != "100500" {
                w.WriteHeader(403)
                body := Body{
                  Error: "unauthorized",
                }
                bodyBytes, _:= json.Marshal(body)
                w.Write(bodyBytes)
                return
              }
              
              if r.Method != "POST" {
                w.WriteHeader(406)
                body := Body {
                Error: "bad method",
              }
              bodyBytes, _:= json.Marshal(body)
              w.Write(bodyBytes)
              return
            }
            
          params := CreateParams{}
          err := params.Bind(r)
          w.Header()["Content-Type"] = []string{"application/json"}
          if err != nil {
            w.WriteHeader(400)
            body := Body {
            Error: err.Error(),
          }
          bodyBytes, _:= json.Marshal(body)
          w.Write(bodyBytes)
          return
        }
        
        result, err := srv.Create(r.Context(), params)
        if err!=nil {
          if apierror, ok := err.(ApiError); ok {
            w.WriteHeader(apierror.HTTPStatus)
            } else {
              w.WriteHeader(500)
            }
            body := Body{
              Error: err.Error(),
            }
            bodyBytes, _:= json.Marshal(body)
            w.Write(bodyBytes)
            return
          }
          body := Body{
            Response: result,
          }
          bodyBytes, _ := json.Marshal(body)
          w.Write(bodyBytes)
          }
// Генерация функции Bind для структуры OtherApi
func (model *OtherApi) Bind(r *http.Request) error {
return nil
}
// Генерация функции Bind для структуры OtherCreateParams
func (model *OtherCreateParams) Bind(r *http.Request) error {
model.Username = r.FormValue("username") 

      if model.Username == "" {
        return fmt.Errorf("username must me not empty")
      }
      
      if len(model.Username) < 3 {
        return fmt.Errorf("username len must be >= 3")
      }
      model.Name = r.FormValue("account_name") 
model.Class = r.FormValue("class") 
switch model.Class {
case "warrior":
case "sorcerer":
case "rouge":
case "":
default:
        return fmt.Errorf("class must be one of [warrior, sorcerer, rouge]")
      }
      if model.Class == "" {
                model.Class = "warrior"
              }
              
            val, err := strconv.Atoi(r.FormValue("level"))
            if err != nil {
              return fmt.Errorf("level must be int")
            }
            model.Level = val
            
      if model.Level > 50 {
        return fmt.Errorf("level must be <= 50")
      }
      
      if model.Level < 1 {
        return fmt.Errorf("level must be >= 1")
      }
      return nil
}
// Генерация функции Bind для структуры OtherUser
func (model *OtherUser) Bind(r *http.Request) error {
return nil
}
//Сгенерированный хэндлер для функции Create

          func (srv *OtherApi) CreateHandler(w http.ResponseWriter, r *http.Request) {
            
              if r.Header.Get("X-Auth") != "100500" {
                w.WriteHeader(403)
                body := Body{
                  Error: "unauthorized",
                }
                bodyBytes, _:= json.Marshal(body)
                w.Write(bodyBytes)
                return
              }
              
              if r.Method != "POST" {
                w.WriteHeader(406)
                body := Body {
                Error: "bad method",
              }
              bodyBytes, _:= json.Marshal(body)
              w.Write(bodyBytes)
              return
            }
            
          params := OtherCreateParams{}
          err := params.Bind(r)
          w.Header()["Content-Type"] = []string{"application/json"}
          if err != nil {
            w.WriteHeader(400)
            body := Body {
            Error: err.Error(),
          }
          bodyBytes, _:= json.Marshal(body)
          w.Write(bodyBytes)
          return
        }
        
        result, err := srv.Create(r.Context(), params)
        if err!=nil {
          if apierror, ok := err.(ApiError); ok {
            w.WriteHeader(apierror.HTTPStatus)
            } else {
              w.WriteHeader(500)
            }
            body := Body{
              Error: err.Error(),
            }
            bodyBytes, _:= json.Marshal(body)
            w.Write(bodyBytes)
            return
          }
          body := Body{
            Response: result,
          }
          bodyBytes, _ := json.Marshal(body)
          w.Write(bodyBytes)
          }

      func (h *MyApi ) ServeHTTP(w http.ResponseWriter, r *http.Request) {
        switch r.URL.Path {
          
          case "/user/profile":
            h.ProfileHandler(w, r)
            
          case "/user/create":
            h.CreateHandler(w, r)
            
        default:
          w.WriteHeader(404)
          body := Body{
            Error: "unknown method",
          }
          bodyBytes, _:= json.Marshal(body)
          w.Write(bodyBytes)
          return
        }
      }
      
      func (h *OtherApi ) ServeHTTP(w http.ResponseWriter, r *http.Request) {
        switch r.URL.Path {
          
          case "/user/create":
            h.CreateHandler(w, r)
            
        default:
          w.WriteHeader(404)
          body := Body{
            Error: "unknown method",
          }
          bodyBytes, _:= json.Marshal(body)
          w.Write(bodyBytes)
          return
        }
      }
      