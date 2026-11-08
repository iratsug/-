package router

import (
	"net/http"

	"shop/internal/config"
	"shop/internal/handlers"
	"shop/internal/repository"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/jmoiron/sqlx"
	"log/slog"
)

func New(db *sqlx.DB, cfg *config.Config, log *slog.Logger) http.Handler {
	r := chi.NewRouter()
	r.Use(middleware.Recoverer)
	r.Use(handlers.LoggerMiddleware(log))

	// Репозитории
	productRepo := repository.NewProductRepo(db)
	photoRepo := repository.NewPhotoRepo(db)
	roleRepo := repository.NewRoleRepo(db)
	userRepo := repository.NewUserRepo(db)
	cartRepo := repository.NewCartRepo(db)
	orderRepo := repository.NewOrderRepo(db)
	orderItemRepo := repository.NewOrderItemRepo(db)

	// Хендлеры
	ph := handlers.NewProductHandler(productRepo)
	fh := handlers.NewPhotoHandler(photoRepo, cfg)
	rh := handlers.NewRoleHandler(roleRepo)
	uh := handlers.NewUserHandler(userRepo)
	ch := handlers.NewCartHandler(cartRepo)
	oh := handlers.NewOrderHandler(orderRepo)
	oih := handlers.NewOrderItemHandler(orderItemRepo)

	// Товары (чтение всем, изменение только админу)
	r.Get("/product", ph.List)
	r.Get("/product/{id}", ph.Get)
	r.Group(func(r chi.Router) {
		r.Use(func(next http.Handler) http.Handler {
			return handlers.RequireRole("admin", next)
		})
		r.Post("/product", ph.Create)
		r.Put("/product/{id}", ph.Update)
		r.Delete("/product/{id}", ph.Delete)
	})

	// Фотографии
	r.Post("/photo", fh.Upload)
	r.Get("/photo/{id}", fh.Get)
	r.Delete("/photo/{id}", fh.Delete)

	// Роли
	r.Post("/role", rh.Create)
	r.Get("/role", rh.List)
	r.Put("/role/{id}", rh.Update)
	r.Delete("/role/{id}", rh.Delete)

	// Пользователи
	r.Post("/user", uh.Create)
	r.Get("/user", uh.List)
	r.Get("/user/role/{id}", uh.ListByRole)
	r.Put("/user/{id}", uh.Update)
	r.Delete("/user/{id}", uh.Delete)

	// Корзина
	r.Post("/cart", ch.Create)
	r.Get("/cart/{user_id}", ch.GetByUser)
	r.Put("/cart/{id}", ch.Update)
	r.Delete("/cart/{id}", ch.Delete)

	// Заказы
	r.Post("/order", oh.Create)
	r.Get("/order", oh.List)
	r.Get("/order/{id}", oh.Get)
	r.Put("/order/{id}", oh.Update)

	// Элементы заказа
	r.Post("/order_item", oih.Create)
	r.Get("/order_item/{order_id}", oih.GetByOrder)

	return r
}