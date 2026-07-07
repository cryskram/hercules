<div align="center">

# Hercules

### A high-performance Go backend for a Bond Wishlist application.

Built using **Go**, **Gin**, **GORM**, **PostgreSQL (Supabase)** and **golang-migrate**.

![Go](https://img.shields.io/badge/Go-1.24+-00ADD8?style=for-the-badge&logo=go)
![Gin](https://img.shields.io/badge/Gin-Web_Framework-00A67E?style=for-the-badge)
![Postgres](https://img.shields.io/badge/PostgreSQL-Supabase-336791?style=for-the-badge&logo=postgresql)

</div>

## Overview

Hercules is a REST API built for managing corporate bond listings and user wishlists.

The project was developed to explore backend engineering using Go while following a layered architecture consisting of:

- Repository Layer
- Service Layer
- Handler Layer
- PostgreSQL
- GORM ORM
- Gin Web Framework

The application supports searching, filtering, sorting and wishlisting bonds while maintaining a clean and scalable architecture.

# Features

## Bonds

- Browse all bonds
- Search bonds
- Filter bonds
- Sort bonds
- Pagination
- Fetch individual bond details

## Wishlists

- Create Wishlist
- Update Wishlist
- Delete Wishlist
- Add Bonds
- Remove Bonds
- View Wishlist Bonds

## Database

- Versioned migrations
- Batch seeding from Excel
- Upsert support
- Foreign Keys
- Cascade Deletes
- Indexed Queries

# Architecture

```
                HTTP Request
                     │
                     ▼
               Gin Handler
                     │
                     ▼
                Service Layer
                     │
                     ▼
              Repository Layer
                     │
                     ▼
                    GORM
                     │
                     ▼
              PostgreSQL (Supabase)
```

# Project Structure

```
hercules/

├── assets/
│   └── BondMaster.xlsx
│
├── cmd/
│   ├── migrate/
│   ├── seed/
│   └── server/
│
├── internal/
│   ├── config/
│   ├── database/
│   ├── dto/
│   ├── handlers/
│   ├── middleware/
│   ├── models/
│   ├── repository/
│   ├── routes/
│   ├── seeder/
│   ├── services/
│   └── utils/
│
├── migrations/
│
├── .env
├── go.mod
└── README.md
```

# Database Schema

## Bonds

```
bonds
│
├── isin (PK)
├── bond_name
├── brand_name
├── yield_pct
├── coupon_rate
├── rating
├── maturity_date
├── ...
```

## Wishlists

```
wishlists
│
├── id (UUID)
├── name
├── description
├── color
├── created_at
└── updated_at
```

## Wishlist Bonds

```
wishlist_bonds
│
├── wishlist_id
├── bond_isin
└── added_at
```

Relationship

```
Wishlist
    │
    │ 1
    ▼
WishlistBond
    ▲
    │
    │ n
Bond
```

# Getting Started

## Clone

```bash
git clone https://github.com/<your-username>/hercules.git

cd hercules
```

## Install Dependencies

```bash
go mod tidy
```

## Environment Variables

Create a `.env` file.

```env
PORT=8080

DATABASE_URL=postgresql://postgres:<password>@<host>:5432/postgres?sslmode=require
```

## Run Migrations

```bash
go run ./cmd/migrate up
```

## Seed Database

```bash
go run ./cmd/seed
```

This imports the bond catalogue from: `assets/BondMaster.xlsx`

## Start Server

```bash
go run ./cmd/server
```

Server starts at

```
http://localhost:8080
```

# API Endpoints

## Bonds

| Method | Endpoint           | Description    |
| ------ | ------------------ | -------------- |
| GET    | `/api/bonds`       | List all bonds |
| GET    | `/api/bonds/:isin` | Bond Details   |

### Query Parameters

| Parameter        | Description                         |
| ---------------- | ----------------------------------- |
| page             | Page Number                         |
| limit            | Page Size                           |
| search           | Search Bond Name / Brand / ISIN     |
| rating           | Filter by Rating                    |
| sector           | Filter by Sector                    |
| coupon_type      | Coupon Type                         |
| nature           | Bond Nature                         |
| payout_frequency | Monthly / Quarterly / Annual        |
| sort             | yield, maturity, rating, investment |
| order            | asc / desc                          |

## Wishlists

| Method | Endpoint             |
| ------ | -------------------- |
| GET    | `/api/wishlists`     |
| POST   | `/api/wishlists`     |
| GET    | `/api/wishlists/:id` |
| PATCH  | `/api/wishlists/:id` |
| DELETE | `/api/wishlists/:id` |

## Wishlist Bonds

| Method | Endpoint                         |
| ------ | -------------------------------- |
| GET    | `/api/wishlists/:id/bonds`       |
| POST   | `/api/wishlists/:id/bonds`       |
| DELETE | `/api/wishlists/:id/bonds/:isin` |

# Example Request

```
GET /api/bonds?page=1&limit=20&sort=yield&order=desc&rating=AAA
```

# Tech Stack

### Language

- Go

### Backend

- Gin
- GORM

### Database

- PostgreSQL
- Supabase

### Migration

- golang-migrate

### Excel Parsing

- excelize

# Testing

The repository contains a `requests.http` file with ready-to-use API requests for:

- Health Check
- Get Bonds
- Search Bonds
- Filter Bonds
- Pagination
- Sorting
- Create Wishlist
- Update Wishlist
- Delete Wishlist
- Add Bond to Wishlist
- Remove Bond from Wishlist

The requests can be executed directly from VS Code using the REST Client extension.
