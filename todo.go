package main

import (
	"html/template"
	"net/http"
)

type Task struct {
	Name string
}
type TodoList struct {
	Tasks []Task
}

func (tl *TodoList) AddTask(name string) {
	task := Task{
		Name: name,
	}
	tl.Tasks = append(tl.Tasks, task)
}

func main() {
	todoList := TodoList{}

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		tmpl := template.Must(template.ParseFiles("todo.html"))
		tmpl.Execute(w, todoList)
	})

	http.HandleFunc("/add", func(w http.ResponseWriter, r *http.Request) {
		if r.Method == "POST" {
			taskName := r.FormValue("task")
			if taskName != "" {
				todoList.AddTask(taskName)
			}
		}
		http.Redirect(w, r, "/", http.StatusSeeOther)
	})

	http.ListenAndServe(":8080", nil)
}
