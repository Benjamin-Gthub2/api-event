SELECT roles.id                               AS role_id,
       roles.id                               AS role_id,
       roles.name                             AS role_name,
       roles.description                      AS role_description,
       roles.enable                           AS role_enable,
       roles.created_at                       AS role_created_at,
       policies.id                            AS policy_id,
       policies.name                          AS policy_name,
       policies.description                   AS policy_description,
       policies.level                         AS policy_level,
       policies.enable                        AS policy_enable,
       policies.created_at                    AS policy_created_at,
       modules.id                             AS module_id,
       modules.name                           AS module_name,
       modules.description                    AS module_description,
       modules.code                           AS module_code,
       modules.icon                           AS module_icon,
       modules.position                       AS module_position,
       modules.created_at                     AS module_created_at,
       users.id                               AS user_id,
       user_roles.id                          AS user_role_id,
       people.id                              AS person_id,
       people.document                        AS person_document,
       people.names                           AS person_names,
       people.surname                         AS person_surname,
       people.last_name                       AS person_last_name,
       people.phone                           AS person_phone,
       people.email                           AS person_email,
       people.gender                          AS person_gender,
       people.enable                          AS person_enable,
       people.created_at                      AS person_created_at,
       document_types.id                      AS document_type_id,
       document_types.number                  AS document_type_number,
       document_types.description             AS document_type_description,
       document_types.abbreviated_description AS document_type_abbreviated_description
FROM core_policies policies
         INNER JOIN core_role_policies role_policies ON policies.id = role_policies.policy_id
         INNER JOIN core_modules modules ON policies.module_id = modules.id
         INNER JOIN core_roles roles ON role_policies.role_id = roles.id
         LEFT JOIN core_user_roles user_roles ON (roles.id = user_roles.role_id AND user_roles.deleted_at IS NULL)
         LEFT JOIN core_users users ON (user_roles.user_id = users.id AND users.deleted_at IS NULL)
         LEFT JOIN hr_people people ON (users.id = people.user_id AND people.deleted_at IS NULL)
         LEFT JOIN core_document_types document_types ON people.type_document_id = document_types.id
WHERE policies.merchant_id = ?
  AND policies.store_id IS NULL
  AND policies.deleted_at IS NULL
  AND role_policies.deleted_at IS NULL
  AND user_roles.deleted_at IS NULL
  AND users.deleted_at IS NULL
  AND people.deleted_at IS NULL;
