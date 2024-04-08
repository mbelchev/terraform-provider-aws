// Code generated by internal/generate/tagstests/main.go; DO NOT EDIT.

package iam_test

import (
	"fmt"
	"testing"

	"github.com/aws/aws-sdk-go-v2/service/iam/types"
	sdkacctest "github.com/hashicorp/terraform-plugin-testing/helper/acctest"
	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
	"github.com/hashicorp/terraform-provider-aws/internal/acctest"
	"github.com/hashicorp/terraform-provider-aws/names"
)

func TestAccIAMVirtualMFADevice_tags(t *testing.T) {
	ctx := acctest.Context(t)
	var v types.VirtualMFADevice
	resourceName := "aws_iam_virtual_mfa_device.test"
	rName := sdkacctest.RandomWithPrefix(acctest.ResourcePrefix)

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(ctx, t) },
		ErrorCheck:               acctest.ErrorCheck(t, names.IAMServiceID),
		ProtoV5ProviderFactories: acctest.ProtoV5ProviderFactories,
		CheckDestroy:             testAccCheckVirtualMFADeviceDestroy(ctx),
		Steps: []resource.TestStep{
			{
				Config: testAccVirtualMFADeviceConfig_tags1(rName, "key1", "value1"),
				Check: resource.ComposeAggregateTestCheckFunc(
					testAccCheckVirtualMFADeviceExists(ctx, resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "tags.%", "1"),
					resource.TestCheckResourceAttr(resourceName, "tags.key1", "value1"),
				),
			},
			{
				ResourceName:      resourceName,
				ImportState:       true,
				ImportStateVerify: true,
				ImportStateVerifyIgnore: []string{
					"base_32_string_seed", "qr_code_png",
				},
			},
			{
				Config: testAccVirtualMFADeviceConfig_tags2(rName, "key1", "value1updated", "key2", "value2"),
				Check: resource.ComposeAggregateTestCheckFunc(
					testAccCheckVirtualMFADeviceExists(ctx, resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "tags.%", "2"),
					resource.TestCheckResourceAttr(resourceName, "tags.key1", "value1updated"),
					resource.TestCheckResourceAttr(resourceName, "tags.key2", "value2"),
				),
			},
			{
				ResourceName:      resourceName,
				ImportState:       true,
				ImportStateVerify: true,
				ImportStateVerifyIgnore: []string{
					"base_32_string_seed", "qr_code_png",
				},
			},
			{
				Config: testAccVirtualMFADeviceConfig_tags1(rName, "key2", "value2"),
				Check: resource.ComposeAggregateTestCheckFunc(
					testAccCheckVirtualMFADeviceExists(ctx, resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "tags.%", "1"),
					resource.TestCheckResourceAttr(resourceName, "tags.key2", "value2"),
				),
			},
			{
				ResourceName:      resourceName,
				ImportState:       true,
				ImportStateVerify: true,
				ImportStateVerifyIgnore: []string{
					"base_32_string_seed", "qr_code_png",
				},
			},
			{
				Config: testAccVirtualMFADeviceConfig_tags0(rName),
				Check: resource.ComposeAggregateTestCheckFunc(
					testAccCheckVirtualMFADeviceExists(ctx, resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "tags.%", "0"),
				),
			},
			{
				ResourceName:      resourceName,
				ImportState:       true,
				ImportStateVerify: true,
				ImportStateVerifyIgnore: []string{
					"base_32_string_seed", "qr_code_png",
				},
			},
		},
	})
}

func TestAccIAMVirtualMFADevice_tags_null(t *testing.T) {
	ctx := acctest.Context(t)
	var v types.VirtualMFADevice
	resourceName := "aws_iam_virtual_mfa_device.test"
	rName := sdkacctest.RandomWithPrefix(acctest.ResourcePrefix)

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(ctx, t) },
		ErrorCheck:               acctest.ErrorCheck(t, names.IAMServiceID),
		ProtoV5ProviderFactories: acctest.ProtoV5ProviderFactories,
		CheckDestroy:             testAccCheckVirtualMFADeviceDestroy(ctx),
		Steps: []resource.TestStep{
			{
				Config: testAccVirtualMFADeviceConfig_tagsNull(rName, "key1"),
				Check: resource.ComposeAggregateTestCheckFunc(
					testAccCheckVirtualMFADeviceExists(ctx, resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "tags.%", "0"),
				),
			},
			{
				ResourceName:      resourceName,
				ImportState:       true,
				ImportStateVerify: true,
				ImportStateVerifyIgnore: []string{
					"base_32_string_seed", "qr_code_png",
				},
			},
			{
				Config:             testAccVirtualMFADeviceConfig_tags0(rName),
				PlanOnly:           true,
				ExpectNonEmptyPlan: false,
			},
		},
	})
}

