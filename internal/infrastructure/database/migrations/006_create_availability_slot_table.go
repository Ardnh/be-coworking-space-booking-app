package migrations

import (
	"gorm.io/gorm"
)

func init() {
	Register(Migration{
		Version: 6,
		Name:    "006_create_availability_slot_table",
		Up: func(db *gorm.DB) error {
			return db.Exec(`
				CREATE TABLE IF NOT EXISTS availability_slots (
				    slot_id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
				    resource_id UUID NOT NULL,
				    date DATE NOT NULL,
				    start_time TIMESTAMP NOT NULL,
				    end_time TIMESTAMP NOT NULL,
				    max_capacity INT NOT NULL,
				    booked_capacity INT NOT NULL DEFAULT 0,
				    is_available BOOLEAN NOT NULL DEFAULT true,
				    price_override DECIMAL(10, 2) DEFAULT 0,
				    created_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
				    updated_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,

				    -- Validasi waktu: end_time harus lebih besar dari start_time
				    CONSTRAINT chk_time_range CHECK (end_time > start_time),

				    -- Validasi capacity: max_capacity harus positif
				    CONSTRAINT chk_max_capacity CHECK (max_capacity > 0),

				    -- Validasi booked_capacity: tidak boleh negatif
				    CONSTRAINT chk_booked_capacity_positive CHECK (booked_capacity >= 0),

				    -- Validasi booked_capacity tidak boleh melebihi max_capacity
				    CONSTRAINT chk_booked_not_exceed_max CHECK (booked_capacity <= max_capacity),

				    -- Validasi price_override tidak boleh negatif
				    CONSTRAINT chk_price_override CHECK (price_override >= 0),

				    -- Validasi date tidak boleh masa lalu
				    CONSTRAINT chk_future_date CHECK (date >= CURRENT_DATE),

				    -- Mencegah duplicate slot untuk resource yang sama di waktu yang sama
				    CONSTRAINT uq_resource_slot UNIQUE (resource_id, date, start_time, end_time),

				    -- Foreign key dengan proper reference
				    FOREIGN KEY (resource_id) REFERENCES resources(resource_id) ON DELETE CASCADE
				);

				-- Index untuk performance query
				CREATE INDEX IF NOT EXISTS idx_availability_slots_resource_id ON availability_slots(resource_id);
				CREATE INDEX IF NOT EXISTS idx_availability_slots_date ON availability_slots(date);
				CREATE INDEX IF NOT EXISTS idx_availability_slots_resource_date ON availability_slots(resource_id, date);
				CREATE INDEX IF NOT EXISTS idx_availability_slots_is_available ON availability_slots(is_available);
			`).Error
		},
		Down: func(db *gorm.DB) error {
			return db.Exec("DROP TABLE IF EXISTS availability_slots CASCADE;").Error
		},
	})
}
