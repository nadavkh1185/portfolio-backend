# Backend Structure - Portfolio CMS (Go + Gin + GORM)

## Overview
Backend ini digunakan untuk CMS portfolio dengan stack:
- Go (Golang)
- Gin (HTTP framework)
- GORM (ORM)
- PostgreSQL (Neon)

---

## Folder Structure

backend/
├── main.go # Entry point aplikasi
├── go.mod
├── go.sum
├── .env # Environment variables

├── config/
│ └── db.go # Koneksi database

├── models/ # Struct GORM (table schema)
│ ├── profile.go
│ ├── about.go
│ ├── skill.go
│ ├── project.go
│ ├── experience.go
│ └── contact.go

├── controllers/ # Logic handler (request/response)
│ ├── profile_controller.go
│ ├── project_controller.go
│ └── ...

├── routes/
│ └── routes.go # Routing endpoint

├── middleware/ # Auth / logging (optional nanti)
│
├── utils/ # Helper (optional)


---

## Architecture Flow

Request → Route → Controller → Model → Database

---

## Notes

- Semua image disimpan sebagai URL (bukan file di DB)
- Table seperti profile, about, contact = singleton (1 data saja)
- Gunakan AutoMigrate untuk generate table
- Jangan over-engineer (tidak pakai repository layer dulu)

---

## Run Project

```bash
go run main.go