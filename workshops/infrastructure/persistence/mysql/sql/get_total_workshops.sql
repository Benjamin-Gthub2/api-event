SELECT COUNT(DISTINCT workshops.id) AS total
FROM workshops workshops
         INNER JOIN workshop_types workshop_types ON workshops.type_id = workshop_types.id
         INNER JOIN events events ON workshops.event_id = events.id
         INNER JOIN users creator_users ON workshops.created_by = creator_users.id
WHERE IF(? IS NULL, TRUE, workshops.event_id = TRIM(?))
  AND IF(? IS NULL, TRUE, workshops.type_id = TRIM(?))
  AND IF(? IS NULL, TRUE, workshops.name COLLATE utf8mb4_general_ci LIKE CONCAT('%', TRIM(?), '%') OR
                          workshops.shortname COLLATE utf8mb4_general_ci LIKE CONCAT('%', TRIM(?), '%') OR
                          workshops.code COLLATE utf8mb4_general_ci LIKE CONCAT('%', TRIM(?), '%') OR
                          workshops.place COLLATE utf8mb4_general_ci LIKE CONCAT('%', TRIM(?), '%'))
  AND workshops.deleted_at IS NULL;
