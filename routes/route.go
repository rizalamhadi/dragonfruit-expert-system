package routes

import (
	"dragon-fruit-expert-system/configs"
	c "dragon-fruit-expert-system/controllers"
	h "dragon-fruit-expert-system/helpers"
	"net/http"
	"time"

	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
	"github.com/go-chi/render"
	"github.com/rs/cors"
)

func Routes() *chi.Mux {

	r := chi.NewRouter()

	CORS := cors.New(cors.Options{
		AllowedOrigins:   []string{"*"},
		AllowedMethods:   []string{"GET", "POST", "PUT", "DELETE", "PATCH", "OPTIONS"},
		AllowedHeaders:   []string{"Accept", "Authorization", "Content-Type", "X-CSRF-Token"},
		AllowCredentials: true,
		MaxAge:           300, // Maximum value not ignored by any of major browsers
	})

	r.Use(CORS.Handler,
		middleware.RequestID,                          // Injects a request ID into the context of each
		middleware.DefaultCompress,                    // Compress results, mostly gzipping assets and json
		middleware.Logger,                             // Log API request c;ass
		middleware.Recoverer,                          // Recover from panics without crashing server
		render.SetContentType(render.ContentTypeJSON), // Set Content-Type headers as application/json
		middleware.Timeout(60*time.Second),
		middleware.StripSlashes,
	)

	db := configs.InitDB()

	questionController := c.NewQuestionController(db)
	penyakitController := c.NewPenyakitController(db)
	gejalaController := c.NewGejalaController(db)
	ruleController := c.NewRuleController(db)
	authController := c.NewAuthController(db)

	r.Post("/question", questionController.CreateQuestion)

	r.Route("/penyakit", func(r chi.Router) {
		r.Get("/", penyakitController.List)
		r.Get("/{id}", penyakitController.Detail)
	})

	r.Route("/gejala", func(r chi.Router) {
		r.Get("/", gejalaController.List)
		r.Get("/{id}", gejalaController.Detail)
	})

	r.Route("/rules", func(r chi.Router) {
		r.Get("/", ruleController.List)
		r.Get("/{id}", ruleController.Detail)
		r.Post("/", ruleController.Create)
	})

	r.Route("/auth", func(r chi.Router) {
		r.Post("/register", authController.Register)
		r.Post("/login", authController.Login)
	})

	r.Get("/", func(w http.ResponseWriter, r *http.Request) {
		h.ResError(w, r, 200, "No error")
	})

	return r
}
