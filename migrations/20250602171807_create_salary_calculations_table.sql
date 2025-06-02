-- +goose Up
-- +goose StatementBegin
CREATE TABLE salary_calculations (
    id BIGINT UNSIGNED AUTO_INCREMENT PRIMARY KEY,
    employee_id BIGINT UNSIGNED NOT NULL,
    calculation_month DATE NOT NULL COMMENT 'First day of the month for calculation (YYYY-MM-01)',
    base_salary DECIMAL(15,2) NOT NULL COMMENT 'Base salary for the month',
    total_working_days INT NOT NULL COMMENT 'Total working days in the month (excluding weekends)',
    absent_days INT NOT NULL DEFAULT 0 COMMENT 'Number of absent working days',
    present_days INT NOT NULL DEFAULT 0 COMMENT 'Number of present working days',
    final_salary DECIMAL(15,2) NOT NULL COMMENT 'Final calculated salary after deductions',
    deduction_amount DECIMAL(15,2) NOT NULL DEFAULT 0.00 COMMENT 'Total deduction amount',
    calculation_formula TEXT COMMENT 'Formula used for calculation (for audit purposes)',
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    modified_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
    deleted_at TIMESTAMP NULL,
    
    FOREIGN KEY (employee_id) REFERENCES employees(id) ON DELETE CASCADE,
    UNIQUE KEY unique_employee_month (employee_id, calculation_month),
    INDEX idx_calculation_month (calculation_month),
    INDEX idx_employee_id (employee_id),
    INDEX idx_deleted_at (deleted_at)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci;
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE salary_calculations;
-- +goose StatementEnd
