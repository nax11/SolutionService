package actions

import (
	"fmt"

	actioncyclicrotation "github.com/nax11/solution_service/actions/action-cyclic-rotation"
	"github.com/nax11/solution_service/models"
	"github.com/pkg/errors"
)

var tasks = map[string]func(sl []int, shift int) []int{
	"Циклическая ротация": actioncyclicrotation.DoTask,
}

func RunAction(taskName string, taskExamples models.TaskExamples) (models.TaskResults, error) {
	taskResults := make(models.TaskResults, len(taskExamples))
	taskAction, ok := tasks[taskName]
	if !ok {
		return nil, errors.Errorf("unsexpected taskName: %v", taskName)
	}
	for i, task := range taskExamples {
		sl, shift, err := prepareParams(task, i)
		if err != nil {
			return nil, err
		}
		result := taskAction(sl, shift)
		taskResults[i] = result
	}
	return taskResults, nil
}

func prepareParams(task []interface{}, line int) ([]int, int, error) {
	if len(task) != 2 {
		return nil, 0, fmt.Errorf("unexpected struct of task examples, fail in line: %v", line)
	}
	sl, ok := task[0].([]interface{})
	if !ok {
		return nil, 0, fmt.Errorf("unexpected param in struct of task examples, fail in line: %v, with value: %v", line, task[0])
	}
	taskSl := make([]int, len(sl))
	for i, item := range sl {
		taskSl[i] = int(item.(float64))
	}
	shift, ok := task[1].(float64)
	if !ok {
		return nil, 0, fmt.Errorf("unexpected param in struct of task examples, fail in line: %v, with shift value: %v", line, task[1])
	}
	return taskSl, int(shift), nil
}
