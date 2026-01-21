# UTILS Directory

## 📖 Filosofi
**"DRY - Don't Repeat Yourself"**. Filosofi utils adalah menyimpan fungsi-fungsi helper yang reusable di satu tempat, sehingga tidak perlu menulis kode yang sama berulang kali di berbagai tempat.

## 🎯 Definisi
Directory yang berisi utility functions dan helper functions yang bisa digunakan di berbagai bagian aplikasi. Functions di sini bersifat generic dan tidak terikat pada business logic tertentu.

## 💡 Tanggung Jawab
- Response formatting
- JWT token generation & validation
- Password hashing & verification
- Data validation
- String manipulation
- Date/time utilities
- File operations
- Random generators

## 📝 Contoh Struktur
```
utils/
├── response.go       # HTTP response helpers
├── jwt.go            # JWT token utilities
├── validator.go      # Custom validators
├── password.go       # Password utilities
└── README.md         # Dokumentasi ini
```

## 🔧 Contoh Kode

### response.go
```go
package utils

import "github.com/gofiber/fiber/v2"

type Response struct {
    Status  string      `json:"status"`
    Message string      `json:"message"`
    Data    interface{} `json:"data,omitempty"`
}

type ErrorDetail struct {
    Field   string `json:"field"`
    Message string `json:"message"`
}

// SuccessResponse - Standard success response
func SuccessResponse(c *fiber.Ctx, code int, message string, data interface{}) error {
    return c.Status(code).JSON(Response{
        Status:  "success",
        Message: message,
        Data:    data,
    })
}

// ErrorResponse - Standard error response
func ErrorResponse(c *fiber.Ctx, code int, message string) error {
    return c.Status(code).JSON(Response{
        Status:  "error",
        Message: message,
    })
}

// ValidationErrorResponse - For validation errors
func ValidationErrorResponse(c *fiber.Ctx, errors interface{}) error {
    return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
        "status":  "error",
        "message": "Validation failed",
        "errors":  errors,
    })
}

// PaginatedResponse - For paginated data
func PaginatedResponse(c *fiber.Ctx, data interface{}, page, pageSize int, total int64) error {
    totalPages := int(total) / pageSize
    if int(total)%pageSize != 0 {
        totalPages++
    }

    return c.Status(fiber.StatusOK).JSON(fiber.Map{
        "status":  "success",
        "message": "Data fetched successfully",
        "data":    data,
        "pagination": fiber.Map{
            "page":        page,
            "page_size":   pageSize,
            "total_items": total,
            "total_pages": totalPages,
        },
    })
}
```

### jwt.go
```go
package utils

import (
    "errors"
    "time"
    "github.com/golang-jwt/jwt/v5"
)

var jwtSecret = []byte("your-secret-key") // Ambil dari config

type Claims struct {
    UserID string `json:"user_id"`
    Email  string `json:"email"`
    Role   string `json:"role"`
    jwt.RegisteredClaims
}

// GenerateJWT - Create new JWT token
func GenerateJWT(userID, email, role string) (string, error) {
    claims := Claims{
        UserID: userID,
        Email:  email,
        Role:   role,
        RegisteredClaims: jwt.RegisteredClaims{
            ExpiresAt: jwt.NewNumericDate(time.Now().Add(24 * time.Hour)), // 24 jam
            IssuedAt:  jwt.NewNumericDate(time.Now()),
            NotBefore: jwt.NewNumericDate(time.Now()),
        },
    }

    token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
    return token.SignedString(jwtSecret)
}

// ValidateJWT - Validate and parse JWT token
func ValidateJWT(tokenString string) (*Claims, error) {
    token, err := jwt.ParseWithClaims(tokenString, &Claims{}, func(token *jwt.Token) (interface{}, error) {
        return jwtSecret, nil
    })

    if err != nil {
        return nil, err
    }

    if claims, ok := token.Claims.(*Claims); ok && token.Valid {
        return claims, nil
    }

    return nil, errors.New("invalid token")
}
```

### validator.go
```go
package utils

import (
    "github.com/go-playground/validator/v10"
)

var validate = validator.New()

// ValidateStruct - Validate struct using tags
func ValidateStruct(s interface{}) error {
    return validate.Struct(s)
}

// GetValidationErrors - Format validation errors
func GetValidationErrors(err error) []ErrorDetail {
    var errors []ErrorDetail

    if validationErrors, ok := err.(validator.ValidationErrors); ok {
        for _, e := range validationErrors {
            errors = append(errors, ErrorDetail{
                Field:   e.Field(),
                Message: getErrorMessage(e),
            })
        }
    }

    return errors
}

func getErrorMessage(e validator.FieldError) string {
    switch e.Tag() {
    case "required":
        return e.Field() + " is required"
    case "email":
        return e.Field() + " must be a valid email"
    case "min":
        return e.Field() + " must be at least " + e.Param() + " characters"
    case "max":
        return e.Field() + " must be at most " + e.Param() + " characters"
    default:
        return e.Field() + " is invalid"
    }
}
```

### password.go
```go
package utils

import (
    "golang.org/x/crypto/bcrypt"
)

// HashPassword - Hash password using bcrypt
func HashPassword(password string) (string, error) {
    bytes, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
    return string(bytes), err
}

// CheckPassword - Verify password against hash
func CheckPassword(password, hash string) bool {
    err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
    return err == nil
}
```

## ⚠️ Best Practices
- ✅ Keep utils functions pure (no side effects)
- ✅ Make functions reusable and generic
- ✅ Add error handling
- ✅ Document functions with comments
- ✅ Use descriptive function names
- ❌ Don't put business logic in utils
- ❌ Don't access database in utils
