package main

import (
	"fmt"
	"log"
	"math/rand"
	"os"
	"time"

	"github.com/joho/godotenv"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

// Define model structures (these should match your actual models)
type Category struct {
	ID   uint   `gorm:"primaryKey"`
	Name string `gorm:"size:50;not null"`
	Slug string `gorm:"size:50;not null;unique"`
}

type Thread struct {
	ID         uint      `gorm:"primaryKey"`
	CategoryID uint      `gorm:"not null"`
	Title      string    `gorm:"size:200;not null"`
	CreatedAt  time.Time `gorm:"not null"`
	UpdatedAt  time.Time `gorm:"not null"`
	LastPostAt time.Time `gorm:"not null"`
	PostCount  int       `gorm:"not null;default:0"`
}

type Post struct {
	ID        uint      `gorm:"primaryKey"`
	ThreadID  uint      `gorm:"not null"`
	Content   string    `gorm:"type:text;not null"`
	AuthorIP  string    `gorm:"type:inet;not null"`
	CreatedAt time.Time `gorm:"not null"`
	IsDeleted bool      `gorm:"not null;default:false"`
}

func main() {
	// Load environment variables
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	// Set up database connection
	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=disable TimeZone=UTC",
		os.Getenv("DB_HOST"),
		os.Getenv("DB_USER"),
		os.Getenv("DB_PASS"),
		os.Getenv("DB_NAME"),
		os.Getenv("DB_PORT"),
	)
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal("Failed to connect to database:", err)
	}

	// Seed data
	seedCategories(db)
	seedThreads(db)
	seedPosts(db)

	log.Println("Seed data creation completed successfully.")
}

func seedCategories(db *gorm.DB) {
	categories := []Category{
		{Name: "General Discussion", Slug: "general-discussion"},
		{Name: "Technology", Slug: "technology"},
		{Name: "Sports", Slug: "sports"},
		{Name: "Entertainment", Slug: "entertainment"},
		{Name: "Science", Slug: "science"},
	}

	for _, category := range categories {
		result := db.Create(&category)
		if result.Error != nil {
			log.Printf("Error creating category %s: %v", category.Name, result.Error)
		}
	}
}

func seedThreads(db *gorm.DB) {
	var categories []Category
	db.Find(&categories)

	for _, category := range categories {
		for i := 0; i < 5; i++ {
			thread := Thread{
				CategoryID: category.ID,
				Title:      fmt.Sprintf("Thread %d in %s", i+1, category.Name),
				CreatedAt:  time.Now(),
				UpdatedAt:  time.Now(),
				LastPostAt: time.Now(),
			}
			result := db.Create(&thread)
			if result.Error != nil {
				log.Printf("Error creating thread in category %s: %v", category.Name, result.Error)
			}
		}
	}
}

func seedPosts(db *gorm.DB) {
	var threads []Thread
	db.Find(&threads)

	for _, thread := range threads {
		for i := 0; i < 10; i++ {
			post := Post{
				ThreadID:  thread.ID,
				Content:   fmt.Sprintf("This is post %d in thread %d", i+1, thread.ID),
				AuthorIP:  fmt.Sprintf("192.168.0.%d", rand.Intn(255)+1),
				CreatedAt: time.Now(),
			}
			result := db.Create(&post)
			if result.Error != nil {
				log.Printf("Error creating post in thread %d: %v", thread.ID, result.Error)
			}
		}
	}
}