func TestAccIAMVirtualMFADevice_tags_AddOnUpdate(t *testing.T) {
	ctx := acctest.Context(t)
	var v types.VirtualMFADevice
	resourceName := "aws_iam_virtual_mfa_device.test"
	rName := sdkacctest.RandomWithPrefix(acctest.ResourcePrefix)

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(ctx, t) },
		ErrorCheck:               acctest.ErrorCheck(t, names.IAMServiceID),
		ProtoV5ProviderFactories: acctest.ProtoV5ProviderFactories,
		CheckDestroy:             testAccCheckVirtualMFADeviceDestroy(ctx),
		Steps: []resource.TestStep{
			{
				Config: testAccVirtualMFADeviceConfig_tags0(rName),
				Check: resource.ComposeAggregateTestCheckFunc(
					testAccCheckVirtualMFADeviceExists(ctx, resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "tags.%", "0"),
				),
			},
			{
				Config: testAccVirtualMFADeviceConfig_tags1(rName, "key1", "value1"),
				Check: resource.ComposeAggregateTestCheckFunc(
					testAccCheckVirtualMFADeviceExists(ctx, resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "tags.%", "1"),
					resource.TestCheckResourceAttr(resourceName, "tags.key1", "value1"),
				),
			},
			{
				ResourceName:      resourceName,
				ImportState:       true,
				ImportStateVerify: true,
				ImportStateVerifyIgnore: []string{
					"base_32_string_seed", "qr_code_png",
				},
			},
		},
	})
}

func TestAccIAMVirtualMFADevice_tags_EmptyTag_OnCreate(t *testing.T) {
	ctx := acctest.Context(t)
	var v types.VirtualMFADevice
	resourceName := "aws_iam_virtual_mfa_device.test"
	rName := sdkacctest.RandomWithPrefix(acctest.ResourcePrefix)

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(ctx, t) },
		ErrorCheck:               acctest.ErrorCheck(t, names.IAMServiceID),
		ProtoV5ProviderFactories: acctest.ProtoV5ProviderFactories,
		CheckDestroy:             testAccCheckVirtualMFADeviceDestroy(ctx),
		Steps: []resource.TestStep{
			{
				Config: testAccVirtualMFADeviceConfig_tags1(rName, "key1", ""),
				Check: resource.ComposeAggregateTestCheckFunc(
					testAccCheckVirtualMFADeviceExists(ctx, resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "tags.%", "1"),
					resource.TestCheckResourceAttr(resourceName, "tags.key1", ""),
				),
			},
			{
				ResourceName:      resourceName,
				ImportState:       true,
				ImportStateVerify: true,
				ImportStateVerifyIgnore: []string{
					"base_32_string_seed", "qr_code_png",
				},
			},
			{
				Config: testAccVirtualMFADeviceConfig_tags0(rName),
				Check: resource.ComposeAggregateTestCheckFunc(
					testAccCheckVirtualMFADeviceExists(ctx, resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "tags.%", "0"),
				),
			},
			{
				ResourceName:      resourceName,
				ImportState:       true,
				ImportStateVerify: true,
				ImportStateVerifyIgnore: []string{
					"base_32_string_seed", "qr_code_png",
				},
			},
		},
	})
}

