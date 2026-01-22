package migrations

import (
	"gorm.io/gorm"
)

func init() {
	Register(Migration{
		Version: 1,
		Name:    "create_blocked_date_table",
		Up: func(db *gorm.DB) error {
			return db.Exec(`
			    CREATE TABLE IF NOT EXISTS blocked_dates (
			        block_id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
			        resource_id UUID NOT NULL,
			        start_date DATE NOT NULL,
			        end_date DATE NOT NULL,
			        reason TEXT NOT NULL,
			        created_by UUID NOT NULL,
			        created_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
			        updated_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
			        deleted_at TIMESTAMP NULL DEFAULT NULL,

			        -- Validasi: end_date harus >= start_date
			        CONSTRAINT chk_date_range CHECK (end_date >= start_date),

			        -- Validasi: reason tidak boleh kosong
			        CONSTRAINT chk_reason_not_empty CHECK (LENGTH(TRIM(reason)) > 0),

			        -- Validasi: start_date tidak boleh terlalu jauh di masa lalu (optional)
			        CONSTRAINT chk_reasonable_start_date CHECK (start_date >= CURRENT_DATE - INTERVAL '30 days'),

			        -- Foreign Keys
			        FOREIGN KEY (created_by) REFERENCES users(user_id) ON DELETE CASCADE,
			        FOREIGN KEY (resource_id) REFERENCES resources(resource_id) ON DELETE CASCADE
			    );

			    -- Index untuk performance query
			    CREATE INDEX IF NOT EXISTS idx_blocked_dates_resource_id ON blocked_dates(resource_id);
			    CREATE INDEX IF NOT EXISTS idx_blocked_dates_start_date ON blocked_dates(start_date);
			    CREATE INDEX IF NOT EXISTS idx_blocked_dates_end_date ON blocked_dates(end_date);
			    CREATE INDEX IF NOT EXISTS idx_blocked_dates_date_range ON blocked_dates(resource_id, start_date, end_date);
			    CREATE INDEX IF NOT EXISTS idx_blocked_dates_deleted_at ON blocked_dates(deleted_at);
			`).Error
		},
		Down: func(db *gorm.DB) error {
			return db.Exec(`DROP TABLE IF EXISTS blocked_dates CASCADE;`).Error
		},
	})
}
