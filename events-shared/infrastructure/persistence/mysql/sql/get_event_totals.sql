SELECT events.id         AS event_id,
       events.total_reg  AS event_total_reg,
       events.total_pay  AS event_total_pay,
       events.total_pres AS event_total_pres
FROM events
WHERE events.deleted_at IS NULL;