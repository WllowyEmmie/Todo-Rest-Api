package main

import (
	"errors"
	"fmt"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type Todo struct {
	Id        int 	`json:"id"`
	Body      string `json:"body"`
	Completed bool   `json:"completed"`
	Important bool   `json:"important"`
	Daily     bool   `json:"daily"`
}

var todolist = []Todo{
	{Id: 1, Body: "Work out", Completed: false, Important: true, Daily: true},
}

func getTodoList(context *gin.Context) {
	context.IndentedJSON(http.StatusOK, todolist)
}
func addToList(context *gin.Context) {
	var newtodo Todo
	if err := context.BindJSON(&newtodo); err != nil {
		return
	}
	todolist = append(todolist, newtodo)
	context.IndentedJSON(http.StatusCreated, newtodo)
}
func getByID(id int) (*Todo, error) {
	for i, item := range todolist {
		if item.Id == id {
			return &todolist[i], nil
		}
	}
	return nil, errors.New("todo not found")
}
func getTodoByParam(context *gin.Context) {
	id := context.Param("id")
	idInt, err := strconv.Atoi(id)
	if err != nil {
		context.IndentedJSON(http.StatusBadRequest, gin.H{"error": "Invalid ID format"})
		return
	}
	searched, err := getByID(idInt)
	if err != nil{
		context.IndentedJSON(http.StatusNotFound, gin.H{"error":  "Todo not found"})
		return
	}
	context.IndentedJSON(http.StatusOK, searched)
}
func getImport()([]Todo, error){
	var importTodos []Todo
	for _, item := range todolist{
		if item.Important{
			importTodos = append(importTodos, item)
		}
	}
	if len(importTodos) == 0 {
		return nil, errors.New("no import todos")
	}
	return importTodos, nil
}
func getDaily()([]Todo, error){
	var dailyTodos []Todo
	for _, item := range todolist{
		if item.Daily{
			dailyTodos = append(dailyTodos, item)
		}
	}
	if len(dailyTodos) == 0 {
		return nil, errors.New("no daily todos")
	}
	return dailyTodos, nil
}
func getCompleted()([]Todo, error){
	var completedTodos []Todo
	for _, item := range todolist{
		if item.Completed{
			completedTodos = append(completedTodos, item)
		}
	}
	if len(completedTodos) == 0 {
		return nil, errors.New("no completed todos")
	}
	return completedTodos, nil
}
func getImportTodo(context *gin.Context){
	importTodos, err := getImport()
	if err != nil{
		context.IndentedJSON(http.StatusNotFound, gin.H{"error":"No Important todos"})
		return
	}
	context.IndentedJSON(http.StatusOK, importTodos)
}
func getDailyTodo(context *gin.Context){
	dailyTodos, err := getDaily()
	if err != nil{
		context.IndentedJSON(http.StatusNotFound, gin.H{"error":"No Daily todos"})
		return
	}
	context.IndentedJSON(http.StatusOK, dailyTodos)
}
func getCompletedTodo(context *gin.Context){
	completedTodos, err := getCompleted()
	if err != nil{
		context.IndentedJSON(http.StatusNotFound, gin.H{"error":"No Completed todos"})
		return
	}
	context.IndentedJSON(http.StatusOK, completedTodos)
}
func getBodyById(context *gin.Context){
	id := context.Param("id")
	idInt, err := strconv.Atoi(id)
	if err != nil {
		context.IndentedJSON(http.StatusBadRequest, gin.H{"error":"Invalid id format"})
		return
	}
	todo, err := getByID(idInt)
	if err != nil{
		context.IndentedJSON(http.StatusNotFound, gin.H{"error":"Id not found"})
		return
	}
	context.IndentedJSON(http.StatusOK, todo.Body)
}
func toggleCompleteStatus(context *gin.Context){
	id := context.Param("id")
	idInt, err := strconv.Atoi(id)
	if err != nil {
		context.IndentedJSON(http.StatusBadRequest, gin.H{"error":"Invalid id format"})
		return
	}
	todo, err := getByID(idInt)
	if err != nil {
		context.IndentedJSON(http.StatusNotFound, gin.H{"error": "todo not found"})
		return
	}
	if todo.Completed {
		todo.Completed = false
	}else{
		todo.Completed = true
	}
}
func toggleDailyStatus(context *gin.Context){
	id := context.Param("id")
	idInt, err := strconv.Atoi(id)
	if err != nil {
		context.IndentedJSON(http.StatusBadRequest, gin.H{"error":"Invalid id format"})
		return
	}
	todo, err := getByID(idInt)
	if err != nil {
		context.IndentedJSON(http.StatusNotFound, gin.H{"error": "todo not found"})
		return
	}
	if todo.Daily {
		todo.Daily = false
	}else{
		todo.Daily = true
	}
}
func toggleImportStatus(context *gin.Context){
	id := context.Param("id")
	idInt, err := strconv.Atoi(id)
	if err != nil {
		context.IndentedJSON(http.StatusBadRequest, gin.H{"error":"Invalid id format"})
		return
	}
	todo, err := getByID(idInt)
	if err != nil {
		context.IndentedJSON(http.StatusNotFound, gin.H{"error": "todo not found"})
		return
	}
	if todo.Important {
		todo.Important = false
	}else{
		todo.Important = true
	}
}
func main() {
	fmt.Println("hello world")
	router := gin.Default()
	router.GET("/todolist", getTodoList)
	router.GET("/todolist/:id", getTodoByParam)
	router.GET("/todolist/body/:id", getBodyById)
	router.GET("/todolist/important", getImportTodo)
	router.GET("/todolist/completed", getCompletedTodo)
	router.GET("/todolist/daily", getDailyTodo)
	router.POST("/todolist", addToList)
	router.PATCH("/todolist/complete/:id", toggleCompleteStatus)
	router.PATCH("/todolist/daily/:id", toggleDailyStatus)
	router.PATCH("/todolist/import/:id", toggleImportStatus)
	router.Run("localhost:8080")
}
