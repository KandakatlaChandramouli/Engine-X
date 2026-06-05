Phase 9 - Recursive Split Propagation

Goal
----
Allow splits to propagate upward through the tree.

Current
-------
Leaf split:
    left
    right
    separator

Root promotion:
    supported

Problem
-------
Parent pages can become full.

Required
--------
Insert separator into parent.

If parent full:
    split parent

Promote middle separator upward.

Repeat until:

    parent exists

or

    new root created

Flow
----
Leaf split
    ->
Parent insert
    ->
Parent split
    ->
Promote separator
    ->
Repeat

Output
------
InternalSplit()
InsertIntoParent()
PropagateSplit()

Result
------
Unlimited tree growth.

Future
------
Phase 10:
Tree search from root.

Phase 11:
Full insert path.
