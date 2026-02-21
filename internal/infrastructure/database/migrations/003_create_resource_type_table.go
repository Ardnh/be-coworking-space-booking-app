package migrations

import "gorm.io/gorm"

func init() {
	Register(Migration{
		Version: 3,
		Name:    "create_resource_type_table",
		Up: func(db *gorm.DB) error {
			return db.Exec(`
				CREATE TABLE IF NOT EXISTS resource_type (
					resource_type_id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
					resource_type_name VARCHAR(255) NOT NULL,
					description TEXT,
					parent_resource_type_id UUID REFERENCES resource_type(resource_type_id),
					icon VARCHAR(255),
					created_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
					updated_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
					deleted_at TIMESTAMP DEFAULT NULL
				);

				CREATE INDEX idx_resources_parent_resource_id ON resource_type(parent_resource_type_id);
				CREATE INDEX idx_resources_resource_name ON resource_type(resource_type_name);
			`).Error
		},
		Down: func(db *gorm.DB) error {
			return db.Exec("DROP TABLE IF EXISTS resource_type CASCADE;").Error
		},
	})
}
