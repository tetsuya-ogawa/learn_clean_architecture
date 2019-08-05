package database

type SqlHandler interface {
  Excute(string, ...interface{}) (Result, error)
  Query(string, ...interface{}) (Row, error)
}

type Result interface {
  LastInsertId() (int, error)
  RowsAffected() (int64, error)
}

type Row interface {
  Scan(...interface{}) error
  Next() bool
  Close() error
}
