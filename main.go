package main

import (
	"database/sql"
	"fmt"
	"html/template"
	"net/http"

	_ "github.com/go-sql-driver/mysql"
)

func main() {
	http.Handle("/css/", http.StripPrefix("/css/", http.FileServer(http.Dir("css"))))
	http.Handle("/img/", http.StripPrefix("/img/", http.FileServer(http.Dir("img"))))
	http.Handle("/js/", http.StripPrefix("/js/", http.FileServer(http.Dir("js"))))
	http.Handle("/lib/", http.StripPrefix("/lib/", http.FileServer(http.Dir("lib"))))
	http.Handle("/scss/", http.StripPrefix("/scss/", http.FileServer(http.Dir("scss"))))

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		if r.Method == "POST" {

			bouttonAbout := r.FormValue("inscription.html")
			if bouttonAbout != "" {

				http.Redirect(w, r, "/"+bouttonAbout, http.StatusSeeOther)
				return
			}
			bouttonBlog := r.FormValue("connexion.html")
			if bouttonBlog != "" {

				http.Redirect(w, r, "/"+bouttonBlog, http.StatusSeeOther)

				return
			}

			Nom := r.FormValue("username")
			email := r.FormValue("email")
			Password := r.FormValue("password")
			Confirm_password := r.FormValue("confirm_password")
			Selected_commune := r.FormValue("selected_commune")

			db, err := sql.Open("mysql", "kedithre:@1996Edith@tcp(127.0.0.1:3306)/dbtest")
			if err != nil {
				http.Error(w, err.Error(), http.StatusInternalServerError)
				return
			}
			defer db.Close()

			err = db.Ping()
			if err != nil {
				panic(err.Error())
			}

			/*// Vérifier que des données valides ont été entrées
			if Nom == "" || Email == "" || Password == "" || Confirm_password == "" || Selected_commune == "" {
				http.Error(w, "Veuillez remplir tous les champs", http.StatusBadRequest)
				return
			}*/

			insertQuery, err := db.Query("INSERT INTO dbtest.electeurs (`username`, `email`, `password`, `confirm_password`,  `selected_commune`) VALUES (?, ?, ?, ?, ?);", Nom, email, Password, Confirm_password, Selected_commune)
			if err != nil {
				http.Error(w, err.Error(), http.StatusInternalServerError)
				return
			}
			defer insertQuery.Close()
			fmt.Println("Connexion réussie à la base de données")
			// Charger et afficher la page d'inscription
			fmt.Println("Données insérées avec succès dans la base de données")

			http.Redirect(w, r, "/connexion", http.StatusSeeOther)
			return
		}

		tmpl, err := template.ParseFiles("index.html")
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		tmpl.Execute(w, nil)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

	})

	http.HandleFunc("/inscription", func(w http.ResponseWriter, r *http.Request) {
		// Charger et afficher la page de connexion
		tmpl, err := template.ParseFiles("inscription.html")
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		tmpl.Execute(w, nil)
	})

	http.HandleFunc("/connexion", func(w http.ResponseWriter, r *http.Request) {
		// Charger et afficher la page de connexion
		tmpl, err := template.ParseFiles("connexion.html")
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		tmpl.Execute(w, nil)
	})

	fmt.Println("http://localhost:8080")

	http.ListenAndServe(":8080", nil)
}
