package handlers

import (
	"log"
	"net/http"
	"product-api/data"
	"regexp"
	"strconv"
)

type Products struct {
	l *log.Logger
}

func NewProduct(l *log.Logger) *Products {
	return &Products{l}
}

func (p *Products) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodGet {
		p.getProducts(rw, r)
		return
	}

	if r.Method == http.MethodPost{
		p.addProduct(rw, r)
		return
	}


	if r.Method == http.MethodPut{

		reg := regexp.MustCompile(`/([0-9]+)`)
		g := reg.FindAllStringSubmatch(r.URL.Path, -1)

		if len(g) != 1{
			p.l.Println("Invalid URI more than id",g)
			http.Error(rw, "Invalid URI", http.StatusBadRequest)
			return
		}

		if len(g[0])!= 2 {
			p.l.Println("Invalid URI more than one capture group",g[0][1])
			http.Error(rw, "Invalid URI", http.StatusBadRequest)
			return
		}
		idString := g[0][1]

		id, err := strconv.Atoi(idString)
		if err != nil {
			http.Error(rw, "Error", http.StatusBadRequest)
		}
		// p.l.Println("Got id:", id)

		p.updateProduct(id,rw,r)
		return
	}

	rw.WriteHeader(http.StatusMethodNotAllowed)
}

func (p *Products) getProducts(rw http.ResponseWriter, _ *http.Request) {
	lp := data.GetProducts()
	// d, err := json.Marshal(data)

	err := lp.ToJSON(rw)
	if err != nil {
		http.Error(rw, "Cannot Mashal the data", http.StatusInternalServerError)
	}
	// rw.Write(d)
}

func (p *Products) addProduct(rw http.ResponseWriter, r *http.Request) {
	
	lp := &data.Product{}
	err := lp.FromJSON(r.Body)
	if err != nil{
		http.Error(rw, "Cannot parse the data", http.StatusBadRequest)
	}

	p.l.Printf("Prod: %#v", lp)
	data.AddProduct(lp)
}

func (p *Products) updateProduct(id int, rw http.ResponseWriter, r *http.Request){

	lp := &data.Product{}
	err := lp.FromJSON(r.Body)
	if err != nil{
		http.Error(rw, "Cannot parse the data", http.StatusBadRequest)
	}

	p.l.Printf("Prod: %#v", lp)
	err = data.UpdateProduct(id, lp)
	if err == data.ErrorProductNotFound{
		http.Error(rw, "Product Not Found", http.StatusNotFound)
		return
	}
	if err != nil{
		http.Error(rw, "Product Not Found", http.StatusInternalServerError)
		return 
	}

	
}