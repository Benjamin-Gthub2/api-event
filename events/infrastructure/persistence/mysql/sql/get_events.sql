SELECT events.id                             AS events_id,
       events.name                           AS events_name,
       events.description                    AS events_description,
       events.code                           AS events_code,
       events.enable                         AS events_enable,
       events.created_at                     AS events_created_at
FROM events events
WHERE events.deleted_at IS NULL
  AND IF(? IS NULL, TRUE, events.enable = ?)
  AND IF(? IS NULL, TRUE, events.name COLLATE utf8mb4_general_ci LIKE CONCAT('%', ?, '%') OR events.description COLLATE utf8mb4_general_ci LIKE CONCAT('%', ?, '%'))
ORDER BY events.created_at DESC
LIMIT ? OFFSET ?;