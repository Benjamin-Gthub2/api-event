UPDATE materials_issued
SET description = ?
WHERE id = ?
  AND deleted_at IS NULL;
