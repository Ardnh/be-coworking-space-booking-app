# My Go Project

Clean Architecture Simple - Fast Development Structure with Fiber

## 📁 Project Structure
```
.
├── cmd/                # Entry point
├── config/             # Configuration
├── models/             # Database models
├── handlers/           # HTTP handlers
├── repositories/       # Data access layer
├── middlewares/        # HTTP middlewares
├── utils/              # Helper functions
└── routes/             # Route definitions
```

## 🚀 Quick Start

1. **Clone & Setup**
```bash
   cd myproject
   cp .env.example .env
   # Edit .env with your configuration
```

2. **Install Dependencies**
```bash
   go mod download
```

3. **Run**
```bash
   go run cmd/main.go
```

4. **Test**
```bash
   curl http://localhost:3000/health
```

## 📖 Documentation

Each folder contains a README.md with:
- Philosophy
- Definition
- Examples
- Best practices

Start reading from:
1. `cmd/README.md` - Entry point
2. `models/README.md` - Data structures
3. `handlers/README.md` - Business logic
4. `routes/README.md` - API endpoints

## 🛠️ Development

### Add New Feature

1. Create model in `models/`
2. Create repository in `repositories/`
3. Create handler in `handlers/`
4. Register routes in `routes/`

## 📚 Tech Stack

- **Framework**: Fiber (Fast Express-like framework)
- **ORM**: GORM
- **Database**: PostgreSQL
- **Auth**: JWT
- **Validation**: go-playground/validator

## 🎯 Why Fiber?

- ⚡ Extremely fast (built on fasthttp)
- 🎨 Express.js-like syntax
- 🔧 Zero memory allocation router
- 💪 Robust middleware support
- 📝 Great documentation

## 🤝 Contributing

1. Read documentation in each folder
2. Follow existing patterns
3. Keep it simple

## 📝 License

MIT
