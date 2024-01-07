package utils

import (
	"strings"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type PreloadEntities map[string]bool

func ApplyDynamicPreloading(db *gorm.DB, ctx *gin.Context, allowedEntities PreloadEntities) *gorm.DB {
	preloadQuery := ctx.DefaultQuery("preload", "") // Read query param "preload" from request
	if preloadQuery == "" {
		return db
	}

	preloadEntities := strings.Split(preloadQuery, ",")

	query := db
	for _, entity := range preloadEntities {
		if allowedEntities[entity] {
			query = query.Preload(entity)
		}
	}

	return query
}
