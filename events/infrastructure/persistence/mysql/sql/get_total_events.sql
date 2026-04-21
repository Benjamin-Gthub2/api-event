SELECT COUNT(*) AS total
FROM core_merchants merchants
WHERE merchants.deleted_at IS NULL
  AND IF(? IS NULL, TRUE, merchants.enable = ?)
  AND IF(? IS NULL, TRUE, merchants.name LIKE CONCAT('%', ?, '%') OR merchants.description LIKE CONCAT('%', ?, '%') OR
                          merchants.document LIKE CONCAT(?, '%'));
