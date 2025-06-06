// Code generated by ent, DO NOT EDIT.

package ent

import (
	"fmt"
	"mceasy/ent/attendance"
	"mceasy/ent/employee"
	"strings"
	"time"

	"entgo.io/ent"
	"entgo.io/ent/dialect/sql"
)

// Attendance is the model entity for the Attendance schema.
type Attendance struct {
	config `json:"-"`
	// ID of the ent.
	ID uint64 `json:"id,omitempty"`
	// CreatedAt holds the value of the "created_at" field.
	CreatedAt time.Time `json:"created_at,omitempty"`
	// ModifiedAt holds the value of the "modified_at" field.
	ModifiedAt time.Time `json:"modified_at,omitempty"`
	// DeletedAt holds the value of the "deleted_at" field.
	DeletedAt time.Time `json:"deleted_at,omitempty"`
	// Foreign key to employees table
	EmployeeID uint64 `json:"employee_id,omitempty"`
	// Date of attendance
	AttendanceDate time.Time `json:"attendance_date,omitempty"`
	// Actual check-in time
	CheckInTime time.Time `json:"check_in_time,omitempty"`
	// Actual check-out time
	CheckOutTime time.Time `json:"check_out_time,omitempty"`
	// Status holds the value of the "status" field.
	Status attendance.Status `json:"status,omitempty"`
	// True for Saturday and Sunday
	IsWeekend bool `json:"is_weekend,omitempty"`
	// Additional notes for attendance
	Notes string `json:"notes,omitempty"`
	// True if manually marked by admin
	MarkedByAdmin bool `json:"marked_by_admin,omitempty"`
	// Edges holds the relations/edges for other nodes in the graph.
	// The values are being populated by the AttendanceQuery when eager-loading is set.
	Edges        AttendanceEdges `json:"edges"`
	selectValues sql.SelectValues
}

// AttendanceEdges holds the relations/edges for other nodes in the graph.
type AttendanceEdges struct {
	// Employee holds the value of the employee edge.
	Employee *Employee `json:"employee,omitempty"`
	// loadedTypes holds the information for reporting if a
	// type was loaded (or requested) in eager-loading or not.
	loadedTypes [1]bool
}

// EmployeeOrErr returns the Employee value or an error if the edge
// was not loaded in eager-loading, or loaded but was not found.
func (e AttendanceEdges) EmployeeOrErr() (*Employee, error) {
	if e.loadedTypes[0] {
		if e.Employee == nil {
			// Edge was loaded but was not found.
			return nil, &NotFoundError{label: employee.Label}
		}
		return e.Employee, nil
	}
	return nil, &NotLoadedError{edge: "employee"}
}

// scanValues returns the types for scanning values from sql.Rows.
func (*Attendance) scanValues(columns []string) ([]any, error) {
	values := make([]any, len(columns))
	for i := range columns {
		switch columns[i] {
		case attendance.FieldIsWeekend, attendance.FieldMarkedByAdmin:
			values[i] = new(sql.NullBool)
		case attendance.FieldID, attendance.FieldEmployeeID:
			values[i] = new(sql.NullInt64)
		case attendance.FieldStatus, attendance.FieldNotes:
			values[i] = new(sql.NullString)
		case attendance.FieldCreatedAt, attendance.FieldModifiedAt, attendance.FieldDeletedAt, attendance.FieldAttendanceDate, attendance.FieldCheckInTime, attendance.FieldCheckOutTime:
			values[i] = new(sql.NullTime)
		default:
			values[i] = new(sql.UnknownType)
		}
	}
	return values, nil
}

