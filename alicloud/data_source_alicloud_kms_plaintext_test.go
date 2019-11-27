package alicloud

import (
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/helper/resource"
)

func TestAccAlicloudKmsPlaintextDataSource(t *testing.T) {
	resourceId := "data.alicloud_kms_plaintext.default"

	resource.Test(t, resource.TestCase{
		PreCheck: func() {
			testAccPreCheck(t)
		},
		// module name
		IDRefreshName: resourceId,
		Providers:     testAccProviders,
		Steps: []resource.TestStep{
			{
				Config: testAccDataSourceKmsPlaintextConfigBasic,
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttrSet(
						resourceId, "plaintext",
					),
					resource.TestCheckResourceAttrSet(
						resourceId, "key_id",
					),
				),
			},
			{
				ResourceName: resourceId,
			},
		},
	})
}

const testAccDataSourceKmsPlaintextConfigBasic = `
resource "alicloud_kms_key" "default" {
    is_enabled = true
}

resource "alicloud_kms_ciphertext" "default" {
	key_id = "${alicloud_kms_key.default.id}"
	plaintext = "plaintext"
}

data "alicloud_kms_plaintext" "default" {
	ciphertext_blob = "${alicloud_kms_ciphertext.default.ciphertext_blob}"
}
`
