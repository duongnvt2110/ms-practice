package gorm_client

import (
	"fmt"
	"log"
	"os"
	"time"

	"ms-practice/pkg/config"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"gorm.io/plugin/dbresolver"
)

// NewGormClient initializes a GORM DB instance with multiple sources and replicas
func NewGormClient(mysqlCfg config.Mysql) (*gorm.DB, error) {
	// Load database config from environment variables
	// Build DSNs for primaries and replicas
	var primaryDialectors []gorm.Dialector
	for _, host := range mysqlCfg.PrimaryHosts {
		if host != "" {
			primaryDialectors = append(primaryDialectors, mysql.Open(fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local",
				mysqlCfg.User, mysqlCfg.Password, host, mysqlCfg.Port, mysqlCfg.DBName)))
		}
	}

	var readDialectors []gorm.Dialector
	for _, host := range mysqlCfg.ReplicaHots {
		if host != "" {
			readDialectors = append(readDialectors, mysql.Open(fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local",
				mysqlCfg.User, mysqlCfg.Password, host, mysqlCfg.Port, mysqlCfg.DBName)))
		}
	}
	if len(primaryDialectors) == 0 {
		return nil, fmt.Errorf("no primary hosts configured for mysql")
	}

	// Configure GORM logging
	gormLogger := logger.New(
		log.New(os.Stdout, "\r\n", log.LstdFlags),
		logger.Config{
			SlowThreshold:             200 * time.Millisecond,
			LogLevel:                  logger.Info,
			IgnoreRecordNotFoundError: true,
			Colorful:                  true,
		},
	)

	// Open initial connection to the first primary database
	db, err := gorm.Open(primaryDialectors[0], &gorm.Config{
		Logger: gormLogger,
	})
	if err != nil {
		return nil, fmt.Errorf("failed to connect to primary database: %w", err)
	}

	// Register dbresolver with multiple sources (write) and replicas (read)
	err = db.Use(dbresolver.Register(dbresolver.Config{
		Sources:  primaryDialectors,         // Multiple primary sources (write)
		Replicas: readDialectors,            // Multiple read replicas (read)
		Policy:   dbresolver.RandomPolicy{}, // Load balance randomly across replicas
	}).
		SetConnMaxLifetime(30 * time.Minute).
		SetMaxIdleConns(10).
		SetMaxOpenConns(100))
	if err != nil {
		return nil, fmt.Errorf("failed to configure dbresolver: %w", err)
	}

	log.Println("Database connected with multiple sources and replicas")
	return db, nil
}
