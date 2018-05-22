package render

import (
	"html/template"
	"net/http"
	"github.com/gorilla/context"
	"github.com/unrolled/render"
)


var Render *render.Render
var fm = template.FuncMap{
	"safe": func(raw string) template.HTML {
		return template.HTML(raw)
	},
}

func Init() {
	debug := true
	Render = render.New(render.Options{
		Directory:     "views",
		Extensions:    []string{".html"},
		Delims:        render.Delims{"{{", "}}"},
		IndentJSON:    false,
		Funcs:         []template.FuncMap{fm},
		IsDevelopment: debug,
	})
}

type ResponseData struct {
	Status  bool   `json:"status"`
	Message string `json:"message"`
	Code    int64 	`json:"code"`
	Data    interface{} `json:"data"`
}

func Put(r *http.Request, key string, val interface{}) {
	m, ok := context.GetOk(r, "DATA_MAP")
	if ok {
		mm := m.(map[string]interface{})
		mm[key] = val
		context.Set(r, "DATA_MAP", mm)
	} else {
		context.Set(r, "DATA_MAP", map[string]interface{}{key: val})
	}
}

func HTML(r *http.Request, w http.ResponseWriter, name string, htmlOpt ...render.HTMLOptions) {
	Render.HTML(w, http.StatusOK, name, context.Get(r, "DATA_MAP"), htmlOpt...)
}

func Respone(w http.ResponseWriter, v interface{}) {
	Render.JSON(w, http.StatusOK, v)
}

func ResponeErr(w http.ResponseWriter, msg string) {
	Render.JSON(w, http.StatusOK, ResponseData{Status:false,Message:msg})
}

// 返回一个文件
func ResponeFile(w http.ResponseWriter, file []byte, filename string,v interface{}) {
	if len(file) == 0{
		Respone(w,v)
	}else {
		w.Header().Add("Content-Type", "application/octet-stream")
		w.Header().Add("content-disposition", "attachment; filename=\""+filename+"\"")
		Render.Data(w, 200, file)
	}
}


// 使用方法
// 导入render 并 render.Init()
