UPDATE workshops
SET type_id   = ?,
    name      = ?,
    shortname = ?,
    code      = ?,
    capacity  =?
WHERE id = ?
  AND deleted_at IS NULL;
