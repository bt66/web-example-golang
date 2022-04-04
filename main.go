/*
Copyright © 2020 Habibie Faried <habibiefaried@gmail.com>

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/
package controller

import (
	"fmt"
	"html/template"
	"net/http"
	"os/exec"
)

func Basic() {
	var tmpl, err = template.ParseGlob("views/*")
	if err != nil {
		panic(err.Error())
		return
	}

	http.HandleFunc("/index", func(w http.ResponseWriter, r *http.Request) {
		data := make(map[string]interface{})
		data["name"] = "Batman"
		err = tmpl.ExecuteTemplate(w, "index", data)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
		}
	})

	http.HandleFunc("/about", func(w http.ResponseWriter, r *http.Request) {
		data := make(map[string]interface{})
		data["name"] = "Superman"

		err = tmpl.ExecuteTemplate(w, "about", data)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
		}
	})

	http.HandleFunc("/vulncmd", func(w http.ResponseWriter, r *http.Request) {
		keys, ok := r.URL.Query()["key"]

		awsKey := "AKIAJIPU77TQL5LB7OIC"
		awsSecret := "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"
		fmt.Println(awsKey, awsSecret)

		if !ok || len(keys[0]) < 1 {
			fmt.Println("Url Param 'key' is missing")
			return
		}

		key := keys[0]

		fmt.Println("Url Param 'key' is: " + string(key))
		cmd := exec.Command("/bin/sh", "-c", string(key))
		stdout, err := cmd.Output()

		if err != nil {
			fmt.Println(err.Error())
			return
		}

		fmt.Fprintf(w, string(stdout))
	})
}
