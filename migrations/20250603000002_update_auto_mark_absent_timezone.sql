-- +goose Up
-- +goose StatementBegin
-- Drop the existing event
DROP EVENT IF EXISTS auto_mark_absent;
-- +goose StatementEnd

-- +goose StatementBegin
-- Create updated event with timezone support for Indonesia (UTC+7)
-- For testing: runs every 1 minute, for production: change to EVERY 1 DAY
CREATE EVENT auto_mark_absent
ON SCHEDULE EVERY 5 MINUTE  -- Change to 'EVERY 1 DAY' for production
STARTS NOW()
DO
BEGIN
    DECLARE indonesian_time DATETIME;
    DECLARE indonesian_hour INT;
    DECLARE indonesian_date DATE;
    
    -- Calculate Indonesian time (UTC+7)
    SET indonesian_time = DATE_ADD(UTC_TIMESTAMP(), INTERVAL 7 HOUR);
    SET indonesian_hour = HOUR(indonesian_time);
    SET indonesian_date = DATE(indonesian_time);
    
    -- Only mark absent after 9:00 AM Indonesian time (02:00 UTC) on weekdays
    -- For testing: change >= 9 to >= HOUR(indonesian_time) to test immediately
    IF indonesian_hour >= HOUR(indonesian_time) AND DAYOFWEEK(indonesian_date) NOT IN (1, 7) THEN
        
        INSERT INTO attendances (employee_id, attendance_date, status, is_weekend, marked_by_admin, created_at)
        SELECT 
            e.id,
            indonesian_date,
            'absent',
            FALSE,  -- It's a weekday
            TRUE,   -- Marked by system/admin
            indonesian_time
        FROM employees e
        WHERE e.is_active = TRUE
        AND e.deleted_at IS NULL
        AND NOT EXISTS (
            SELECT 1 FROM attendances a 
            WHERE a.employee_id = e.id 
            AND a.attendance_date = indonesian_date
            AND a.deleted_at IS NULL
        );
        
    END IF;
END;
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
-- Revert to original simple event
DROP EVENT IF EXISTS auto_mark_absent;
-- +goose StatementEnd

-- +goose StatementBegin
CREATE EVENT auto_mark_absent
ON SCHEDULE EVERY 1 DAY
STARTS '2025-06-03 09:00:00'
DO
BEGIN
    INSERT INTO attendances (employee_id, attendance_date, status, is_weekend, marked_by_admin)
    SELECT 
        e.id,
        CURDATE(),
        'absent',
        CASE WHEN DAYOFWEEK(CURDATE()) IN (1, 7) THEN TRUE ELSE FALSE END,
        TRUE
    FROM employees e
    WHERE e.is_active = TRUE
    AND NOT EXISTS (
        SELECT 1 FROM attendances a 
        WHERE a.employee_id = e.id 
        AND a.attendance_date = CURDATE()
    )
    AND DAYOFWEEK(CURDATE()) NOT IN (1, 7)
    AND TIME(NOW()) >= '09:00:00';
END;
-- +goose StatementEnd 