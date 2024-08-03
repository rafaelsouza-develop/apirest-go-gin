package main

import (
	"apirest-go-gin/controllers"
	"apirest-go-gin/database"
	"apirest-go-gin/models"
	"encoding/json"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
	"io"
	"net/http"
	"net/http/httptest"
	"strconv"
	"testing"
)

var idAluno int

func SetupRouter() *gin.Engine {
	gin.SetMode(gin.ReleaseMode)
	r := gin.Default()
	return r
}

func CriaAlunoMock() {
	aluno := models.Aluno{
		Nome: "Nome Aluno",
		CPF:  "12312312311",
		RG:   "123123123",
	}
	database.DB.Create(&aluno)
	idAluno = int(aluno.ID)
}

func DeletaAlunoMock() {
	var aluno models.Aluno
	database.DB.Delete(&aluno, idAluno)
}

func TestQuandoChamaSaudacaoDeveRetornarStatusCode200(t *testing.T) {
	r := SetupRouter()
	r.GET("/:nome", controllers.Saudacao)
	req, _ := http.NewRequest("GET", "/rafael", nil)
	response := httptest.NewRecorder()
	r.ServeHTTP(response, req)

	assert.Equal(t, http.StatusOK, response.Code, "Deveriam ser iguais")
	mockResposta := `{"API diz:":"E ai rafael, tudo blz?"}`
	respostaBody, _ := io.ReadAll(response.Body)
	assert.Equal(t, mockResposta, string(respostaBody))

}

func TestListandoTodosAlunosHandler(t *testing.T) {
	database.ConectaComBancoDeDados()
	CriaAlunoMock()
	defer DeletaAlunoMock()
	r := SetupRouter()
	r.GET("/alunos", controllers.ExibeTodosOsAlunos)

	req, _ := http.NewRequest("GET", "/alunos", nil)
	response := httptest.NewRecorder()
	r.ServeHTTP(response, req)

	assert.Equal(t, http.StatusOK, response.Code, "Deveriam ser iguais")
	fmt.Println(response.Body)
}

func TestBuscaAlunoPorCpf(t *testing.T) {
	database.ConectaComBancoDeDados()
	CriaAlunoMock()
	defer DeletaAlunoMock()
	r := SetupRouter()
	r.GET("/alunos/cpf/:cpf", controllers.BuscaAlunoPorCPF)

	req, _ := http.NewRequest("GET", "/alunos/cpf/12312312311", nil)
	response := httptest.NewRecorder()

	r.ServeHTTP(response, req)

	assert.Equal(t, http.StatusOK, response.Code, "Deveriam ser iguais")

}

func TestBuscaAlunoPorId(t *testing.T) {
	database.ConectaComBancoDeDados()
	CriaAlunoMock()
	defer DeletaAlunoMock()
	r := SetupRouter()

	r.GET("/alunos/:id", controllers.BuscarAlunoPorId)
	pathBusca := "/alunos/" + strconv.Itoa(idAluno)
	req, _ := http.NewRequest("GET", pathBusca, nil)
	response := httptest.NewRecorder()
	r.ServeHTTP(response, req)

	var alunoMock models.Aluno
	json.Unmarshal(response.Body.Bytes(), &alunoMock)
	fmt.Println(alunoMock.Nome)

	assert.Equal(t, "Nome Aluno", alunoMock.Nome, "Deveriam ser iguais")
	assert.Equal(t, "12312312311", alunoMock.CPF, "Deveriam ser iguais")
	assert.Equal(t, "123123123", alunoMock.RG, "Deveriam ser iguais")
	assert.Equal(t, http.StatusOK, response.Code, "Deveriam ser iguais")
}
