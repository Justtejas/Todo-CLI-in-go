package todo

import (
	"encoding/json"
	"errors"
	"os"
	"time" 
)

type item struct {
	Task      string
	Done      bool
	CreatedAt time.Time
	CompletedAt time.Time
}

type Todo []item

func (t *Todo)  Add(task string){
	todo := item{
		Task: task,
		Done: false,
		CreatedAt: time.Now(),
		CompletedAt: time.Time{},
	}
	*t = append(*t, todo)
}

func (t *Todo) Complete(index int) error{
	ls := *t
	if index <= 0 || index > len(ls){
		return errors.New("invalid index")
	}

	ls[index-1].CompletedAt = time.Now()
	ls[index-1].Done = true

	return nil
}

func (t *Todo) Delete(index int) error {
	ls := *t
	if index <= 0 || index > len(ls){
		return errors.New("invalid index")
	}

	*t = append(ls[:index-1],ls[index:]...)
	return nil
}


func (t *Todo) Load(filename string) error{
	file, err := os.ReadFile(filename)
	if err!=nil {
		if errors.Is(err, os.ErrNotExist) {
			return nil
		}
		return err
	}

	if len(file) == 0 {
		return err
	}

	err = json.Unmarshal(file,t)
	if err!=nil {
		return err
	}

	return nil
}


func (t *Todo) Store(filename string) error {
	data, err := json.Marshal(t)
	if err!= nil {
		return err
	}

	return os.WriteFile(filename,data,0644)
}