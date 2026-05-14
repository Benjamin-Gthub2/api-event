SELECT COUNT(DISTINCT attendances.id) AS total
FROM attendances
         INNER JOIN workshops workshops ON attendances.workshop_id = workshops.id
         INNER JOIN workshop_types workshop_types ON workshops.type_id = workshop_types.id
         INNER JOIN events events ON workshops.event_id = events.id
         INNER JOIN people beneficiaries ON attendances.beneficiary_id = beneficiaries.id
         INNER JOIN document_types document_types ON beneficiaries.type_document_id = document_types.id
         INNER JOIN users creator_users ON attendances.created_by = creator_users.id
WHERE IF(? IS NULL, TRUE, events.id = TRIM(?))
  AND IF(? IS NULL, TRUE, workshops.id = TRIM(?))
  AND IF(? IS NULL, TRUE, beneficiaries.id = TRIM(?))
  AND IF(? IS NULL, TRUE, DATE(attendances.created_at) BETWEEN ? AND ?)
  AND attendances.deleted_at IS NULL
  AND IF(? IS NULL, TRUE, beneficiaries.names COLLATE utf8mb4_general_ci LIKE CONCAT('%', TRIM(?), '%') OR
                          beneficiaries.surname COLLATE utf8mb4_general_ci LIKE CONCAT('%', TRIM(?), '%') OR
                          beneficiaries.last_name COLLATE utf8mb4_general_ci LIKE CONCAT('%', TRIM(?), '%') OR
                          beneficiaries.document COLLATE utf8mb4_general_ci LIKE CONCAT('%', TRIM(?), '%'));
