Phase 8 - Root Promotion

Goal
----
Convert first leaf split into a real tree.

Flow
----
Leaf full
    ->
Leaf split
    ->
Allocate right leaf
    ->
Create root
    ->
Root.LeftChild = left leaf
    ->
Insert separator -> right leaf
    ->
Meta.RootPageID = root

Result
------
        Root
       /    \
   LeafA   LeafB

Output
------
PromoteRoot()

Future
------
Phase 9:
Recursive split propagation.

Phase 10:
Full B+Tree insert path.