func TestAccIAMVirtualMFADevice_tags_EmptyTag_OnUpdate_Add(t *testing.T) {
	ctx := acctest.Context(t)
	var v types.VirtualMFADevice
	resourceName := "aws_iam_virtual_mfa_device.test"
	rName := sdkacctest.RandomWithPrefix(acctest.ResourcePrefix)

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(ctx, t) },
		ErrorCheck:               acctest.ErrorCheck(t, names.IAMServiceID),
		ProtoV5ProviderFactories: acctest.ProtoV5ProviderFactories,
		CheckDestroy:             testAccCheckVirtualMFADeviceDestroy(ctx),
		Steps: []resource.TestStep{
			{
				Config: testAccVirtualMFADeviceConfig_tags1(rName, "key1", "value1"),
				Check: resource.ComposeAggregateTestCheckFunc(
					testAccCheckVirtualMFADeviceExists(ctx, resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "tags.%", "1"),
					resource.TestCheckResourceAttr(resourceName, "tags.key1", "value1"),
				),
			},
			{
				Config: testAccVirtualMFADeviceConfig_tags2(rName, "key1", "value1", "key2", ""),
				Check: resource.ComposeAggregateTestCheckFunc(
					testAccCheckVirtualMFADeviceExists(ctx, resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "tags.%", "2"),
					resource.TestCheckResourceAttr(resourceName, "tags.key1", "value1"),
					resource.TestCheckResourceAttr(resourceName, "tags.key2", ""),
				),
			},
			{
				ResourceName:      resourceName,
				ImportState:       true,
				ImportStateVerify: true,
				ImportStateVerifyIgnore: []string{
					"base_32_string_seed", "qr_code_png",
				},
			},
			{
				Config: testAccVirtualMFADeviceConfig_tags1(rName, "key1", "value1"),
				Check: resource.ComposeAggregateTestCheckFunc(
					testAccCheckVirtualMFADeviceExists(ctx, resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "tags.%", "1"),
					resource.TestCheckResourceAttr(resourceName, "tags.key1", "value1"),
				),
			},
			{
				ResourceName:      resourceName,
				ImportState:       true,
				ImportStateVerify: true,
				ImportStateVerifyIgnore: []string{
					"base_32_string_seed", "qr_code_png",
				},
			},
		},
	})
}

func TestAccIAMVirtualMFADevice_tags_EmptyTag_OnUpdate_Replace(t *testing.T) {
	ctx := acctest.Context(t)
	var v types.VirtualMFADevice
	resourceName := "aws_iam_virtual_mfa_device.test"
	rName := sdkacctest.RandomWithPrefix(acctest.ResourcePrefix)

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(ctx, t) },
		ErrorCheck:               acctest.ErrorCheck(t, names.IAMServiceID),
		ProtoV5ProviderFactories: acctest.ProtoV5ProviderFactories,
		CheckDestroy:             testAccCheckVirtualMFADeviceDestroy(ctx),
		Steps: []resource.TestStep{
			{
				Config: testAccVirtualMFADeviceConfig_tags1(rName, "key1", "value1"),
				Check: resource.ComposeAggregateTestCheckFunc(
					testAccCheckVirtualMFADeviceExists(ctx, resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "tags.%", "1"),
					resource.TestCheckResourceAttr(resourceName, "tags.key1", "value1"),
				),
			},
			{
				Config: testAccVirtualMFADeviceConfig_tags1(rName, "key1", ""),
				Check: resource.ComposeAggregateTestCheckFunc(
					testAccCheckVirtualMFADeviceExists(ctx, resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "tags.%", "1"),
					resource.TestCheckResourceAttr(resourceName, "tags.key1", ""),
				),
			},
			{
				ResourceName:      resourceName,
				ImportState:       true,
				ImportStateVerify: true,
				ImportStateVerifyIgnore: []string{
					"base_32_string_seed", "qr_code_png",
				},
			},
		},
	})
}

