# Engine-X Architecture

Reserved Pages

0 -> Meta A
1 -> Meta B
2 -> Freelist

User Pages Begin At

3

Allocator Strategy

1. Reuse freelist pages first.
2. Extend file only when freelist empty.

Crash Recovery

1. WAL replay.
2. Read newest valid meta page.
3. Restore root page.
4. Restore freelist.

Future Phases

Phase 3
- Meta page
- Freelist
- Allocator

Phase 4
- B+Tree Leaf Nodes
- Binary Search
- Ordered Inserts

Phase 5
- Internal Nodes
- Splits
- Root Promotion

Phase 6
- Transactions
- MVCC
- Snapshots
