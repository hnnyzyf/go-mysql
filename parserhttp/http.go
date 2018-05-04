package parserhttp

import "net/http"
import "encoding/json"
import "errors"
import "runtime"
import "fmt"

//用户请求的格式
type request struct {
	Host string `json:"host"`
	Port string `json:"port"`
	DB   string `json:"db"`
	Sql  string `json:"sql"`
}

//返回的结果的格式
type response struct {
	Status     bool         `json:"status"`
	Message    string       `json:"message"`
	Data       *parseResult `json:"data"`
	Stacktrace string       `json:"stacktrace"`
}

var errorResponse = &response{Status: false, Message: "", Data: nil}

//错误定义
var RequestError = errors.New("请检查请求的Json格式")

var User = ""
var PassWord = ""

//处理请求的函数
func HandleIndex(w http.ResponseWriter, r *http.Request) {
	var ask request
	err := json.NewDecoder(r.Body).Decode(&ask)

	if err != nil {
		errorResponse.Message = RequestError.Error()
		jsonResponse(w, errorResponse)
		return
	}

	// 定义recover方法，在后面程序出现异常的时候就会捕获
	defer func() {
		if r := recover(); r != nil {
			// 这里可以对异常进行一些处理和捕获
			errorResponse.Message = fmt.Sprintln(r)
			errorResponse.Data = &parseResult{Sql: ask.Sql}
			//记录当前goroutine的调用栈
			errorResponse.Stacktrace = StackTrace(false)
			jsonResponse(w, errorResponse)
		}
	}()

	res, err := parseStmt(&ask)

	if err != nil {
		errorResponse.Message = err.Error()
		errorResponse.Data = res
		jsonResponse(w, errorResponse)
		return
	}

	result := &response{Status: true, Message: "", Data: res}
	jsonResponse(w, result)

}

func jsonResponse(w http.ResponseWriter, object *response) {
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(object)
}

func ServerHttp() {
	mux := http.NewServeMux()
	mux.Handle("/", http.HandlerFunc(HandleIndex))
	http.ListenAndServe(":8000", mux)
}

func StackTrace(all bool) string {
	// Reserve 10K buffer at first
	buf := make([]byte, 1024)
	for {
		size := runtime.Stack(buf, all)
		// The size of the buffer may be not enough to hold the stacktrace,
		// so double the buffer size
		if size == len(buf) {
			buf = make([]byte, len(buf)<<1)
			continue
		}
		return string(buf[0:size])
	}
}
