package main

import (
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"
)

type article struct {
	Id      uint   `json:"id" form:"id"`
	Title   string `json:"title" form:"title"`
	Content string `json:"content" form:"content"`
}

// dummy data
var data = []article{
	article{1, "lorem", "lorem"},
	article{2, "ipsum", "ipsum"},
	article{3, "abcd", "abcd"},
	article{3, "efgh", "efgh"},
}

func main() {
	// create a new echo instance
	e := echo.New()
	// define endpoint / Route / to handler function
	e.GET("/articles", GetArticleController)
	e.POST("/articles", PostArticleController)
	e.PUT("/articles/:article_id", PutArticleController)
	e.GET("/articles/:article_id", GetArticleByIdController)
	// start the server, and log if it fails
	e.Logger.Fatal(e.Start(":8080"))
}

func GetArticleController(c echo.Context) error {
	page := c.QueryParam("page")
	status := c.QueryParam("status")

	return c.JSON(http.StatusOK, map[string]any{
		"message": "success",
		"page":    page,
		"status":  status,
		"data":    data,
	})
}

func PostArticleController(c echo.Context) error {
	title := c.FormValue("title")
	content := c.FormValue("content")

	var responseData = map[string]any{
		"title":   title,
		"content": content,
		"message": "success insert data",
		"status":  true,
	}
	return c.JSON(http.StatusCreated, responseData)
}

func GetArticleByIdController(c echo.Context) error {
	id := c.Param("article_id")
	idConv, errConv := strconv.Atoi(id)
	if errConv != nil {
		return c.JSON(http.StatusBadRequest, map[string]any{
			"message": "error convert param. id should be number",
			"status":  "failed",
		})
	}
	var data = article{
		Id:      uint(idConv),
		Title:   "naruto",
		Content: "naruto oke",
	}
	return c.JSON(http.StatusOK, data)
}

func PutArticleController(c echo.Context) error {
	id := c.Param("article_id")
	idConv, errConv := strconv.Atoi(id)
	if errConv != nil {
		return c.JSON(http.StatusBadRequest, map[string]any{
			"message": "error convert param. id should be number",
			"status":  "failed",
		})
	}

	articleInput := article{}
	errBind := c.Bind(&articleInput)
	if errBind != nil {
		return c.JSON(http.StatusBadRequest, map[string]any{
			"message": "error bind data",
			"status":  "failed",
		})
	}
	return c.JSON(http.StatusOK, map[string]any{
		"databind": articleInput,
		"id_param": idConv,
	})
}
