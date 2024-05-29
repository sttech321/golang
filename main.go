package main

import (
    "net/http"
    "strconv"

    "github.com/labstack/echo/v4"
    "github.com/labstack/echo/v4/middleware"
)

type User struct {
    ID    int    `json:"id"`
    Name  string `json:"name"`
    Email string `json:"email"`
}

var (
    users = map[int]*User{}
    seq   = 1
)

func createUser(c echo.Context) error {
    user := &User{
        ID: seq,
    }
    if err := c.Bind(user); err != nil {
        return err
    }
    users[user.ID] = user
    seq++
    return c.JSON(http.StatusCreated, user)
}

func getUser(c echo.Context) error {
    id, _ := strconv.Atoi(c.Param("id"))
    user, ok := users[id]
    if !ok {
        return c.JSON(http.StatusNotFound, echo.Map{"error": "User not found"})
    }
    return c.JSON(http.StatusOK, user)
}

func updateUser(c echo.Context) error {
    id, _ := strconv.Atoi(c.Param("id"))
    user, ok := users[id]
    if !ok {
        return c.JSON(http.StatusNotFound, echo.Map{"error": "User not found"})
    }
    if err := c.Bind(user); err != nil {
        return err
    }
    users[id] = user
    return c.JSON(http.StatusOK, user)
}

func deleteUser(c echo.Context) error {
    id, _ := strconv.Atoi(c.Param("id"))
    if _, ok := users[id]; !ok {
        return c.JSON(http.StatusNotFound, echo.Map{"error": "User not found"})
    }
    delete(users, id)
    return c.NoContent(http.StatusNoContent)
}

func listUsers(c echo.Context) error {
    return c.JSON(http.StatusOK, users)
}

func main() {
    e := echo.New()

    // Middleware
    e.Use(middleware.Logger())
    e.Use(middleware.Recover())

    // Routes
    e.POST("/users", createUser)
    e.GET("/users/:id", getUser)
    e.PUT("/users/:id", updateUser)
    e.DELETE("/users/:id", deleteUser)
    e.GET("/users", listUsers)

    // Serve static files
    e.Static("/", "public")

    // Start server
    e.Logger.Fatal(e.Start(":8080"))
}
