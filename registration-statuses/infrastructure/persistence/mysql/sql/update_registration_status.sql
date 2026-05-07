UPDATE registration_statuses
SET code        = ?,
    description = ?,
    position    = ?,
    enable      = ?
WHERE id = ?
  AND deleted_at IS NULL;
