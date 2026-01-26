package migrations

import "gorm.io/gorm"

func init() {
	Register(Migration{
		Version: 1,
		Name:    "create_booking_history_table",
		Up: func(db *gorm.DB) error {
			return db.Exec(`
			    CREATE TABLE IF NOT EXISTS booking_history (
			        history_id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
			        booking_id UUID NOT NULL,
			        user_id INTEGER NOT NULL,
			        old_status VARCHAR(255) NOT NULL,
			        new_status VARCHAR(255) NOT NULL,
			        changed_by UUID NOT NULL,
			        change_reason TEXT NOT NULL,
			        changed_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
			        metadata JSONB,

			        -- Foreign Key Constraints
			        CONSTRAINT fk_history_booking
			            FOREIGN KEY (booking_id)
			            REFERENCES bookings(booking_id)
			            ON DELETE CASCADE,

			        CONSTRAINT fk_history_user
			            FOREIGN KEY (user_id)
			            REFERENCES users(id)
			            ON DELETE RESTRICT,

			        CONSTRAINT fk_history_changed_by
			            FOREIGN KEY (changed_by)
			            REFERENCES users(id)
			            ON DELETE RESTRICT,

			        -- Check Constraints
			        CONSTRAINT chk_status_values
			            CHECK (
			                old_status IN ('pending', 'confirmed', 'cancelled', 'completed', 'no_show', 'rescheduled') AND
			                new_status IN ('pending', 'confirmed', 'cancelled', 'completed', 'no_show', 'rescheduled')
			            ),

			        CONSTRAINT chk_status_changed
			            CHECK (old_status != new_status),

			        CONSTRAINT chk_change_reason_not_empty
			            CHECK (TRIM(change_reason) != ''),

			        -- Prevent invalid status transitions (opsional, sesuaikan dengan business logic)
			        CONSTRAINT chk_valid_transitions
			            CHECK (
			                -- Completed/cancelled tidak bisa diubah lagi
			                (old_status NOT IN ('completed', 'cancelled')) OR
			                -- Kecuali edge case tertentu yang Anda izinkan
			                (old_status = 'cancelled' AND new_status = 'pending')
			            )
			    );

			    -- Indexes untuk performa query
			    CREATE INDEX IF NOT EXISTS idx_booking_history_booking_id
			        ON booking_history(booking_id, changed_at DESC);

			    CREATE INDEX IF NOT EXISTS idx_booking_history_user_id
			        ON booking_history(user_id, changed_at DESC);

			    CREATE INDEX IF NOT EXISTS idx_booking_history_changed_by
			        ON booking_history(changed_by, changed_at DESC);

			    CREATE INDEX IF NOT EXISTS idx_booking_history_changed_at
			        ON booking_history(changed_at DESC);

			    CREATE INDEX IF NOT EXISTS idx_booking_history_status_transition
			        ON booking_history(old_status, new_status);

			    -- GIN index untuk metadata JSONB (jika perlu search dalam metadata)
			    CREATE INDEX IF NOT EXISTS idx_booking_history_metadata
			        ON booking_history USING GIN (metadata);

			    -- Composite index untuk audit queries
			    CREATE INDEX IF NOT EXISTS idx_booking_history_audit
			        ON booking_history(booking_id, changed_at DESC, new_status);
			`).Error
		},
		Down: func(db *gorm.DB) error {
			return db.Exec("DROP TABLE IF EXISTS booking_history").Error
		},
	})
}
