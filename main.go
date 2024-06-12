package main

import (
	"fmt"
	"log"
	"os"

	"github.com/gofiber/fiber/v2"
	"github.com/joho/godotenv"
)

func main(){
	fmt.Println("hello world")

	var x int = 5
	var p *int = &x

	println(*p,p,&x)

	type Todo struct {
		ID int `json:"id"`
		Completed bool `json:"completed"`
		Body string `json:"body"`
	}



	//variables
	// var myName string = "janith"
	// const mySecondName string = "janith"
	// myThirdName := "janith"  //default and easy way

	// fmt.Println(myName)
	// fmt.Println(mySecondName)
	// fmt.Println(myThirdName)

	//make a new app
	app := fiber.New()

	err := godotenv.Load(".env")
	if err != nil {
		log.Fatal("error loading .env file")
	}

	//get PORT from the .env file
	PORT := os.Getenv("PORT")

	todos := []Todo{}

	//get request
	app.Get("/",func(c *fiber.Ctx) error{
		return c.Status(200).JSON(todos)
	})

	

	// Create a Todo
	app.Post("/api/todo", func(c *fiber.Ctx) error {

		todo := &Todo{}

		//get the body from the request
		if err := c.BodyParser(todo); err != nil {
			return err
		}

		if todo.Body == ""{
			return c.Status(400).JSON(fiber.Map{"error" : "Todo body is missing"})
		}

		todo.ID = len(todos) + 1
		todos = append(todos, *todo)

		return c.Status(200).JSON(todo)
	})

	// Update a Todo
	 app.Patch("/api/todos/:id", func(c *fiber.Ctx) error {

		id := c.Params("id")

		for i, todo := range todos {
			if fmt.Sprint(todo.ID) == id {  //sprint is to convert int to string
				todos[i].Completed = true
				return c.Status(200).JSON(todos[i])
			}
		}

		return c.Status(404).JSON(fiber.Map{"error":"Todo not found"})
	 })

	 // Delete a Todo
	 app.Delete("/api/todos/:id", func(c *fiber.Ctx) error{
		id := c.Params("id")

		for i, todo := range todos {
			if fmt.Sprint(todo.ID) == id {
				todos = append(todos[:i], todos[i+1:]...)
				return c.Status(200).JSON(fiber.Map{"success":true})
			}
		}

		return c.Status(400).JSON(fiber.Map{"error" : "Todo bis not found"})
	 })

	 

	log.Fatal(app.Listen(":"+PORT))

}