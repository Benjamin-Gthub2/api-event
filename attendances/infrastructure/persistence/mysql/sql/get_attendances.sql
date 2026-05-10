SELECT attendances.id                         AS attendance_id,
       attendances.created_at                 AS attendance_created_at,
       workshops.id                           AS workshop_id,
       workshops.name                         AS workshop_name,
       workshops.shortname                    AS workshop_shortname,
       workshops.code                         AS workshop_code,
       workshops.capacity                     AS workshop_capacity,
       workshops.created_at                   AS workshop_created_at,
       workshop_types.id                      AS workshop_type_id,
       workshop_types.code                    AS workshop_type_code,
       workshop_types.description             AS workshop_type_description,
       events.id                              AS events_id,
       events.name                            AS events_name,
       events.code                            AS events_code,
       beneficiaries.id                       AS beneficiary_id,
       beneficiaries.document                 AS beneficiary_document,
       beneficiaries.names                    AS beneficiary_names,
       beneficiaries.surname                  AS beneficiary_surname,
       beneficiaries.last_name                AS beneficiary_last_name,
       document_types.id                      AS beneficiary_document_type_id,
       document_types.description             AS beneficiary_document_type_description,
       document_types.abbreviated_description AS beneficiary_document_type_abbreviated_description,
       document_types.enable                  AS beneficiary_document_type_enable,
       creator_users.id                       AS created_by_id,
       creator_users.username                 AS created_by_username
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
LIMIT ? OFFSET ?;
