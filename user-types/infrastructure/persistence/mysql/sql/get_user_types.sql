SELECT id,
       description,
       code,
       enable,
       created_at
FROM user_types
WHERE deleted_at IS NULL
  AND IF(? IS NULL, TRUE, description COLLATE utf8mb4_general_ci LIKE CONCAT('%', TRIM(?), '%') OR
                          code COLLATE utf8mb4_general_ci LIKE CONCAT('%', TRIM(?), '%'))
ORDER BY created_at
LIMIT ? OFFSET ?;
