SELECT COUNT(DISTINCT attendances.id) AS total
FROM attendances
         INNER JOIN users creator_users ON attendances.created_by = creator_users.id
WHERE IF(? IS NULL, TRUE, DATE(attendances.created_at) BETWEEN ? AND ?)
  AND attendances.deleted_at IS NULL;
