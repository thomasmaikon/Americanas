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
		"name": "TesteDelete",
		"climate": "arid",
		"terrain": "desert"
		}`).Expect(t).Status(http.StatusCreated).End()

	apitest.New("Delete Planet").Handler(app.Route).Delete("/api/remove").QueryParams(map[string]string{"name": "TesteDelete"}).Expect(t).Body(`{"Removed":true}`).Status(http.StatusOK).End()
}

func Test_DeletePlanetWithErrorName(t *testing.T) {
	app := NewApp().DefineRoutes()
	apitest.New("Create Planet").Handler(app.Route).Post("/api/create").JSON(`{
		"name": "Teste10",
		"climate": "arid",
		"terrain": "desert"
		}`).Expect(t).Status(http.StatusCreated).End()

	apitest.New("Delete Planet").Handler(app.Route).Delete("/api/remove").QueryParams(map[string]string{"name": "ErrorName"}).Expect(t).Body(`{"Removed":false}`).Status(http.StatusOK).End()
}

func Test_FindPlanet(t *testing.T) {
	app := NewApp().DefineRoutes()
	apitest.New("Create Planet").Handler(app.Route).Post("/api/create").JSON(`{
		"name": "Tatooine",
		"climate": "arid",
		"terrain": "desert"
	}`).Expect(t).Status(http.StatusCreated).End()

	apitest.New("Lists planets").Handler(NewApp().DefineRoutes().Route).Get("/api/list/").QueryParams(map[string]string{"name": "Tatooine"}).Expect(t).Body(`{
		"Id": "630fa686fe2e7f586b09b61c",
		"Name": "Tatooine",
		"Climate": "arid",
		"Terrain": "desert",
		"Films": [
			"https://swapi.dev/api/films/1/",
			"https://swapi.dev/api/films/3/",
			"https://swapi.dev/api/films/4/",
			"https://swapi.dev/api/films/5/",
			"https://swapi.dev/api/films/6/"
		]
	}`).Status(http.StatusOK).End()
}

func Test_CreatePlanetAlredyExist(t *testing.T) {
	app := NewApp().DefineRoutes()
	apitest.New("Create Planet").Handler(app.Route).Post("/api/create").JSON(`{
		"name": "Existente1",
		"climate": "arid",
		"terrain": "desert"
	}`).Expect(t).Status(http.StatusCreated).End()

	apitest.New("Create Planet").Handler(app.Route).Post("/api/create").JSON(`{
		"name": "Existente1",
		"climate": "arid",
		"terrain": "desert"
	}`).Expect(t).Body(`{ "Info": "Planeta ja existente"}`).End()

}
