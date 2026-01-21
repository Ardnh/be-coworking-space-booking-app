# MODELS Directory

## 📖 Filosofi
**"Models are the source of truth about your data"**. Models merepresentasikan struktur data aplikasi dan business entities. Filosofinya adalah memiliki single source of truth untuk struktur data yang digunakan di seluruh aplikasi.

## 🎯 Definisi
Directory yang berisi definisi struct Go yang merepresentasikan tabel database dan business entities. Menggunakan GORM tags untuk mapping ke database dan JSON tags untuk API responses.

## 💡 Tanggung Jawab
- Definisi struktur data (entities)
- Database schema mapping (via GORM tags)
- JSON serialization rules
- Data validation rules
- Business rules di level entity
- Relationships antar entities

## 📝 Contoh Struktur
```
models/
├── user.go           # User entity
├── product.go        # Product entity
├── order.go          # Order entity
├── common.go         # Shared models (timestamps, pagination)
└── README.md         # Dokumentasi ini
```

## 🔧 Contoh Kode

### user.go
```go
package models

import (
    "time"
    "github.com/google/uuid"
    "gorm.io/gorm"
)

type User struct {
    ID        uuid.UUID `gorm:"type:uuid;primaryKey" json:"id"`
    Email     string    `gorm:"uniqueIndex;not null" json:"email"`
    Name      string    `gorm:"not null" json:"name"`
    Password  string    `gorm:"not null" json:"-"` // Hidden from JSON
    Role      string    `gorm:"type:varchar(20);default:'user'" json:"role"`
    IsActive  bool      `gorm:"default:true" json:"is_active"`
    CreatedAt time.Time `gorm:"autoCreateTime" json:"created_at"`
    UpdatedAt time.Time `gorm:"autoUpdateTime" json:"updated_at"`

    // Relationships
    Orders []Order `gorm:"foreignKey:UserID" json:"orders,omitempty"`
}

// BeforeCreate - GORM hook
func (u *User) BeforeCreate(tx *gorm.DB) error {
    u.ID = uuid.New()
    return nil
}

// TableName - Custom table name
func (User) TableName() string {
    return "users"
}

// Business logic methods
func (u *User) IsAdmin() bool {
    return u.Role == "admin"
}

func (u *User) CanDelete() bool {
    return !u.IsAdmin() // Admin tidak bisa dihapus
}
```

### product.go
```go
package models

import (
    "errors"
    "time"
    "github.com/google/uuid"
)

type Product struct {
    ID          uuid.UUID `gorm:"type:uuid;primaryKey" json:"id"`
    Name        string    `gorm:"not null" json:"name"`
    Description string    `gorm:"type:text" json:"description"`
    Price       float64   `gorm:"type:decimal(10,2);not null" json:"price"`
    Stock       int       `gorm:"default:0" json:"stock"`
    CategoryID  uuid.UUID `gorm:"type:uuid" json:"category_id"`
    IsActive    bool      `gorm:"default:true" json:"is_active"`
    CreatedAt   time.Time `gorm:"autoCreateTime" json:"created_at"`
    UpdatedAt   time.Time `gorm:"autoUpdateTime" json:"updated_at"`

    // Relationships
    Category Category `gorm:"foreignKey:CategoryID" json:"category,omitempty"`
}

// Business methods
func (p *Product) IsAvailable() bool {
    return p.IsActive && p.Stock > 0
}

func (p *Product) ReduceStock(quantity int) error {
    if p.Stock < quantity {
        return errors.New("insufficient stock")
    }
    p.Stock -= quantity
    return nil
}
```

### order.go
```go
package models

import (
    "time"
    "github.com/google/uuid"
)

type OrderStatus string

const (
    OrderPending   OrderStatus = "pending"
    OrderPaid      OrderStatus = "paid"
    OrderShipped   OrderStatus = "shipped"
    OrderDelivered OrderStatus = "delivered"
    OrderCancelled OrderStatus = "cancelled"
)

type Order struct {
    ID         uuid.UUID   `gorm:"type:uuid;primaryKey" json:"id"`
    UserID     uuid.UUID   `gorm:"type:uuid;not null" json:"user_id"`
    TotalPrice float64     `gorm:"type:decimal(10,2);not null" json:"total_price"`
    Status     OrderStatus `gorm:"type:varchar(20);default:'pending'" json:"status"`
    CreatedAt  time.Time   `gorm:"autoCreateTime" json:"created_at"`
    UpdatedAt  time.Time   `gorm:"autoUpdateTime" json:"updated_at"`

    // Relationships
    User       User        `gorm:"foreignKey:UserID" json:"user,omitempty"`
    OrderItems []OrderItem `gorm:"foreignKey:OrderID" json:"order_items,omitempty"`
}

type OrderItem struct {
    ID        uuid.UUID `gorm:"type:uuid;primaryKey" json:"id"`
    OrderID   uuid.UUID `gorm:"type:uuid;not null" json:"order_id"`
    ProductID uuid.UUID `gorm:"type:uuid;not null" json:"product_id"`
    Quantity  int       `gorm:"not null" json:"quantity"`
    Price     float64   `gorm:"type:decimal(10,2);not null" json:"price"`

    // Relationships
    Product Product `gorm:"foreignKey:ProductID" json:"product,omitempty"`
}
```

## ⚠️ Best Practices
- ✅ Gunakan UUID untuk primary key (lebih secure)
- ✅ Selalu set `json:"-"` untuk password
- ✅ Gunakan pointer (*) untuk optional fields
- ✅ Buat method untuk business logic di model
- ✅ Gunakan GORM hooks (BeforeCreate, AfterCreate, dll)
- ❌ Jangan taruh database query di models
- ❌ Jangan taruh HTTP logic di models

## 🎓 Tips
- Gunakan `omitempty` untuk relationships agar JSON tidak bloat
- Buat constants untuk enum values (OrderStatus)
- Models hanya tentang data structure, bukan data access
