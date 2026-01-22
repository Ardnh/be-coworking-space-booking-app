package migrations

import "gorm.io/gorm"

func init() {
	Register(Migration{
		Version: 1,
		Name:    "create_category_table",
		Up: func(db *gorm.DB) error {
			return db.Exec(`
				CREATE TABLE IF NOT EXISTS categories (
					category_id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
					category_name VARCHAR(255) NOT NULL,
					description TEXT,
					parent_category_id UUID REFERENCES categories(category_id),
					icon VARCHAR(255),
					created_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
					updated_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
					deleted_at TIMESTAMP DEFAULT NULL
				);

				CREATE INDEX idx_categories_parent_category_id ON categories(parent_category_id);
				CREATE INDEX idx_categories_category_name ON categories(category_name);
			`).Error
		},
		Down: func(db *gorm.DB) error {
			return db.Exec("DROP TABLE IF EXISTS categories CASCADE;").Error
		},
	})
}
