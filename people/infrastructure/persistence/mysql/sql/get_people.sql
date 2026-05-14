SELECT people.id                             AS people_id,
       people.document                       AS people_document,
       people.names                          AS people_name,
       people.surname                        AS people_surname,
       people.last_name                      AS people_last_name,
       people.phone                          AS people_phone,
       people.email                          AS people_email,
       people.gender                         AS people_gender,
       people.enable                         AS people_enable,
       people.created_at                     AS people_created_at,
       users.id                              AS user_id,
       users.username                        AS user_username,
       users.created_at                      AS user_created_at,
       document_type.id                      AS document_type_id,
       document_type.description             AS document_type_description,
       document_type.abbreviated_description AS document_type_abbreviated_description
FROM people people
         LEFT JOIN users users ON people.user_id = users.id
         INNER JOIN document_types document_type ON document_type.id = people.type_document_id
WHERE people.deleted_at IS NULL
  AND users.deleted_at IS NULL
  AND document_type.deleted_at IS NULL
  AND IF(? IS NULL, TRUE, people.names COLLATE utf8mb4_general_ci LIKE CONCAT('%', TRIM(?), '%') OR
                          people.surname COLLATE utf8mb4_general_ci LIKE CONCAT('%', TRIM(?), '%') OR
                          people.last_name COLLATE utf8mb4_general_ci LIKE CONCAT('%', TRIM(?), '%'))
  AND IF(? IS NULL, TRUE, document_type.id LIKE CONCAT('%', TRIM(?), '%'))
  AND IF(? IS NULL, TRUE, people.document LIKE CONCAT(TRIM(?), '%'))
  AND IF(? IS NULL, TRUE, people.names COLLATE utf8mb4_general_ci LIKE CONCAT('%', TRIM(?), '%') OR
                          people.surname COLLATE utf8mb4_general_ci LIKE CONCAT('%', TRIM(?), '%') OR
                          people.last_name COLLATE utf8mb4_general_ci LIKE CONCAT('%', TRIM(?), '%') OR
                          people.document COLLATE utf8mb4_general_ci LIKE CONCAT('%', TRIM(?), '%'))
ORDER BY people.names, people.surname, people.last_name
LIMIT ? OFFSET ?;
