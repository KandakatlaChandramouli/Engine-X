Phase 7.5 - Internal Node Layout Fix

Problem
-------
Current internal entries store:

    key -> child

This cannot represent:

    left child
    separator
    right child

Required
--------
Internal node must support:

    children = keys + 1

Layout
------
LeftChild

Slot:
    separator key
    right child

Search
------
key < first separator
    -> LeftChild

key >= separator
    -> slot child

Future
------
Phase 8:
Root promotion

Phase 9:
Recursive split propagation
