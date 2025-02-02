---
layout: page
title: proxmox_virtual_environment_haresources
parent: Data Sources
subcategory: Virtual Environment
description: |-
  Retrieves the list of High Availability resources.
---

# Data Source: proxmox_virtual_environment_haresources

Retrieves the list of High Availability resources.

## Example Usage

```terraform
// This will fetch the set of all HA resource identifiers.
data "proxmox_virtual_environment_haresources" "example_all" {}

// This will fetch the set of HA resource identifiers that correspond to virtual machines.
data "proxmox_virtual_environment_haresources" "example_vm" {
  type = "vm"
}

output "data_proxmox_virtual_environment_haresources" {
  value = {
    all = data.proxmox_virtual_environment_haresources.example_all.resource_ids
    vms = data.proxmox_virtual_environment_haresources.example_vm.resource_ids
  }
}
```

<!-- schema generated by tfplugindocs -->
## Schema

### Optional

- `type` (String) The type of High Availability resources to fetch (`vm` or `ct`). All resources will be fetched if this option is unset.

### Read-Only

- `id` (String) The ID of this resource.
- `resource_ids` (Set of String) The identifiers of the High Availability resources.
