package migrations

import "gorm.io/gorm"

func init() {
	Register(Migration{
		Version: 5,
		Name:    "create_pricing_rules_table",
		Up: func(db *gorm.DB) error {
			return db.Exec(`
		        CREATE TABLE IF NOT EXISTS pricing_rules (
		            rule_id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
		            resource_id UUID NOT NULL,
		            day_of_week SMALLINT NOT NULL,
		            time_start TIME NOT NULL,
		            time_end TIME NOT NULL,
		            price_multiplier DECIMAL(5, 2),
		            fixed_price DECIMAL(10, 2),
		            is_active BOOLEAN NOT NULL DEFAULT true,
		            start_date DATE,
		            end_date DATE,
		            created_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,

		            CONSTRAINT chk_day_of_week CHECK (day_of_week BETWEEN 0 AND 6),
		            CONSTRAINT chk_time_range CHECK (time_end > time_start),
		            CONSTRAINT chk_price CHECK (price_multiplier IS NOT NULL OR fixed_price IS NOT NULL),

					-- Foreign Keys
		            FOREIGN KEY (resource_id) REFERENCES resources(resource_id)
		        );
		        CREATE INDEX idx_pricing_rules_resource_id ON pricing_rules(resource_id);
		        CREATE INDEX idx_pricing_rules_day_of_week ON pricing_rules(day_of_week);
		    `).Error
		},
		Down: func(db *gorm.DB) error {
			return db.Exec("DROP TABLE pricing_rules CASCADE;").Error
		},
	})
}
