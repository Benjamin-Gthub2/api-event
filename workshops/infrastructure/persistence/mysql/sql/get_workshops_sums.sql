SELECT workshops.id                          AS workshop_id,
       workshops.name                        AS workshop_name,
       workshops.start_date                  AS workshop_start_date,
       workshops.end_date                    AS workshop_end_date,
       workshops.place                       AS workshop_place,
       workshops.capacity                    AS workshop_capacity,
       workshops.total_pres                  AS workshop_total_pres,
       workshop_speakers.degree_abbreviation AS workshop_speaker_degree_abbreviation,
       speakers.id                           AS speaker_id,
       speakers.names                        AS speaker_name,
       speakers.surname                      AS speaker_surname,
       speakers.last_name                    AS speaker_last_name
FROM workshops workshops
         LEFT JOIN workshop_speakers workshop_speakers ON workshops.id = workshop_speakers.workshop_id
         LEFT JOIN people speakers ON workshop_speakers.speaker_id = speakers.id
WHERE IF(? IS NULL, TRUE, workshops.id = ?)
  AND IF(? IS NULL, TRUE, workshops.name LIKE CONCAT('%', ?, '%') OR workshops.code LIKE CONCAT('%', ?, '%') OR
                          workshops.place LIKE CONCAT('%', ?, '%'))
  AND IF(? IS NULL, TRUE, workshops.start_date >= ?)
  AND IF(? IS NULL, TRUE, workshops.end_date <= ?)
  AND workshops.deleted_at IS NULL
ORDER BY workshops.start_date, workshops.name;