package application

import (
	"github.com/akmittal/optimg/server/pkg/config"
	"github.com/akmittal/optimg/server/pkg/db"
	"github.com/akmittal/optimg/server/pkg/image"
	"github.com/akmittal/optimg/server/pkg/operation"
	"github.com/akmittal/optimg/server/pkg/router"
	"github.com/akmittal/optimg/server/pkg/server"
)

type Application struct {
	DB     *db.DB
	Cfg    *config.Config
	Router *router.Router
	Server *server.Server
}

func Get() (*Application, error) {
	cfg := config.Get()
	db, err := db.Get(cfg.GetDBConnStr())

	if err != nil {
		return nil, err
	}
	router, err := router.Get()
	server := server.Get(cfg.GetAppHost(), router)

	return &Application{
		DB:     db,
		Cfg:    cfg,
		Router: router,
		Server: server,
	}, nil
}

func (a *Application) Start() error {
	a.RegisterRoutes()
	// Migrate the schema
	a.DB.Client.AutoMigrate(&image.Image{}, &operation.Operation{}, &operation.Transformation{})

	return a.Server.Start()

}
func (a *Application) StartStaticServer(path string) error {
	a.Router.Handle("/public/{path:*}", image.ImageServer(path))
	return a.Server.Start()
}

func (a *Application) RegisterRoutes() {

	// a.Router.Get("/api/public", user.UserController())
	a.Router.Post("/api/optimize", operation.HandleStartOperation(a.DB.Client))
	a.Router.Get("/api/gallery", image.HandleGallery(a.DB.Client))
	a.Router.Get("/api/image", image.ImageController(a.DB.Client))
	a.Router.Handle("/api/static/source/*", image.HandleStaticSource())
	a.Router.Handle("/api/static/dest/*", image.HandleStaticDest())
	a.Router.Handle("/api/stats", operation.HandleStats(a.DB.Client))
	a.Router.Handle("/public/{path:*}", image.ImageServer("/Users/amittal"))

}
