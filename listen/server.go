package listen

import (
	"encoding/json"
	"gocarbe/shared"
	"io/ioutil"
	"net/http"
)

func requestDispatcher(w http.ResponseWriter, r *http.Request) {
	jsonBody, err := ioutil.ReadAll(r.Body)
	err = r.Body.Close()
	if err != nil {
		shared.LOG(err.Error())
	}

	request := make(Request)
	err = json.Unmarshal(jsonBody, &request)
	if err != nil {
		shared.LOG(err.Error())
	}

	response := handleRequest(request)

	jsonBody, err = json.Marshal(response)
	if err != nil {
		shared.LOG(err.Error())
	}

	_, err = w.Write(jsonBody)
	if err != nil {
		shared.LOG(err.Error())
	}
}

func Start() {
	http.HandleFunc("/", requestDispatcher)

	if err := http.ListenAndServe(":5000", nil); err != nil {
		shared.LOG(err.Error())
		panic(err)
	}
}
