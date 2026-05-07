SELECT COUNT(*) AS total
FROM people people
         LEFT JOIN users users ON people.user_id = users.id
         INNER JOIN document_types document_type ON document_type.id = people.type_document_id
WHERE people.deleted_at IS NULL
  AND users.deleted_at IS NULL
  AND document_type.deleted_at IS NULL
  AND IF(? IS NULL, TRUE, people.names LIKE CONCAT('%', TRIM(?), '%') OR
                          people.surname LIKE CONCAT('%', TRIM(?), '%') OR
                          people.last_name LIKE CONCAT('%', TRIM(?), '%'))
  AND IF(? IS NULL, TRUE, document_type.id LIKE CONCAT('%', TRIM(?), '%'))
  AND IF(? IS NULL, TRUE, people.document LIKE CONCAT(TRIM(?), '%'));
