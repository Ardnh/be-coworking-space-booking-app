package migrations

import "gorm.io/gorm"

func init() {
	Register(Migration{
		Version: 1,
		Name:    "create_payment_table",
		Up: func(db *gorm.DB) error {
			return db.Exec(`
			    CREATE TABLE IF NOT EXISTS payments (
			        payment_id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
			        booking_id UUID NOT NULL,
			        user_id INTEGER NOT NULL REFERENCES users(id) ON DELETE RESTRICT,
			        amount DECIMAL(10, 2) NOT NULL,
			        payment_method VARCHAR(50) NOT NULL,
			        payment_date TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
			        payment_status VARCHAR(20) NOT NULL,
			        notes TEXT,
			        created_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
			        updated_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,

			        -- Foreign Key Constraints
			        CONSTRAINT fk_booking
			            FOREIGN KEY (booking_id)
			            REFERENCES bookings(booking_id)
			            ON DELETE CASCADE,

			        -- Check Constraints
			        CONSTRAINT chk_amount_positive
			            CHECK (amount > 0),

			        CONSTRAINT chk_payment_status
			            CHECK (payment_status IN ('pending', 'completed', 'failed', 'refunded', 'cancelled')),

			        CONSTRAINT chk_payment_method
			            CHECK (payment_method IN ('credit_card', 'debit_card', 'bank_transfer', 'ewallet', 'cash')),

			        CONSTRAINT chk_dates
			            CHECK (updated_at >= created_at)
			    );

			    -- Indexes untuk performa query
			    CREATE INDEX IF NOT EXISTS idx_payments_booking_id ON payments(booking_id);
			    CREATE INDEX IF NOT EXISTS idx_payments_user_id ON payments(user_id);
			    CREATE INDEX IF NOT EXISTS idx_payments_status ON payments(payment_status);
			    CREATE INDEX IF NOT EXISTS idx_payments_date ON payments(payment_date DESC);
			    CREATE INDEX IF NOT EXISTS idx_payments_created_at ON payments(created_at DESC);
			`).Error
		},
		Down: func(db *gorm.DB) error {
			return db.Exec("DROP TABLE IF EXISTS payments").Error
		},
	})
}
