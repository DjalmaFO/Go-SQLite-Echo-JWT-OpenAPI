package main

import (
	"fmt"
	"log"
	"net/http"
	"strconv"

	"github.com/dgrijalva/jwt-go"
	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
)

func configurarRotas() {
	e.POST("/login", func(c echo.Context) error {
		user := new(User)

		err := c.Bind(&user)
		if err != nil {
			log.Println(err.Error())
			return c.JSON(http.StatusBadRequest, "Não autorizado")
		}

		err = user.ValidateUser()
		if err != nil {
			log.Println(err.Error())
			return c.JSON(http.StatusUnauthorized, map[string]interface{}{"msg": "Usuário ou senha inválido"})
		}

		token, err := user.GenerateToken()
		if err != nil {
			return c.JSON(http.StatusBadRequest, map[string]interface{}{"msg": fmt.Sprintf("Erro inesperado: %s", err.Error())})
		}

		return c.JSON(http.StatusOK, map[string]interface{}{"token": token})
	})

	r := e.Group("/r")
	r.Use(middleware.JWT([]byte(jwtSecret)))

	r.GET("/users", func(c echo.Context) error {
		users := new(Users)

		if err := users.GetAllUsers(); err != nil {
			log.Println(err.Error())
			return c.JSON(http.StatusBadRequest, map[string]interface{}{"msg": "Falha ao obter dados"})
		}

		return c.JSON(http.StatusOK, users)
	})

	r.GET("/user/:id", func(c echo.Context) error {
		//Exemplo de como obter dados do Claims
		info := c.Get("user").(*jwt.Token)
		claims := info.Claims.(jwt.MapClaims)
		infoID := claims["id"].(string)
		infoName := claims["name"].(string)

		user := new(User)
		id, err := strconv.Atoi(c.Param("id"))
		if err != nil {
			return c.JSON(http.StatusBadRequest, err.Error())
		}
		user.ID = id

		if err := user.GerUser(); err != nil {
			log.Println(err.Error())
			return c.JSON(http.StatusBadRequest, "Falha ao obter dados")
		}

		return c.JSON(http.StatusOK,
			map[string]interface{}{"consulted_by: ": map[string]interface{}{"ID": infoID, "Name": infoName}, "result": user})
	})

	r.POST("/user", func(c echo.Context) error {
		user := new(User)

		if err := c.Bind(user); err != nil {
			log.Println(err.Error())
			return c.JSON(http.StatusBadRequest, err.Error())
		}

		if err := user.CreateNewUser(); err != nil {
			log.Println(err.Error())
			return c.JSON(http.StatusBadRequest, map[string]interface{}{"msg": fmt.Sprintf("Erro inesperado: %s", err.Error())})
		}

		return c.JSON(http.StatusOK, map[string]interface{}{"msg": "Dados armazenados com sucesso"})
	})
}
