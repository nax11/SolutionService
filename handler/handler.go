package handler

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"path/filepath"

	"github.com/gorilla/mux"
	"github.com/nax11/solution_service/actions"
	"github.com/nax11/solution_service/models"
	"github.com/pkg/errors"
)

const (
	taskRequestURL  = "https://kuvaev-ituniversity.vps.elewise.com/tasks/%v"
	taskResponseURL = "https://kuvaev-ituniversity.vps.elewise.com/solution"
)

func Perform(w http.ResponseWriter, req *http.Request) {
	params := mux.Vars(req)
	taskName := params["taskName"]

	taskExamples, err := getTaskExamples(taskName, true)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	taskResults, err := actions.RunAction(taskName, *taskExamples)
	if err != nil {
		http.Error(w, errors.Wrap(err, "fail calculate examples").Error(), http.StatusBadRequest)
		return
	}
	respBody, err := buildResponse(taskName, *taskExamples, taskResults)
	if err != nil {
		http.Error(w, errors.Wrap(err, "fail build response").Error(), http.StatusBadRequest)
		return
	}

	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	_, err = w.Write(respBody)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}

}

func getTaskExamples(taskName string, fromFile bool) (*models.TaskExamples, error) {
	if fromFile {
		gp := filepath.Join("models", "example.json")
		body, _ := ioutil.ReadFile(gp)
		var taskExamples models.TaskExamples
		json.Unmarshal(body, &taskExamples)
		return &taskExamples, nil
	}
	url := fmt.Sprintf(taskRequestURL, taskName)
	resp, err := http.Get(url)
	if err != nil {
		return nil, errors.Wrap(err, "get task fail")
	}

	defer resp.Body.Close()
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, errors.Wrap(err, "fail read body")
	}
	var taskExamples models.TaskExamples
	err = json.Unmarshal(body, &taskExamples)
	if err != nil {
		return nil, errors.Wrap(err, "fail parse body")
	}

	return &taskExamples, nil
}

func putTaskResult(body []byte) error {
	_, err := http.Post(taskResponseURL, "application/json", bytes.NewBuffer(body))
	if err != nil {
		return errors.Wrap(err, "send task results fail")
	}
	return nil
}

func buildResponse(taskName string, taskExamples models.TaskExamples, taskResults models.TaskResults) ([]byte, error) {
	res := models.ServiceResponse{
		User: "Nick",
		Task: taskName,
		Results: models.Results{
			Payload: taskExamples,
			Results: taskResults,
		},
	}
	return json.Marshal(res)
}
