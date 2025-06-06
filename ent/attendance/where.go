// Code generated by ent, DO NOT EDIT.

package attendance

import (
	"mceasy/ent/predicate"
	"time"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
)

// ID filters vertices based on their ID field.
func ID(id uint64) predicate.Attendance {
	return predicate.Attendance(sql.FieldEQ(FieldID, id))
}

// IDEQ applies the EQ predicate on the ID field.
func IDEQ(id uint64) predicate.Attendance {
	return predicate.Attendance(sql.FieldEQ(FieldID, id))
}

// IDNEQ applies the NEQ predicate on the ID field.
func IDNEQ(id uint64) predicate.Attendance {
	return predicate.Attendance(sql.FieldNEQ(FieldID, id))
}

// IDIn applies the In predicate on the ID field.
func IDIn(ids ...uint64) predicate.Attendance {
	return predicate.Attendance(sql.FieldIn(FieldID, ids...))
}

// IDNotIn applies the NotIn predicate on the ID field.
func IDNotIn(ids ...uint64) predicate.Attendance {
	return predicate.Attendance(sql.FieldNotIn(FieldID, ids...))
}

// IDGT applies the GT predicate on the ID field.
func IDGT(id uint64) predicate.Attendance {
	return predicate.Attendance(sql.FieldGT(FieldID, id))
}

// IDGTE applies the GTE predicate on the ID field.
func IDGTE(id uint64) predicate.Attendance {
	return predicate.Attendance(sql.FieldGTE(FieldID, id))
}

// IDLT applies the LT predicate on the ID field.
func IDLT(id uint64) predicate.Attendance {
	return predicate.Attendance(sql.FieldLT(FieldID, id))
}

// IDLTE applies the LTE predicate on the ID field.
func IDLTE(id uint64) predicate.Attendance {
	return predicate.Attendance(sql.FieldLTE(FieldID, id))
}

// CreatedAt applies equality check predicate on the "created_at" field. It's identical to CreatedAtEQ.
func CreatedAt(v time.Time) predicate.Attendance {
	return predicate.Attendance(sql.FieldEQ(FieldCreatedAt, v))
}

// ModifiedAt applies equality check predicate on the "modified_at" field. It's identical to ModifiedAtEQ.
func ModifiedAt(v time.Time) predicate.Attendance {
	return predicate.Attendance(sql.FieldEQ(FieldModifiedAt, v))
}

// DeletedAt applies equality check predicate on the "deleted_at" field. It's identical to DeletedAtEQ.
func DeletedAt(v time.Time) predicate.Attendance {
	return predicate.Attendance(sql.FieldEQ(FieldDeletedAt, v))
}

// EmployeeID applies equality check predicate on the "employee_id" field. It's identical to EmployeeIDEQ.
func EmployeeID(v uint64) predicate.Attendance {
	return predicate.Attendance(sql.FieldEQ(FieldEmployeeID, v))
}

// AttendanceDate applies equality check predicate on the "attendance_date" field. It's identical to AttendanceDateEQ.
func AttendanceDate(v time.Time) predicate.Attendance {
	return predicate.Attendance(sql.FieldEQ(FieldAttendanceDate, v))
}

// CheckInTime applies equality check predicate on the "check_in_time" field. It's identical to CheckInTimeEQ.
func CheckInTime(v time.Time) predicate.Attendance {
	return predicate.Attendance(sql.FieldEQ(FieldCheckInTime, v))
}

// CheckOutTime applies equality check predicate on the "check_out_time" field. It's identical to CheckOutTimeEQ.
func CheckOutTime(v time.Time) predicate.Attendance {
	return predicate.Attendance(sql.FieldEQ(FieldCheckOutTime, v))
}

// IsWeekend applies equality check predicate on the "is_weekend" field. It's identical to IsWeekendEQ.
func IsWeekend(v bool) predicate.Attendance {
	return predicate.Attendance(sql.FieldEQ(FieldIsWeekend, v))
}

// Notes applies equality check predicate on the "notes" field. It's identical to NotesEQ.
func Notes(v string) predicate.Attendance {
	return predicate.Attendance(sql.FieldEQ(FieldNotes, v))
}

