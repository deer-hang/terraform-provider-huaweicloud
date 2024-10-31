package cph

import (
	"context"
	"fmt"
	"strings"
	"time"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/validation"

	"github.com/chnsz/golangsdk"

	"github.com/huaweicloud/terraform-provider-huaweicloud/huaweicloud/config"
	"github.com/huaweicloud/terraform-provider-huaweicloud/huaweicloud/utils"
)

var PhoneResetNonUpdatableParams = []string{"image_id", "phones", "phones.*.phone_id", "phones.*.property"}

// @API CPH POST /v1/{project_id}/cloud-phone/phones/batch-reset
// @API CPH GET /v1/{project_id}/cloud-phone/phones/{phone_id}
func ResourcePhoneReset() *schema.Resource {
	return &schema.Resource{
		CreateContext: resourcePhoneResetCreate,
		UpdateContext: resourcePhoneResetUpdate,
		ReadContext:   resourcePhoneResetRead,
		DeleteContext: resourcePhoneResetDelete,

		Timeouts: &schema.ResourceTimeout{
			Create: schema.DefaultTimeout(30 * time.Minute),
		},

		CustomizeDiff: config.FlexibleForceNew(PhoneResetNonUpdatableParams),

		Schema: map[string]*schema.Schema{
			"region": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
				ForceNew: true,
			},
			"phones": {
				Type:        schema.TypeList,
				Required:    true,
				MaxItems:    1,
				Description: `Specifies the CPH phones.`,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"phone_id": {
							Type:        schema.TypeString,
							Required:    true,
							Description: `Specifies the phone ID.`,
						},
						"property": {
							Type:        schema.TypeString,
							Optional:    true,
							Description: `Specifies the phone property.`,
						},
					},
				},
			},
			"image_id": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: `Specifies the image ID of the CPH phone.`,
			},
			"enable_force_new": {
				Type:         schema.TypeString,
				Optional:     true,
				ValidateFunc: validation.StringInSlice([]string{"true", "false"}, false),
				Description:  utils.SchemaDesc("", utils.SchemaDescInput{Internal: true}),
			},
		},
	}
}

func resourcePhoneResetCreate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	cfg := meta.(*config.Config)
	region := cfg.GetRegion(d)

	client, err := cfg.NewServiceClient("cph", region)
	if err != nil {
		return diag.Errorf("error creating CPH client: %s", err)
	}

	// createPhoneRestart: create CPH phone reset
	createPhoneResetHttpUrl := "v1/{project_id}/cloud-phone/phones/batch-reset"
	createPhoneResetPath := client.Endpoint + createPhoneResetHttpUrl
	createPhoneResetPath = strings.ReplaceAll(createPhoneResetPath, "{project_id}", client.ProjectID)

	createPhoneResetOpt := golangsdk.RequestOpts{
		KeepResponseBody: true,
	}

	createPhoneResetOpt.JSONBody = utils.RemoveNil(map[string]interface{}{
		"image_id": utils.ValueIgnoreEmpty(d.Get("image_id")),
		"phones":   d.Get("phones"),
	})
	createPhoneResetResp, err := client.Request("POST", createPhoneResetPath, &createPhoneResetOpt)
	if err != nil {
		return diag.Errorf("error creating CPH phone reset: %s", err)
	}

	resp, err := utils.FlattenResponse(createPhoneResetResp)
	if err != nil {
		return diag.FromErr(err)
	}
	id := utils.PathSearch("jobs|[0].phone_id", resp, "").(string)
	if id == "" {
		return diag.Errorf("Unable to find the phone ID from the API response")
	}
	d.SetId(id)

	errorCode := utils.PathSearch("jobs|[0].error_code", resp, "").(string)
	if errorCode != "" {
		errorMsg := utils.PathSearch("jobs|[0].error_msg", resp, "").(string)
		return diag.Errorf("failed to reset CPH phone (phone_id: %s), error_code: %s, error_msg: %s", id, errorCode, errorMsg)
	}

	err = checkPhoneResetStatus(ctx, client, d.Id(), d.Timeout(schema.TimeoutCreate))
	if err != nil {
		return diag.FromErr(err)
	}

	return nil
}

func resourcePhoneResetRead(_ context.Context, _ *schema.ResourceData, _ interface{}) diag.Diagnostics {
	return nil
}

func resourcePhoneResetUpdate(_ context.Context, _ *schema.ResourceData, _ interface{}) diag.Diagnostics {
	return nil
}

func resourcePhoneResetDelete(_ context.Context, _ *schema.ResourceData, _ interface{}) diag.Diagnostics {
	errorMsg := "Deleting CPH action resource is not supported. The resource is only removed from the state."
	return diag.Diagnostics{
		diag.Diagnostic{
			Severity: diag.Warning,
			Summary:  errorMsg,
		},
	}
}

func checkPhoneResetStatus(ctx context.Context, client *golangsdk.ServiceClient, id string, timeout time.Duration) error {
	stateConf := &resource.StateChangeConf{
		Pending:      []string{"PENDING"},
		Target:       []string{"COMPLETED"},
		Refresh:      phoneResetStateRefreshFunc(client, id),
		Timeout:      timeout,
		PollInterval: 10 * timeout,
		Delay:        10 * time.Second,
	}
	_, err := stateConf.WaitForStateContext(ctx)
	if err != nil {
		return fmt.Errorf("error waiting for CPH phone reset to be completed: %s", err)
	}
	return nil
}

func phoneResetStateRefreshFunc(client *golangsdk.ServiceClient, id string) resource.StateRefreshFunc {
	return func() (interface{}, string, error) {
		getPhoneRespBody, err := getPhoneDetail(client, id)
		if err != nil {
			return nil, "ERROR", err
		}

		// Status is 2, indicates the phone is running normally.
		phoneStatus := utils.PathSearch("status", getPhoneRespBody, float64(0)).(float64)
		if int(phoneStatus) == 2 {
			return getPhoneRespBody, "COMPLETED", nil
		}
		if int(phoneStatus) == -5 {
			return getPhoneRespBody, "ERROR", fmt.Errorf("failed to reset cloud phone: %s", id)
		}
		return getPhoneRespBody, "PENDING", nil
	}
}