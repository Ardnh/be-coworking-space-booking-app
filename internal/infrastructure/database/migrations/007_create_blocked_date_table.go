package migrations

import (
	"gorm.io/gorm"
)

func init() {
	Register(Migration{
		Version: 7,
		Name:    "create_blocked_date_table",
		Up: func(db *gorm.DB) error {
			return db.Exec(`
			    CREATE TABLE IF NOT EXISTS blocked_dates (
			        block_id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
			        resource_id UUID NOT NULL,
			        date DATE NOT NULL,
					time_from TIME NOT NULL,
					time_to TIME NOT NULL,
			        reason TEXT NOT NULL,
			        created_by UUID NOT NULL,
			        created_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
			        updated_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
			        deleted_at TIMESTAMP NULL DEFAULT NULL,

			        -- Validasi: reason tidak boleh kosong
			        CONSTRAINT chk_reason_not_empty CHECK (LENGTH(TRIM(reason)) > 0),

			        -- Validasi: date tidak boleh di masa lalu (optional)
			        CONSTRAINT chk_reasonable_start_date CHECK (date >= CURRENT_DATE),

			        -- Foreign Keys
			        FOREIGN KEY (created_by) REFERENCES users(user_id) ON DELETE CASCADE,
			        FOREIGN KEY (resource_id) REFERENCES resources(resource_id) ON DELETE CASCADE
			    );

			    -- Index untuk performance query
			    CREATE INDEX IF NOT EXISTS idx_blocked_dates_resource_id ON blocked_dates(resource_id);
			    CREATE INDEX IF NOT EXISTS idx_blocked_dates_start_date ON blocked_dates(date);
			    CREATE INDEX IF NOT EXISTS idx_blocked_dates_deleted_at ON blocked_dates(deleted_at);
			`).Error
		},
		Down: func(db *gorm.DB) error {
			return db.Exec(`DROP TABLE IF EXISTS blocked_dates CASCADE;`).Error
		},
	})
}