// MarkedByAdmin applies equality check predicate on the "marked_by_admin" field. It's identical to MarkedByAdminEQ.
func MarkedByAdmin(v bool) predicate.Attendance {
	return predicate.Attendance(sql.FieldEQ(FieldMarkedByAdmin, v))
}

// CreatedAtEQ applies the EQ predicate on the "created_at" field.
func CreatedAtEQ(v time.Time) predicate.Attendance {
	return predicate.Attendance(sql.FieldEQ(FieldCreatedAt, v))
}

// CreatedAtNEQ applies the NEQ predicate on the "created_at" field.
func CreatedAtNEQ(v time.Time) predicate.Attendance {
	return predicate.Attendance(sql.FieldNEQ(FieldCreatedAt, v))
}

// CreatedAtIn applies the In predicate on the "created_at" field.
func CreatedAtIn(vs ...time.Time) predicate.Attendance {
	return predicate.Attendance(sql.FieldIn(FieldCreatedAt, vs...))
}

// CreatedAtNotIn applies the NotIn predicate on the "created_at" field.
func CreatedAtNotIn(vs ...time.Time) predicate.Attendance {
	return predicate.Attendance(sql.FieldNotIn(FieldCreatedAt, vs...))
}

// CreatedAtGT applies the GT predicate on the "created_at" field.
func CreatedAtGT(v time.Time) predicate.Attendance {
	return predicate.Attendance(sql.FieldGT(FieldCreatedAt, v))
}

// CreatedAtGTE applies the GTE predicate on the "created_at" field.
func CreatedAtGTE(v time.Time) predicate.Attendance {
	return predicate.Attendance(sql.FieldGTE(FieldCreatedAt, v))
}

// CreatedAtLT applies the LT predicate on the "created_at" field.
func CreatedAtLT(v time.Time) predicate.Attendance {
	return predicate.Attendance(sql.FieldLT(FieldCreatedAt, v))
}

// CreatedAtLTE applies the LTE predicate on the "created_at" field.
func CreatedAtLTE(v time.Time) predicate.Attendance {
	return predicate.Attendance(sql.FieldLTE(FieldCreatedAt, v))
}

// ModifiedAtEQ applies the EQ predicate on the "modified_at" field.
func ModifiedAtEQ(v time.Time) predicate.Attendance {
	return predicate.Attendance(sql.FieldEQ(FieldModifiedAt, v))
}

// ModifiedAtNEQ applies the NEQ predicate on the "modified_at" field.
func ModifiedAtNEQ(v time.Time) predicate.Attendance {
	return predicate.Attendance(sql.FieldNEQ(FieldModifiedAt, v))
}

// ModifiedAtIn applies the In predicate on the "modified_at" field.
func ModifiedAtIn(vs ...time.Time) predicate.Attendance {
	return predicate.Attendance(sql.FieldIn(FieldModifiedAt, vs...))
}

// ModifiedAtNotIn applies the NotIn predicate on the "modified_at" field.
func ModifiedAtNotIn(vs ...time.Time) predicate.Attendance {
	return predicate.Attendance(sql.FieldNotIn(FieldModifiedAt, vs...))
}

// ModifiedAtGT applies the GT predicate on the "modified_at" field.
func ModifiedAtGT(v time.Time) predicate.Attendance {
	return predicate.Attendance(sql.FieldGT(FieldModifiedAt, v))
}

// ModifiedAtGTE applies the GTE predicate on the "modified_at" field.
func ModifiedAtGTE(v time.Time) predicate.Attendance {
	return predicate.Attendance(sql.FieldGTE(FieldModifiedAt, v))
}

// ModifiedAtLT applies the LT predicate on the "modified_at" field.
func ModifiedAtLT(v time.Time) predicate.Attendance {
	return predicate.Attendance(sql.FieldLT(FieldModifiedAt, v))
}

// ModifiedAtLTE applies the LTE predicate on the "modified_at" field.
func ModifiedAtLTE(v time.Time) predicate.Attendance {
	return predicate.Attendance(sql.FieldLTE(FieldModifiedAt, v))
}

