// Generated by PMS #496
package gaussdb

import (
	"context"

	"github.com/hashicorp/go-multierror"
	"github.com/hashicorp/go-uuid"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/tidwall/gjson"

	"github.com/huaweicloud/terraform-provider-huaweicloud/huaweicloud/config"
	"github.com/huaweicloud/terraform-provider-huaweicloud/huaweicloud/helper/httphelper"
	"github.com/huaweicloud/terraform-provider-huaweicloud/huaweicloud/helper/schemas"
	"github.com/huaweicloud/terraform-provider-huaweicloud/huaweicloud/utils"
)

func DataSourceGaussdbOpengaussRestorableInstances() *schema.Resource {
	return &schema.Resource{
		ReadContext: dataSourceGaussdbOpengaussRestorableInstancesRead,

		Schema: map[string]*schema.Schema{
			"region": {
				Type:        schema.TypeString,
				Optional:    true,
				Computed:    true,
				Description: `Specifies the region in which to query the resource. If omitted, the provider-level region will be used.`,
			},
			"source_instance_id": {
				Type:        schema.TypeString,
				Required:    true,
				Description: `Specifies the ID of the GaussDB OpenGauss instance to be restored.`,
			},
			"backup_id": {
				Type:         schema.TypeString,
				Optional:     true,
				AtLeastOneOf: []string{"backup_id", "restore_time"},
				Description:  `Specifies the instance backup ID.`,
			},
			"restore_time": {
				Type:         schema.TypeString,
				Optional:     true,
				AtLeastOneOf: []string{"backup_id", "restore_time"},
				Description:  `Specifies the time point of data restoration in the UNIX timestamp format.`,
			},
			"instances": {
				Type:        schema.TypeList,
				Computed:    true,
				Description: `Indicates the instances that can be used for backups and restorations.`,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"instance_id": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: `Indicates the instance ID.`,
						},
						"instance_name": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: `Indicates the instance name.`,
						},
						"instance_mode": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: `Indicates the instance model.`,
						},
						"volume_type": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: `Indicates the storage type.`,
						},
						"data_volume_size": {
							Type:        schema.TypeFloat,
							Computed:    true,
							Description: `Indicates the storage space, in GB`,
						},
						"version": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: `Indicates the instance version`,
						},
						"mode": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: `Indicates the deployment model.`,
						},
					},
				},
			},
		},
	}
}

type OpengaussRestorableInstancesDSWrapper struct {
	*schemas.ResourceDataWrapper
	Config *config.Config
}

func newOpengaussRestorableInstancesDSWrapper(d *schema.ResourceData, meta interface{}) *OpengaussRestorableInstancesDSWrapper {
	return &OpengaussRestorableInstancesDSWrapper{
		ResourceDataWrapper: schemas.NewSchemaWrapper(d),
		Config:              meta.(*config.Config),
	}
}

func dataSourceGaussdbOpengaussRestorableInstancesRead(_ context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	wrapper := newOpengaussRestorableInstancesDSWrapper(d, meta)
	lisResInsRst, err := wrapper.ListRestorableInstances()
	if err != nil {
		return diag.FromErr(err)
	}

	id, err := uuid.GenerateUUID()
	if err != nil {
		return diag.FromErr(err)
	}
	d.SetId(id)

	err = wrapper.listRestorableInstancesToSchema(lisResInsRst)
	if err != nil {
		return diag.FromErr(err)
	}

	return nil
}

// @API GaussDB GET /v3/{project_id}/restorable-instances
func (w *OpengaussRestorableInstancesDSWrapper) ListRestorableInstances() (*gjson.Result, error) {
	client, err := w.NewClient(w.Config, "opengauss")
	if err != nil {
		return nil, err
	}

	uri := "/v3/{project_id}/restorable-instances"
	params := map[string]any{
		"source_instance_id": w.Get("source_instance_id"),
		"backup_id":          w.Get("backup_id"),
		"restore_time":       w.Get("restore_time"),
	}
	params = utils.RemoveNil(params)
	return httphelper.New(client).
		Method("GET").
		URI(uri).
		Query(params).
		OffsetPager("instances", "offset", "limit", 0).
		Request().
		Result()
}

func (w *OpengaussRestorableInstancesDSWrapper) listRestorableInstancesToSchema(body *gjson.Result) error {
	d := w.ResourceData
	mErr := multierror.Append(nil,
		d.Set("region", w.Config.GetRegion(w.ResourceData)),
		d.Set("instances", schemas.SliceToList(body.Get("instances"),
			func(instances gjson.Result) any {
				return map[string]any{
					"instance_id":      instances.Get("instance_id").Value(),
					"instance_name":    instances.Get("instance_name").Value(),
					"instance_mode":    instances.Get("instance_mode").Value(),
					"volume_type":      instances.Get("volume_type").Value(),
					"data_volume_size": instances.Get("data_volume_size").Value(),
					"version":          instances.Get("version").Value(),
					"mode":             instances.Get("mode").Value(),
				}
			},
		)),
	)
	return mErr.ErrorOrNil()
}
