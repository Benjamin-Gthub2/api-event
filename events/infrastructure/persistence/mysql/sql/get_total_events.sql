SELECT COUNT(*) AS total
FROM events events
WHERE events.deleted_at IS NULL
  AND IF(? IS NULL, TRUE, events.enable = ?)
  AND IF(? IS NULL, TRUE, events.name COLLATE utf8mb4_general_ci LIKE CONCAT('%', ?, '%') OR events.description COLLATE utf8mb4_general_ci LIKE CONCAT('%', ?, '%'));
