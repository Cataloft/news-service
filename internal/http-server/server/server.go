package server

import (
	"news/internal/config"
	"news/internal/http-server/handlers/news/edit"
	"news/internal/http-server/handlers/news/list"
	"news/internal/repository/postgres"

	"github.com/gofiber/fiber/v2"
)

type Server struct {
	App *fiber.App
	DB  *postgres.Postgres
	cfg *config.Config
}

func New(db *postgres.Postgres, cfg *config.Config) *Server {
	app := fiber.New()

	return &Server{
		App: app,
		DB:  db,
		cfg: cfg,
	}
}

func (s *Server) Start() error {
	s.initHandlers()
	return s.App.Listen(s.cfg.Address)
}

func (s *Server) initHandlers() {
	s.App.Post("/edit/:id", edit.UpdateNews(s.DB))
	s.App.Get("/list", list.GetListNews(s.DB))
}