// DeletedAtEQ applies the EQ predicate on the "deleted_at" field.
func DeletedAtEQ(v time.Time) predicate.Attendance {
	return predicate.Attendance(sql.FieldEQ(FieldDeletedAt, v))
}

// DeletedAtNEQ applies the NEQ predicate on the "deleted_at" field.
func DeletedAtNEQ(v time.Time) predicate.Attendance {
	return predicate.Attendance(sql.FieldNEQ(FieldDeletedAt, v))
}

// DeletedAtIn applies the In predicate on the "deleted_at" field.
func DeletedAtIn(vs ...time.Time) predicate.Attendance {
	return predicate.Attendance(sql.FieldIn(FieldDeletedAt, vs...))
}

// DeletedAtNotIn applies the NotIn predicate on the "deleted_at" field.
func DeletedAtNotIn(vs ...time.Time) predicate.Attendance {
	return predicate.Attendance(sql.FieldNotIn(FieldDeletedAt, vs...))
}

// DeletedAtGT applies the GT predicate on the "deleted_at" field.
func DeletedAtGT(v time.Time) predicate.Attendance {
	return predicate.Attendance(sql.FieldGT(FieldDeletedAt, v))
}

// DeletedAtGTE applies the GTE predicate on the "deleted_at" field.
func DeletedAtGTE(v time.Time) predicate.Attendance {
	return predicate.Attendance(sql.FieldGTE(FieldDeletedAt, v))
}

// DeletedAtLT applies the LT predicate on the "deleted_at" field.
func DeletedAtLT(v time.Time) predicate.Attendance {
	return predicate.Attendance(sql.FieldLT(FieldDeletedAt, v))
}

// DeletedAtLTE applies the LTE predicate on the "deleted_at" field.
func DeletedAtLTE(v time.Time) predicate.Attendance {
	return predicate.Attendance(sql.FieldLTE(FieldDeletedAt, v))
}

// DeletedAtIsNil applies the IsNil predicate on the "deleted_at" field.
func DeletedAtIsNil() predicate.Attendance {
	return predicate.Attendance(sql.FieldIsNull(FieldDeletedAt))
}

// DeletedAtNotNil applies the NotNil predicate on the "deleted_at" field.
func DeletedAtNotNil() predicate.Attendance {
	return predicate.Attendance(sql.FieldNotNull(FieldDeletedAt))
}

// EmployeeIDEQ applies the EQ predicate on the "employee_id" field.
func EmployeeIDEQ(v uint64) predicate.Attendance {
	return predicate.Attendance(sql.FieldEQ(FieldEmployeeID, v))
}

// EmployeeIDNEQ applies the NEQ predicate on the "employee_id" field.
func EmployeeIDNEQ(v uint64) predicate.Attendance {
	return predicate.Attendance(sql.FieldNEQ(FieldEmployeeID, v))
}

// EmployeeIDIn applies the In predicate on the "employee_id" field.
func EmployeeIDIn(vs ...uint64) predicate.Attendance {
	return predicate.Attendance(sql.FieldIn(FieldEmployeeID, vs...))
}

// EmployeeIDNotIn applies the NotIn predicate on the "employee_id" field.
func EmployeeIDNotIn(vs ...uint64) predicate.Attendance {
	return predicate.Attendance(sql.FieldNotIn(FieldEmployeeID, vs...))
}

// AttendanceDateEQ applies the EQ predicate on the "attendance_date" field.
func AttendanceDateEQ(v time.Time) predicate.Attendance {
	return predicate.Attendance(sql.FieldEQ(FieldAttendanceDate, v))
}

// AttendanceDateNEQ applies the NEQ predicate on the "attendance_date" field.
func AttendanceDateNEQ(v time.Time) predicate.Attendance {
	return predicate.Attendance(sql.FieldNEQ(FieldAttendanceDate, v))
}

// AttendanceDateIn applies the In predicate on the "attendance_date" field.
func AttendanceDateIn(vs ...time.Time) predicate.Attendance {
	return predicate.Attendance(sql.FieldIn(FieldAttendanceDate, vs...))
}

// AttendanceDateNotIn applies the NotIn predicate on the "attendance_date" field.
func AttendanceDateNotIn(vs ...time.Time) predicate.Attendance {
	return predicate.Attendance(sql.FieldNotIn(FieldAttendanceDate, vs...))
}

