package main

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"time"
	"winartodev/coba-mongodb/presentation"
	"winartodev/coba-mongodb/repository"
	"winartodev/coba-mongodb/usecase"

	"github.com/joeshaw/envdecode"
	"github.com/julienschmidt/httprouter"
	"github.com/subosito/gotenv"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"
)

type config struct {
	Host     string `env:"HOST, default=localhost"`
	Port     string `env:"PORT, default=8080"`
	Database struct {
		Host string `env:"MONGO_HOST, default=localhost"`
		Port string `env:"MONGO_PORT, default=27017"`
		Name string `env:"MONGO_DB_NAME, required"`
	}
}

func main() {
	// setup env
	var cfg config
	gotenv.Load(".env")

	if err := envdecode.Decode(&cfg); err != nil {
		panic(err)
	}

	// connect into mongodb server
	ctx, cancel := context.WithTimeout(context.Background(), 2*time.Second)
	defer cancel()
	mongoCfg := fmt.Sprintf("mongodb://%s:%s", cfg.Database.Host, cfg.Database.Port)
	client, err := mongo.Connect(ctx, options.Client().ApplyURI(mongoCfg))
	if err != nil {
		panic(err)
	}

	if err = client.Ping(ctx, readpref.Primary()); err != nil {
		panic(err)
	}

	database := client.Database(cfg.Database.Name)

	// initialize repo
	categoryRepo := repository.NewCategoryRepository(database)
	productRepo := repository.NewProductRepository(database)

	// initialize usecase
	categoryUsecase := usecase.NewCategoryUsecase(categoryRepo)
	productUsecase := usecase.NewProductUsecase(productRepo)

	// initialise handler
	categoryPresenter := presentation.NewCategoryPresentation(categoryUsecase)
	productPresenter := presentation.NewProductPresentation(productUsecase)

	// REST method
	r := httprouter.New()
	// category
	r.GET("/categories", categoryPresenter.FindAll)
	r.GET("/categories/:slug", categoryPresenter.FindOne)
	r.POST("/categories", categoryPresenter.Insert)
	r.PUT("/categories/:slug", categoryPresenter.Update)
	r.DELETE("/categories/:slug", categoryPresenter.Delete)

	// product
	r.GET("/products", productPresenter.FindAll)
	r.POST("/products", productPresenter.Insert)

	r.GET("/healthz", healthz)

	// configure server
	s := &http.Server{
		Addr:    fmt.Sprintf(":%s", cfg.Port),
		Handler: r,
	}

	// open connection
	log.Printf("listen and serve on http://localhost:%v\n", cfg.Port)
	if serr := s.ListenAndServe(); serr != nil {
		log.Fatal(serr)
	}
}

// healthz is function to check connection is ready will return OK
func healthz(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	w.Header().Set("Content-Type", "application/json")
	fmt.Fprintln(w, "OK")
}
