package ovh

import (
	"fmt"
	"sort"
	"strconv"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/ovh/terraform-provider-ovh/ovh/helpers/hashcode"
)

// Allowed networks

func dataSourceDedicatedCloudAllowedNetworks() *schema.Resource {
	return &schema.Resource{
		Read: dataSourceDedicatedCloudAllowedNetworksRead,
		Schema: map[string]*schema.Schema{
			"service_name": {
				Type:     schema.TypeString,
				Required: true,
			},
			"allowed_networks": {
				Type:     schema.TypeList,
				Computed: true,
				Elem: &schema.Schema{
					Type: schema.TypeInt,
				},
			},
		},
	}
}

func dataSourceDedicatedCloudAllowedNetworksRead(d *schema.ResourceData, meta interface{}) error {
	config := meta.(*Config)
	serviceName := d.Get("service_name").(string)
	result := make([]int, 0)

	endpoint := fmt.Sprintf("/dedicatedCloud/%s/allowedNetwork", serviceName)

	err := config.OVHClient.Get(endpoint, &result)
	if err != nil {
		return fmt.Errorf("Error calling GET %s:\n\t %q", endpoint, err)
	}

	sort.Ints(result)
	var stringResults []string
	for _, i := range result {
		stringResults = append(stringResults, strconv.Itoa(i))
	}
	d.SetId(hashcode.Strings(stringResults))
	d.Set("allowed_networks", result)

	return nil
}

// Allowed network

func dataSourceDedicatedCloudAllowedNetwork() *schema.Resource {
	return &schema.Resource{
		Read: dataSourceDedicatedCloudAllowedNetworkRead,
		Schema: map[string]*schema.Schema{
			"service_name": {
				Type:     schema.TypeString,
				Required: true,
			},
			"network_access_id": {
				Type:     schema.TypeInt,
				Required: true,
			},
			"description": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"network": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"state": {
				Type:     schema.TypeString,
				Computed: true,
			},
		},
	}
}

func dataSourceDedicatedCloudAllowedNetworkRead(d *schema.ResourceData, meta interface{}) error {
	config := meta.(*Config)
	serviceName := d.Get("service_name").(string)
	allowedNetworkId := d.Get("network_access_id").(int)
	allowedNetwork := &DedicatedCloudAllowedNetwork{}

	endpoint := fmt.Sprintf("/dedicatedCloud/%s/allowedNetwork/%d", serviceName, allowedNetworkId)

	err := config.OVHClient.Get(
		endpoint,
		&allowedNetwork,
	)
	if err != nil {
		return fmt.Errorf("Error calling GET %s:\n\t %q", endpoint, err)
	}

	d.SetId(fmt.Sprintf("%s/%d", d.Get("service_name"), d.Get("network_access_id").(int)))
	d.Set("description", *allowedNetwork.Description)
	d.Set("network", *allowedNetwork.Network)
	d.Set("network_access_id", d.Get("network_access_id").(int))
	d.Set("state", *allowedNetwork.State)

	return nil
}

// Allowed network tasks

func dataSourceDedicatedCloudAllowedNetworkTasks() *schema.Resource {
	return &schema.Resource{
		Read: dataSourceDedicatedCloudAllowedNetworksRead,
		Schema: map[string]*schema.Schema{
			"service_name": {
				Type:     schema.TypeString,
				Required: true,
			},
			"network_access_id": {
				Type:     schema.TypeInt,
				Required: true,
			},
			"tasks": {
				Type:     schema.TypeList,
				Computed: true,
				Elem: &schema.Schema{
					Type: schema.TypeInt,
				},
			},
		},
	}
}

