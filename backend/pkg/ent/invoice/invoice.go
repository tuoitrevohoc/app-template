// Code generated by ent, DO NOT EDIT.

package invoice

const (
	// Label holds the string label denoting the invoice type in the database.
	Label = "invoice"
	// FieldID holds the string denoting the id field in the database.
	FieldID = "id"
	// FieldTitle holds the string denoting the title field in the database.
	FieldTitle = "title"
	// FieldLeetCodeLink holds the string denoting the leet_code_link field in the database.
	FieldLeetCodeLink = "leet_code_link"
	// FieldInvoicedTo holds the string denoting the invoiced_to field in the database.
	FieldInvoicedTo = "invoiced_to"
	// Table holds the table name of the invoice in the database.
	Table = "invoices"
)

// Columns holds all SQL columns for invoice fields.
var Columns = []string{
	FieldID,
	FieldTitle,
	FieldLeetCodeLink,
	FieldInvoicedTo,
}

// ValidColumn reports if the column name is valid (part of the table columns).
func ValidColumn(column string) bool {
	for i := range Columns {
		if column == Columns[i] {
			return true
		}
	}
	return false
}
