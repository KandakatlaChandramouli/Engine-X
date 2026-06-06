package txn

import (
        "bytes"
        "encoding/binary"
)

type UpdateRecord struct {
        TxID   uint64
        PageID uint64
        Offset uint16

        Before []byte
        After  []byte
}

func EncodeUpdateRecord(
        r UpdateRecord,
) []byte {

        var buf bytes.Buffer

        binary.Write(
                &buf,
                binary.LittleEndian,
                r.TxID,
        )

        binary.Write(
                &buf,
                binary.LittleEndian,
                r.PageID,
        )

        binary.Write(
                &buf,
                binary.LittleEndian,
                r.Offset,
        )

        binary.Write(
                &buf,
                binary.LittleEndian,
                uint32(len(r.Before)),
        )

        buf.Write(
                r.Before,
        )

        binary.Write(
                &buf,
                binary.LittleEndian,
                uint32(len(r.After)),
        )

        buf.Write(
                r.After,
        )

        return buf.Bytes()
}

func DecodeUpdateRecord(
        data []byte,
) (
        UpdateRecord,
        error,
) {

        var r UpdateRecord

        reader :=
                bytes.NewReader(
                        data,
                )

        err :=
                binary.Read(
                        reader,
                        binary.LittleEndian,
                        &r.TxID,
                )

        if err != nil {
                return r,
                        err
        }

        err =
                binary.Read(
                        reader,
                        binary.LittleEndian,
                        &r.PageID,
                )

        if err != nil {
                return r,
                        err
        }

        err =
                binary.Read(
                        reader,
                        binary.LittleEndian,
                        &r.Offset,
                )

        if err != nil {
                return r,
                        err
        }

        var beforeLen uint32

        err =
                binary.Read(
                        reader,
                        binary.LittleEndian,
                        &beforeLen,
                )

        if err != nil {
                return r,
                        err
        }

        r.Before =
                make(
                        []byte,
                        beforeLen,
                )

        _, err =
                reader.Read(
                        r.Before,
                )

        if err != nil {
                return r,
                        err
        }

        var afterLen uint32

        err =
                binary.Read(
                        reader,
                        binary.LittleEndian,
                        &afterLen,
                )

        if err != nil {
                return r,
                        err
        }

        r.After =
                make(
                        []byte,
                        afterLen,
                )

        _, err =
                reader.Read(
                        r.After,
                )

        return r,
                err
}
