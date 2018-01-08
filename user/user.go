package user

import (
	"log"
	"net/http"
	"time"

	"github.com/labstack/echo"
)

type User struct {
	Id         string `json:"id"`
	Name       string `json:"name" xml:"name" form:"name" query:"name"`
	Email      string `json:"email" xml:"email" form:"email" query:"email"`
	CreateDate string `json:"createDate"`
}

type Users []User

// GetUser method
// @author  Teerapong, Singthong
// @description This method is return all users as users json
func GetUser(c echo.Context) error {
	log.Println("getUser")

	id := c.Param("id")

	// TODO: this should connnect to database
	// to retrieve all users
	resp := new(User)
	resp.Id = id
	resp.Name = "Teerapong Singthong"
	resp.Email = "st.teerapong@goofaceter.com"

	return c.JSON(http.StatusOK, resp)
}

// QueryUser method
// @author  Teerapong, Singthong
// @description This method is return all users as users json
func QueryUser(c echo.Context) error {
	team := c.QueryParam("team")
	member := c.QueryParam("member")
	resp := "team: " + team + ", member: " + member
	return c.String(http.StatusOK, resp)
}

// GetUsers method
// @author  Teerapong, Singthong
// @description This method is return all users as users json
func GetUsers(c echo.Context) error {
	log.Println("getUsers")

	// TODO: this is placeholder of golang
	// https://medium.com/@Martynas/formatting-date-and-time-in-golang-5816112bf098
	cDate := time.Now().Format("2006-01-02")

	var users Users
	users = append(users, User{
		Id:         "1",
		Name:       "Teerapong Singthong",
		Email:      "st.teerapong@gmail.com",
		CreateDate: cDate,
	})

	return c.JSON(http.StatusOK, users)
}

// CreateUser method
// @author Teerapong, Singthong
// @description This method is create a new user, and return response as user json
func CreateUser(c echo.Context) (err error) {
	log.Println("createUser")

	// TODO: insert data to mongo
	t := time.Now()
	user := new(User)
	user.CreateDate = t.String()

	if err = c.Bind(user); err != nil {
		return
	}

	return c.JSON(http.StatusOK, user)
}
