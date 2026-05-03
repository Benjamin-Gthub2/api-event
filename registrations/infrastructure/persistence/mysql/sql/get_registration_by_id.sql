SELECT registrations.id                                     AS registration_id,
       registrations.created_at                             AS registration_created_at,
       statuses.id                                          AS status_id,
       statuses.code                                        AS status_code,
       statuses.description                                 AS status_description,
       statuses.position                                    AS status_position,
       statuses.enable                                      AS status_enable,
       statuses.created_at                                  AS status_created_at,
       sessions.id                                          AS session_id,
       sessions.start_date                                  AS session_start_date,
       sessions.end_date                                    AS session_end_date,
       sessions.created_at                                  AS session_created_at,
       workshops.id                                         AS workshop_id,
       workshops.name                                       AS workshop_name,
       beneficiaries.id                                     AS beneficiary_id,
       beneficiaries.document                               AS beneficiary_document,
       beneficiaries.names                                  AS beneficiary_names,
       beneficiaries.surname                                AS beneficiary_surname,
       beneficiaries.last_name                              AS beneficiary_last_name,
       beneficiaries_users.id                               AS beneficiary_user_id,
       beneficiaries_users.username                         AS beneficiary_username,
       beneficiaries_user_types.id                          AS beneficiary_user_type_id,
       beneficiaries_user_types.description                 AS beneficiary_user_type_description,
       beneficiaries_user_types.code                        AS beneficiary_user_type_code,
       beneficiaries_user_types.created_at                  AS beneficiary_user_type_created_at,
       beneficiaries_document_types.id                      AS beneficiary_document_type_id,
       beneficiaries_document_types.description             AS beneficiary_document_type_description,
       beneficiaries_document_types.abbreviated_description AS beneficiary_document_type_abbreviated_description,
       beneficiaries_document_types.enable                  AS beneficiary_document_type_enable,
       creators.id                                          AS creator_id,
       creators.document                                    AS creator_document,
       creators.names                                       AS creator_names,
       creators.surname                                     AS creator_surname,
       creators.last_name                                   AS creator_last_name,
       creator_users.id                                     AS creator_user_id,
       creator_users.username                               AS creator_username,
       creator_user_types.id                                AS creator_user_type_id,
       creator_user_types.description                       AS creator_user_type_description,
       creator_user_types.code                              AS creator_user_type_code,
       creator_user_types.created_at                        AS creator_user_type_created_at,
       creators_document_types.id                           AS creator_document_type_id,
       creators_document_types.description                  AS creator_document_type_description,
       creators_document_types.abbreviated_description      AS creator_document_type_abbreviated_description,
       creators_document_types.enable                       AS creator_document_type_enable
FROM registrations registrations
         INNER JOIN registration_statuses statuses
                    ON registrations.status_id = statuses.id
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
WHERE registrations.id = ?
  AND registrations.deleted_at IS NULL;