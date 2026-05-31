package main

import (
	"log"
	"net/http"

	"red_social/internal/config"
	"red_social/internal/handlers"
	"red_social/internal/middleware"
	"red_social/internal/repositories"
	"red_social/internal/services"
)

func main() {
	cfg, err := config.Load()
	if err != nil {
		log.Fatalf("Error loading config: %v", err)
	}

	db, err := repositories.NewConnection(cfg.DatabaseURL)
	if err != nil {
		log.Fatalf("Error connecting to database: %v", err)
	}
	defer db.Close()

	userRepo := repositories.NewUserRepository(db)
	sessionRepo := repositories.NewSessionRepository(db)
	postRepo := repositories.NewPostRepository(db)

	authService := services.NewAuthService(userRepo)
	postService := services.NewPostService(postRepo)
	userService := services.NewUserService(userRepo)

	authHandler := handlers.NewAuthHandler(authService, sessionRepo)
	postHandler := handlers.NewPostHandler(postService, userService)
	userHandler := handlers.NewUserHandler(userService)
	authMiddleware := middleware.NewAuthMiddleware(sessionRepo, userRepo)

	mux := http.NewServeMux()

	mux.HandleFunc("GET /health", func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
		w.Write([]byte("ok"))
	})

	mux.HandleFunc("GET /register", authHandler.RegisterGET)
	mux.HandleFunc("POST /register", authHandler.RegisterPOST)
	mux.HandleFunc("GET /login", authHandler.LoginGET)
	mux.HandleFunc("POST /login", authHandler.LoginPOST)

	mux.Handle("POST /logout", authMiddleware.RequireAuth(http.HandlerFunc(authHandler.LogoutPOST)))

	mux.Handle("GET /", authMiddleware.RequireAuth(http.HandlerFunc(postHandler.HandleGetTimeline)))
	mux.Handle("POST /posts", authMiddleware.RequireAuth(http.HandlerFunc(postHandler.HandleCreatePost)))
	mux.Handle("POST /posts/{id}/like", authMiddleware.RequireAuth(http.HandlerFunc(postHandler.HandleToggleLike)))

	mux.Handle("GET /profile", authMiddleware.RequireAuth(http.HandlerFunc(userHandler.HandleGetProfile)))
	mux.Handle("POST /profile/edit", authMiddleware.RequireAuth(http.HandlerFunc(userHandler.HandleEditProfile)))

	mux.Handle("GET /static/", http.StripPrefix("/static/", http.FileServer(http.Dir("./static"))))

	addr := ":" + cfg.Port
	log.Printf("Servidor escuchando en %s", addr)
	if err := http.ListenAndServe(addr, mux); err != nil {
		log.Fatalf("Error al iniciar el servidor: %v", err)
	}
}
