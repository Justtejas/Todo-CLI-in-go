# Todo-CLI-in-go
This is a command line todo list app made in Golang to learn how command line apps work with flags and can take/pipe output from other commands.
This was implemented for the sole purpose of learning Golang and how command line apps work.
## Clone the repo
```
git clone https://github.com/Justtejas/Todo-CLI-in-go
```
- Open the cloned folder in terminal
- run ``` go build ./todo/ ```
- run ``` ./todo ``` with the any of the following flags:
  - -list
  - -add
  - -complete
  - -del
- To list all the tasks:
  - ``` ./todo -list ```
- To add a task:
  - ``` ./todo -add "New Task" ``` or ``` echo "New Task" | ./todo -add ```
- To mark a task as complete:
  - ``` ./todo -complete=1 ```, where 1 is the index of the task
- To delete a task from the list
  - ``` ./todo -del=1 ```, where 1 is the index of the task
