/*
An implementation of a solution to the "Angry Professor" problem as a JSON service.
See: https://www.hackerrank.com/challenges/angry-professor
*/
package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

type AttInfo struct {
	NumStudents   int
	MinAttendance int
	ArrivalTimes  []int
}

type AngryInfo struct {
	NumClasses  int
	Attendances []AttInfo
}

func (a AngryInfo) checkCancelledClasses() []string {
	result := make([]string, a.NumClasses)
	for i, att := range a.Attendances {
		k := att.MinAttendance
		for _, t := range att.ArrivalTimes {
			if t <= 0 {
				k--
			}
		}
		result[i] = "YES"
		if k <= 0 {
			result[i] = "NO"
		}
	}
	return result
}

func handleAngryProfessor(w http.ResponseWriter, r *http.Request) {
	var angry AngryInfo
	var err error
	var body []byte

	if body, err = ioutil.ReadAll(r.Body); err == nil {
		if err = json.Unmarshal(body, &angry); err == nil {
			fmt.Printf("Err nil, json")
			result := angry.checkCancelledClasses()
			if body, err = json.Marshal(result); err == nil {
				w.Header().Add("Content-Type", "application/json")
				w.Write(body)
				return
			}
		}
	}

	w.WriteHeader(500)
	w.Write([]byte(err.Error()))
}

func main() {
	http.HandleFunc("/angry", handleAngryProfessor)
	http.ListenAndServe(":8080", nil)
}
