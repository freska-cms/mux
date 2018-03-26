package mux

import (
	//"encoding/json"
	"net/http"
	"io/ioutil"

)

func  HandleJson(r *http.Request, params *RequestParams)error{
	content, err := ioutil.ReadAll(r.Body)
	if err != nil {
		return err

	}
	params.json = content
	return nil
}
