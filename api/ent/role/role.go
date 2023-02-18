// Code generated by ent, DO NOT EDIT.

package role

const (
	// Label holds the string label denoting the role type in the database.
	Label = "role"
	// FieldID holds the string denoting the id field in the database.
	FieldID = "id"
	// FieldName holds the string denoting the name field in the database.
	FieldName = "name"
	// FieldDescription holds the string denoting the description field in the database.
	FieldDescription = "description"
	// EdgeUsers holds the string denoting the users edge name in mutations.
	EdgeUsers = "users"
	// EdgePermissions holds the string denoting the permissions edge name in mutations.
	EdgePermissions = "permissions"
	// Table holds the table name of the role in the database.
	Table = "roles"
	// UsersTable is the table that holds the users relation/edge.
	UsersTable = "users"
	// UsersInverseTable is the table name for the User entity.
	// It exists in this package in order to avoid circular dependency with the "user" package.
	UsersInverseTable = "users"
	// UsersColumn is the table column denoting the users relation/edge.
	UsersColumn = "user_role"
	// PermissionsTable is the table that holds the permissions relation/edge. The primary key declared below.
	PermissionsTable = "role_permissions"
	// PermissionsInverseTable is the table name for the Permission entity.
	// It exists in this package in order to avoid circular dependency with the "permission" package.
	PermissionsInverseTable = "permissions"
)

// Columns holds all SQL columns for role fields.
var Columns = []string{
	FieldID,
	FieldName,
	FieldDescription,
}

var (
	// PermissionsPrimaryKey and PermissionsColumn2 are the table columns denoting the
	// primary key for the permissions relation (M2M).
	PermissionsPrimaryKey = []string{"role_id", "permission_id"}
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