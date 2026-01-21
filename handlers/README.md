# HANDLERS Directory

## 📖 Filosofi
**"Handlers are the gateway between HTTP and your business logic"**. Filosofinya adalah memisahkan HTTP concerns (request/response handling) dari business logic. Handlers bertanggung jawab menerima HTTP request, validasi input, memanggil repository, dan mengembalikan HTTP response.

## 🎯 Definisi
Directory yang berisi HTTP handlers (atau controllers dalam MVC pattern). Handlers menerima HTTP requests, memproses data, berinteraksi dengan repositories, dan mengembalikan HTTP responses menggunakan Fiber.

## 💡 Tanggung Jawab
- Parse HTTP request (JSON body, query params, URL params)
- Validasi input
- Call repository methods (business logic)
- Handle errors
- Format & return HTTP response
- Set appropriate HTTP status codes

## 📝 Contoh Struktur
```
handlers/
├── user_handler.go      # User CRUD operations
├── product_handler.go   # Product operations
├── auth_handler.go      # Authentication (login, register)
└── README.md            # Dokumentasi ini
```

## 🔧 Contoh Kode

### user_handler.go
```go
package handlers

import (
    "myproject/models"
    "myproject/repositories"
    "myproject/utils"
    "github.com/gofiber/fiber/v2"
    "github.com/google/uuid"
    "golang.org/x/crypto/bcrypt"
)

type UserHandler struct {
    repo *repositories.UserRepository
}

func NewUserHandler(repo *repositories.UserRepository) *UserHandler {
    return &UserHandler{repo: repo}
}

// Create - POST /users
func (h *UserHandler) Create(c *fiber.Ctx) error {
    // 1. Parse request body
    var input struct {
        Email    string `json:"email" validate:"required,email"`
        Name     string `json:"name" validate:"required"`
        Password string `json:"password" validate:"required,min=6"`
    }

    if err := c.BodyParser(&input); err != nil {
        return utils.ErrorResponse(c, fiber.StatusBadRequest, "Invalid input: "+err.Error())
    }

    // 2. Validate input
    if err := utils.ValidateStruct(input); err != nil {
        return utils.ValidationErrorResponse(c, err)
    }

    // 3. Check if email exists
    existing, _ := h.repo.GetByEmail(input.Email)
    if existing != nil {
        return utils.ErrorResponse(c, fiber.StatusConflict, "Email already exists")
    }

    // 4. Hash password
    hashedPassword, err := bcrypt.GenerateFromPassword([]byte(input.Password), bcrypt.DefaultCost)
    if err != nil {
        return utils.ErrorResponse(c, fiber.StatusInternalServerError, "Failed to process password")
    }

    // 5. Create user
    user := models.User{
        ID:       uuid.New(),
        Email:    input.Email,
        Name:     input.Name,
        Password: string(hashedPassword),
        Role:     "user",
        IsActive: true,
    }

    if err := h.repo.Create(&user); err != nil {
        return utils.ErrorResponse(c, fiber.StatusInternalServerError, "Failed to create user")
    }

    // 6. Return response (hide password)
    user.Password = ""
    return utils.SuccessResponse(c, fiber.StatusCreated, "User created successfully", user)
}

// GetByID - GET /users/:id
func (h *UserHandler) GetByID(c *fiber.Ctx) error {
    // 1. Get ID from URL param
    id := c.Params("id")

    // 2. Validate UUID
    if _, err := uuid.Parse(id); err != nil {
        return utils.ErrorResponse(c, fiber.StatusBadRequest, "Invalid user ID")
    }

    // 3. Get user from database
    user, err := h.repo.GetByID(id)
    if err != nil {
        return utils.ErrorResponse(c, fiber.StatusNotFound, "User not found")
    }

    // 4. Hide password
    user.Password = ""

    // 5. Return response
    return utils.SuccessResponse(c, fiber.StatusOK, "User fetched successfully", user)
}

// GetAll - GET /users
func (h *UserHandler) GetAll(c *fiber.Ctx) error {
    // 1. Parse query parameters
    role := c.Query("role")        // ?role=admin
    isActive := c.Query("is_active") // ?is_active=true

    // 2. Build filter
    filter := make(map[string]interface{})
    if role != "" {
        filter["role"] = role
    }
    if isActive == "true" {
        filter["is_active"] = true
    } else if isActive == "false" {
        filter["is_active"] = false
    }

    // 3. Get users
    users, err := h.repo.GetAll(filter)
    if err != nil {
        return utils.ErrorResponse(c, fiber.StatusInternalServerError, "Failed to fetch users")
    }

    // 4. Hide passwords
    for i := range users {
        users[i].Password = ""
    }

    // 5. Return response
    return utils.SuccessResponse(c, fiber.StatusOK, "Users fetched successfully", fiber.Map{
        "users": users,
        "total": len(users),
    })
}

// Update - PUT /users/:id
func (h *UserHandler) Update(c *fiber.Ctx) error {
    // 1. Get ID
    id := c.Params("id")

    // 2. Check if user exists
    user, err := h.repo.GetByID(id)
    if err != nil {
        return utils.ErrorResponse(c, fiber.StatusNotFound, "User not found")
    }

    // 3. Parse update data
    var input struct {
        Name     string `json:"name"`
        IsActive *bool  `json:"is_active"` // Pointer untuk optional
    }

    if err := c.BodyParser(&input); err != nil {
        return utils.ErrorResponse(c, fiber.StatusBadRequest, "Invalid input")
    }

    // 4. Update fields
    if input.Name != "" {
        user.Name = input.Name
    }
    if input.IsActive != nil {
        user.IsActive = *input.IsActive
    }

    // 5. Save to database
    if err := h.repo.Update(user); err != nil {
        return utils.ErrorResponse(c, fiber.StatusInternalServerError, "Failed to update user")
    }

    // 6. Return response
    user.Password = ""
    return utils.SuccessResponse(c, fiber.StatusOK, "User updated successfully", user)
}

// Delete - DELETE /users/:id
func (h *UserHandler) Delete(c *fiber.Ctx) error {
    // 1. Get ID
    id := c.Params("id")

    // 2. Check if user exists
    user, err := h.repo.GetByID(id)
    if err != nil {
        return utils.ErrorResponse(c, fiber.StatusNotFound, "User not found")
    }

    // 3. Business rule: Admin tidak bisa dihapus
    if user.IsAdmin() {
        return utils.ErrorResponse(c, fiber.StatusForbidden, "Cannot delete admin user")
    }

    // 4. Delete user
    if err := h.repo.Delete(id); err != nil {
        return utils.ErrorResponse(c, fiber.StatusInternalServerError, "Failed to delete user")
    }

    // 5. Return response
    return utils.SuccessResponse(c, fiber.StatusOK, "User deleted successfully", nil)
}
```

