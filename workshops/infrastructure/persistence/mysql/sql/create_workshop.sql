INSERT INTO workshops(id,
                      type_id,
                      name,
                      shortname,
                      code,
                      capacity,
                      total_pres,
                      start_date,
                      end_date,
                      place,
                      event_id,
                      created_by,
                      created_at)
VALUES (?, ?, ?, ?, ?, ?, 0, ?, ?, ?, ?, ?, ?);
