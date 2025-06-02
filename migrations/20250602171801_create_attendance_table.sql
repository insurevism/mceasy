-- +goose Up
-- +goose StatementBegin
CREATE TABLE attendance (
    id BIGINT UNSIGNED AUTO_INCREMENT PRIMARY KEY,
    employee_id BIGINT UNSIGNED NOT NULL,
    attendance_date DATE NOT NULL,
    check_in_time TIME NULL COMMENT 'Actual check-in time',
    check_out_time TIME NULL COMMENT 'Actual check-out time',
    status ENUM('present', 'absent', 'late', 'half_day') NOT NULL DEFAULT 'absent',
    is_weekend BOOLEAN NOT NULL DEFAULT FALSE COMMENT 'True for Saturday and Sunday',
    notes TEXT NULL COMMENT 'Additional notes for attendance',
    marked_by_admin BOOLEAN NOT NULL DEFAULT FALSE COMMENT 'True if manually marked by admin',
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    modified_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
    deleted_at TIMESTAMP NULL,
    
    FOREIGN KEY (employee_id) REFERENCES employees(id) ON DELETE CASCADE,
    UNIQUE KEY unique_employee_date (employee_id, attendance_date),
    INDEX idx_attendance_date (attendance_date),
    INDEX idx_employee_id (employee_id),
    INDEX idx_status (status),
    INDEX idx_is_weekend (is_weekend),
    INDEX idx_deleted_at (deleted_at)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci;
-- +goose StatementEnd

-- +goose StatementBegin
CREATE EVENT auto_mark_absent
ON SCHEDULE EVERY 1 DAY
STARTS '2025-06-03 09:00:00'
DO
BEGIN
    -- Mark employees as absent if they haven't checked in by 9:00 AM on weekdays
    INSERT INTO attendance (employee_id, attendance_date, status, is_weekend, marked_by_admin)
    SELECT 
        e.id,
        CURDATE(),
        'absent',
        CASE WHEN DAYOFWEEK(CURDATE()) IN (1, 7) THEN TRUE ELSE FALSE END,
        TRUE
    FROM employees e
    WHERE e.is_active = TRUE
    AND NOT EXISTS (
        SELECT 1 FROM attendance a 
        WHERE a.employee_id = e.id 
        AND a.attendance_date = CURDATE()
    )
    AND DAYOFWEEK(CURDATE()) NOT IN (1, 7) -- Exclude Sunday (1) and Saturday (7)
    AND TIME(NOW()) >= '09:00:00';
END;
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP EVENT IF EXISTS auto_mark_absent;
-- +goose StatementEnd

-- +goose StatementBegin
DROP TABLE attendance;
-- +goose StatementEnd
