package controllers

import (
	"gitlab.com/kiplexlab/go-microservice-boilerplate/internal/forms"
	"gitlab.com/kiplexlab/go-microservice-boilerplate/internal/models"

	"net/http"

	"github.com/gin-gonic/gin"
)

//TodoController ...
type TodoController struct{}

var todoModel = new(models.TodoModel)
var todoForm = new(forms.TodoForm)

//Create ...
func (ctrl TodoController) Create(c *gin.Context) {

	var form forms.CreateTodoForm
	if validationErr := c.ShouldBindJSON(&form); validationErr != nil {
		message := todoForm.Create(validationErr)
		c.AbortWithStatusJSON(http.StatusNotAcceptable, gin.H{"message": message})
		return
	}

	id, err := todoModel.Create(c.Request.Context(), &models.Todo{
		Title:   form.Title,
		Content: form.Content,
	})

	if err != nil {
		c.AbortWithStatusJSON(http.StatusNotAcceptable, gin.H{"message": "Todo could not be created"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Todo created", "id": id})
}

//All ...
// func (ctrl ArticleController) All(c *gin.Context) {
// 	userID := getUserID(c)

// 	results, err := articleModel.All(userID)
// 	if err != nil {
// 		c.AbortWithStatusJSON(http.StatusNotAcceptable, gin.H{"Message": "Could not get articles"})
// 		return
// 	}

// 	c.JSON(http.StatusOK, gin.H{"results": results})
// }

// //One ...
// func (ctrl ArticleController) One(c *gin.Context) {
// 	userID := getUserID(c)

// 	id := c.Param("id")

// 	getID, err := strconv.ParseInt(id, 10, 64)
// 	if getID == 0 || err != nil {
// 		c.AbortWithStatusJSON(http.StatusNotFound, gin.H{"Message": "Invalid parameter"})
// 		return
// 	}

// 	data, err := articleModel.One(userID, getID)
// 	if err != nil {
// 		c.AbortWithStatusJSON(http.StatusNotFound, gin.H{"Message": "Article not found"})
// 		return
// 	}

// 	c.JSON(http.StatusOK, gin.H{"data": data})
// }

// //Update ...
// func (ctrl ArticleController) Update(c *gin.Context) {
// 	userID := getUserID(c)

// 	id := c.Param("id")

// 	getID, err := strconv.ParseInt(id, 10, 64)
// 	if getID == 0 || err != nil {
// 		c.AbortWithStatusJSON(http.StatusNotFound, gin.H{"Message": "Invalid parameter"})
// 		return
// 	}

// 	var form forms.CreateAtricleForm

// 	if validationErr := c.ShouldBindJSON(&form); validationErr != nil {
// 		message := articleForm.Create(validationErr)
// 		c.AbortWithStatusJSON(http.StatusNotAcceptable, gin.H{"message": message})
// 		return
// 	}

// 	err = articleModel.Update(userID, getID, form)
// 	if err != nil {
// 		c.AbortWithStatusJSON(http.StatusNotAcceptable, gin.H{"Message": "Article could not be updated"})
// 		return
// 	}

// 	c.JSON(http.StatusOK, gin.H{"message": "Article updated"})
// }

// //Delete ...
// func (ctrl ArticleController) Delete(c *gin.Context) {
// 	userID := getUserID(c)

// 	id := c.Param("id")

// 	getID, err := strconv.ParseInt(id, 10, 64)
// 	if getID == 0 || err != nil {
// 		c.AbortWithStatusJSON(http.StatusNotFound, gin.H{"Message": "Invalid parameter"})
// 		return
// 	}

// 	err = articleModel.Delete(userID, getID)
// 	if err != nil {
// 		c.AbortWithStatusJSON(http.StatusNotAcceptable, gin.H{"Message": "Article could not be deleted"})
// 		return
// 	}

// 	c.JSON(http.StatusOK, gin.H{"message": "Article deleted"})

// }
