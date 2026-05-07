UPDATE registrations
SET status_id = (SELECT id FROM registration_statuses WHERE code = ? AND deleted_at IS NULL LIMIT 1)
WHERE id = ?
  AND deleted_at IS NULL;
