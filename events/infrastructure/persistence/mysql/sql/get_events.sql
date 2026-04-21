SELECT events.id                             AS events_id,
       events.name                           AS events_name,
       events.description                    AS events_description,
       events.code                           AS events_code,
       events.phone                          AS events_phone,
       events.document                       AS events_document,
       events.address                        AS events_address,
       events.industry                       AS events_industry,
       events.enable                         AS events_enable,
       events.created_at                     AS events_created_at,
       events_files.id                       AS events_files_id,
       events_files.name                     AS events_files_name,
       CAST(events_files.weight AS CHAR(25)) AS events_files_weight,
       events_files.url                      AS events_files_url,
       events_files.created_at               AS events_files_created_at
FROM events events
         LEFT JOIN core_event_files events_files ON events.id = events_files.event_id
WHERE events.deleted_at IS NULL
  AND events_files.deleted_at IS NULL
  AND IF(? IS NULL, TRUE, events.enable = ?)
  AND IF(? IS NULL, TRUE, events.name LIKE CONCAT('%', ?, '%') OR events.description LIKE CONCAT('%', ?, '%') OR
                          events.document LIKE CONCAT(?, '%'))
ORDER BY events.name, events_files.created_at DESC
LIMIT ? OFFSET ?;