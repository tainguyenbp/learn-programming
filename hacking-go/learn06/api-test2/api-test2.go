package main

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

type user struct {
	User_ID string  `json:"user_id"`
	Title   string  `json:"title"`
	Name    string  `json:"name"`
	Role    string  `json:"role"`
	Teams   string  `json:"teams"`
	Salary  float64 `json:"salary"`
}

// users slice to seed record user data.
var users = []user{
	{User_ID: "1", Title: "Platform Engineer", Name: "Tai Nguyen", Role: "Platform", Teams: "SRE Teams", Salary: 56.99},
	{User_ID: "2", Title: "System Engineer", Name: "TaiNN", Role: "System", Teams: "SRE Teams", Salary: 17.99},
	{User_ID: "3", Title: "Linux Engineer", Name: "t2jnn", Role: "Linux", Teams: "SRE Teams", Salary: 39.99},
	{User_ID: "4", Title: "System Admin", Name: "T2jNN", Role: "System", Teams: "SRE Teams", Salary: 39.99},
	{User_ID: "5", Title: "Cloud Engineer", Name: "l30nt2j", Role: "Cloud", Teams: "SRE Teams", Salary: 39.99},
	{User_ID: "6", Title: "DevOps Engineer", Name: "tainnsre", Role: "DevOps", Teams: "SRE Teams", Salary: 39.99},
	{User_ID: "7", Title: "SysOps Engineer", Name: "t2jnnsr3", Role: "SysOps", Teams: "SRE Teams", Salary: 20.99},
}

func getUsers(c *gin.Context) {
	c.IndentedJSON(http.StatusOK, users)
}

// postUsers adds an user from JSON received in the request body.
func postUsers(c *gin.Context) {
	var newUser user

	// Call BindJSON to bind the received JSON to
	// newUser.
	if err := c.BindJSON(&newUser); err != nil {
		return
	}

	// Add the new user to the slice.
	users = append(users, newUser)
	c.IndentedJSON(http.StatusCreated, newUser)
}

// getUserByuser_id locates the user whose user_id value matches the id
// parameter sent by the client, then returns that user as a response.
func getUserByuser_id(c *gin.Context) {
	user_id := c.Param("user_id")

	// Loop over the list of users, looking for
	// an user whose user_id value matches the parameter.
	for _, a := range users {
		if a.User_ID == user_id {
			c.IndentedJSON(http.StatusOK, a)
			fmt.Println("status 200 for this user ", a.User_ID)
			return
		}
	}
	c.IndentedJSON(http.StatusNotFound, gin.H{"message": "user not found"})
}

func main() {
	router := gin.Default()
	router.GET("/users", getUsers)
	router.POST("/users", postUsers)
	router.GET("/users/:user_id", getUserByuser_id)

	router.Run("localhost:8080")
}
