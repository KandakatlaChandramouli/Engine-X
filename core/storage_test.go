package core

import (
        "os"
        "testing"
)

func TestStorageOpenClose(
        t *testing.T,
) {
        path := "test.db"

        defer os.Remove(path)

        s, err := Open(
                path,
                4096,
        )

        if err != nil {
                t.Fatal(err)
        }

        if err := s.Close(); err != nil {
                t.Fatal(err)
        }
}

func TestStoragePageCount(
        t *testing.T,
) {
        path := "test_count.db"

        defer os.Remove(path)

        s, err := Open(
                path,
                4096,
        )

        if err != nil {
                t.Fatal(err)
        }

        defer s.Close()

        if s.PageCount() == 0 {
                t.Fatal()
        }
}

func TestStoragePageAccess(
        t *testing.T,
) {
        path := "test_page.db"

        defer os.Remove(path)

        s, err := Open(
                path,
                4096,
        )

        if err != nil {
                t.Fatal(err)
        }

        defer s.Close()

        page := s.Page(0)

        if len(page) != 4096 {
                t.Fatalf(
                        "expected 4096 got %d",
                        len(page),
                )
        }
}

func TestStorageSync(
        t *testing.T,
) {
        path := "test_sync.db"

        defer os.Remove(path)

        s, err := Open(
                path,
                4096,
        )

        if err != nil {
                t.Fatal(err)
        }

        defer s.Close()

        if err := s.Sync(); err != nil {
                t.Fatal(err)
        }
}

func TestStorageInvalidSize(
        t *testing.T,
) {
        _, err :=
                Open(
                        "bad.db",
                        0,
                )

        if err == nil {
                t.Fatal()
        }
}

func TestStorageOpenInvalidPath(
        t *testing.T,
) {
        _, err :=
                Open(
                        "/root/does-not-exist/db",
                        4096,
                )

        if err == nil {
                t.Fatal()
        }
}

func TestStorageDoubleClose(
        t *testing.T,
) {
        path := "double_close.db"

        defer os.Remove(path)

        db, err :=
                Open(
                        path,
                        4096,
                )

        if err != nil {
                t.Fatal(err)
        }

        if err := db.Close(); err != nil {
                t.Fatal(err)
        }

        _ = db.Close()
}
