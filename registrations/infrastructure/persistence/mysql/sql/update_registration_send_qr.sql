UPDATE registrations
SET send_qr = TRUE
WHERE id = ?
  AND deleted_at IS NULL;
