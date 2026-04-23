UPDATE event_types
SET code        = ?,
    description = ?,
    enable      = ?
WHERE id = ?
  AND deleted_at IS NULL;
