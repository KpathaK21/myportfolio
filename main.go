package main

import (
    "html/template"
    "net/http"
    "log"
)

type Project struct {
	Name		string
    Language 	string
    URL         string
}

func main() {
    projects := []Project{
        {"Chatbot", "Go", "https://github.com/cs272-f24/project05-KpathaK21/tree/cr"},
        {"Chatbot", "Python", "https://kunjpathak.pythonanywhere.com"},
        {"AI - Decision Making Algorithm", "Python", "https://github.com/USFCS-Foundations-Spring25/assignment-2-search-KpathaK21"},
        {"AI - Clustering and Naive Bayes", "Python", "https://github.com/USFCS-Foundations-Spring25/assignment-3-clustering-and-naive-bayes-KpathaK21"},
        {"ML - Climate Analysis", "Python/Google Colab", "https://colab.research.google.com/drive/1pDD31CV8AKUtzlpoSFSP5FZaWzU95B2p"},
        {"ML - Clustering Analysis", "Python/Google Colab", "https://colab.research.google.com/drive/1aBL3saEI_axNEvF8We3HoSAS-M4vomEY"},
        {"Web Crawler", "Go", "https://github.com/cs272-f24/project04-KpathaK21/tree/cr"},
        {"Learning Management System\n (Senior Project)", "GO/HTML/CSS/JavaScript", "https://github.com/MattLADS/website/tree/KpathaK"},
        {"64-bit Processor", "Assembly(RISC-V)/C", "https://github.com/cs315-s24/project07-KpathaK21"},
        {"HTTP Server", "C", "https://github.com/cs221-03-f23/project05-KpathaK21"},
        {"Heap Allocator(xv-6)", "C", "https://github.com/USF-CS326-S25/project01-KpathaK21/tree/main/xv6-riscv"},

        
    }


    http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
        tmpl := template.Must(template.ParseFiles("templates/index.html"))
        tmpl.Execute(w, nil)
    })

    http.HandleFunc("/projects", func(w http.ResponseWriter, r *http.Request) {
        tmpl := template.Must(template.ParseFiles("templates/projects.html"))
        tmpl.Execute(w, map[string]interface{}{
            "Projects": projects,
        })
    })

    http.HandleFunc("/contacts", func(w http.ResponseWriter, r *http.Request) {
        tmpl := template.Must(template.ParseFiles("templates/contacts.html"))
        tmpl.Execute(w, nil)
    })

    http.HandleFunc("/about", func(w http.ResponseWriter, r *http.Request) {
        tmpl := template.Must(template.ParseFiles("templates/about.html"))
        tmpl.Execute(w, nil)
    })

    http.Handle("/static/", http.StripPrefix("/static/", http.FileServer(http.Dir("static"))))

    log.Println("Running on http://localhost:8080")
    log.Fatal(http.ListenAndServe(":8080", nil))
}
