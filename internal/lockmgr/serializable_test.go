package lockmgr

import "testing"

func TestSerializableReadRead(
        t *testing.T,
) {

        lm := New()

        if !SerializableRead(
                lm,
                1,
                100,
        ) {
                t.Fatal()
        }

        if !SerializableRead(
                lm,
                2,
                100,
        ) {
                t.Fatal()
        }
}

func TestSerializableWriteBlocked(
        t *testing.T,
) {

        lm := New()

        SerializableRead(
                lm,
                1,
                100,
        )

        if SerializableWrite(
                lm,
                2,
                100,
        ) {
                t.Fatal()
        }
}

func TestSerializableWriteExclusive(
        t *testing.T,
) {

        lm := New()

        if !SerializableWrite(
                lm,
                1,
                100,
        ) {
                t.Fatal()
        }

        if SerializableRead(
                lm,
                2,
                100,
        ) {
                t.Fatal()
        }
}
