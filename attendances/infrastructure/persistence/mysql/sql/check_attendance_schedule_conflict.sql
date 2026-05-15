SELECT COUNT(*)
FROM attendances a
         INNER JOIN workshops w ON w.id = a.workshop_id AND w.deleted_at IS NULL
WHERE a.beneficiary_id = ?
  AND a.deleted_at IS NULL
  AND a.workshop_id != ?
  AND w.start_date = (SELECT start_date FROM workshops WHERE id = ? AND deleted_at IS NULL);
