SELECT COUNT(*) AS total
FROM events events
WHERE events.deleted_at IS NULL
  AND IF(? IS NULL, TRUE, events.enable = ?)
  AND IF(? IS NULL, TRUE, events.name LIKE CONCAT('%', ?, '%') OR events.description LIKE CONCAT('%', ?, '%'));