### auth_handler.go
```go
package handlers

import (
    "myproject/models"
    "myproject/repositories"
    "myproject/utils"
    "github.com/gofiber/fiber/v2"
    "golang.org/x/crypto/bcrypt"
)

type AuthHandler struct {
    userRepo *repositories.UserRepository
}

func NewAuthHandler(userRepo *repositories.UserRepository) *AuthHandler {
    return &AuthHandler{userRepo: userRepo}
}

// Login - POST /auth/login
func (h *AuthHandler) Login(c *fiber.Ctx) error {
    var input struct {
        Email    string `json:"email" validate:"required,email"`
        Password string `json:"password" validate:"required"`
    }

    if err := c.BodyParser(&input); err != nil {
        return utils.ErrorResponse(c, fiber.StatusBadRequest, "Invalid input")
    }

    // Validate
    if err := utils.ValidateStruct(input); err != nil {
        return utils.ValidationErrorResponse(c, err)
    }

    // Get user by email
    user, err := h.userRepo.GetByEmail(input.Email)
    if err != nil {
        return utils.ErrorResponse(c, fiber.StatusUnauthorized, "Invalid credentials")
    }

    // Check password
    if err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(input.Password)); err != nil {
        return utils.ErrorResponse(c, fiber.StatusUnauthorized, "Invalid credentials")
    }

    // Check if active
    if !user.IsActive {
        return utils.ErrorResponse(c, fiber.StatusForbidden, "Account is inactive")
    }

    // Generate JWT token
    token, err := utils.GenerateJWT(user.ID.String(), user.Email, user.Role)
    if err != nil {
        return utils.ErrorResponse(c, fiber.StatusInternalServerError, "Failed to generate token")
    }

    return utils.SuccessResponse(c, fiber.StatusOK, "Login successful", fiber.Map{
        "token": token,
        "user": fiber.Map{
            "id":    user.ID,
            "email": user.Email,
            "name":  user.Name,
            "role":  user.Role,
        },
    })
}
```

## 🎯 HTTP Status Codes

| Status Code | Kapan Digunakan |
|------------|-----------------|
| 200 OK | Success GET, PUT, PATCH |
| 201 Created | Success POST (create new resource) |
| 204 No Content | Success DELETE |
| 400 Bad Request | Invalid input, validation error |
| 401 Unauthorized | Missing/invalid token |
| 403 Forbidden | Valid token, tapi tidak punya akses |
| 404 Not Found | Resource tidak ditemukan |
| 409 Conflict | Duplicate data (email exists) |
| 500 Internal Server Error | Server error |

## ⚠️ Best Practices
- ✅ Gunakan `c.BodyParser()` untuk parse JSON body
- ✅ Always validate input dengan validator
- ✅ Return appropriate HTTP status codes
- ✅ Hide sensitive data (password) dari response
- ✅ Use consistent response format
- ✅ Handle all possible errors
- ✅ Return error untuk propagate ke error handler
- ❌ Jangan expose internal error messages ke client

## 🎓 Tips Fiber
- `c.Params("id")` untuk URL parameters
- `c.Query("key")` untuk query strings
- `c.BodyParser(&struct)` untuk JSON body
- `return c.Status(code).JSON(data)` untuk response
- Return error akan di-handle oleh error handler middleware
