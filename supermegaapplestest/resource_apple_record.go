package apples_old

import (
	"context"
	"log"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	client2 "github.com/ifreespam/terraform-provider-supermegaapplestest/client"
)

func resourceAppleRecord() *schema.Resource {
	return &schema.Resource{
		CreateContext: resourceAppleRecordCreate,
		UpdateContext: resourceAppleRecordUpdate,
		ReadContext:   resourceAppleRecordRead,
		DeleteContext: resourceAppleRecordDelete,

		Importer: &schema.ResourceImporter{
			//State:        schema.ImportStatePassthrough,
			StateContext: schema.ImportStatePassthroughContext,
		},

		Schema: map[string]*schema.Schema{
			"name": {
				Type:        schema.TypeString,
				Required:    true,
				Description: "The name of an apple",
			},
			"sort": {
				Type:        schema.TypeString,
				Required:    true,
				Description: "The sort of an apple",
			},
			"count_items": {
				Type:        schema.TypeInt,
				Optional:    true,
				Description: "The count of an apple (UPDATED)",
				Default:     10,
			},
		},
	}
}

func resourceAppleRecordCreate(ctx context.Context, data *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(*client2.TheClient)

	name := data.Get("name").(string)
	sort := data.Get("sort").(string)
	countItems := data.Get("count_items").(int)

	appleItem := client2.Apple{
		Name:  name,
		Sort:  sort,
		Count: countItems,
	}

	appleHash := client2.HashApple(&appleItem)

	log.Println("[DEBUG] =============== CREATE " + appleHash)

	client.AddApple(&appleItem)

	data.SetId(appleHash)

	return nil
}

func resourceAppleRecordUpdate(ctx context.Context, data *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(*client2.TheClient)

	id := data.Id()

	log.Println("[DEBUG] =============== UPDATE", id)

	currentApple, err := client.GetAppleByHash(id)

	if err != nil {
		return diag.FromErr(err)
	}

	if data.HasChanges("name") {
		log.Println("[DEBUG] =============== UPDATE name changed")
		currentApple.Name = data.Get("na	me").(string)
	}

	if data.HasChanges("sort") {
		log.Println("[DEBUG] =============== UPDATE sort changed")
		currentApple.Sort = data.Get("sort").(string)
	}

	if data.HasChanges("count_items") {
		log.Println("[DEBUG] =============== UPDATE count_items changed")
		currentApple.Count = data.Get("count_items").(int)
	}

	client.DeleteAppleByHash(id)
	client.AddApple(currentApple)

	newId := client2.HashAppleV3(currentApple)
	data.SetId(newId)

	return nil
}

func resourceAppleRecordRead(ctx context.Context, data *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(*client2.TheClient)

	id := data.Id()

	log.Println("[DEBUG] =============== READ " + id)

	apple, err := client.GetAppleByHash(id)

	if err != nil {
		return diag.FromErr(err)
	}

	_ = data.Set("name", apple.Name)
	_ = data.Set("sort", apple.Sort)
	_ = data.Set("count_items", apple.Count)

	return nil
}

func resourceAppleRecordDelete(ctx context.Context, data *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(*client2.TheClient)

	id := data.Id()

	log.Println("[DEBUG] =============== DELETE " + id)

	client.DeleteAppleByHash(id)

	return nil
}
