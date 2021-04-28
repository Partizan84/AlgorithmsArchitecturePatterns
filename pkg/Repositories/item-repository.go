package repositories

type myHandler struct {

	func getItems(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(items)
	}

	func getItem(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		params := mux.Vars(r)
		for _, item := range items {
			if item.ID == params["id"] {
				json.NewEncoder(w).Encode(item)
				return
			}
		}
		json.NewEncoder(w).Encode(&Item{})
	}

	func postItem(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		var item Item
		_ = json.NewDecoder(r.Body).Decode(&item)
		item.ID = strconv.Itoa(rand.Intn(1000000))
		items = append(items, item)
		json.NewEncoder(w).Encode(item)
	}

	func deleteItem(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		params := mux.Vars(r)
		for index, item := range items {
			if item.ID == params["id"] {
				items = append(items[:index], items[index+1:]...)
				break
			}
		}
		json.NewEncoder(w).Encode(items)
	}

}

func (h myHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte(`Структура объединяющая handler-функции`))
}
