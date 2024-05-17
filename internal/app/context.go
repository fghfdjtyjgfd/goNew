package app

import (
	"basic-api/internal/env"

	"go.mongodb.org/mongo-driver/mongo"
	"gorm.io/gorm"
)

// Context app context
type Context struct {
	Config  *env.Config
	Db      *gorm.DB
	MongoDb *mongo.Database
}
