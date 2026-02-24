package migrations

import "gorm.io/gorm"

func init() {
	Register(Migration{
		Version: 4,
		Name:    "create_resources_table",
		Up: func(db *gorm.DB) error {
			return db.Exec(`
			    CREATE TABLE IF NOT EXISTS resources (
			        resource_id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
			        vendor_id UUID NOT NULL,
			        resource_name VARCHAR(255) NOT NULL,
			        resource_type_id UUID NOT NULL,
			        description TEXT,
			        capacity INT NOT NULL,
					operation_time_from TIME NOT NULL,
					operation_time_to TIME NOT NULL,
					end_date DATE,
			        price_per_unit DECIMAL(10, 2) NOT NULL,
			        images JSONB,
			        location VARCHAR(255),
			        status VARCHAR(20) NOT NULL DEFAULT 'active',
			        CONSTRAINT chk_status CHECK (status IN ('active', 'inactive')),
			        created_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
			        updated_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
			        deleted_at TIMESTAMP DEFAULT NULL,
			        FOREIGN KEY (vendor_id) REFERENCES vendors(vendor_id),
			        FOREIGN KEY (resource_type_id) REFERENCES resource_type(resource_type_id)
			    );
			    CREATE INDEX IF NOT EXISTS idx_resources_vendor_id ON resources(vendor_id);
			    CREATE INDEX IF NOT EXISTS idx_resources_resource_type_id ON resources(resource_type_id);
			    CREATE INDEX IF NOT EXISTS idx_resources_resource_name ON resources(resource_name);
			    CREATE INDEX IF NOT EXISTS idx_resources_status ON resources(status);
			`).Error
		},
		Down: func(db *gorm.DB) error {
			return db.Exec("DROP TABLE IF EXISTS resources CASCADE;").Error
		},
	})
}
