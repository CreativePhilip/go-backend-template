select id, name, slug, organization_member.created_at as created_at
from organization_member
         inner join organization o on organization_member.organization_id = o.id
where user_id = $1