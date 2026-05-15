SELECT COUNT(*)
FROM attendances
WHERE workshop_id = ?
  AND beneficiary_id = ?
  AND deleted_at IS NULL;
