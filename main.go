package main

import (
	"azul/game"
	"encoding/json"
	"fmt"
	"html/template"
	"log"
	"log/slog"
	"net/http"
	"os"
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

/*
func GameToJson(g game.Game) string {
	b, err := json.Marshal(g)
	if err != nil {
		return ""
	}
	return string(b)
}
*/

func main() {
	port := "3000"
	fmt.Printf("serving on port: %s\n", port)
	logger := slog.New(slog.NewTextHandler(os.Stdout, nil))
	gm := game.NewGameManager()

	tmpl := template.Must(template.New("game").
		ParseFiles(
			"frontend/templates/game.html",
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
		g.Center.RED = 5
		g.Center.BLUE = 3

		err := g.Players[0].SetPatternline(3, 4, game.RED)
		err = g.Players[0].SetPatternline(1, 1, game.BLUE)
		check(err)

		g.Players[0].PlaceTileWall(0, 2)
		jsonBytes, err := json.Marshal(g)
		if err != nil {
			fmt.Println("Error encoding JSON:", err)
			return
		}

		fmt.Println(string(jsonBytes))
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
		idStr := chi.URLParam(r, "ID")
		id, _ := strconv.ParseUint(idStr, 10, 64)
		g := gm.Games[id]

		var move game.Move

		err := json.NewDecoder(r.Body).Decode(&move)

		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}

		if move.IsValid(&g, logger) {
			g.HandleMove(move, logger)
		}

		gm.Games[id] = g

	})

	http.ListenAndServe(":"+port, r)

}
