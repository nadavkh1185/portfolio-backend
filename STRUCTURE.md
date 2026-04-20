# Backend Structure - Portfolio CMS (Go + Gin + GORM)

## Overview
Backend ini digunakan untuk CMS portfolio dengan stack:
- Go (Golang)
- Gin (HTTP framework)
- GORM (ORM)
- PostgreSQL (Neon)

# Backend Structure - Portfolio CMS (Go + Gin + GORM)

    ## Overview
    Backend ini digunakan untuk CMS portfolio dengan stack:
    - Go (Golang)
    - Gin (HTTP framework)
    - GORM (ORM)
    - PostgreSQL (Neon)

    ## Folder Structure

    backend/
    │   .env
    │   .gitignore
    │   go.mod
    │   go.sum
    │   main.go        # Entry point aplikasi
    │   STRUCTURE.md

    ├── config/
    │   └── db.go     # Koneksi database

    ├── controllers/  # Logic handler (request/response)
    │   ├── about_controller.go
    │   ├── auth_controller.go
    │   ├── contact_controller.go
    │   ├── experience_controller.go
    │   ├── profile_controller.go
    │   ├── project_controller.go
    │   ├── skill_controller.go
    │   └── upload_controller.go

    ├── middleware/   # Auth / logging (optional)
    │   └── auth_middleware.go

    ├── models/       # Struct GORM (table schema)
    │   ├── about.go
    │   ├── contact.go
    │   ├── experience.go
    │   ├── profile.go
    │   ├── project.go
    │   ├── skill.go
    │   └── user.go

    ├── routes/       # Routing endpoint
    │   └── routes.go

    ├── uploads/      # Berkas upload (uploaded files storage)

    └── utils/        # Helper utilities
        └── hash.go

    ## Architecture Flow

    Request → Route → Controller → Model → Database

    ## Notes

    - Semua image disimpan sebagai URL (bukan file di DB)
    - Table seperti profile, about, contact = singleton (1 data saja)
    - Gunakan AutoMigrate untuk generate table
    - Jangan over-engineer (tidak pakai repository layer dulu)

    ## Run Project 
    go run main.go
    