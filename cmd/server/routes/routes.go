package routes

import (
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type Router interface {
	MapRoutes()
}

type router struct {
	r  *gin.Engine
	rg *gin.RouterGroup
	db *gorm.DB
}

func NewRouter(r *gin.Engine, db *gorm.DB) Router {
	return &router{r: r, db: db}
}

func (r *router) MapRoutes() {
	r.usersRoutes()
}

func (r *router) usersRoutes() {
}

