# mock

A mock library implementing [github.com/hanzo-ds/go/lib/driver](https://pkg.go.dev/github.com/hanzo-ds/go/lib/driver),
for testing code that talks to Datastore without a server.

## Install

```sh
go get github.com/hanzo-ds/mock
```

## Usage

```go
package main

import (
	"context"
	"testing"

	datastore "github.com/hanzo-ds/go"
	mock "github.com/hanzo-ds/mock"
)

type Video struct {
	Title string
}

func fetchVideos(conn datastore.Conn) (*Video, error) {
	var v Video
	if err := conn.QueryRow(context.Background(), "SELECT title FROM videos LIMIT 1").Scan(&v.Title); err != nil {
		return nil, err
	}
	return &v, nil
}

func TestFetchVideos(t *testing.T) {
	conn, err := mock.NewDatastoreNative(nil)
	if err != nil {
		t.Fatal(err)
	}

	row := mock.NewRow(
		[]mock.ColumnType{{Name: "title", Type: "String"}},
		[]any{"a title"},
	)
	conn.ExpectQueryRow("SELECT title FROM videos LIMIT 1").WillReturnRow(row)

	if _, err := fetchVideos(conn); err != nil {
		t.Fatal(err)
	}
	if err := conn.ExpectationsWereMet(); err != nil {
		t.Fatal(err)
	}
}
```

See the package documentation at [pkg.go.dev](https://pkg.go.dev/github.com/hanzo-ds/mock).

## Provenance

A de-branded fork of the upstream mock library, cut from its v0.13.0 release —
the latest upstream release tag. It is built on [github.com/hanzo-ds/go](https://github.com/hanzo-ds/go)
and [sqlmock](https://github.com/DATA-DOG/go-sqlmock).
