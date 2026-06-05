Phase 7 - Internal Pages

Goals
-----
Store separator keys.
Store child page IDs.
Support binary search over separators.
Support insertion after leaf split.

Internal Layout
---------------
slot:
    key
    childPageID

Search
------
key < separator -> left child
key >= separator -> right child

Output
------
InternalSearch()
InternalInsert()

Future
------
Phase 8:
Root promotion.

Phase 9:
Recursive split propagation.

