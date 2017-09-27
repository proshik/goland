package main

import (
	"fmt"
	"html/template"
	"net/http"
	"strconv"
)

type Word struct {
	Value     string
	Translate string
}

type Group struct {
	Num   int
	Title string
	Words []Word
}

var t map[string]*template.Template

var groups []Group

func init() {
	t = make(map[string]*template.Template)
	temp := template.Must(template.ParseFiles("./template/template/base.html", "./template/template/index.html"))
	t["index.html"] = temp
	temp = template.Must(template.ParseFiles("./template/template/base.html", "./template/template/groups.html"))
	t["groups.html"] = temp
	temp = template.Must(template.ParseFiles("./template/template/base.html", "./template/template/rules.html"))
	t["rules.html"] = temp
	temp = template.Must(template.ParseFiles("./template/template/base.html", "./template/template/words.html"))
	t["words.html"] = temp
}

func displayIndex(w http.ResponseWriter, r *http.Request) {
	t["index.html"].ExecuteTemplate(w, "base", nil)
}

func displayGroups(w http.ResponseWriter, r *http.Request) {
	t["groups.html"].ExecuteTemplate(w, "base", groups)
}

func displayRules(w http.ResponseWriter, r *http.Request) {
	t["rules.html"].ExecuteTemplate(w, "base", groups)
}

func displayWords(w http.ResponseWriter, r *http.Request) {
	num, err := strconv.Atoi(r.URL.Query().Get("group"))
	if err != nil {
		errorHandler(w, r, http.StatusBadRequest)
		return
	}
	for _, g := range groups {
		if g.Num == num {
			t["words.html"].ExecuteTemplate(w, "base", g)
			break
		}
	}
	//errorHandler(w, r, http.StatusNotFound)
}

func errorHandler(w http.ResponseWriter, r *http.Request, status int) {
	w.WriteHeader(status)
	if status == http.StatusNotFound {
		fmt.Fprint(w, "custom 404")
	}
}

func main() {
	groups = append(groups,
		Group{1, "Предметы и явления", []Word{{"angle", "угол"}, {"ant", "муравей"}}})
	groups = append(groups,
		Group{2, "Общие ", []Word{{"account", "счет"}, {"act", "действия"}}})

	dir := http.Dir("./template/")
	handler := http.FileServer(dir)
	http.Handle("/static/", handler)

	http.HandleFunc("/", displayIndex)
	http.HandleFunc("/groups", displayGroups)
	http.HandleFunc("/rules", displayRules)
	http.HandleFunc("/words", displayWords)

	http.ListenAndServe(":8080", nil)
}
