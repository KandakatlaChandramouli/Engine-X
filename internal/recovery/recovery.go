package recovery

import (
        "io"
        "os"

        "engine-x/wal"
)

type RecoveryRecord struct {
        Header  wal.RecordHeader
        Payload []byte
}

type RecoveryReport struct {
        Records uint64
        LastLSN uint64
}

func Scan(
        path string,
) (
        []RecoveryRecord,
        RecoveryReport,
        error,
) {

        f, err :=
                os.Open(path)

        if err != nil {
                return nil,
                        RecoveryReport{},
                        err
        }

        defer f.Close()

        var out []RecoveryRecord

        var report RecoveryReport

        for {

                hdr,
                payload,
                err :=
                        wal.ReadRecord(
                                f,
                        )

                if err == io.EOF {
                        break
                }

                if err != nil {
                        return nil,
                                RecoveryReport{},
                                err
                }

                out =
                        append(
                                out,
                                RecoveryRecord{
                                        Header: hdr,
                                        Payload: payload,
                                },
                        )

                report.Records++
                report.LastLSN = hdr.LSN
        }

        return out,
                report,
                nil
}
