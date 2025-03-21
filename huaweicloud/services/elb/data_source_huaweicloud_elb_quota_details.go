// Generated by PMS #580
package elb

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

func DataSourceElbQuotaDetails() *schema.Resource {
	return &schema.Resource{
		ReadContext: dataSourceElbQuotaDetailsRead,

		Schema: map[string]*schema.Schema{
			"region": {
				Type:        schema.TypeString,
				Optional:    true,
				Computed:    true,
				Description: `Specifies the region in which to query the resource. If omitted, the provider-level region will be used.`,
			},
			"quota_key": {
				Type:        schema.TypeList,
				Optional:    true,
				Elem:        &schema.Schema{Type: schema.TypeString},
				Description: `Specifies the resource type.`,
			},
			"quotas": {
				Type:        schema.TypeList,
				Computed:    true,
				Description: `Indicates the resource quotas.`,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"quota_key": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: `Indicates the resource type.`,
						},
						"quota_limit": {
							Type:        schema.TypeInt,
							Computed:    true,
							Description: `Indicates the total quota.`,
						},
						"used": {
							Type:        schema.TypeInt,
							Computed:    true,
							Description: `Indicates the used quota.`,
						},
						"unit": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: `Indicates the quota unit.`,
						},
					},
				},
			},
		},
	}
}

type QuotaDetailsDSWrapper struct {
	*schemas.ResourceDataWrapper
	Config *config.Config
}

func newQuotaDetailsDSWrapper(d *schema.ResourceData, meta interface{}) *QuotaDetailsDSWrapper {
	return &QuotaDetailsDSWrapper{
		ResourceDataWrapper: schemas.NewSchemaWrapper(d),
		Config:              meta.(*config.Config),
	}
}

func dataSourceElbQuotaDetailsRead(_ context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	wrapper := newQuotaDetailsDSWrapper(d, meta)
	listQuotaDetailsRst, err := wrapper.ListQuotaDetails()
	if err != nil {
		return diag.FromErr(err)
	}

	id, err := uuid.GenerateUUID()
	if err != nil {
		return diag.FromErr(err)
	}
	d.SetId(id)

	err = wrapper.listQuotaDetailsToSchema(listQuotaDetailsRst)
	if err != nil {
		return diag.FromErr(err)
	}

	return nil
}

// @API ELB GET /v3/{project_id}/elb/quotas/details
func (w *QuotaDetailsDSWrapper) ListQuotaDetails() (*gjson.Result, error) {
	client, err := w.NewClient(w.Config, "elb")
	if err != nil {
		return nil, err
	}

	uri := "/v3/{project_id}/elb/quotas/details"
	params := map[string]any{
		"quota_key": w.ListToArray("quota_key"),
	}
	params = utils.RemoveNil(params)
	return httphelper.New(client).
		Method("GET").
		URI(uri).
		Query(params).
		Request().
		Result()
}

func (w *QuotaDetailsDSWrapper) listQuotaDetailsToSchema(body *gjson.Result) error {
	d := w.ResourceData
	mErr := multierror.Append(nil,
		d.Set("region", w.Config.GetRegion(w.ResourceData)),
		d.Set("quotas", schemas.SliceToList(body.Get("quotas"),
			func(quotas gjson.Result) any {
				return map[string]any{
					"quota_key":   quotas.Get("quota_key").Value(),
					"quota_limit": quotas.Get("quota_limit").Value(),
					"used":        quotas.Get("used").Value(),
					"unit":        quotas.Get("unit").Value(),
				}
			},
		)),
	)
	return mErr.ErrorOrNil()
}
