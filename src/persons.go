package main

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/go-chi/chi"
	"github.com/go-chi/render"
	"gitlab.com/idoko/bucketeer/db"
	//"gitlab.com/idoko/bucketeer/db"
	//"gitlab.com/idoko/bucketeer/models"
	"net/http"
	"strconv"
	//"gitlab.com/idoko/bucketeer/db"
	//"gitlab.com/idoko/bucketeer/models"
)

var itemIDKey = "itemID"

func persons(router chi.Router) {
	println("persons yra")
	router.Get("/v1/persons", getAllItems)
	router.Post("/v1/persons", createItem)
	router.Route("/v1/persons/{itemId}", func(router chi.Router) {
		router.Use(ItemContext)
		router.Get("/", getItem)
		router.Patch("/", updateItem)
		router.Delete("/", deleteItem)
	})
}
func ItemContext(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		itemId := chi.URLParam(r, "itemId")
		if itemId == "" {
			render.Render(w, r, ErrorRenderer(fmt.Errorf("item ID is required")))
			return
		}
		Id, err := strconv.Atoi(itemId)
		if err != nil {
			render.Render(w, r, ErrorRenderer(fmt.Errorf("invalid item ID")))
		}
		ctx := context.WithValue(r.Context(), itemIDKey, Id)
		next.ServeHTTP(w, r.WithContext(ctx))
	})
}
func del(b []byte) []byte {
	kk := 1
	for k, v := range b {
		//print(k, string(v)+" ")
		//println(strconv.Atoi(string(v)))
		if v == 0 {
			kk = k
			break
		}
	}
	println("?", kk)
	return b[0:kk]
}
func createItem(w http.ResponseWriter, r *http.Request) {
	item := &Person{}
	println("here")
	b := make([]byte, 100, 100)
	r.Body.Read(b)
	b = del(b)
	err := json.Unmarshal(b, &item)
	if err != nil {
		println(err.Error())
	}
	println(item.Work)
	println("!!!2", string(b))
	//if err := render.Bind(r, item); err != nil {
	//	render.Render(w, r, ErrBadRequest)
	//	return
	//}
	//println(item.to_String())
	if err, _ := dbInstance.AddItem(item); err != nil {
		render.Render(w, r, ErrorRenderer(err))
		return
	} else {
		println("yra202")
		println(item.Id)
		//var h1 *http.Header
		r.Header.Set("location", "/api/v1/persons/"+strconv.Itoa(item.Id))
		render.Respond(w, r, http.Response{
			Status:     "201",
			StatusCode: 201,
			Header:     r.Header,
		})
	}
	//if err := render.Render(w, r, item); err != nil {
	//	render.Render(w, r, ServerErrorRenderer(err))
	//	return
	//}
}

func getAllItems(w http.ResponseWriter, r *http.Request) {
	items, err := dbInstance.GetAllItems()
	if err != nil {
		render.Render(w, r, ServerErrorRenderer(err))
		return
	}
	render.JSON(w, r, items)
}

func getItem(w http.ResponseWriter, r *http.Request) {
	itemID := r.Context().Value(itemIDKey).(int)
	item, err := dbInstance.GetItemById(itemID)
	if err != nil {
		if err == db.ErrNoMatch {
			render.Render(w, r, ErrNotFound)
		} else {
			render.Render(w, r, ErrorRenderer(err))
		}
		return
	}
	render.JSON(w, r, item)
}

func deleteItem(w http.ResponseWriter, r *http.Request) {
	itemId := r.Context().Value(itemIDKey).(int)
	err := dbInstance.DeleteItem(itemId)
	if err != nil {
		if err == db.ErrNoMatch {
			render.Render(w, r, ErrNotFound)
		} else {
			render.Render(w, r, ServerErrorRenderer(err))
		}
		return
	}
}
func updateItem(w http.ResponseWriter, r *http.Request) {
	itemId := r.Context().Value(itemIDKey).(int)
	itemData := Person{}
	if err := render.Bind(r, &itemData); err != nil {
		render.Render(w, r, ErrBadRequest)
		return
	}
	item, err := dbInstance.UpdateItem(itemId, itemData)
	if err != nil {
		if err == db.ErrNoMatch {
			render.Render(w, r, ErrNotFound)
		} else {
			render.Render(w, r, ServerErrorRenderer(err))
		}
		return
	}
	if err := render.Render(w, r, &item); err != nil {
		render.Render(w, r, ServerErrorRenderer(err))
		println("ne_yra201")
		return
	} else {
		println("yra201")
	}
}

//$ curl -X POST -d '{"id" : 1 , "name": "42" , "age: 42, "adress":"42", "work":"42"}' localhost:8080/persons/
//{
//"args": {},
//"data": "",
//"files": {},
//"form": {
//"{\"answer\":42}": ""
//},
//"headers": {
//"Accept": "*/*",
//"Content-Length": "13",
//"Content-Type": "application/x-www-form-urlencoded",
//"Host": "httpbin.org",
//"User-Agent": "curl/7.58.0",
//"X-Amzn-Trace-Id": "Root=1-5ee8e3fd-8437029087be44707bd15320",
//"X-B3-Parentspanid": "2a739cfc42d28236",
//"X-B3-Sampled": "0",
//"X-B3-Spanid": "8bdf030613bb9c8d",
//"X-B3-Traceid": "75d84f317abad5232a739cfc42d28236",
//"X-Envoy-External-Address": "10.100.91.201",
//"X-Forwarded-Client-Cert": "By=spiffe://cluster.local/ns/httpbin-istio/sa/httpbin;Hash=ea8c3d70befa0d73aa0f07fdb74ec4700d42a72889a04630741193548f1a7ae1;Subject=\"\";URI=spiffe://cluster.local/ns/istio-system/sa/istio-ingressgateway-service-account"
//},
//"json": null,
//"origin": "69.84.111.39,10.100.91.201",
//"url": "http://httpbin.org/post"
//}
