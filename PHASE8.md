Phase 8 - Root Promotion

Goal
----
Create a new root when the existing root splits.

Flow
----
Root Leaf Full
    ->
LeafSplit
    ->
Create Internal Root
    ->
Insert Separator Key
    ->
Point To Left/Right Children

Result
------
Tree height becomes 2.

Requirements
------------
- Preserve existing root page.
- Allocate new root page.
- Update meta root page id.
- Support future recursive promotion.

Future
------
Phase 9:
Internal page split propagation.

Phase 10:
Recursive B+Tree growth.
