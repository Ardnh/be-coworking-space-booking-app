package migrations

import "gorm.io/gorm"

func init() {
	Register(Migration{
		Version: 2,
		Name:    "create_vendor_table",
		Up: func(db *gorm.DB) error {
			return db.Exec(`
				CREATE TABLE vendors (
					vendor_id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
					owner_user_id UUID NOT NULL,
					vendor_name VARCHAR(255) NOT NULL,
					address VARCHAR(255) NOT NULL,
					city VARCHAR(255) NOT NULL,
					phone_number VARCHAR(20) NOT NULL,
					email VARCHAR(255) UNIQUE NOT NULL,
					description TEXT NOT NULL,
					rating DECIMAL(10,2) NOT NULL DEFAULT 0,
					status VARCHAR(20) NOT NULL DEFAULT 'active',
					CONSTRAINT chk_status CHECK (status IN ('active', 'inactive')),
					created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
					updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
					deleted_at TIMESTAMP DEFAULT NULL
				);

				CREATE INDEX idx_vendors_deleted_at ON vendors(deleted_at);
				CREATE INDEX idx_vendors_name ON vendors(vendor_name);
				CREATE INDEX idx_vendors_email ON vendors(email);
			`).Error
		},
		Down: func(db *gorm.DB) error {
			return db.Exec("DROP TABLE IF EXISTS users CASCADE;").Error
		},
	})
}
