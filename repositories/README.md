# REPOSITORIES Directory

## 📖 Filosofi
**"Abstract your data access layer"**. Repository pattern memisahkan business logic dari data access logic. Filosofinya adalah membuat abstraction layer antara aplikasi dan database, sehingga mudah untuk switch database atau mock data untuk testing.

## 🎯 Definisi
Directory yang berisi kode untuk berinteraksi dengan database. Repositories bertanggung jawab untuk semua operasi CRUD (Create, Read, Update, Delete) dan query kompleks ke database.

## 💡 Tanggung Jawab
- CRUD operations (Create, Read, Update, Delete)
- Database queries (SELECT, INSERT, UPDATE, DELETE)
- Filtering, sorting, pagination
- Transactions
- Complex joins dan relations
- Query optimization

## 📝 Contoh Struktur
```
repositories/
├── user_repository.go       # User data access
├── product_repository.go    # Product data access
├── order_repository.go      # Order data access
└── README.md                # Dokumentasi ini
```

## 🔧 Contoh Kode

### user_repository.go
```go
package repositories

import (
    "context"
    "myproject/models"
    "gorm.io/gorm"
)

type UserRepository struct {
    db *gorm.DB
}

func NewUserRepository(db *gorm.DB) *UserRepository {
    return &UserRepository{db: db}
}

// Create - Insert new user
func (r *UserRepository) Create(user *models.User) error {
    return r.db.Create(user).Error
}

// GetByID - Get user by ID
func (r *UserRepository) GetByID(id string) (*models.User, error) {
    var user models.User
    err := r.db.First(&user, "id = ?", id).Error
    if err != nil {
        return nil, err
    }
    return &user, nil
}

// GetByEmail - Get user by email
func (r *UserRepository) GetByEmail(email string) (*models.User, error) {
    var user models.User
    err := r.db.Where("email = ?", email).First(&user).Error
    if err != nil {
        if err == gorm.ErrRecordNotFound {
            return nil, nil // Return nil if not found
        }
        return nil, err
    }
    return &user, nil
}

// GetAll - Get all users with filters
func (r *UserRepository) GetAll(filter map[string]interface{}) ([]models.User, error) {
    var users []models.User
    query := r.db

    // Apply filters dynamically
    if role, ok := filter["role"].(string); ok && role != "" {
        query = query.Where("role = ?", role)
    }

    if isActive, ok := filter["is_active"].(bool); ok {
        query = query.Where("is_active = ?", isActive)
    }

    err := query.Find(&users).Error
    return users, err
}

// GetWithPagination - Get users with pagination
func (r *UserRepository) GetWithPagination(page, pageSize int) ([]models.User, int64, error) {
    var users []models.User
    var total int64

    // Count total
    r.db.Model(&models.User{}).Count(&total)

    // Get paginated data
    offset := (page - 1) * pageSize
    err := r.db.Offset(offset).Limit(pageSize).Find(&users).Error

    return users, total, err
}

// Update - Update user
func (r *UserRepository) Update(user *models.User) error {
    return r.db.Save(user).Error
}

// Delete - Delete user
func (r *UserRepository) Delete(id string) error {
    return r.db.Delete(&models.User{}, "id = ?", id).Error
}

// Search - Full text search
func (r *UserRepository) Search(keyword string) ([]models.User, error) {
    var users []models.User
    searchPattern := "%" + keyword + "%"
    err := r.db.Where("name ILIKE ? OR email ILIKE ?", searchPattern, searchPattern).Find(&users).Error
    return users, err
}
```

## ⚠️ Best Practices
- ✅ Always use context for cancellation
- ✅ Use transactions untuk operasi yang saling terkait
- ✅ Handle `gorm.ErrRecordNotFound` explicitly
- ✅ Use prepared statements (avoid SQL injection)
- ✅ Index kolom yang sering di-query
- ✅ Use pagination untuk large datasets
- ❌ Jangan expose repository errors langsung ke client
- ❌ Jangan query di loop (N+1 problem) - use Preload
- ❌ Jangan fetch all data tanpa limit
