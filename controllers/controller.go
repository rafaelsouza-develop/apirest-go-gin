package controllers

import (
	"apirest-go-gin/database"
	"apirest-go-gin/models"
	"github.com/gin-gonic/gin"
	"net/http"
)

func ExibeTodosOsAlunos(context *gin.Context) {
	var alunos = []models.Aluno{}
	database.DB.Find(&alunos)
	context.JSON(http.StatusOK, alunos)
}

func BuscarAlunoPorId(context *gin.Context) {
	var aluno models.Aluno
	id := context.Params.ByName("id")

	database.DB.First(&aluno, id)

	if aluno.ID == 0 {
		context.JSON(http.StatusNotFound, gin.H{"message": "student not found"})
		return
	}
	context.JSON(http.StatusOK, aluno)
}

func Saudacao(c *gin.Context) {
	nome := c.Params.ByName("nome")
	c.JSON(200, gin.H{
		"API diz:": "E ai " + nome + ", tudo blz?",
	})

}

func CriaNovoAluno(context *gin.Context) {
	var aluno models.Aluno
	if err := context.ShouldBindJSON(&aluno); err != nil {
		context.JSON(http.StatusBadRequest, gin.H{
			"message": err.Error(),
		})
		return
	}
	if err := models.ValidadeDadosAluno(&aluno); err != nil {
		context.JSON(http.StatusBadRequest, gin.H{
			"message": err.Error(),
		})
		return
	}
	database.DB.Create(&aluno)
	context.JSON(http.StatusOK, aluno)
}

func DeletarAluno(context *gin.Context) {

	var aluno models.Aluno
	id := context.Params.ByName("id")

	database.DB.First(&aluno, id)

	if aluno.ID == 0 {
		context.JSON(http.StatusNotFound, gin.H{"message": "student not found"})
		return
	}
	database.DB.Delete(&aluno)
	context.JSON(http.StatusOK, aluno)

}

func AtualizarDadosAluno(context *gin.Context) {
	var aluno models.Aluno
	id := context.Params.ByName("id")

	database.DB.First(&aluno, id)

	if aluno.ID == 0 {
		context.JSON(http.StatusNotFound, gin.H{"message": "student not found"})
		return
	}

	if err := context.ShouldBindJSON(&aluno); err != nil {
		context.JSON(http.StatusBadRequest, gin.H{
			"message": err.Error(),
		})
		return
	}
	if err := models.ValidadeDadosAluno(&aluno); err != nil {
		context.JSON(http.StatusBadRequest, gin.H{
			"message": err.Error(),
		})
		return
	}

	database.DB.Model(&aluno).UpdateColumns(aluno)

	context.JSON(http.StatusOK, aluno)
}

func BuscaAlunoPorCPF(context *gin.Context) {
	var aluno models.Aluno
	cpf := context.Param("cpf")
	database.DB.Where(&models.Aluno{CPF: cpf}).First(&aluno)

	if aluno.ID == 0 {
		context.JSON(http.StatusNotFound, gin.H{"message": "student not found"})
		return
	}
	context.JSON(http.StatusOK, aluno)
}
