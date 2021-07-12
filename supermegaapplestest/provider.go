package apples_old

import (
	"context"
	"log"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	client2 "github.com/ifreespam/terraform-provider-supermegaapplestest/client"
)

// Provider - apples OLD
func Provider() *schema.Provider {
	return &schema.Provider{
		Schema: map[string]*schema.Schema{
			"user_name": {
				Type:        schema.TypeString,
				Required:    true,
				Description: "User name identifier",
			},
		},
		ResourcesMap: map[string]*schema.Resource{
			"apples_record": resourceAppleRecord(),
		},
		ConfigureContextFunc: configureContext,
	}
}

func configureContext(ctx context.Context, data *schema.ResourceData) (interface{}, diag.Diagnostics) {
	var diags diag.Diagnostics

	userName := data.Get("user_name").(string)
	log.Println("[DEBUG] =============== configureContext " + userName)

	client := client2.NewTheClient(userName)
	return client, diags
}
