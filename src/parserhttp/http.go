package parserhttp

import "net/http"
import "encoding/json"
import "errors"

//用户请求的格式
type request struct {
	Host string `json:"host"`
	Port string `json:"port"`
	DB   string `json:"db"`
	Sql  string `json:"sql"`
}

//返回的结果的格式
type response struct {
	Status  bool         `json:"status"`
	Message string       `json:"message"`
	Data    *parseResult `json:"data"`
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
