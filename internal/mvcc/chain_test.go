package mvcc

import "testing"

func TestReadVersion(
        t *testing.T,
) {

        v1 := &Version{
                BeginTS: 1,
                Value: []byte("v1"),
        }

        v2 := &Version{
                BeginTS: 10,
                Value: []byte("v2"),
        }

        head :=
                InsertVersion(
                        v1,
                        v2,
                )

        result :=
                ReadVersion(
                        head,
                        NewSnapshot(15),
                )

        if string(result.Value) != "v2" {
                t.Fatal()
        }
}

func TestOldSnapshot(
        t *testing.T,
) {

        v1 := &Version{
                BeginTS: 1,
                Value: []byte("old"),
        }

        v2 := &Version{
                BeginTS: 10,
                Value: []byte("new"),
        }

        head :=
                InsertVersion(
                        v1,
                        v2,
                )

        result :=
                ReadVersion(
                        head,
                        NewSnapshot(5),
                )

        if string(result.Value) != "old" {
                t.Fatal()
        }
}
