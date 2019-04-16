package cas

import (
	"github.com/vmware/terraform-provider-cas/sdk"

	"github.com/hashicorp/terraform/helper/schema"
)

// nicsSchema returns the schema to use for the nics property
func nicsSchema(isRequired bool) *schema.Schema {
	return &schema.Schema{
		Type:     schema.TypeList,
		Required: isRequired,
		Optional: !isRequired,
		Elem: &schema.Resource{
			Schema: map[string]*schema.Schema{
				"name": &schema.Schema{
					Type:     schema.TypeString,
					Optional: true,
				},
				"description": &schema.Schema{
					Type:     schema.TypeString,
					Optional: true,
				},
				"device_index": &schema.Schema{
					Type:     schema.TypeInt,
					Optional: true,
				},
				"network_id": &schema.Schema{
					Type:     schema.TypeString,
					Required: true,
				},
				"addresses": &schema.Schema{
					Type:     schema.TypeList,
					Optional: true,
					Elem: &schema.Schema{
						Type: schema.TypeString,
					},
				},
				"security_group_ids": &schema.Schema{
					Type:     schema.TypeList,
					Optional: true,
					Elem: &schema.Schema{
						Type: schema.TypeString,
					},
				},
				"custom_properties": &schema.Schema{
					Type:     schema.TypeMap,
					Optional: true,
				},
			},
		},
	}
}

func expandNics(configNics []interface{}) []tango.Nic {
	nics := make([]tango.Nic, 0, len(configNics))

	for _, configNic := range configNics {
		nicMap := configNic.(map[string]interface{})

		nic := tango.Nic{
			NetworkID: nicMap["network_id"].(string),
		}

		if v, ok := nicMap["name"].(string); ok && v != "" {
			nic.Name = v
		}

		if v, ok := nicMap["description"].(string); ok && v != "" {
			nic.Description = v
		}

		if v, ok := nicMap["device_index"].(int); ok && v != 0 {
			nic.DeviceIndex = v
		}

		if v, ok := nicMap["addresses"].([]interface{}); ok && len(v) != 0 {
			addresses := make([]string, 0)

			for _, value := range v {
				addresses = append(addresses, value.(string))
			}

			nic.Addresses = addresses
		}

		if v, ok := nicMap["security_group_ids"].([]interface{}); ok && len(v) != 0 {
			securityGroupIds := make([]string, 0)

			for _, value := range v {
				securityGroupIds = append(securityGroupIds, value.(string))
			}

			nic.SecurityGroupIDs = securityGroupIds
		}

		nic.CustomProperties = expandCustomProperties(nicMap["custom_properties"].(map[string]interface{}))

		nics = append(nics, nic)
	}

	return nics
}