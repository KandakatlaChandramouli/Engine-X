package recovery

import (
        "encoding/binary"
        "os"
)

type Checkpoint struct {
        LastLSN uint64
}

func WriteCheckpoint(
        path string,
        lsn uint64,
) error {

        f,
        err :=
                os.Create(
                        path,
                )

        if err != nil {
                return err
        }

        defer f.Close()

        var buf [8]byte

        binary.LittleEndian.PutUint64(
                buf[:],
                lsn,
        )

        _, err =
                f.Write(
                        buf[:],
                )

        return err
}

func ReadCheckpoint(
        path string,
) (
        Checkpoint,
        error,
) {

        f,
        err :=
                os.Open(
                        path,
                )

        if err != nil {
                return Checkpoint{},
                        err
        }

        defer f.Close()

        var buf [8]byte

        _, err =
                f.Read(
                        buf[:],
                )

        if err != nil {
                return Checkpoint{},
                        err
        }

        return Checkpoint{
                LastLSN:
                        binary.LittleEndian.Uint64(
                                buf[:],
                        ),
        }, nil
}
