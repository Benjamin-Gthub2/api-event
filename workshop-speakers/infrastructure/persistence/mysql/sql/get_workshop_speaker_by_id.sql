SELECT workshop_speakers.id                   AS workshop_speaker_id,
       workshop_speakers.degree_abbreviation   AS workshop_speaker_degree_abbreviation,
       workshop_speakers.created_at            AS workshop_speaker_created_at,
       workshops.id                 AS workshop_id,
       workshops.name               AS workshop_name,
       workshops.shortname          AS workshop_shortname,
       speaker_people.id            AS speaker_id,
       speaker_people.names         AS speaker_names,
       speaker_people.surname       AS speaker_surname,
       speaker_people.last_name     AS speaker_last_name,
       speaker_people.document      AS speaker_document,
       creator_users.id             AS created_by_id,
       creator_users.username       AS created_by_username
FROM workshop_speakers
         INNER JOIN workshops ON workshop_speakers.workshop_id = workshops.id
         INNER JOIN people speaker_people ON workshop_speakers.speaker_id = speaker_people.id
         INNER JOIN users creator_users ON workshop_speakers.created_by = creator_users.id
WHERE workshop_speakers.id = ?
  AND workshop_speakers.deleted_at IS NULL;
