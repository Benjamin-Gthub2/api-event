UPDATE sessions
SET start_date = ?,
    end_date   = ?
WHERE id = ?
  AND deleted_at IS NULL;
