package responser

import (
	"fmt"
	"net/http"
)

func defaultAllCorsSetting(res http.ResponseWriter) {
	res.Header().Set("Access-Control-Allow-Origin", "*")
	res.Header().Set("Access-Control-Max-Age", "15")
	res.Header().Set("Access-Control-Allow-Methods", "POST, GET, OPTIONS")
}

func ResponseJSON(res http.ResponseWriter, stateCode int, msg string, arg ...string) {
	defaultAllCorsSetting(res)
	res.WriteHeader(stateCode)
	res.Header().Set("Content-Type", "application/json")
	var data string
	if len(arg) == 0 {
		data = fmt.Sprintf(msg)
	} else {
		data = fmt.Sprintf(msg, arg)
	}
	res.Write([]byte(data))
}
