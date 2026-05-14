SELECT COUNT(DISTINCT workshop_types.id) AS total
FROM workshop_types workshop_types
         INNER JOIN users creator_users ON workshop_types.created_by = creator_users.id
WHERE workshop_types.deleted_at IS NULL
  AND IF(? IS NULL, TRUE, workshop_types.description COLLATE utf8mb4_general_ci LIKE CONCAT('%', TRIM(?), '%') OR
                          workshop_types.code COLLATE utf8mb4_general_ci LIKE CONCAT('%', TRIM(?), '%'));
