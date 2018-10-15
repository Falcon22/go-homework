package main

import (
	"fmt"
	//	"go/ast"
	"go/parser"
	"go/token"
	"log"
	"os"
	//"reflect"
	//"strings"
	//"text/template"
	"text/template"
	"go/ast"
	"reflect"
	"strings"
	"strconv"
	"encoding/json"
)

type NameTpl struct {
	Name string
}

type Rout struct {
	HandlerName string
	UrlName     string
}


type ServParamTpl struct{
	Name string
	Urls []Rout
}

type ValidField struct {
	Name string
	ParamName string
	Required bool
	Default string
	DefaultNotEmpty bool
	Min int
	Max int
	Values []string
	ValuesWithСomma string
	ValuesNotEmpty bool
	IsString bool
}

type CheckStruct struct {
	NameStruct string
	Fields []ValidField
}

type HandlerStruct struct {
	NameStruct string
	NameFunc string
	Auth bool
	Post bool
	Params string
}

type ApiGen struct {
	Url string `json:"url"`
	Auth bool `json:"auth"`
	Method string `json:"method"`
}

var (
	handlerTpl = template.Must(template.New("handlerTpl").Parse(`
func (srv *{{.NameStruct}}) {{.NameFunc}}Handler(w http.ResponseWriter, r *http.Request) { {{if .Auth}}
	if r.Header.Get("X-Auth") != "100500" {
		w.WriteHeader(403)
		body := Body{
			Error: "unauthorized",
		}
		bodyBytes, _ := json.Marshal(body)
		w.Write(bodyBytes)
		return
	}
{{end}}{{if .Post}}
	if r.Method != "POST" {
		w.WriteHeader(406)
		body := Body{
			Error: "bad method",
		}
		bodyBytes, _ := json.Marshal(body)
		w.Write(bodyBytes)
		return
	}
{{end}}
	params := {{.Params}}{}
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
	result, err := srv.{{.NameFunc}}(r.Context(), params)
	if err != nil {
		apierror, ok := err.(ApiError)
		if ok {
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
`))

	checkTpl = template.Must(template.New("checkTpl").Parse(`
{{define "strSwitchTpl"}}	switch model.{{.Name}} {
	{{range .Values}}case "{{.}}":
	{{end}}case "":
	default:
		return fmt.Errorf("{{.ParamName}} must be one of [{{.ValuesWithСomma}}]")
	}{{end}}{{define "intSwitchTpl"}}switch model.{{.Name}} {
	{{range .Values}}case {{.}}:
	{{end}}case 0:
	default: 
	return fmt.Errorf("class must be one of {{.Values}}")
}{{end}}{{define "validFieldTpl"}}{{if .IsString}}	model.{{.Name}} = r.FormValue("{{.ParamName}}"){{else}}	val, err := strconv.Atoi(r.FormValue("{{.ParamName}}"))
	if err != nil {
		return fmt.Errorf("{{.ParamName}} must be int")
	}
	model.{{.Name}} = val{{end}}{{if .Required}}
	if model.{{.Name}} == {{if .IsString}}""{{else}} 0 {{end}} {
		return fmt.Errorf("{{.ParamName}} must me not empty")
	}{{end}}{{if .DefaultNotEmpty}}
	if model.{{.Name}} == {{if .IsString}}""{{else}} 0 {{end}} {
		model.{{.Name}} = {{if .IsString}}"{{.Default}}"{{else}} {{.Default}} {{end}}
	}
{{end}}{{if eq .Min -1}}{{else}}	
	if {{if .IsString}}len(model.{{.Name}}){{else}}model.{{.Name}}{{end}} < {{.Min}} {
		return fmt.Errorf("{{.ParamName}}{{if .IsString}} len{{end}} must be >= {{.Min}}")
	}
{{end}}{{if eq .Max -1}}{{else}}	if {{if .IsString}}len(model.{{.Name}}){{else}}model.{{.Name}}{{end}} > {{.Max}} {
		return fmt.Errorf("{{.ParamName}}{{if .IsString}} len{{end}} must be <= {{.Max}}")
	}
{{end}}{{if .ValuesNotEmpty}}{{if .IsString}}{{template "strSwitchTpl" .}}{{else}}{{template "intSwitchTpl" .}}{{end}}
{{end}}
{{end}}func (model *{{.NameStruct}}) Check(r *http.Request) error {
{{range .Fields}}{{template "validFieldTpl" .}}{{end}}	return nil
}
`))

	serveTpl = template.Must(template.New("serveTpl").Parse(`
{{define "urlTpl"}}	case "{{.UrlName}}":
		h.{{.HandlerName}}Handler(w, r)
{{end}}
func (h *{{.Name}}) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	switch r.URL.Path {
{{range .Urls}}{{template "urlTpl" .}}{{end}}	default:
		w.WriteHeader(404)
		body := Body{
			Error: "unknown method",
		}
		bodyBytes, _ := json.Marshal(body)
		w.Write(bodyBytes)
		return
	}
}`))
)

