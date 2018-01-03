package mux

import (
	"encoding/json"
	"net/http"
)

func  HandleJson(r *http.Request, params *RequestParams){
	var data map[string]interface{}
	decoder := json.NewDecoder(r.Body)
	err := decoder.Decode(&data)

	if err != nil {
		panic(err)
	}
	// Add the form values
	for k, v := range data {
		newval := v.(string)
		var value []string
		value = append(value, newval)
		params.Add(k, value)
	}
}
