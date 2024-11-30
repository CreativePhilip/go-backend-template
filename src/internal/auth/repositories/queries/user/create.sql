insert into app_user (first_name, last_name, email, password, is_staff, last_logged_in, created_at)
values ($1, $2, $3, $4, $5, $6, $7)
returning id, last_logged_in, created_at