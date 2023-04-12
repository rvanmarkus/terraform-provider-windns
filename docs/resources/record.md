---
# generated by https://github.com/hashicorp/terraform-plugin-docs
page_title: "windns_record Resource - terraform-provider-windns"
subcategory: ""
description: |-
  windns_record manages DNS Records in a Windows DNS Server.
---

# windns_record (Resource)

`windns_record` manages DNS Records in a Windows DNS Server.



<!-- schema generated by tfplugindocs -->
## Schema

### Required

- `name` (String) The name of the dns records.
- `records` (Set of String) A list of records.
- `type` (String) The type of the dns records. (AAAA, A, CNAME, TXT or PTR)
- `zone_name` (String) The zone name for the dns records.

### Read-Only

- `id` (String) The ID of this resource.

