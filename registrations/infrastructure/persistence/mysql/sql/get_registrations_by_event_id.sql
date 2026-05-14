SELECT registrations.id                                     AS registration_by_event_id,
       registrations.send_qr                                AS registration_by_event_send_qr,
       registrations.send_certificate                       AS registration_by_event_send_certificate,
       registrations.created_at                             AS registration_by_event_created_at,
       statuses.id                                          AS status_by_event_id,
       statuses.code                                        AS status_by_event_code,
       statuses.description                                 AS status_by_event_description,
       statuses.position                                    AS status_by_event_position,
       statuses.enable                                      AS status_by_event_enable,
       statuses.created_at                                  AS status_by_event_created_at,
       events.id                                            AS event_by_event_id,
       events.name                                          AS event_by_event_name,
       events.description                                   AS event_by_event_description,
       events.created_at                                    AS event_by_event_created_at,
       beneficiaries.id                                     AS beneficiary_by_event_id,
       beneficiaries.document                               AS beneficiary_by_event_document,
       beneficiaries.names                                  AS beneficiary_by_event_names,
       beneficiaries.surname                                AS beneficiary_by_event_surname,
       beneficiaries.last_name                              AS beneficiary_by_event_last_name,
       beneficiaries.phone                                  AS beneficiary_by_event_phone,
       beneficiaries_document_types.id                      AS beneficiary_document_type_by_event_id,
       beneficiaries_document_types.description             AS beneficiary_document_type_by_event_description,
       beneficiaries_document_types.abbreviated_description AS beneficiary_document_type_by_event_abbreviated_description,
       beneficiaries_document_types.enable                  AS beneficiary_document_type_by_event_enable
FROM registrations registrations
         INNER JOIN registration_statuses statuses
                    ON registrations.status_id = statuses.id
         INNER JOIN events events ON registrations.event_id = events.id
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
WHERE registrations.deleted_at IS NULL
  AND beneficiaries.deleted_at IS NULL
  AND beneficiaries.user_id IS NULL
  AND registrations.event_id = ?
  AND IF(? IS NULL, TRUE, beneficiaries.names COLLATE utf8mb4_general_ci LIKE CONCAT('%', TRIM(?), '%') OR
                          beneficiaries.surname COLLATE utf8mb4_general_ci LIKE CONCAT('%', TRIM(?), '%') OR
                          beneficiaries.last_name COLLATE utf8mb4_general_ci LIKE CONCAT('%', TRIM(?), '%') OR
                          beneficiaries.document COLLATE utf8mb4_general_ci LIKE CONCAT('%', TRIM(?), '%'))
ORDER BY beneficiaries.names, beneficiaries.surname, beneficiaries.last_name;