func TestAccIAMVirtualMFADevice_tags_DefaultTags_providerOnly(t *testing.T) {
	ctx := acctest.Context(t)
	var v types.VirtualMFADevice
	resourceName := "aws_iam_virtual_mfa_device.test"
	rName := sdkacctest.RandomWithPrefix(acctest.ResourcePrefix)

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(ctx, t) },
		ErrorCheck:               acctest.ErrorCheck(t, names.IAMServiceID),
		ProtoV5ProviderFactories: acctest.ProtoV5ProviderFactories,
		CheckDestroy:             testAccCheckVirtualMFADeviceDestroy(ctx),
		Steps: []resource.TestStep{
			{
				Config: acctest.ConfigCompose(
					acctest.ConfigDefaultTags_Tags1("key1", "value1"),
					testAccVirtualMFADeviceConfig_tags0(rName),
				),
				Check: resource.ComposeAggregateTestCheckFunc(
					testAccCheckVirtualMFADeviceExists(ctx, resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "tags.%", "0"),
					resource.TestCheckResourceAttr(resourceName, "tags_all.%", "1"),
					resource.TestCheckResourceAttr(resourceName, "tags_all.key1", "value1"),
				),
			},
			{
				ResourceName:      resourceName,
				ImportState:       true,
				ImportStateVerify: true,
				ImportStateVerifyIgnore: []string{
					"base_32_string_seed", "qr_code_png",
				},
			},
			{
				Config: acctest.ConfigCompose(
					acctest.ConfigDefaultTags_Tags2("key1", "value1updated", "key2", "value2"),
					testAccVirtualMFADeviceConfig_tags0(rName),
				),
				Check: resource.ComposeAggregateTestCheckFunc(
					testAccCheckVirtualMFADeviceExists(ctx, resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "tags.%", "0"),
					resource.TestCheckResourceAttr(resourceName, "tags_all.%", "2"),
					resource.TestCheckResourceAttr(resourceName, "tags_all.key1", "value1updated"),
					resource.TestCheckResourceAttr(resourceName, "tags_all.key2", "value2"),
				),
			},
			{
				ResourceName:      resourceName,
				ImportState:       true,
				ImportStateVerify: true,
				ImportStateVerifyIgnore: []string{
					"base_32_string_seed", "qr_code_png",
				},
			},
			{
				Config: acctest.ConfigCompose(
					acctest.ConfigDefaultTags_Tags1("key2", "value2"),
					testAccVirtualMFADeviceConfig_tags0(rName),
				),
				Check: resource.ComposeAggregateTestCheckFunc(
					testAccCheckVirtualMFADeviceExists(ctx, resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "tags.%", "0"),
					resource.TestCheckResourceAttr(resourceName, "tags_all.%", "1"),
					resource.TestCheckResourceAttr(resourceName, "tags_all.key2", "value2"),
				),
			},
			{
				ResourceName:      resourceName,
				ImportState:       true,
				ImportStateVerify: true,
				ImportStateVerifyIgnore: []string{
					"base_32_string_seed", "qr_code_png",
				},
			},
			{
				Config: acctest.ConfigCompose(
					acctest.ConfigDefaultTags_Tags0(),
					testAccVirtualMFADeviceConfig_tags0(rName),
				),
				Check: resource.ComposeAggregateTestCheckFunc(
					testAccCheckVirtualMFADeviceExists(ctx, resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "tags.%", "0"),
					resource.TestCheckResourceAttr(resourceName, "tags_all.%", "0"),
				),
			},
			{
				ResourceName:      resourceName,
				ImportState:       true,
				ImportStateVerify: true,
				ImportStateVerifyIgnore: []string{
					"base_32_string_seed", "qr_code_png",
				},
			},
		},
	})
}

