package main

import (
	"context"
	"encoding/json"
	"log"
	"time"

	"github.com/redis/go-redis/v9"
)

// Product represents a catalog item
type Product struct {
	ID          string    `json:"id"`
	Name        string    `json:"name"`
	Description string    `json:"description"`
	Price       float64   `json:"price"`
	Category    string    `json:"category"`
	CreatedAt   time.Time `json:"created_at"`
	UpdatedAt   time.Time `json:"updated_at"`
}

// CatalogService manages product catalog with Redis caching
type CatalogService struct {
	db    ProductDB // Interface for database operations
	cache *redis.Client
	ctx   context.Context
}

// NewCatalogService creates a new catalog service
func NewCatalogService(db ProductDB, redisAddr string) (*CatalogService, error) {
	rdb := redis.NewClient(&redis.Options{
		Addr:     redisAddr,
		Password: "", // no password set
		DB:       0,  // use default DB
	})

	// Test connection
	ctx := context.Background()
	_, err := rdb.Ping(ctx).Result()
	if err != nil {
		return nil, err
	}

	return &CatalogService{
		db:    db,
		cache: rdb,
		ctx:   ctx,
	}, nil
}

// GetProduct retrieves a product by ID (cache-aside pattern)
func (s *CatalogService) GetProduct(id string) (*Product, error) {
	// Try cache first
	val, err := s.cache.Get(s.ctx, "product:"+id).Result()
	if err == nil {
		// Cache hit
		var product Product
		if err := json.Unmarshal([]byte(val), &product); err != nil {
			return nil, err
		}
		log.Printf("Cache hit for product %s\n", id)
		return &product, nil
	}

	if err != redis.Nil {
		// Unexpected error
		return nil, err
	}

	// Cache miss - get from DB
	product, err := s.db.GetProduct(id)
	if err != nil {
		return nil, err
	}

	// Store in cache
	data, err := json.Marshal(product)
	if err != nil {
		return nil, err
	}

	// Cache for 1 hour
	if err := s.cache.Set(s.ctx, "product:"+id, data, time.Hour).Err(); err != nil {
		log.Printf("Failed to cache product %s: %v\n", id, err)
	} else {
		log.Printf("Cached product %s\n", id)
	}

	return product, nil
}

// SaveProduct saves a product (write-through cache)
func (s *CatalogService) SaveProduct(product *Product) error {
	// Update timestamps
	now := time.Now()
	if product.CreatedAt.IsZero() {
		product.CreatedAt = now
	}
	product.UpdatedAt = now

	// Save to DB
	if err := s.db.SaveProduct(product); err != nil {
		return err
	}

	// Update cache
	data, err := json.Marshal(product)
	if err != nil {
		return err
	}

	return s.cache.Set(s.ctx, "product:"+product.ID, data, time.Hour).Err()
}

// GetProductsByCategory retrieves products by category (using sorted sets)
func (s *CatalogService) GetProductsByCategory(category string) ([]*Product, error) {
	cacheKey := "category:" + category

	// Try cache first
	productIDs, err := s.cache.SMembers(s.ctx, cacheKey).Result()
	if err != nil && err != redis.Nil {
		return nil, err
	}

	if len(productIDs) > 0 {
		// Cache hit - get all products using pipeline
		pipe := s.cache.Pipeline()
		cmds := make(map[string]*redis.StringCmd)

		for _, id := range productIDs {
			cmds[id] = pipe.Get(s.ctx, "product:"+id)
		}

		_, err := pipe.Exec(s.ctx)
		if err != nil && err != redis.Nil {
			return nil, err
		}

		// Collect results
		products := make([]*Product, 0, len(productIDs))
		for _, id := range productIDs {
			if val, err := cmds[id].Result(); err == nil {
				var product Product
				if err := json.Unmarshal([]byte(val), &product); err == nil {
					products = append(products, &product)
				}
			}
		}

		if len(products) > 0 {
			log.Printf("Cache hit for category %s\n", category)
			return products, nil
		}
	}

	// Cache miss or incomplete cache - get from DB
	products, err := s.db.GetProductsByCategory(category)
	if err != nil {
		return nil, err
	}

	// Update cache using pipeline
	pipe := s.cache.Pipeline()

	// Add to category set
	for _, product := range products {
		data, err := json.Marshal(product)
		if err != nil {
			continue
		}

		pipe.SAdd(s.ctx, cacheKey, product.ID)
		pipe.Set(s.ctx, "product:"+product.ID, data, time.Hour)
	}

	// Set category expiration
	pipe.Expire(s.ctx, cacheKey, time.Hour)

	if _, err := pipe.Exec(s.ctx); err != nil {
		log.Printf("Failed to cache category %s: %v\n", category, err)
	} else {
		log.Printf("Cached category %s with %d products\n", category, len(products))
	}

	return products, nil
}

// Close closes the Redis connection
func (s *CatalogService) Close() error {
	return s.cache.Close()
}
