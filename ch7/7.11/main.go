package main

import (
	"fmt"
	"log"
	"net/http"
	"strconv"
	"sync"
)

/*
	Add additional handlers so that clients can create, read, update, and delete
	database entries. For example, a request of the form /update?item=sock&price=6
	will update the price of an item in the inventory and report and error if the
	item does not exist or if the price is invalid. (Warning: this change introduces
	concurrent variable updates.)
*/
const url = "localhost:8000"

func main() {
	r := &repo{db: db{"shoes": 50, "socks": 5}}

	http.HandleFunc("/", r.list)
	http.HandleFunc("/list", r.list)
	http.HandleFunc("/price", r.price)
	http.HandleFunc("/create", r.create)
	http.HandleFunc("/read", r.read)
	http.HandleFunc("/update", r.update)
	http.HandleFunc("/delete", r.delete)

	fmt.Printf("Server listening on http://%s\n", url)
	log.Fatal(http.ListenAndServe(url, nil))
}

type dollars float32

func (d dollars) String() string { return fmt.Sprintf("$%.2f", d) }

type db map[string]dollars

type repo struct {
	db   db
	lock sync.RWMutex
}

func (r *repo) list(w http.ResponseWriter, req *http.Request) {
	r.lock.RLock()
	defer r.lock.RUnlock()

	for item, price := range r.db {
		fmt.Fprintf(w, "%s: %s\n", item, price)
	}
}

func (r *repo) price(w http.ResponseWriter, req *http.Request) {
	item := req.URL.Query().Get("item")
	r.lock.RLock()
	defer r.lock.RUnlock()

	if price, ok := r.db[item]; ok {
		fmt.Fprintf(w, "%s\n", price)
	} else {
		w.WriteHeader(http.StatusNotFound) // 404
		fmt.Fprintf(w, "no such item: %q\n", item)
	}
}

func (r *repo) create(w http.ResponseWriter, req *http.Request) {
	item := req.URL.Query().Get("item")
	price := req.URL.Query().Get("price")
	if item == "" || price == "" {
		w.WriteHeader(http.StatusBadRequest)
		fmt.Fprintf(w, "missing parameters")
		return
	}
	p, err := strconv.ParseFloat(price, 32)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		fmt.Fprintf(w, "%s is not a valid price\n", price)
		return
	}
	r.lock.Lock()
	defer r.lock.Unlock()
	if _, ok := r.db[item]; !ok {
		r.db[item] = dollars(p)
		fmt.Fprintf(w, fmt.Sprintf("%s created priced %s\n", item, dollars(p)))
	} else {
		w.WriteHeader(http.StatusConflict)
		fmt.Fprintf(w, "%s already exists", item)
	}
}

func (r *repo) read(w http.ResponseWriter, req *http.Request) {
	item := req.URL.Query().Get("item")
	if item == "" {
		w.WriteHeader(http.StatusBadRequest)
		fmt.Fprintf(w, "missing item paramter")
		return
	}
	r.lock.RLock()
	defer r.lock.RUnlock()
	price, ok := r.db[item]
	if !ok {
		w.WriteHeader(http.StatusNotFound)
		return
	}
	fmt.Fprintf(w, "%s: %s\n", item, price)
}

func (r *repo) update(w http.ResponseWriter, req *http.Request) {
	item := req.URL.Query().Get("item")
	price := req.URL.Query().Get("price")
	if item == "" || price == "" {
		w.WriteHeader(http.StatusBadRequest)
		fmt.Fprintf(w, "missing parameters")
		return
	}
	p, err := strconv.ParseFloat(price, 32)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		fmt.Fprintf(w, "%s is not a valid price\n", price)
		return
	}
	r.lock.Lock()
	defer r.lock.Unlock()
	if _, ok := r.db[item]; !ok {
		w.WriteHeader(http.StatusNotFound)
		return
	}
	r.db[item] = dollars(p)
	fmt.Fprintf(w, "%s updated: %s\n", item, dollars(p))
}

func (r *repo) delete(w http.ResponseWriter, req *http.Request) {
	item := req.URL.Query().Get("item")
	if item == "" {
		w.WriteHeader(http.StatusBadRequest)
		fmt.Fprintf(w, "missing item paramter")
		return
	}
	r.lock.Lock()
	defer r.lock.Unlock()
	if _, ok := r.db[item]; !ok {
		w.WriteHeader(http.StatusNotFound)
		return
	}
	delete(r.db, item)
	fmt.Fprintf(w, "%s has been deleted\n", item)
}