func TestAccIAMVirtualMFADevice_tags_DefaultTags_nonOverlapping(t *testing.T) {
	ctx := acctest.Context(t)
	var v types.VirtualMFADevice
	resourceName := "aws_iam_virtual_mfa_device.test"
	rName := sdkacctest.RandomWithPrefix(acctest.ResourcePrefix)

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(ctx, t) },
		ErrorCheck:               acctest.ErrorCheck(t, names.IAMServiceID),
		ProtoV5ProviderFactories: acctest.ProtoV5ProviderFactories,
		CheckDestroy:             testAccCheckVirtualMFADeviceDestroy(ctx),
		Steps: []resource.TestStep{
			{
				Config: acctest.ConfigCompose(
					acctest.ConfigDefaultTags_Tags1("providerkey1", "providervalue1"),
					testAccVirtualMFADeviceConfig_tags1(rName, "resourcekey1", "resourcevalue1"),
				),
				Check: resource.ComposeAggregateTestCheckFunc(
					testAccCheckVirtualMFADeviceExists(ctx, resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "tags.%", "1"),
					resource.TestCheckResourceAttr(resourceName, "tags.resourcekey1", "resourcevalue1"),
					resource.TestCheckResourceAttr(resourceName, "tags_all.%", "2"),
					resource.TestCheckResourceAttr(resourceName, "tags_all.providerkey1", "providervalue1"),
					resource.TestCheckResourceAttr(resourceName, "tags_all.resourcekey1", "resourcevalue1"),
				),
			},
			{
				ResourceName:      resourceName,
				ImportState:       true,
				ImportStateVerify: true,
				ImportStateVerifyIgnore: []string{
					"base_32_string_seed", "qr_code_png",
				},
			},
			{
				Config: acctest.ConfigCompose(
					acctest.ConfigDefaultTags_Tags1("providerkey1", "providervalue1updated"),
					testAccVirtualMFADeviceConfig_tags2(rName, "resourcekey1", "resourcevalue1updated", "resourcekey2", "resourcevalue2"),
				),
				Check: resource.ComposeAggregateTestCheckFunc(
					testAccCheckVirtualMFADeviceExists(ctx, resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "tags.%", "2"),
					resource.TestCheckResourceAttr(resourceName, "tags.resourcekey1", "resourcevalue1updated"),
					resource.TestCheckResourceAttr(resourceName, "tags.resourcekey2", "resourcevalue2"),
					resource.TestCheckResourceAttr(resourceName, "tags_all.%", "3"),
					resource.TestCheckResourceAttr(resourceName, "tags_all.providerkey1", "providervalue1updated"),
					resource.TestCheckResourceAttr(resourceName, "tags_all.resourcekey1", "resourcevalue1updated"),
					resource.TestCheckResourceAttr(resourceName, "tags_all.resourcekey2", "resourcevalue2"),
				),
			},
			{
				ResourceName:      resourceName,
				ImportState:       true,
				ImportStateVerify: true,
				ImportStateVerifyIgnore: []string{
					"base_32_string_seed", "qr_code_png",
				},
			},
			{
				Config: acctest.ConfigCompose(
					acctest.ConfigDefaultTags_Tags0(),
					testAccVirtualMFADeviceConfig_tags0(rName),
				),
				Check: resource.ComposeAggregateTestCheckFunc(
					testAccCheckVirtualMFADeviceExists(ctx, resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "tags.%", "0"),
					resource.TestCheckResourceAttr(resourceName, "tags_all.%", "0"),
				),
			},
			{
				ResourceName:      resourceName,
				ImportState:       true,
				ImportStateVerify: true,
				ImportStateVerifyIgnore: []string{
					"base_32_string_seed", "qr_code_png",
				},
			},
		},
	})
}

