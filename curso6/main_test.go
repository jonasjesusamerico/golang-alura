package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"golang-alura/curso6/controllers"
	"golang-alura/curso6/database"
	"golang-alura/curso6/models"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"strconv"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
)

var ID int

func SetupDasRotasTeste() *gin.Engine {
	gin.SetMode(gin.ReleaseMode)
	rotas := gin.Default()
	return rotas
}

func CriaAlunoMock() {
	aluno := models.Aluno{Nome: "Nome do aluno Teste", CPF: "12345678901", RG: "123456789"}
	database.DB.Create(&aluno)
	ID = int(aluno.ID)
}

func DeletaAlunoMock() {
	var aluno models.Aluno
	database.DB.Delete(&aluno, ID)
}

func TestVerificaStatusCodeDaSaudacaoComParametro(t *testing.T) {
	r := SetupDasRotasTeste()
	r.GET("/:nome", controllers.Saudacao)

	req, _ := http.NewRequest("GET", "/jonas", nil)
	resposta := httptest.NewRecorder()
	r.ServeHTTP(resposta, req)
	assert.Equal(t, http.StatusOK, resposta.Code, "Deveriam ser iguais")
	mockDaResposta := `{"API Diz":"E ai jonas, tudo beleza?"}`
	respostaBody, _ := ioutil.ReadAll(resposta.Body)
	assert.Equal(t, mockDaResposta, string(respostaBody))
}

func TestListandoTodosOsAlunosHandler(t *testing.T) {
	database.ConectaComBancoDados()
	CriaAlunoMock()
	defer DeletaAlunoMock()
	r := SetupDasRotasTeste()
	r.GET("/alunos", controllers.ExibeTodosAlunos)
	req, _ := http.NewRequest("GET", "/alunos", nil)
	resposta := httptest.NewRecorder()
	r.ServeHTTP(resposta, req)
	assert.Equal(t, http.StatusOK, resposta.Code)
}

func TestBuscaAlunoPorCpfHandler(t *testing.T) {
	database.ConectaComBancoDados()
	CriaAlunoMock()
	defer DeletaAlunoMock()
	r := SetupDasRotasTeste()
	r.GET("/alunos/cpf/:cpf", controllers.BuscaAlunoPorCpf)
	req, _ := http.NewRequest("GET", "/alunos/cpf/12345678901", nil)
	resposta := httptest.NewRecorder()
	r.ServeHTTP(resposta, req)
	assert.Equal(t, http.StatusOK, resposta.Code)
}

func TestBuscaAlunoPorIDHandler(t *testing.T) {
	database.ConectaComBancoDados()
	CriaAlunoMock()
	defer DeletaAlunoMock()
	r := SetupDasRotasTeste()
	r.GET("/alunos/:id", controllers.BuscaAlunoPorId)
	req, _ := http.NewRequest("GET", "/alunos/"+strconv.Itoa(ID), nil)
	resposta := httptest.NewRecorder()
	r.ServeHTTP(resposta, req)

	var alunoMock models.Aluno
	json.Unmarshal(resposta.Body.Bytes(), &alunoMock)
	fmt.Println(alunoMock.Nome)

	assert.Equal(t, "Nome do aluno Teste", alunoMock.Nome)
	assert.Equal(t, "12345678901", alunoMock.CPF)
	assert.Equal(t, "123456789", alunoMock.RG)
	assert.Equal(t, http.StatusOK, resposta.Code)
}

func TestDelataAlunoHandler(t *testing.T) {
	database.ConectaComBancoDados()
	CriaAlunoMock()
	r := SetupDasRotasTeste()
	r.DELETE("/alunos/:id", controllers.DeleteAluno)
	req, _ := http.NewRequest("DELETE", "/alunos/"+strconv.Itoa(ID), nil)
	resposta := httptest.NewRecorder()
	r.ServeHTTP(resposta, req)

	assert.Equal(t, http.StatusOK, resposta.Code)

}

func TestEditaUmAlunoHandler(t *testing.T) {
	database.ConectaComBancoDados()
	CriaAlunoMock()
	defer DeletaAlunoMock()
	r := SetupDasRotasTeste()
	r.PATCH("/alunos/:id", controllers.EditaAluno)
	aluno := models.Aluno{Nome: "Nome do aluno Teste", CPF: "47345678901", RG: "700456789"}
	valorJson, _ := json.Marshal(aluno)
	req, _ := http.NewRequest("PATCH", "/alunos/"+strconv.Itoa(ID), bytes.NewBuffer(valorJson))
	resposta := httptest.NewRecorder()
	r.ServeHTTP(resposta, req)

	var alunoMockAtualizado models.Aluno
	json.Unmarshal(resposta.Body.Bytes(), &alunoMockAtualizado)

	assert.Equal(t, "47345678901", alunoMockAtualizado.CPF)
	assert.Equal(t, "700456789", alunoMockAtualizado.RG)
	assert.Equal(t, "Nome do aluno Teste", alunoMockAtualizado.Nome)
}
