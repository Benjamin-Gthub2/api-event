UPDATE workshops
SET deleted_at = ?,
    deleted_by = ?
WHERE id = ?
  AND deleted_at IS NULL;
