package responser

import (
	"fmt"
	"net/http"
)

func Response(res http.ResponseWriter, stateCode int, msg string, arg ...string) {
	res.WriteHeader(stateCode)
	var data string
	if len(arg) == 0 {
		data = fmt.Sprintf(msg)
	} else {
		data = fmt.Sprintf(msg,arg)
	}
	res.Write([]byte(data))
}
