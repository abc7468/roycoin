package explorer

import (
	"fmt"
	"html/template"
	"log"
	"net/http"

	"github.com/abc7468/roycoin/blockchain"
)

var templates *template.Template // template를 한번에 로드하여 define으로 가져올 수 있도록 하기 위한 변수
const (
	templateDir string = "explorer/templates/"
)

type homeData struct {
	PageTitle string
	Blocks    []*blockchain.Block
}

func home(rw http.ResponseWriter, r *http.Request) {
	data := homeData{"Home", nil}
	templates.ExecuteTemplate(rw, "home", data)
}
func add(rw http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case "GET":
		templates.ExecuteTemplate(rw, "add", nil)
	case "POST":
		r.ParseForm()
		data := r.Form.Get("blockData")
		blockchain.Blockchain().AddBlock(data)
		http.Redirect(rw, r, "/", http.StatusPermanentRedirect)
	}
}

func Start(port int) {
	handler := http.NewServeMux()
	templates = template.Must(template.ParseGlob(templateDir + "pages/*.gohtml"))
	templates = template.Must(templates.ParseGlob(templateDir + "partials/*.gohtml")) // **/*.go.html와 같은 문법은 go에서 지원하지 않음
	//	http.HandleFunc("/",func(rw http.ResponseWriter, r *http.Request) {})
	handler.HandleFunc("/", home)
	handler.HandleFunc("/add", add)
	fmt.Printf("Listening on http://localhost:%d\n", port)
	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%d", port), handler))
}
