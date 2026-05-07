SELECT users.id                               AS user_id,
       users.username                         AS user_name,
       users.created_at                       AS user_created_at,
       users.type_id                          AS user_type_id,
       users.description                      AS user_type_description,
       users.code                             AS user_type_code,
       user_roles.id                          AS user_role_id,
       roles.id                               AS role_id,
       roles.name                             AS role_name,
       roles.description                      AS role_description,
       roles.enable                           AS role_enable,
       roles.created_at                       AS role_created_at,
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
       document_types.abbreviated_description AS document_type_abbreviated_description,
       document_types.enable                  AS document_type_enable,
       document_types.created_at              AS document_type_created_at
FROM (SELECT users.id,
             MAX(users.created_at) AS max_created_at,
             users.username,
             users.created_at,
             types.id              AS type_id,
             types.description,
             types.code
      FROM (SELECT * FROM users WHERE deleted_at IS NULL) users
               INNER JOIN user_types types ON users.type_id = types.id
               LEFT JOIN user_roles user_roles ON user_roles.user_id = users.id AND user_roles.deleted_at IS NULL
               LEFT JOIN roles roles ON user_roles.role_id = roles.id
      WHERE IF(? IS NULL, TRUE, types.id = TRIM(?))
        AND IF(? IS NULL, TRUE, users.username LIKE CONCAT('%', TRIM(?), '%'))
        AND IF(? = '', TRUE, FIND_IN_SET(roles.id, TRIM(?)))
      GROUP BY users.id
      ORDER BY max_created_at DESC
      LIMIT ? OFFSET ?) users
         LEFT JOIN user_roles user_roles ON user_roles.user_id = users.id AND user_roles.deleted_at IS NULL
         LEFT JOIN roles roles ON user_roles.role_id = roles.id
         LEFT JOIN people people ON users.id = people.user_id
         LEFT JOIN document_types document_types ON people.type_document_id = document_types.id;