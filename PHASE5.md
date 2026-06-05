Phase 5 - Ordered Leaf Insert

Goal:
- Maintain sorted slot order.
- Preserve existing payload layout.
- Keep Insert() generic.
- Implement ordering only in leaf layer.

Constraints:
- No page splits yet.
- No internal nodes yet.
- No duplicate update path yet.
- No MVCC.
- No transactions.

Success Criteria:
- Keys remain sorted regardless of insertion order.
- LeafSearch() remains O(log n).
- Existing Get() remains unchanged.
