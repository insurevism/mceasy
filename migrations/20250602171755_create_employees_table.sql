-- +goose Up
-- +goose StatementBegin
CREATE TABLE employees (
    id BIGINT UNSIGNED AUTO_INCREMENT PRIMARY KEY,
    employee_id VARCHAR(20) NOT NULL UNIQUE COMMENT 'Unique employee identifier like EMP-0001, EMP-0002',
    full_name VARCHAR(255) NOT NULL,
    email VARCHAR(255) NOT NULL UNIQUE,
    phone VARCHAR(20),
    position VARCHAR(100),
    department VARCHAR(100),
    hire_date DATE NOT NULL,
    base_salary DECIMAL(15,2) NOT NULL DEFAULT 10000000.00 COMMENT 'Base salary in IDR',
    is_active BOOLEAN NOT NULL DEFAULT TRUE,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    modified_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
    deleted_at TIMESTAMP NULL,
    
    INDEX idx_employee_id (employee_id),
    INDEX idx_email (email),
    INDEX idx_is_active (is_active),
    INDEX idx_deleted_at (deleted_at)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci;

-- Create a trigger to auto-generate employee_id
DELIMITER $$
CREATE TRIGGER before_insert_employees
    BEFORE INSERT ON employees
    FOR EACH ROW
BEGIN
    IF NEW.employee_id IS NULL OR NEW.employee_id = '' THEN
        SET @next_id = (SELECT COALESCE(MAX(CAST(SUBSTRING(employee_id, 5) AS UNSIGNED)), 0) + 1 FROM employees WHERE employee_id REGEXP '^EMP-[0-9]+$');
        SET NEW.employee_id = CONCAT('EMP-', LPAD(@next_id, 4, '0'));
    END IF;
END$$
DELIMITER ;
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TRIGGER IF EXISTS before_insert_employees;
DROP TABLE employees;
-- +goose StatementEnd
