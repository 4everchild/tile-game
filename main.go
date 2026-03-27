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
			g.AdvanceGame()

			// temporarily hardcoded cpu moves
			g.MakeCpuMoves(logger)
			//

			w.Header().Set("Content-Type", "application/json")
			w.WriteHeader(http.StatusOK)
			//response := map[string]string{"message": "move received"}
			json.NewEncoder(w).Encode(g)
		} else {
			w.Header().Set("Content-Type", "application/json")
			w.WriteHeader(http.StatusBadRequest)
			response := map[string]string{"message": "move not valid"}
			json.NewEncoder(w).Encode(response)
		}

		gm.Games[id] = g

	})

	http.ListenAndServe(":"+port, r)

}