func main() {
	//structsForServe := make ([]string, 0)
	handlers := make(map[string][]Rout)
	fset := token.NewFileSet()
	node, err := parser.ParseFile(fset, os.Args[1], nil, parser.ParseComments)
	if err != nil {
		log.Fatal(err)
	}

	out, _ := os.Create(os.Args[2])

	fmt.Fprintln(out, `package main`)
	fmt.Fprintln(out)
	fmt.Fprintln(out, `import "fmt"`)
	fmt.Fprintln(out, `import "net/http"`)
	fmt.Fprintln(out, `import "strconv"`)
	fmt.Fprintln(out, `import "encoding/json"`)
	fmt.Fprintln(out, `
type Body struct {
	Error    string      "json:\"error\""
	Response interface{} "json:\"response,omitempty\""
}`)
	for _, f := range node.Decls {
		switch decl := f.(type) {
		case *ast.GenDecl:
			switch decl.Tok {
			case token.TYPE:
				for _, spec := range decl.Specs {
					currType, ok := spec.(*ast.TypeSpec)
					if !ok {
						continue
					}
					currStruct, ok := currType.Type.(*ast.StructType)
					if !ok {
						continue
					}

					var tmpStruct CheckStruct
					tmpStruct.NameStruct = currType.Name.Name
					//fmt.Fprintf(out, "Struct: %s\n", currType.Name.Name)
					for _, field := range currStruct.Fields.List {

						fieldName := field.Names[0].Name
						fieldIdent, ok := field.Type.(*ast.Ident)
						if !ok {
							continue
						}
						fieldType := fieldIdent.Name
						if field.Tag != nil {
							tag := reflect.StructTag(field.Tag.Value[1: len(field.Tag.Value)-1])
							validator := tag.Get("apivalidator")
							if validator != "" {
								var tmp ValidField
								tmp.Name = fieldName
								if (fieldType == "string") {
									tmp.IsString = true
								}
								tmp.ParamName = strings.ToLower(fieldName)
								tmp.Max, tmp.Min = -1, -1
								tags := strings.Split(validator, ",")
								for _, tag := range tags {
									values := strings.Split(tag, "=")
									if len(values) > 1 {
										switch values[0] {
										case "paramname":
											tmp.ParamName = values[1]
										case "enum":
											tmp.Values = strings.Split(values[1], "|")
											tmp.ValuesWithСomma = strings.Join(tmp.Values, ", ")
											tmp.ValuesNotEmpty = true;
										case "default":
											tmp.Default = values[1]
											tmp.DefaultNotEmpty = true;
										case "max":
											tmp.Max, err = strconv.Atoi(values[1])
											if err != nil {
												tmp.Max = -1
											}
										case "min":
											tmp.Min, err = strconv.Atoi(values[1])
											if err != nil {
												tmp.Min = -1
											}
										}
									} else {
										if values[0] == "required" {
											tmp.Required = true
										}
									}
								}
								tmpStruct.Fields = append(tmpStruct.Fields, tmp)
							}
						}
					}
					checkTpl.Execute(out, tmpStruct)
				}
			}
		case *ast.FuncDecl:
			var tmpHandler HandlerStruct
			tmpHandler.NameFunc = decl.Name.Name
			if decl.Type != nil {
				paramList := decl.Type.Params.List
				if len(paramList) > 1 {
					typeName, ok := paramList[1].Type.(*ast.Ident)
					if !ok {
						continue
					}
					tmpHandler.Params = typeName.Name
				}
			}

			var tmpApi ApiGen
			if decl.Doc != nil {
				for _, comment := range decl.Doc.List {
					//fmt.Fprintf(out, "Comment: %s\n", comment.Text)
					commentText := comment.Text[14:]
					err = json.Unmarshal([]byte(commentText), &tmpApi)
					if err != nil {
						continue
					}
					tmpHandler.Auth = tmpApi.Auth
					if tmpApi.Method == "POST" {
						tmpHandler.Post = true
					}
				}
			} else {
				continue
			}

			fieldList := decl.Recv
			if (fieldList != nil) {
				fieldType, ok := fieldList.List[0].Type.(*ast.StarExpr)
				if ok {
					structName := fieldType.X.(*ast.Ident).Name
					handlers[structName] = append(handlers[structName], Rout{decl.Name.Name, tmpApi.Url})
					tmpHandler.NameStruct = structName
				}
			}
		handlerTpl.Execute(out, tmpHandler)
		case *ast.BadDecl:
			fmt.Printf("SELECT BadDecl\n")
		}
		//fmt.Fprintf(out, "\n")
	}

	fmt.Fprintf(out, ``)

	for h := range handlers {
		serveTpl.Execute(out, ServParamTpl{h, handlers[h]})
		//for _, r := range handlers[h] {
		//	urlTpl.Execute(out, r)
		//}
		//fmt.Fprintf(out,"\n}\n")
	}

}
