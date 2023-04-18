package main

import (
	"log"
	"net/http"
)

func main() {
	http.HandleFunc("/todos", handleTodos)
	http.HandleFunc("/todos/", handleTodo)

	log.Fatal(http.ListenAndServe(":8080", nil))
}

type Todo struct {
    ID        string `json:"id"`
    Title     string `json:"title"`
    Completed bool   `json:"completed"`
}

func handleTodos(w http.ResponseWriter, r *http.Request) {
    if r.Method == "GET" {
        // list all todos, e.g. from a database
        todos := []Todo{
            {
                ID:        "1",
                Title:     "Buy groceries",
                Completed: false,
            },
            {
                ID:        "2",
                Title:     "Finish homework",
                Completed: true,
            },
        }

        jsonBytes, err := json.Marshal(todos)
        if err != nil {
            http.Error(w, err.Error(), http.StatusInternalServerError)
            return
        }

        w.Header().Set("Content-Type", "application/json")
        w.Write(jsonBytes)
    } else if r.Method == "POST" {
        var todo Todo
        err := json.NewDecoder(r.Body).Decode(&todo)
        if err != nil {
            http.Error(w, err.Error(), http.StatusBadRequest)
            return
        }

        // create a new todo, e.g. in a database
        todo.ID = strconv.Itoa(len(todos) + 1) // generate a unique ID
        todos = append(todos, todo)

        jsonBytes, err := json.Marshal(todo)
        if err != nil {
            http.Error(w, err.Error(), http.StatusInternalServerError)
            return
        }

        w.Header().Set("Content-Type", "application/json")
        w.WriteHeader(http.StatusCreated)
        w.Write(jsonBytes)
    } else {
        http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
    }
}
  func handleTodo(w http.ResponseWriter, r *http.Request) {
    switch r.Method {
    case "GET":
        getTodo(w, r)
    case "PUT":
        updateTodo(w, r)
    case "DELETE":
        deleteTodo(w, r)
    default:
        http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
    }
}


