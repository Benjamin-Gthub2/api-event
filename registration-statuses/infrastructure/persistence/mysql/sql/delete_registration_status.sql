UPDATE registration_statuses
SET deleted_at = ?
WHERE id = ?
  AND deleted_at IS NULL;
