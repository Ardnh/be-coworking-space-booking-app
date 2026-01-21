package migrations

import "gorm.io/gorm"

func init() {
	Register(Migration{
		Version: 1,
		Name:    "create_users_table",
		Up: func(db *gorm.DB) error {
			return db.Exec(`
				CREATE TABLE IF NOT EXISTS users(
				    user_id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
				    email VARCHAR(255) UNIQUE NOT NULL,
				    password VARCHAR(255) NOT NULL,
					username VARCHAR(255) NOT NULL,
				    full_name VARCHAR(255) NOT NULL,
				    phone_number VARCHAR(20) NOT NULL,
				    user_type VARCHAR(20) NOT NULL,
				    profile_image VARCHAR(255) NOT NULL,
				    status VARCHAR(20) NOT NULL DEFAULT 'active',
				    CONSTRAINT chk_status CHECK (status IN ('active', 'inactive', 'suspended', 'banned', 'pending')),
				    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
				    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
				    deleted_at TIMESTAMP DEFAULT NULL
				);

				CREATE INDEX idx_users_deleted_at ON users(deleted_at);
				CREATE INDEX idx_users_username ON users(username);
				CREATE INDEX idx_users_email ON users(email);
			`).Error
		},
		Down: func(db *gorm.DB) error {
			return db.Exec("DROP TABLE IF EXISTS users CASCADE;").Error
		},
	})
}
