package stubs

import "fmt"

type PostgresqlStubs struct {
}

// CreateUp Create up migration content.
func (receiver PostgresqlStubs) CreateUp(fileName, table string) string {
	return fmt.Sprintf(`CREATE TABLE %v (
  id SERIAL PRIMARY KEY NOT NULL,
  created_at timestamp NOT NULL,
  updated_at timestamp NOT NULL,
  deleted_at timestamp NULL
);
`, table)
}

// CreateDown Create down migration content.
func (receiver PostgresqlStubs) CreateDown(fileName, table string) string {
	return fmt.Sprintf(`DROP TABLE IF EXISTS %v;`, table)
}

// UpdateUp Update up migration content.
func (receiver PostgresqlStubs) UpdateUp(fileName, table string) string {
	return fmt.Sprintf(`ALTER TABLE %v ADD column varchar(255) NOT NULL;`, table)
}

// UpdateDown Update down migration content.
func (receiver PostgresqlStubs) UpdateDown(fileName, table string) string {
	return fmt.Sprintf(`ALTER TABLE %v DROP COLUMN column;`, table)
}
