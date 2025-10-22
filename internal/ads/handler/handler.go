package handler

// func Routes() chi.Router {
// r := chi.NewRouter()
// r.Get("/", listAds)
// r.Post("/", createAd)

// return r
// }

// func List(w http.ResponseWriter, r *http.Request) {
// mock data
// ads := []map[string]interface{}{
// 	{"id": 1, "title": "Test Ad", "price": 99.99},
// }

// ads := shared.RespondJSON(w, http.StatusOK, ads)
// }

// func Create(w http.ResponseWriter, r *http.Request) {
// mock data
// var ad struct {
// 	Title string  `json:"title"`
// 	Price float64 `json:"price"`
// }

// if err := json.NewDecoder(r.Body).Decode(&ad); err != nil {
// 	log.Error().Err(err).Msg("invalid request body")
// 	shared.RespondError(w, http.StatusBadRequest, "invalid body")
// 	return
// }

// ad, err :=
// token, err := h.svc.Login(ctx, req.Email, req.Password)
// shared.RespondJSON(w, http.StatusCreated, ad)
// }
