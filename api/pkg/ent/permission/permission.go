// Code generated by ent, DO NOT EDIT.

package permission

const (
	// Label holds the string label denoting the permission type in the database.
	Label = "permission"
	// FieldID holds the string denoting the id field in the database.
	FieldID = "id"
	// FieldName holds the string denoting the name field in the database.
	FieldName = "name"
	// FieldDescription holds the string denoting the description field in the database.
	FieldDescription = "description"
	// EdgeRoles holds the string denoting the roles edge name in mutations.
	EdgeRoles = "roles"
	// Table holds the table name of the permission in the database.
	Table = "permissions"
	// RolesTable is the table that holds the roles relation/edge. The primary key declared below.
	RolesTable = "role_permissions"
	// RolesInverseTable is the table name for the Role entity.
	// It exists in this package in order to avoid circular dependency with the "role" package.
	RolesInverseTable = "roles"
)

// Columns holds all SQL columns for permission fields.
var Columns = []string{
	FieldID,
	FieldName,
	FieldDescription,
}

var (
	// RolesPrimaryKey and RolesColumn2 are the table columns denoting the
	// primary key for the roles relation (M2M).
	RolesPrimaryKey = []string{"role_id", "permission_id"}
)

// ValidColumn reports if the column name is valid (part of the table columns).
func ValidColumn(column string) bool {
	for i := range Columns {
		if column == Columns[i] {
			return true
		}
	}
	return false
}