package main

import (
	"azul/game"
	"fmt"
	"html/template"
	"log"
	"net/http"
	"strconv"
	"unsafe"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
)

func temp_add(a, b int) int {
	return a + b
}

func check(err error) {
	if err != nil {
		log.Fatal(err.Error())
	}
}

func MakeNewGame(seed uint64) game.Game {
	g := game.NewGame(seed)
	return g
}

func main() {

	gm := game.NewGameManager()

	tmpl := template.Must(template.New("game").
		Funcs(template.FuncMap{
			"add": temp_add,
		}).
		ParseFiles(
			"frontend/templates/center.html",
			"frontend/templates/game.html",
			"frontend/templates/factory_display.html",
			"frontend/templates/player.html",

			"frontend/templates/sack.html",
			"frontend/templates/discarded.html",
			"frontend/templates/seed.html",
		))

	//fmt.Println(g.state)

	r := chi.NewRouter()

	r.Use(middleware.Logger)

	fs_static := http.FileServer(http.Dir("./frontend/static"))
	r.Handle("/static/*", http.StripPrefix("/static", fs_static))

	r.Get("/", func(w http.ResponseWriter, r *http.Request) {
		http.ServeFile(w, r, "./frontend/home/index.html")
	})

	r.Post("/games", func(w http.ResponseWriter, r *http.Request) {
		id := gm.CreateGame()
		g := gm.Games[id]
		//

		//
		gm.Games[id] = g
		fmt.Println("created game: ", id)
		fmt.Println(unsafe.Sizeof(g))

		url := fmt.Sprintf("/games/%d", id)

		http.Redirect(w, r, url, http.StatusSeeOther)

	})

	r.Get("/games/{ID}", func(w http.ResponseWriter, r *http.Request) {
		idStr := chi.URLParam(r, "ID")
		id, _ := strconv.ParseUint(idStr, 10, 64)

		g := gm.Games[id]

		err1 := tmpl.ExecuteTemplate(w, "game", g)
		check(err1)
	})

	r.Post("/games/{ID}/move", func(w http.ResponseWriter, r *http.Request) {

	})

	http.ListenAndServe(":3000", r)

}
