module backend/DBAdapter

go 1.24.5

replace backend/DBConnection => ../DBConnection

require (
	backend/DBConnection v0.0.0-00010101000000-000000000000
	github.com/google/uuid v1.6.0
)
