package recovery

import (
        "os"
        "testing"

        "engine-x/wal"
)

func TestScan(
        t *testing.T,
) {

        path := "recovery_test.wal"

        defer os.Remove(path)

        w,
                err :=
                wal.OpenWAL(path)

        if err != nil {
                t.Fatal(err)
        }

        _, _ =
                w.Append(
                        1,
                        []byte("a"),
                )

        _, _ =
                w.Append(
                        2,
                        []byte("b"),
                )

        if err := w.Sync(); err != nil {
                t.Fatal(err)
        }

        records,
        report,
        err :=
                Scan(path)

        if err != nil {
                t.Fatal(err)
        }

        if len(records) != 2 {
                t.Fatal()
        }

        if report.Records != 2 {
                t.Fatal()
        }

        if report.LastLSN != 2 {
                t.Fatal()
        }
}