// assignValues assigns the values that were returned from sql.Rows (after scanning)
// to the Attendance fields.
func (a *Attendance) assignValues(columns []string, values []any) error {
	if m, n := len(values), len(columns); m < n {
		return fmt.Errorf("mismatch number of scan values: %d != %d", m, n)
	}
	for i := range columns {
		switch columns[i] {
		case attendance.FieldID:
			value, ok := values[i].(*sql.NullInt64)
			if !ok {
				return fmt.Errorf("unexpected type %T for field id", value)
			}
			a.ID = uint64(value.Int64)
		case attendance.FieldCreatedAt:
			if value, ok := values[i].(*sql.NullTime); !ok {
				return fmt.Errorf("unexpected type %T for field created_at", values[i])
			} else if value.Valid {
				a.CreatedAt = value.Time
			}
		case attendance.FieldModifiedAt:
			if value, ok := values[i].(*sql.NullTime); !ok {
				return fmt.Errorf("unexpected type %T for field modified_at", values[i])
			} else if value.Valid {
				a.ModifiedAt = value.Time
			}
		case attendance.FieldDeletedAt:
			if value, ok := values[i].(*sql.NullTime); !ok {
				return fmt.Errorf("unexpected type %T for field deleted_at", values[i])
			} else if value.Valid {
				a.DeletedAt = value.Time
			}
		case attendance.FieldEmployeeID:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for field employee_id", values[i])
			} else if value.Valid {
				a.EmployeeID = uint64(value.Int64)
			}
		case attendance.FieldAttendanceDate:
			if value, ok := values[i].(*sql.NullTime); !ok {
				return fmt.Errorf("unexpected type %T for field attendance_date", values[i])
			} else if value.Valid {
				a.AttendanceDate = value.Time
			}
		case attendance.FieldCheckInTime:
			if value, ok := values[i].(*sql.NullTime); !ok {
				return fmt.Errorf("unexpected type %T for field check_in_time", values[i])
			} else if value.Valid {
				a.CheckInTime = value.Time
			}
		case attendance.FieldCheckOutTime:
			if value, ok := values[i].(*sql.NullTime); !ok {
				return fmt.Errorf("unexpected type %T for field check_out_time", values[i])
			} else if value.Valid {
				a.CheckOutTime = value.Time
			}
		case attendance.FieldStatus:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field status", values[i])
			} else if value.Valid {
				a.Status = attendance.Status(value.String)
			}
		case attendance.FieldIsWeekend:
			if value, ok := values[i].(*sql.NullBool); !ok {
				return fmt.Errorf("unexpected type %T for field is_weekend", values[i])
			} else if value.Valid {
				a.IsWeekend = value.Bool
			}
		case attendance.FieldNotes:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field notes", values[i])
			} else if value.Valid {
				a.Notes = value.String
			}
		case attendance.FieldMarkedByAdmin:
			if value, ok := values[i].(*sql.NullBool); !ok {
				return fmt.Errorf("unexpected type %T for field marked_by_admin", values[i])
			} else if value.Valid {
				a.MarkedByAdmin = value.Bool
			}
		default:
			a.selectValues.Set(columns[i], values[i])
		}
	}
	return nil
}

// Value returns the ent.Value that was dynamically selected and assigned to the Attendance.
// This includes values selected through modifiers, order, etc.
func (a *Attendance) Value(name string) (ent.Value, error) {
	return a.selectValues.Get(name)
}

// QueryEmployee queries the "employee" edge of the Attendance entity.
func (a *Attendance) QueryEmployee() *EmployeeQuery {
	return NewAttendanceClient(a.config).QueryEmployee(a)
}

// Update returns a builder for updating this Attendance.
// Note that you need to call Attendance.Unwrap() before calling this method if this Attendance
// was returned from a transaction, and the transaction was committed or rolled back.
func (a *Attendance) Update() *AttendanceUpdateOne {
	return NewAttendanceClient(a.config).UpdateOne(a)
}

// Unwrap unwraps the Attendance entity that was returned from a transaction after it was closed,
// so that all future queries will be executed through the driver which created the transaction.
func (a *Attendance) Unwrap() *Attendance {
	_tx, ok := a.config.driver.(*txDriver)
	if !ok {
		panic("ent: Attendance is not a transactional entity")
	}
	a.config.driver = _tx.drv
	return a
}

// String implements the fmt.Stringer.
func (a *Attendance) String() string {
	var builder strings.Builder
	builder.WriteString("Attendance(")
	builder.WriteString(fmt.Sprintf("id=%v, ", a.ID))
	builder.WriteString("created_at=")
	builder.WriteString(a.CreatedAt.Format(time.ANSIC))
	builder.WriteString(", ")
	builder.WriteString("modified_at=")
	builder.WriteString(a.ModifiedAt.Format(time.ANSIC))
	builder.WriteString(", ")
	builder.WriteString("deleted_at=")
	builder.WriteString(a.DeletedAt.Format(time.ANSIC))
	builder.WriteString(", ")
	builder.WriteString("employee_id=")
	builder.WriteString(fmt.Sprintf("%v", a.EmployeeID))
	builder.WriteString(", ")
	builder.WriteString("attendance_date=")
	builder.WriteString(a.AttendanceDate.Format(time.ANSIC))
	builder.WriteString(", ")
	builder.WriteString("check_in_time=")
	builder.WriteString(a.CheckInTime.Format(time.ANSIC))
	builder.WriteString(", ")
	builder.WriteString("check_out_time=")
	builder.WriteString(a.CheckOutTime.Format(time.ANSIC))
	builder.WriteString(", ")
	builder.WriteString("status=")
	builder.WriteString(fmt.Sprintf("%v", a.Status))
	builder.WriteString(", ")
	builder.WriteString("is_weekend=")
	builder.WriteString(fmt.Sprintf("%v", a.IsWeekend))
	builder.WriteString(", ")
	builder.WriteString("notes=")
	builder.WriteString(a.Notes)
	builder.WriteString(", ")
	builder.WriteString("marked_by_admin=")
	builder.WriteString(fmt.Sprintf("%v", a.MarkedByAdmin))
	builder.WriteByte(')')
	return builder.String()
}

// Attendances is a parsable slice of Attendance.
type Attendances []*Attendance
