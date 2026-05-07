SELECT COUNT(DISTINCT registrations.id) AS total
FROM registrations registrations
         INNER JOIN sessions sessions ON registrations.session_id = sessions.id
         INNER JOIN workshops workshops ON sessions.workshop_id = workshops.id
         INNER JOIN people beneficiaries
                    ON registrations.beneficiary_id = beneficiaries.id
         LEFT JOIN users beneficiaries_users
                   ON beneficiaries.user_id = beneficiaries_users.id
         LEFT JOIN user_types beneficiaries_user_types
                   ON beneficiaries_user_types.id = beneficiaries_users.type_id
         INNER JOIN document_types beneficiaries_document_types
                    ON beneficiaries.type_document_id = beneficiaries_document_types.id
         INNER JOIN users creator_users
                    ON registrations.created_by = creator_users.id
         INNER JOIN user_types creator_user_types
                    ON creator_user_types.id = creator_users.type_id
         INNER JOIN people creators
                    ON creators.user_id = creator_users.id
         INNER JOIN document_types creators_document_types
                    ON creators.type_document_id = creators_document_types.id
WHERE IF(? IS NULL, TRUE, DATE(registrations.created_at) BETWEEN ? AND ?)
  AND IF(? IS NULL, TRUE, registrations.created_by = TRIM(?))
  AND registrations.deleted_at IS NULL;