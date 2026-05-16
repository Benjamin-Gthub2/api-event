SELECT workshop_id
FROM attendances
WHERE id = ?
  AND deleted_at IS NULL;