func TestAccIAMVirtualMFADevice_tags_DefaultTags_overlapping(t *testing.T) {
	ctx := acctest.Context(t)
	var v types.VirtualMFADevice
	resourceName := "aws_iam_virtual_mfa_device.test"
	rName := sdkacctest.RandomWithPrefix(acctest.ResourcePrefix)

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(ctx, t) },
		ErrorCheck:               acctest.ErrorCheck(t, names.IAMServiceID),
		ProtoV5ProviderFactories: acctest.ProtoV5ProviderFactories,
		CheckDestroy:             testAccCheckVirtualMFADeviceDestroy(ctx),
		Steps: []resource.TestStep{
			{
				Config: acctest.ConfigCompose(
					acctest.ConfigDefaultTags_Tags1("overlapkey1", "providervalue1"),
					testAccVirtualMFADeviceConfig_tags1(rName, "overlapkey1", "resourcevalue1"),
				),
				Check: resource.ComposeAggregateTestCheckFunc(
					testAccCheckVirtualMFADeviceExists(ctx, resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "tags.%", "1"),
					resource.TestCheckResourceAttr(resourceName, "tags.overlapkey1", "resourcevalue1"),
					resource.TestCheckResourceAttr(resourceName, "tags_all.%", "1"),
					resource.TestCheckResourceAttr(resourceName, "tags_all.overlapkey1", "resourcevalue1"),
				),
			},
			{
				ResourceName:      resourceName,
				ImportState:       true,
				ImportStateVerify: true,
				ImportStateVerifyIgnore: []string{
					"base_32_string_seed", "qr_code_png",
				},
			},
			{
				Config: acctest.ConfigCompose(
					acctest.ConfigDefaultTags_Tags2("overlapkey1", "providervalue1", "overlapkey2", "providervalue2"),
					testAccVirtualMFADeviceConfig_tags2(rName, "overlapkey1", "resourcevalue1", "overlapkey2", "resourcevalue2"),
				),
				Check: resource.ComposeAggregateTestCheckFunc(
					testAccCheckVirtualMFADeviceExists(ctx, resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "tags.%", "2"),
					resource.TestCheckResourceAttr(resourceName, "tags.overlapkey1", "resourcevalue1"),
					resource.TestCheckResourceAttr(resourceName, "tags.overlapkey2", "resourcevalue2"),
					resource.TestCheckResourceAttr(resourceName, "tags_all.%", "2"),
					resource.TestCheckResourceAttr(resourceName, "tags_all.overlapkey1", "resourcevalue1"),
					resource.TestCheckResourceAttr(resourceName, "tags_all.overlapkey2", "resourcevalue2"),
				),
			},
			{
				ResourceName:      resourceName,
				ImportState:       true,
				ImportStateVerify: true,
				ImportStateVerifyIgnore: []string{
					"base_32_string_seed", "qr_code_png",
				},
			},
			{
				Config: acctest.ConfigCompose(
					acctest.ConfigDefaultTags_Tags1("overlapkey1", "providervalue1"),
					testAccVirtualMFADeviceConfig_tags1(rName, "overlapkey1", "resourcevalue2"),
				),
				Check: resource.ComposeAggregateTestCheckFunc(
					testAccCheckVirtualMFADeviceExists(ctx, resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "tags.%", "1"),
					resource.TestCheckResourceAttr(resourceName, "tags.overlapkey1", "resourcevalue2"),
					resource.TestCheckResourceAttr(resourceName, "tags_all.%", "1"),
					resource.TestCheckResourceAttr(resourceName, "tags_all.overlapkey1", "resourcevalue2"),
				),
			},
			{
				ResourceName:      resourceName,
				ImportState:       true,
				ImportStateVerify: true,
				ImportStateVerifyIgnore: []string{
					"base_32_string_seed", "qr_code_png",
				},
			},
		},
	})
}

func TestAccIAMVirtualMFADevice_tags_DefaultTags_updateToProviderOnly(t *testing.T) {
	ctx := acctest.Context(t)
	var v types.VirtualMFADevice
	resourceName := "aws_iam_virtual_mfa_device.test"
	rName := sdkacctest.RandomWithPrefix(acctest.ResourcePrefix)

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(ctx, t) },
		ErrorCheck:               acctest.ErrorCheck(t, names.IAMServiceID),
		ProtoV5ProviderFactories: acctest.ProtoV5ProviderFactories,
		CheckDestroy:             testAccCheckVirtualMFADeviceDestroy(ctx),
		Steps: []resource.TestStep{
			{
				Config: testAccVirtualMFADeviceConfig_tags1(rName, "key1", "value1"),
				Check: resource.ComposeAggregateTestCheckFunc(
					testAccCheckVirtualMFADeviceExists(ctx, resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "tags.%", "1"),
					resource.TestCheckResourceAttr(resourceName, "tags.key1", "value1"),
					resource.TestCheckResourceAttr(resourceName, "tags_all.%", "1"),
					resource.TestCheckResourceAttr(resourceName, "tags_all.key1", "value1"),
				),
			},
			{
				Config: acctest.ConfigCompose(
					acctest.ConfigDefaultTags_Tags1("key1", "value1"),
					testAccVirtualMFADeviceConfig_tags0(rName),
				),
				Check: resource.ComposeAggregateTestCheckFunc(
					testAccCheckVirtualMFADeviceExists(ctx, resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "tags.%", "0"),
					resource.TestCheckResourceAttr(resourceName, "tags_all.%", "1"),
					resource.TestCheckResourceAttr(resourceName, "tags_all.key1", "value1"),
				),
			},
			{
				ResourceName:      resourceName,
				ImportState:       true,
				ImportStateVerify: true,
				ImportStateVerifyIgnore: []string{
					"base_32_string_seed", "qr_code_png",
				},
			},
		},
	})
}

