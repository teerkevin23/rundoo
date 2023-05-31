package main

import (
	"context"
	"net/http"
	"fmt"
	//"io/ioutil"
	"strings"
	"encoding/json"

	"github.com/teerkevin23/rundoo/cmd/web/domain"
	"github.com/teerkevin23/rundoo/internal/response"
)

func (app *application) home(w http.ResponseWriter, r *http.Request) {
	data := app.newTemplateData(r)

	err := response.Page(w, http.StatusOK, data, "pages/home.tmpl")
	if err != nil {
		app.serverError(w, r, err)
	}
}

func (app *application) createProduct(w http.ResponseWriter, r *http.Request) {
	c := context.Background()
	//w.Header().Set("Access-Control-Allow-Origin", "*")
	fmt.Printf("server: %s /\n", r.Method)
	fmt.Printf("server: query id: %s\n", r.URL.Query().Get("id"))
	fmt.Printf("server: content-type: %s\n", r.Header.Get("content-type"))
	fmt.Printf("server: headers:\n")
	for headerName, headerValue := range r.Header {
		fmt.Printf("\t%s = %s\n", headerName, strings.Join(headerValue, ", "))
	}

	var product domain.Product
	defer r.Body.Close()
	err := json.NewDecoder(r.Body).Decode(&product)
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
	}
	fmt.Println("1", err)
	fmt.Println("kevin", product)



	err = app.ProductUsecase.Create(c, &product)
	if err != nil {
		fmt.Println("Some error here")
		// should not be 500...
		app.serverError(w, r, err)
	}

	//productMarshall, err := json.Marshal(product)
	//fmt.Println(productMarshall, err)
	//fmt.Println(product)
	fmt.Fprintf(w, "")
}

func (app *application) getProducts(w http.ResponseWriter, r *http.Request) {
	c := context.Background()
	fmt.Printf("server: %s /\n", r.Method)
	fmt.Printf("server: query id: %s\n", r.URL.Query().Get("id"))
	fmt.Printf("server: content-type: %s\n", r.Header.Get("content-type"))
	fmt.Printf("server: headers:\n")
	for headerName, headerValue := range r.Header {
		fmt.Printf("\t%s = %s\n", headerName, strings.Join(headerValue, ", "))
	}
	fmt.Println(c)
	query := r.URL.Query()
	fmt.Println("query", query)
	filter := query["filter"][0]
	if len(filter) == 0 {
		filter = ""
	}
	fmt.Println("....", filter)

	productList, err := app.ProductUsecase.Get(c, filter)
	if err != nil {
		app.serverError(w, r, err)
	}
	fmt.Println("before ret", productList)

	response, err := json.Marshal(productList)
	if err != nil {
		fmt.Println("something wrong with marshal")
	}
	w.WriteHeader(200)
	w.Write(response)
}
