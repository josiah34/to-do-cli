package main

func main() {
	todos := Todos{}
	storage := newStorage[Todos]("todos.json")
	storage.load(&todos)
	cmdFlags := NewCmdFlags()
	cmdFlags.Execute(&todos)
	storage.Save(todos)

}