func TestAccIAMVirtualMFADevice_tags_DefaultTags_updateToResourceOnly(t *testing.T) {
	ctx := acctest.Context(t)
	var v types.VirtualMFADevice
	resourceName := "aws_iam_virtual_mfa_device.test"
	rName := sdkacctest.RandomWithPrefix(acctest.ResourcePrefix)

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(ctx, t) },
		ErrorCheck:               acctest.ErrorCheck(t, names.IAMServiceID),
		ProtoV5ProviderFactories: acctest.ProtoV5ProviderFactories,
		CheckDestroy:             testAccCheckVirtualMFADeviceDestroy(ctx),
		Steps: []resource.TestStep{
			{
				Config: acctest.ConfigCompose(
					acctest.ConfigDefaultTags_Tags1("key1", "value1"),
					testAccVirtualMFADeviceConfig_tags0(rName),
				),
				Check: resource.ComposeAggregateTestCheckFunc(
					testAccCheckVirtualMFADeviceExists(ctx, resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "tags.%", "0"),
					resource.TestCheckResourceAttr(resourceName, "tags_all.%", "1"),
					resource.TestCheckResourceAttr(resourceName, "tags_all.key1", "value1"),
				),
			},
			{
				Config: testAccVirtualMFADeviceConfig_tags1(rName, "key1", "value1"),
				Check: resource.ComposeAggregateTestCheckFunc(
					testAccCheckVirtualMFADeviceExists(ctx, resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "tags.%", "1"),
					resource.TestCheckResourceAttr(resourceName, "tags.key1", "value1"),
					resource.TestCheckResourceAttr(resourceName, "tags_all.%", "1"),
					resource.TestCheckResourceAttr(resourceName, "tags_all.key1", "value1"),
				),
			},
			{
				ResourceName:      resourceName,
				ImportState:       true,
				ImportStateVerify: true,
				ImportStateVerifyIgnore: []string{
					"base_32_string_seed", "qr_code_png",
				},
			},
		},
	})
}

func TestAccIAMVirtualMFADevice_tags_DefaultTags_emptyResourceTag(t *testing.T) {
	ctx := acctest.Context(t)
	var v types.VirtualMFADevice
	resourceName := "aws_iam_virtual_mfa_device.test"
	rName := sdkacctest.RandomWithPrefix(acctest.ResourcePrefix)

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(ctx, t) },
		ErrorCheck:               acctest.ErrorCheck(t, names.IAMServiceID),
		ProtoV5ProviderFactories: acctest.ProtoV5ProviderFactories,
		CheckDestroy:             testAccCheckVirtualMFADeviceDestroy(ctx),
		Steps: []resource.TestStep{
			{
				Config: acctest.ConfigCompose(
					acctest.ConfigDefaultTags_Tags1("key1", "value1"),
					testAccVirtualMFADeviceConfig_tags1(rName, "key1", ""),
				),
				Check: resource.ComposeAggregateTestCheckFunc(
					testAccCheckVirtualMFADeviceExists(ctx, resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "tags.%", "1"),
					resource.TestCheckResourceAttr(resourceName, "tags.key1", ""),
					resource.TestCheckResourceAttr(resourceName, "tags_all.%", "1"),
					resource.TestCheckResourceAttr(resourceName, "tags_all.key1", ""),
				),
			},
			{
				ResourceName:      resourceName,
				ImportState:       true,
				ImportStateVerify: true,
				ImportStateVerifyIgnore: []string{
					"base_32_string_seed", "qr_code_png",
				},
			},
		},
	})
}

