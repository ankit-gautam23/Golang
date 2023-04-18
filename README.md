I will be creating a REST API using the Go programming language and the JSON data format. I will use the net/http package that comes with Go to create a web server and handle HTTP requests.

First, let's define our data model. We will create a simple Todo struct that will have an ID, a Title, and a Completed flag:
```
type Todo struct {
    ID        string `json:"id"`
    Title     string `json:"title"`
    Completed bool   `json:"completed"`
}
```
Next, let's define our API endpoints for the CRUD operations:

```
func main() {
    http.HandleFunc("/todos", handleTodos)
    http.HandleFunc("/todos/", handleTodo)
    http.ListenAndServe(":8080", nil)
}
```
I defined two endpoints here, one for handling a list of todos (/todos) and one for handling individual todos (/todos/{id}).

Now, let's implement the handlers for these endpoints. First, let's implement the handler for the /todos endpoint:

```
func handleTodos(w http.ResponseWriter, r *http.Request) {
    switch r.Method {
    case "GET":
        getTodos(w, r)
    case "POST":
        createTodo(w, r)
    default:
        http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
    }
}
```
This handler function will switch on the HTTP method and call a different function based on the method. I have implemented the GET and POST methods here, but you could also implement PUT and DELETE methods.

Let's implement the getTodos and createTodo functions:
```
func getTodos(w http.ResponseWriter, r *http.Request) {
    todos := []Todo{
        {ID: "1", Title: "Buy groceries", Completed: false},
        {ID: "2", Title: "Clean house", Completed: true},
        {ID: "3", Title: "Take out trash", Completed: false},
    }

    jsonBytes, err := json.Marshal(todos)
    if err != nil {
        http.Error(w, err.Error(), http.StatusInternalServerError)
        return
    }

    w.Header().Set("Content-Type", "application/json")
    w.Write(jsonBytes)
}

func createTodo(w http.ResponseWriter, r *http.Request) {
    var todo Todo
    err := json.NewDecoder(r.Body).Decode(&todo)
    if err != nil {
        http.Error(w, err.Error(), http.StatusBadRequest)
        return
    }

    // do something with the new todo, like add it to a database
    // ...

    w.WriteHeader(http.StatusCreated)
}
```
The getTodos function returns a hardcoded list of todos as JSON data. The createTodo function decodes JSON data from the request body and creates a new Todo object. We have not implemented any database functionality in this example, but you could easily add a database package and replace the commented code with database calls.