// AttendanceDateGT applies the GT predicate on the "attendance_date" field.
func AttendanceDateGT(v time.Time) predicate.Attendance {
	return predicate.Attendance(sql.FieldGT(FieldAttendanceDate, v))
}

// AttendanceDateGTE applies the GTE predicate on the "attendance_date" field.
func AttendanceDateGTE(v time.Time) predicate.Attendance {
	return predicate.Attendance(sql.FieldGTE(FieldAttendanceDate, v))
}

// AttendanceDateLT applies the LT predicate on the "attendance_date" field.
func AttendanceDateLT(v time.Time) predicate.Attendance {
	return predicate.Attendance(sql.FieldLT(FieldAttendanceDate, v))
}

// AttendanceDateLTE applies the LTE predicate on the "attendance_date" field.
func AttendanceDateLTE(v time.Time) predicate.Attendance {
	return predicate.Attendance(sql.FieldLTE(FieldAttendanceDate, v))
}

// CheckInTimeEQ applies the EQ predicate on the "check_in_time" field.
func CheckInTimeEQ(v time.Time) predicate.Attendance {
	return predicate.Attendance(sql.FieldEQ(FieldCheckInTime, v))
}

// CheckInTimeNEQ applies the NEQ predicate on the "check_in_time" field.
func CheckInTimeNEQ(v time.Time) predicate.Attendance {
	return predicate.Attendance(sql.FieldNEQ(FieldCheckInTime, v))
}

// CheckInTimeIn applies the In predicate on the "check_in_time" field.
func CheckInTimeIn(vs ...time.Time) predicate.Attendance {
	return predicate.Attendance(sql.FieldIn(FieldCheckInTime, vs...))
}

// CheckInTimeNotIn applies the NotIn predicate on the "check_in_time" field.
func CheckInTimeNotIn(vs ...time.Time) predicate.Attendance {
	return predicate.Attendance(sql.FieldNotIn(FieldCheckInTime, vs...))
}

// CheckInTimeGT applies the GT predicate on the "check_in_time" field.
func CheckInTimeGT(v time.Time) predicate.Attendance {
	return predicate.Attendance(sql.FieldGT(FieldCheckInTime, v))
}

// CheckInTimeGTE applies the GTE predicate on the "check_in_time" field.
func CheckInTimeGTE(v time.Time) predicate.Attendance {
	return predicate.Attendance(sql.FieldGTE(FieldCheckInTime, v))
}

// CheckInTimeLT applies the LT predicate on the "check_in_time" field.
func CheckInTimeLT(v time.Time) predicate.Attendance {
	return predicate.Attendance(sql.FieldLT(FieldCheckInTime, v))
}

// CheckInTimeLTE applies the LTE predicate on the "check_in_time" field.
func CheckInTimeLTE(v time.Time) predicate.Attendance {
	return predicate.Attendance(sql.FieldLTE(FieldCheckInTime, v))
}

// CheckInTimeIsNil applies the IsNil predicate on the "check_in_time" field.
func CheckInTimeIsNil() predicate.Attendance {
	return predicate.Attendance(sql.FieldIsNull(FieldCheckInTime))
}

// CheckInTimeNotNil applies the NotNil predicate on the "check_in_time" field.
func CheckInTimeNotNil() predicate.Attendance {
	return predicate.Attendance(sql.FieldNotNull(FieldCheckInTime))
}

// CheckOutTimeEQ applies the EQ predicate on the "check_out_time" field.
func CheckOutTimeEQ(v time.Time) predicate.Attendance {
	return predicate.Attendance(sql.FieldEQ(FieldCheckOutTime, v))
}

// CheckOutTimeNEQ applies the NEQ predicate on the "check_out_time" field.
func CheckOutTimeNEQ(v time.Time) predicate.Attendance {
	return predicate.Attendance(sql.FieldNEQ(FieldCheckOutTime, v))
}

// CheckOutTimeIn applies the In predicate on the "check_out_time" field.
func CheckOutTimeIn(vs ...time.Time) predicate.Attendance {
	return predicate.Attendance(sql.FieldIn(FieldCheckOutTime, vs...))
}

