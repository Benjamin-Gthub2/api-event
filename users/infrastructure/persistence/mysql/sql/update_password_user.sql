UPDATE core_users
SET core_users.password_hash = TRIM(?)
WHERE id = ?
  AND deleted_at IS NULL;
