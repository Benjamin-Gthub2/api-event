UPDATE workshops
SET type_id    = ?,
    name       = ?,
    shortname  = ?,
    code       = ?,
    capacity   = ?,
    start_date = ?,
    end_date   = ?,
    place      = ?
WHERE id = ?
  AND deleted_at IS NULL;