// CheckOutTimeNotIn applies the NotIn predicate on the "check_out_time" field.
func CheckOutTimeNotIn(vs ...time.Time) predicate.Attendance {
	return predicate.Attendance(sql.FieldNotIn(FieldCheckOutTime, vs...))
}

// CheckOutTimeGT applies the GT predicate on the "check_out_time" field.
func CheckOutTimeGT(v time.Time) predicate.Attendance {
	return predicate.Attendance(sql.FieldGT(FieldCheckOutTime, v))
}

// CheckOutTimeGTE applies the GTE predicate on the "check_out_time" field.
func CheckOutTimeGTE(v time.Time) predicate.Attendance {
	return predicate.Attendance(sql.FieldGTE(FieldCheckOutTime, v))
}

// CheckOutTimeLT applies the LT predicate on the "check_out_time" field.
func CheckOutTimeLT(v time.Time) predicate.Attendance {
	return predicate.Attendance(sql.FieldLT(FieldCheckOutTime, v))
}

// CheckOutTimeLTE applies the LTE predicate on the "check_out_time" field.
func CheckOutTimeLTE(v time.Time) predicate.Attendance {
	return predicate.Attendance(sql.FieldLTE(FieldCheckOutTime, v))
}

// CheckOutTimeIsNil applies the IsNil predicate on the "check_out_time" field.
func CheckOutTimeIsNil() predicate.Attendance {
	return predicate.Attendance(sql.FieldIsNull(FieldCheckOutTime))
}

// CheckOutTimeNotNil applies the NotNil predicate on the "check_out_time" field.
func CheckOutTimeNotNil() predicate.Attendance {
	return predicate.Attendance(sql.FieldNotNull(FieldCheckOutTime))
}

// StatusEQ applies the EQ predicate on the "status" field.
func StatusEQ(v Status) predicate.Attendance {
	return predicate.Attendance(sql.FieldEQ(FieldStatus, v))
}

// StatusNEQ applies the NEQ predicate on the "status" field.
func StatusNEQ(v Status) predicate.Attendance {
	return predicate.Attendance(sql.FieldNEQ(FieldStatus, v))
}

// StatusIn applies the In predicate on the "status" field.
func StatusIn(vs ...Status) predicate.Attendance {
	return predicate.Attendance(sql.FieldIn(FieldStatus, vs...))
}

// StatusNotIn applies the NotIn predicate on the "status" field.
func StatusNotIn(vs ...Status) predicate.Attendance {
	return predicate.Attendance(sql.FieldNotIn(FieldStatus, vs...))
}

// IsWeekendEQ applies the EQ predicate on the "is_weekend" field.
func IsWeekendEQ(v bool) predicate.Attendance {
	return predicate.Attendance(sql.FieldEQ(FieldIsWeekend, v))
}

// IsWeekendNEQ applies the NEQ predicate on the "is_weekend" field.
func IsWeekendNEQ(v bool) predicate.Attendance {
	return predicate.Attendance(sql.FieldNEQ(FieldIsWeekend, v))
}

// NotesEQ applies the EQ predicate on the "notes" field.
func NotesEQ(v string) predicate.Attendance {
	return predicate.Attendance(sql.FieldEQ(FieldNotes, v))
}

// NotesNEQ applies the NEQ predicate on the "notes" field.
func NotesNEQ(v string) predicate.Attendance {
	return predicate.Attendance(sql.FieldNEQ(FieldNotes, v))
}

// NotesIn applies the In predicate on the "notes" field.
func NotesIn(vs ...string) predicate.Attendance {
	return predicate.Attendance(sql.FieldIn(FieldNotes, vs...))
}

// NotesNotIn applies the NotIn predicate on the "notes" field.
func NotesNotIn(vs ...string) predicate.Attendance {
	return predicate.Attendance(sql.FieldNotIn(FieldNotes, vs...))
}

// NotesGT applies the GT predicate on the "notes" field.
func NotesGT(v string) predicate.Attendance {
	return predicate.Attendance(sql.FieldGT(FieldNotes, v))
}

