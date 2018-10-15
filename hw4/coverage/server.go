package main

import (
	"net/http"
	"os"
	"io/ioutil"
	"encoding/xml"
	"strings"
	"sort"
	"io"
	"strconv"
	"encoding/json"
	"time"
)

var filePath string

type UserXML struct {
	Id            int `xml:"id"`
	Guid          string `xml:"guid"`
	IsActive      bool `xml:"isActive"`
	Balance       string `xml:"balance"`
	Picture       string `xml:"picture"`
	Age           int `xml:"age"`
	EyeColor      string `xml:"eyeColor"`
	FirstName     string `xml:"first_name"`
	LastName      string `xml:"last_name"`
	Gender        string `xml:"gender"`
	Company       string `xml:"compane"`
	Email         string `xml:"email"`
	Phone         string `xml:"phone"`
	Address       string `xml:"address"`
	About         string `xml:"about"`
	Registered    string `xml:"registered"`
	FavoriteFruit string `xml:"favoriteFruit"`
}

func Min(a, b int)(c int) {
	if a < b {
		return a
	} else {
		return b
	}
}

func (this *Users) ConvertToUser() ([]User) {
	res := make([]User, len(this.List))
	for i := range this.List {
		res[i].Id = this.List[i].Id
		res[i].Name = this.List[i].FirstName + " " + this.List[i].LastName
		res[i].Age = this.List[i].Age
		res[i].Gender = this.List[i].Gender
		res[i].About = this.List[i].About
	}
	return res
}

type Users struct {
	Version string    `xml:"version,attr"`
	List    []UserXML `xml:"row"`
}

type UsersId []User

func (this UsersId) Len() int {
	return len(this)
}

func (this UsersId) Less(i, j int) bool {
	return this[i].Id < this[j].Id
}

func (this UsersId) Swap(i, j int) {
	this[i], this[j] = this[j], this[i]
}

type UsersName []User

func (this UsersName) Len() int {
	return len(this)
}

func (this UsersName) Less(i, j int) bool {
	return this[i].Name < this[j].Name
}

func (this UsersName) Swap(i, j int) {
	this[i], this[j] = this[j], this[i]
}

type UsersAge []User

func (this UsersAge) Len() int {
	return len(this)
}

func (this UsersAge) Less(i, j int) bool {
	return this[i].Age < this[j].Age
}

func (this UsersAge) Swap(i, j int) {
	this[i], this[j] = this[j], this[i]
}

func BadServer(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusInternalServerError)
}

func BadServerJson(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusBadRequest)
	io.WriteString(w, "!!!")
}

func BadJsonServer(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
	io.WriteString(w, "!!!")
}

func BadServerUnknow(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusBadRequest)
	io.WriteString(w, `{"status": 500, "error": "???"}`)
}

func BadServerSleep(w http.Response, r *http.Request) {
	time.Sleep(20 * time.Second);
}

func SearchServer(w http.ResponseWriter, r *http.Request) {
	filePath := r.Header.Get("AccessToken")

	if !strings.Contains(filePath, ".xml") {
		w.WriteHeader(http.StatusUnauthorized)
		return
	}

	limit, err := strconv.Atoi(r.FormValue("limit"))
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	offset, err := strconv.Atoi(r.FormValue("offset"))
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	query := strings.ToLower(r.FormValue("query"))
	queryField := strings.ToLower(r.FormValue("order_field"))

	orderBy, err := strconv.Atoi(r.FormValue("order_by"))
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	file, err := os.Open(filePath)
	defer file.Close()
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}


	fileData, err := ioutil.ReadAll(file)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	data := new(Users)
	err = xml.Unmarshal(fileData, &data)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	res := new(Users)
	if len(query) == 0 {
		res = data
	} else {
		for _, u := range data.List {
			if strings.EqualFold(query, strings.ToLower(u.FirstName)) ||
			   strings.EqualFold(query, strings.ToLower(u.LastName)) ||
			   strings.Contains(query, strings.ToLower(u.About)) {
			   	res.List = append(res.List, u)
			}
		}
	}

	out := res.ConvertToUser()

	if orderBy != 0 {
		switch queryField {
		case "id":
			if orderBy > 0 {
				sort.Sort(UsersId(out))
			} else {
				sort.Sort(sort.Reverse(UsersId(out)))
			}
		case "age":
			if orderBy > 0 {
				sort.Sort(UsersAge(out))
			} else {
				sort.Sort(sort.Reverse(UsersAge(out)))
			}
		case "name":
			if orderBy > 0 {
				sort.Sort(UsersName(out))
			} else {
				sort.Sort(sort.Reverse(UsersName(out)))
			}
		default:
			if len(queryField) == 0 {
				sort.Sort(UsersName(out))
			} else {
				w.WriteHeader(http.StatusBadRequest)
				io.WriteString(w, `{"status": 400, "error": "ErrorBadOrderField"}`)
				return
			}
		}
	}

	if offset >= len(out) {
		empt, _ := json.Marshal(out[:0])
		io.WriteString(w, string(empt))
		return
	}

	out = out[offset : Min(offset + limit, len(out))]
	dat, _ := json.Marshal(out)

	w.WriteHeader(http.StatusOK)
	io.WriteString(w, string(dat))
}

