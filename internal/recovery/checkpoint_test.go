package recovery

import (
        "os"
        "testing"
)

func TestCheckpointRoundTrip(
        t *testing.T,
) {

        path := "checkpoint.dat"

        defer os.Remove(path)

        err :=
                WriteCheckpoint(
                        path,
                        12345,
                )

        if err != nil {
                t.Fatal(err)
        }

        cp,
        err :=
                ReadCheckpoint(
                        path,
                )

        if err != nil {
                t.Fatal(err)
        }

        if cp.LastLSN != 12345 {
                t.Fatal()
        }
}
