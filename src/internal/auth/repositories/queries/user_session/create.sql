insert into user_session (user_id, cookie_value, expires_at, created_at)
values ($1, $2, $3, $4)
returning id