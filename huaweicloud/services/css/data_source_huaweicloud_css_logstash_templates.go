// Generated by PMS #172
package css

import (
	"context"

	"github.com/hashicorp/go-multierror"
	"github.com/hashicorp/go-uuid"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/tidwall/gjson"

	"github.com/huaweicloud/terraform-provider-huaweicloud/huaweicloud/config"
	"github.com/huaweicloud/terraform-provider-huaweicloud/huaweicloud/helper/filters"
	"github.com/huaweicloud/terraform-provider-huaweicloud/huaweicloud/helper/httphelper"
	"github.com/huaweicloud/terraform-provider-huaweicloud/huaweicloud/helper/schemas"
	"github.com/huaweicloud/terraform-provider-huaweicloud/huaweicloud/utils"
)

func DataSourceCssLogstashTemplates() *schema.Resource {
	return &schema.Resource{
		ReadContext: dataSourceCssLogstashTemplatesRead,

		Schema: map[string]*schema.Schema{
			"region": {
				Type:        schema.TypeString,
				Optional:    true,
				Computed:    true,
				Description: `Specifies the region in which to query the resource. If omitted, the provider-level region will be used.`,
			},
			"type": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: `Specifies the type of the CSS logstash configuration template.`,
			},
			"name": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: `Specifies the name of the template.`,
			},
			"system_templates": {
				Type:        schema.TypeList,
				Computed:    true,
				Description: `The list of the system templates.`,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"id": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: `The ID of the system template.`,
						},
						"name": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: `The name of the system template.`,
						},
						"conf_content": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: `The configuration file content of the system template.`,
						},
						"desc": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: `The description of the system template.`,
						},
					},
				},
			},
			"custom_templates": {
				Type:        schema.TypeList,
				Computed:    true,
				Description: `The list of the custom templates.`,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"id": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: `The ID of the custom template.`,
						},
						"name": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: `The name of the custom template.`,
						},
						"conf_content": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: `The configuration file content of the custom template.`,
						},
						"desc": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: `The description of the custom template.`,
						},
					},
				},
			},
		},
	}
}

type LogstashTemplatesDSWrapper struct {
	*schemas.ResourceDataWrapper
	Config *config.Config
}

func newLogstashTemplatesDSWrapper(d *schema.ResourceData, meta interface{}) *LogstashTemplatesDSWrapper {
	return &LogstashTemplatesDSWrapper{
		ResourceDataWrapper: schemas.NewSchemaWrapper(d),
		Config:              meta.(*config.Config),
	}
}

func dataSourceCssLogstashTemplatesRead(_ context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	wrapper := newLogstashTemplatesDSWrapper(d, meta)
	lisTemRst, err := wrapper.ListTemplates()
	if err != nil {
		return diag.FromErr(err)
	}

	id, err := uuid.GenerateUUID()
	if err != nil {
		return diag.FromErr(err)
	}
	d.SetId(id)

	err = wrapper.listTemplatesToSchema(lisTemRst)
	if err != nil {
		return diag.FromErr(err)
	}

	return nil
}

// @API CSS GET /v1.0/{project_id}/lgsconf/template
func (w *LogstashTemplatesDSWrapper) ListTemplates() (*gjson.Result, error) {
	client, err := w.NewClient(w.Config, "css")
	if err != nil {
		return nil, err
	}

	uri := "/v1.0/{project_id}/lgsconf/template"
	params := map[string]any{
		"type": w.Get("type"),
	}
	params = utils.RemoveNil(params)
	return httphelper.New(client).
		Method("GET").
		URI(uri).
		Query(params).
		Filter(
			filters.New().From("systemTemplates").
				Where("name", "=", w.Get("name")),
		).
		Filter(
			filters.New().From("customTemplates").
				Where("name", "=", w.Get("name")),
		).
		OkCode(200).
		Request().
		Result()
}

func (w *LogstashTemplatesDSWrapper) listTemplatesToSchema(body *gjson.Result) error {
	d := w.ResourceData
	mErr := multierror.Append(nil,
		d.Set("region", w.Config.GetRegion(w.ResourceData)),
		d.Set("system_templates", schemas.SliceToList(body.Get("systemTemplates"),
			func(systemTemplates gjson.Result) any {
				return map[string]any{
					"id":           systemTemplates.Get("id").Value(),
					"name":         systemTemplates.Get("name").Value(),
					"conf_content": systemTemplates.Get("confContent").Value(),
					"desc":         systemTemplates.Get("desc").Value(),
				}
			},
		)),
		d.Set("custom_templates", schemas.SliceToList(body.Get("customTemplates"),
			func(customTemplates gjson.Result) any {
				return map[string]any{
					"id":           customTemplates.Get("id").Value(),
					"name":         customTemplates.Get("name").Value(),
					"conf_content": customTemplates.Get("confContent").Value(),
					"desc":         customTemplates.Get("desc").Value(),
				}
			},
		)),
	)
	return mErr.ErrorOrNil()
}