// NotesGTE applies the GTE predicate on the "notes" field.
func NotesGTE(v string) predicate.Attendance {
	return predicate.Attendance(sql.FieldGTE(FieldNotes, v))
}

// NotesLT applies the LT predicate on the "notes" field.
func NotesLT(v string) predicate.Attendance {
	return predicate.Attendance(sql.FieldLT(FieldNotes, v))
}

// NotesLTE applies the LTE predicate on the "notes" field.
func NotesLTE(v string) predicate.Attendance {
	return predicate.Attendance(sql.FieldLTE(FieldNotes, v))
}

// NotesContains applies the Contains predicate on the "notes" field.
func NotesContains(v string) predicate.Attendance {
	return predicate.Attendance(sql.FieldContains(FieldNotes, v))
}

// NotesHasPrefix applies the HasPrefix predicate on the "notes" field.
func NotesHasPrefix(v string) predicate.Attendance {
	return predicate.Attendance(sql.FieldHasPrefix(FieldNotes, v))
}

// NotesHasSuffix applies the HasSuffix predicate on the "notes" field.
func NotesHasSuffix(v string) predicate.Attendance {
	return predicate.Attendance(sql.FieldHasSuffix(FieldNotes, v))
}

// NotesIsNil applies the IsNil predicate on the "notes" field.
func NotesIsNil() predicate.Attendance {
	return predicate.Attendance(sql.FieldIsNull(FieldNotes))
}

// NotesNotNil applies the NotNil predicate on the "notes" field.
func NotesNotNil() predicate.Attendance {
	return predicate.Attendance(sql.FieldNotNull(FieldNotes))
}

// NotesEqualFold applies the EqualFold predicate on the "notes" field.
func NotesEqualFold(v string) predicate.Attendance {
	return predicate.Attendance(sql.FieldEqualFold(FieldNotes, v))
}

// NotesContainsFold applies the ContainsFold predicate on the "notes" field.
func NotesContainsFold(v string) predicate.Attendance {
	return predicate.Attendance(sql.FieldContainsFold(FieldNotes, v))
}

// MarkedByAdminEQ applies the EQ predicate on the "marked_by_admin" field.
func MarkedByAdminEQ(v bool) predicate.Attendance {
	return predicate.Attendance(sql.FieldEQ(FieldMarkedByAdmin, v))
}

// MarkedByAdminNEQ applies the NEQ predicate on the "marked_by_admin" field.
func MarkedByAdminNEQ(v bool) predicate.Attendance {
	return predicate.Attendance(sql.FieldNEQ(FieldMarkedByAdmin, v))
}

// HasEmployee applies the HasEdge predicate on the "employee" edge.
func HasEmployee() predicate.Attendance {
	return predicate.Attendance(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.Edge(sqlgraph.M2O, true, EmployeeTable, EmployeeColumn),
		)
		sqlgraph.HasNeighbors(s, step)
	})
}

// HasEmployeeWith applies the HasEdge predicate on the "employee" edge with a given conditions (other predicates).
func HasEmployeeWith(preds ...predicate.Employee) predicate.Attendance {
	return predicate.Attendance(func(s *sql.Selector) {
		step := newEmployeeStep()
		sqlgraph.HasNeighborsWith(s, step, func(s *sql.Selector) {
			for _, p := range preds {
				p(s)
			}
		})
	})
}

// And groups predicates with the AND operator between them.
func And(predicates ...predicate.Attendance) predicate.Attendance {
	return predicate.Attendance(func(s *sql.Selector) {
		s1 := s.Clone().SetP(nil)
		for _, p := range predicates {
			p(s1)
		}
		s.Where(s1.P())
	})
}

// Or groups predicates with the OR operator between them.
func Or(predicates ...predicate.Attendance) predicate.Attendance {
	return predicate.Attendance(func(s *sql.Selector) {
		s1 := s.Clone().SetP(nil)
		for i, p := range predicates {
			if i > 0 {
				s1.Or()
			}
			p(s1)
		}
		s.Where(s1.P())
	})
}

// Not applies the not operator on the given predicate.
func Not(p predicate.Attendance) predicate.Attendance {
	return predicate.Attendance(func(s *sql.Selector) {
		p(s.Not())
	})
}