func TestAccIAMVirtualMFADevice_tags_DefaultTags_nullOverlappingResourceTag(t *testing.T) {
	ctx := acctest.Context(t)
	var v types.VirtualMFADevice
	resourceName := "aws_iam_virtual_mfa_device.test"
	rName := sdkacctest.RandomWithPrefix(acctest.ResourcePrefix)

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(ctx, t) },
		ErrorCheck:               acctest.ErrorCheck(t, names.IAMServiceID),
		ProtoV5ProviderFactories: acctest.ProtoV5ProviderFactories,
		CheckDestroy:             testAccCheckVirtualMFADeviceDestroy(ctx),
		Steps: []resource.TestStep{
			{
				Config: acctest.ConfigCompose(
					acctest.ConfigDefaultTags_Tags1("key1", "providervalue1"),
					testAccVirtualMFADeviceConfig_tagsNull(rName, "key1"),
				),
				Check: resource.ComposeAggregateTestCheckFunc(
					testAccCheckVirtualMFADeviceExists(ctx, resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "tags.%", "0"),
					resource.TestCheckResourceAttr(resourceName, "tags_all.%", "1"),
					resource.TestCheckResourceAttr(resourceName, "tags_all.key1", "providervalue1"),
				),
			},
			{
				ResourceName:      resourceName,
				ImportState:       true,
				ImportStateVerify: true,
				ImportStateVerifyIgnore: []string{
					"base_32_string_seed", "qr_code_png",
				},
			},
		},
	})
}

func TestAccIAMVirtualMFADevice_tags_DefaultTags_nullNonOverlappingResourceTag(t *testing.T) {
	ctx := acctest.Context(t)
	var v types.VirtualMFADevice
	resourceName := "aws_iam_virtual_mfa_device.test"
	rName := sdkacctest.RandomWithPrefix(acctest.ResourcePrefix)

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(ctx, t) },
		ErrorCheck:               acctest.ErrorCheck(t, names.IAMServiceID),
		ProtoV5ProviderFactories: acctest.ProtoV5ProviderFactories,
		CheckDestroy:             testAccCheckVirtualMFADeviceDestroy(ctx),
		Steps: []resource.TestStep{
			{
				Config: acctest.ConfigCompose(
					acctest.ConfigDefaultTags_Tags1("providerkey1", "providervalue1"),
					testAccVirtualMFADeviceConfig_tagsNull(rName, "resourcekey1"),
				),
				Check: resource.ComposeAggregateTestCheckFunc(
					testAccCheckVirtualMFADeviceExists(ctx, resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "tags.%", "0"),
					resource.TestCheckResourceAttr(resourceName, "tags_all.%", "1"),
					resource.TestCheckResourceAttr(resourceName, "tags_all.providerkey1", "providervalue1"),
				),
			},
			{
				ResourceName:      resourceName,
				ImportState:       true,
				ImportStateVerify: true,
				ImportStateVerifyIgnore: []string{
					"base_32_string_seed", "qr_code_png",
				},
			},
		},
	})
}

func testAccVirtualMFADeviceConfig_tags0(rName string) string {
	return fmt.Sprintf(`
resource "aws_iam_virtual_mfa_device" "test" {
  virtual_mfa_device_name = %[1]q

}
`, rName)
}

func testAccVirtualMFADeviceConfig_tags1(rName, tagKey1, tagValue1 string) string {
	return fmt.Sprintf(`
resource "aws_iam_virtual_mfa_device" "test" {
  virtual_mfa_device_name = %[1]q

  tags = {
    %[2]q = %[3]q
  }
}
`, rName, tagKey1, tagValue1)
}

func testAccVirtualMFADeviceConfig_tags2(rName, tagKey1, tagValue1, tagKey2, tagValue2 string) string {
	return fmt.Sprintf(`
resource "aws_iam_virtual_mfa_device" "test" {
  virtual_mfa_device_name = %[1]q

  tags = {
    %[2]q = %[3]q
    %[4]q = %[5]q
  }
}
`, rName, tagKey1, tagValue1, tagKey2, tagValue2)
}

func testAccVirtualMFADeviceConfig_tagsNull(rName, tagKey1 string) string {
	return fmt.Sprintf(`
resource "aws_iam_virtual_mfa_device" "test" {
  virtual_mfa_device_name = %[1]q

  tags = {
    %[2]q = null
  }
}
`, rName, tagKey1)
}
