package main

import (
	"net/http"
	"testing"

	"github.com/steinfletcher/apitest"
)

// necessario colocar o dockerTest para nao considerar o banco local no momento de executar

func Test_ListPLanets(t *testing.T) {
	apitest.New("List planets").Handler(NewApp().DefineRoutes().Route).Get("/api/list").Expect(t).Body(`{"Info":[]}`).Status(http.StatusOK).End()
}

func Test_CreatePlanet(t *testing.T) {
	apitest.New("Create Planet").Handler(NewApp().DefineRoutes().Route).Post("/api/create").JSON(`{
	"name": "Teste",
    "climate": "arid",
    "terrain": "desert"
	}`).Expect(t).Body(`{"Info": "planeta cadastrado com sucesso"}`).Status(http.StatusCreated).End()
}

func Test_DeletePlanet(t *testing.T) {
	app := NewApp().DefineRoutes()
	apitest.New("Create Planet").Handler(app.Route).Post("/api/create").JSON(`{
		"name": "Teste",
		"climate": "arid",
		"terrain": "desert"
		}`).Expect(t).End()

	apitest.New("Delete Planet").Handler(app.Route).Delete("/api/remove").QueryParams(map[string]string{"name": "Teste"}).Expect(t).Body(`{"Removed":true}`).Status(http.StatusOK).End()
}

func Test_DeletePlanetWithErrorName(t *testing.T) {
	app := NewApp().DefineRoutes()
	apitest.New("Create Planet").Handler(app.Route).Post("/api/create").JSON(`{
		"name": "Teste10",
		"climate": "arid",
		"terrain": "desert"
		}`).Expect(t).End()

	apitest.New("Delete Planet").Handler(app.Route).Delete("/api/remove").QueryParams(map[string]string{"name": "ErrorName"}).Expect(t).Body(`{"Removed":false}`).Status(http.StatusOK).End()
}

func Test_ListCreatedPlanet(t *testing.T) {
	app := NewApp().DefineRoutes()
	apitest.New("Create Planet").Handler(app.Route).Post("/api/create").JSON(`{
		"name": "Teste60",
		"climate": "arid",
		"terrain": "desert"
	}`).Expect(t).End()

	apitest.New("List planets").Handler(NewApp().DefineRoutes().Route).Get("/api/list").Expect(t).Body(`{
		"Info": [
			{
				"Name": "Teste",
				"Climate": "arid",
				"Terrain": "desert"
			}
		]
	}`).Status(http.StatusOK).End()
}