func dataSourceDedicatedCloudAllowedNetworkTasksRead(d *schema.ResourceData, meta interface{}) error {
	config := meta.(*Config)
	serviceName := d.Get("service_name").(string)
	allowedNetworkId := d.Get("network_access_id").(string)
	result := make([]int, 0)

	endpoint := fmt.Sprintf("/dedicatedCloud/%s/allowedNetwork/%s/task", serviceName, allowedNetworkId)

	err := config.OVHClient.Get(endpoint, &result)
	if err != nil {
		return fmt.Errorf("Error calling GET %s:\n\t %q", endpoint, err)
	}

	sort.Ints(result)
	var stringResults []string
	for _, i := range result {
		stringResults = append(stringResults, strconv.Itoa(i))
	}
	d.SetId(hashcode.Strings(stringResults))
	d.Set("tasks", result)

	return nil
}

// Allowed network task

func dataSourceDedicatedCloudAllowedNetworkTask() *schema.Resource {
	return &schema.Resource{
		Read: dataSourceDedicatedCloudAllowedNetworkTaskRead,
		Schema: map[string]*schema.Schema{
			"service_name": {
				Type:     schema.TypeString,
				Required: true,
			},
			"network_access_id": {
				Type:     schema.TypeInt,
				Required: true,
			},
			"task_id": {
				Type:     schema.TypeInt,
				Required: true,
			},
			"created_by": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"created_from": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"datacenter_id": {
				Type:     schema.TypeInt,
				Computed: true,
			},
			"description": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"end_date": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"execution_date": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"filer_id": {
				Type:     schema.TypeInt,
				Computed: true,
			},
			"host_id": {
				Type:     schema.TypeInt,
				Computed: true,
			},
			"last_modification_date": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"maintenance_date_from": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"maintenance_date_to": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"name": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"network": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"order_id": {
				Type:     schema.TypeInt,
				Computed: true,
			},
			"parent_task_id": {
				Type:     schema.TypeInt,
				Computed: true,
			},
			"progress": {
				Type:     schema.TypeInt,
				Computed: true,
			},
			"state": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"user_id": {
				Type:     schema.TypeInt,
				Computed: true,
			},
			"vlan_id": {
				Type:     schema.TypeInt,
				Computed: true,
			},
		},
	}
}

func dataSourceDedicatedCloudAllowedNetworkTaskRead(d *schema.ResourceData, meta interface{}) error {
	config := meta.(*Config)
	serviceName := d.Get("service_name").(string)
	allowedNetworkId := d.Get("network_access_id").(int)
	taskId := d.Get("task_id").(int)
	task := &DedicatedCloudTask{}

	endpoint := fmt.Sprintf("/dedicatedCloud/%s/allowedNetwork/%d/task/%d", serviceName, allowedNetworkId, taskId)
	err := config.OVHClient.Get(
		endpoint,
		&task,
	)
	if err != nil {
		return fmt.Errorf("Error calling GET %s:\n\t %q", endpoint, err)
	}

	d.SetId(fmt.Sprintf("%s/%d/%d", d.Get("service_name"), d.Get("network_access_id"), d.Get("task_id")))
	d.Set("created_by", task.CreatedBy)
	d.Set("created_from", task.CreatedFrom)
	d.Set("datacenter_id", task.DatacenterId)
	d.Set("description", task.Description)
	d.Set("end_date", task.EndDate)
	d.Set("execution_date", task.ExecutionDate)
	d.Set("filer_id", task.FilerId)
	d.Set("host_id", task.HostId)
	d.Set("last_modification_date", task.LastModificationDate)
	d.Set("maintenance_date_from", task.MaintenanceDateFrom)
	d.Set("maintenance_date_to", task.MaintenanceDateTo)
	d.Set("name", task.Name)
	d.Set("network", task.Network)
	d.Set("network_access_id", task.NetworkAccessId)
	d.Set("order_id", task.OrderId)
	d.Set("parent_task_id", task.ParentTaskId)
	d.Set("progress", task.Progress)
	d.Set("state", task.State)
	d.Set("task_id", task.TaskId)
	d.Set("type", task.Type)
	d.Set("user_id", task.UserId)
	d.Set("vlan_id", task.VlanId)

	return nil
}
