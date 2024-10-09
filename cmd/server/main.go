package main

if err := database.RunMigrations(db); err != nil {
    log.Fatalf("Failed to run database migrations: %v", err)
}
