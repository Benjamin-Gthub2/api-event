SELECT COUNT(DISTINCT event_types.id) AS total
FROM event_types event_types
         INNER JOIN users creator_users ON event_types.created_by = creator_users.id
WHERE event_types.deleted_at IS NULL
  AND IF(? IS NULL, TRUE, event_types.description COLLATE utf8mb4_general_ci LIKE CONCAT('%', TRIM(?), '%') OR
                          event_types.code COLLATE utf8mb4_general_ci LIKE CONCAT('%', TRIM(?), '%'));
