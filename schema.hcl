schema "public" {}

table "app_user" {
  schema = schema.public

  column "id" {
    null = false
    type = bigserial
  }

  column "email" {
    null = false
    type = varchar(255)
  }

  column "first_name" {
    null = false
    type = varchar(255)
  }

  column "last_name" {
    null = false
    type = varchar(255)
  }

  column "password" {
    null = false
    type = bytea
  }

  column "is_staff" {
    null = false
    type = boolean
    default = false
  }

  column "last_logged_in" {
    null = false
    type = timestamp
  }

  column "created_at" {
    null = false
    type = timestamp
  }

  primary_key {
    columns = [column.id]
  }

  index "app_user_email_unique" {
    unique = true
    columns = [column.email]
  }
}


table "user_session" {
  schema = schema.public

  column "id" {
    null = false
    type = bigserial
  }

  column "user_id" {
    null = false
    type = bigint
  }

  column "cookie_value" {
    null = false
    type = varchar(511)
  }

  column "expires_at" {
    null = false
    type = timestamp
  }

  column "created_at" {
    null = false
    type = timestamp
  }

  primary_key {
    columns = [column.id]
  }

  foreign_key "user_id" {
    columns = [column.user_id]
    ref_columns = [table.app_user.column.id]
    on_update = NO_ACTION
    on_delete = CASCADE
  }

  index "idx__cookie_value" {
    columns = [column.cookie_value]
  }

  index "idx__expires_at" {
    columns = [column.expires_at]
  }
}


table "organization" {
  schema = schema.public

  column "id" {
    null = false
    type = bigserial
  }

  column "name" {
    null = false
    type = varchar(255)
  }

  column "slug" {
    null = false
    type = varchar(255)
  }

  column "created_at" {
    null = false
    type = timestamp
  }

  primary_key {
    columns = [column.id]
  }
}

table "organization_member" {
  schema = schema.public

  column "user_id" {
    null = false
    type = bigserial
  }

  column "organization_id" {
    null = false
    type = bigserial
  }

  column "role" {
    null = false
    type = varchar(127)
  }

  column "created_at" {
    null = false
    type = timestamp
  }

  primary_key {
    columns = [column.user_id, column.organization_id]
  }

  foreign_key "user_id" {
    columns = [column.user_id]
    ref_columns = [table.app_user.column.id]
    on_update = NO_ACTION
    on_delete = CASCADE
  }

  foreign_key "organization_id" {
    columns = [column.organization_id]
    ref_columns = [table.organization.column.id]
    on_update = NO_ACTION
    on_delete = CASCADE
  }
}
