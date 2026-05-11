SELECT COUNT(DISTINCT workshop_speakers.id) AS total
FROM workshop_speakers
         INNER JOIN workshops ON workshop_speakers.workshop_id = workshops.id
         INNER JOIN people speaker_people ON workshop_speakers.speaker_id = speaker_people.id
         INNER JOIN users creator_users ON workshop_speakers.created_by = creator_users.id
WHERE IF(? IS NULL, TRUE, workshop_speakers.workshop_id = TRIM(?))
  AND IF(? IS NULL, TRUE, workshop_speakers.speaker_id = TRIM(?))
  AND workshop_speakers.deleted_at IS NULL;
