package migrations

import "gorm.io/gorm"

func init() {
	Register(Migration{
		Version: 8,
		Name:    "create_bookings_table",
		Up: func(db *gorm.DB) error {
			return db.Exec(`
			    CREATE TABLE IF NOT EXISTS bookings (
			        booking_id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
			        user_id UUID NOT NULL,
			        resource_id UUID NOT NULL,
					seats INT NOT NULL,
			        total_price DECIMAL(10, 2) NOT NULL,
			        booking_status VARCHAR(20) NOT NULL DEFAULT 'pending',
			        payment_status VARCHAR(20) NOT NULL DEFAULT 'pending',
			        booking_code VARCHAR(20) NOT NULL UNIQUE,
			        created_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
			        updated_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,

			        -- Validasi: booking_status hanya boleh nilai tertentu
			        CONSTRAINT chk_booking_status CHECK (
			            booking_status IN ('pending', 'confirmed', 'cancelled', 'completed', 'rejected')
			        ),

			        -- Validasi: payment_status hanya boleh nilai tertentu
			        CONSTRAINT chk_payment_status CHECK (
			            payment_status IN ('pending', 'paid', 'failed', 'refunded', 'cancelled')
			        ),

			        -- Validasi: total_price harus positif
			        CONSTRAINT chk_positive_price CHECK (total_price > 0),

			        -- Validasi: booking_code format (contoh: BKG-XXXXXX)
			        CONSTRAINT chk_booking_code_format CHECK (
			            booking_code ~ '^[A-Z0-9]{3,20}$'
			        ),

			        -- Foreign Keys
			        FOREIGN KEY (user_id) REFERENCES users(user_id) ON DELETE CASCADE,
			        FOREIGN KEY (resource_id) REFERENCES resources(resource_id) ON DELETE RESTRICT
			    );

			    -- Index untuk performance
			    CREATE INDEX IF NOT EXISTS idx_bookings_user_id ON bookings(user_id);
			    CREATE INDEX IF NOT EXISTS idx_bookings_resource_id ON bookings(resource_id);
			    CREATE INDEX IF NOT EXISTS idx_bookings_booking_status ON bookings(booking_status);
			    CREATE INDEX IF NOT EXISTS idx_bookings_payment_status ON bookings(payment_status);
			    CREATE INDEX IF NOT EXISTS idx_bookings_booking_code ON bookings(booking_code);
			    CREATE INDEX IF NOT EXISTS idx_bookings_created_at ON bookings(created_at);
			`).Error
		},
		Down: func(db *gorm.DB) error {
			return db.Exec(`DROP TABLE IF EXISTS bookings CASCADE;`).Error
		},
	})
}
