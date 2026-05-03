SELECT COUNT(DISTINCT materials_issued.id) AS total
FROM materials_issued
         INNER JOIN users creator_users ON materials_issued.created_by = creator_users.id
WHERE IF(? IS NULL, TRUE, DATE(materials_issued.created_at) BETWEEN ? AND ?)
  AND materials_issued.deleted_at IS NULL;
