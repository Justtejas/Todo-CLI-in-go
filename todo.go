package todo

import (
	"encoding/json"
	"errors"
	"fmt"
	"os"
	"time"

	"github.com/alexeyco/simpletable"
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


func (t *Todo) Print()  {
	
	table := simpletable.New()

	table.Header = &simpletable.Header{
		Cells: []*simpletable.Cell{
			{Align: simpletable.AlignCenter, Text: "Index"},
			{Align: simpletable.AlignCenter, Text: "Task"},
			{Align: simpletable.AlignCenter, Text: "Done"},
			{Align: simpletable.AlignCenter, Text: "Created At"},
			{Align: simpletable.AlignCenter, Text: "Completed At"},
		},
	}	

	var cells [][]*simpletable.Cell

	for idx, item := range *t {
		idx++
		task := item.Task
		done := "No"
		if item.Done {
			task = fmt.Sprintf("\u2705 %s", item.Task)
			done = "Yes"
		}
		cells = append(cells, []*simpletable.Cell{
			{Text: fmt.Sprintf("%d", idx)},
			{Text: task},
			{Text: done},
			{Text: item.CreatedAt.Format(time.RFC822)},
			{Text: item.CompletedAt.Format(time.RFC822)},
		})
	}

	table.Body = &simpletable.Body{Cells: cells}

	table.Footer = &simpletable.Footer{Cells: []*simpletable.Cell{
		{Align: simpletable.AlignCenter, Span: 5, Text: fmt.Sprintf("You have %d pending task(s)", t.CountPending())},
	}}

	table.SetStyle(simpletable.StyleUnicode)

	table.Println()
}

func (t *Todo) CountPending() int {
	total := 0
	for _, item := range *t {
		if !item.Done {
			total++
		}
	}

	return total
}