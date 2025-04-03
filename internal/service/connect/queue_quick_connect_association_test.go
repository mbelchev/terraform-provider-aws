// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package connect_test

import (
	"context"
	"fmt"
	"testing"

	"github.com/aws/aws-sdk-go/service/connect"
	"github.com/hashicorp/aws-sdk-go-base/v2/awsv1shim/v2/tfawserr"
	sdkacctest "github.com/hashicorp/terraform-plugin-testing/helper/acctest"
	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
	"github.com/hashicorp/terraform-plugin-testing/terraform"
	"github.com/hashicorp/terraform-provider-aws/internal/acctest"
	"github.com/hashicorp/terraform-provider-aws/internal/conns"
	tfconnect "github.com/hashicorp/terraform-provider-aws/internal/service/connect"
)

func testAccQueueQuickConnectAssociation_basic(t *testing.T) {
	ctx := acctest.Context(t)
	var v connect.DescribeQueueOutput
	rName := sdkacctest.RandomWithPrefix("resource-test-terraform")
	rName2 := sdkacctest.RandomWithPrefix("resource-test-terraform")
	rName3 := sdkacctest.RandomWithPrefix("resource-test-terraform")
	resourceName := "aws_connect_queue_quick_connect_association.test"
	queueResourceName := "aws_connect_queue.test"
	queueDatasourceName := "data.aws_connect_queue.test"
	phoneNumber := "+12345678912"

	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(ctx, t) },
		ErrorCheck:               acctest.ErrorCheck(t, connect.EndpointsID),
		ProtoV5ProviderFactories: acctest.ProtoV5ProviderFactories,
		CheckDestroy:             testAccCheckQueueQuickConnectAssociationDestroy(ctx),
		Steps: []resource.TestStep{
			{
				Config: testAccQueueQuickConnectAssociationConfig_basic(rName, rName2, rName3, phoneNumber),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckQueueQuickConnectAssociationExists(ctx, resourceName),
					testAccCheckQueueExists(ctx, queueResourceName, &v),
					resource.TestCheckResourceAttrSet(resourceName, "id"),
					resource.TestCheckResourceAttrPair(resourceName, "instance_id", "aws_connect_instance.test", "id"),
					resource.TestCheckResourceAttrPair(resourceName, "queue_id", queueResourceName, "queue_id"),
					resource.TestCheckResourceAttr(queueResourceName, "quick_connect_ids.#", "0"),
					resource.TestCheckResourceAttrPair(queueDatasourceName, "quick_connect_ids.0", "aws_connect_quick_connect.test", "quick_connect_id"),
				),
			},
			{
				ResourceName:      resourceName,
				ImportState:       true,
				ImportStateVerify: true,
			},
		},
	})
}

func testAccQueueQuickConnectAssociation_disappears(t *testing.T) {
	ctx := acctest.Context(t)
	rName := sdkacctest.RandomWithPrefix("resource-test-terraform")
	rName2 := sdkacctest.RandomWithPrefix("resource-test-terraform")
	rName3 := sdkacctest.RandomWithPrefix("resource-test-terraform")
	resourceName := "aws_connect_queue_quick_connect_association.test"
	phoneNumber := "+12345678912"

	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(ctx, t) },
		ErrorCheck:               acctest.ErrorCheck(t, connect.EndpointsID),
		ProtoV5ProviderFactories: acctest.ProtoV5ProviderFactories,
		CheckDestroy:             testAccCheckQueueQuickConnectAssociationDestroy(ctx),
		Steps: []resource.TestStep{
			{
				Config: testAccQueueQuickConnectAssociationConfig_basic(rName, rName2, rName3, phoneNumber),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckQueueQuickConnectAssociationExists(ctx, resourceName),
					acctest.CheckResourceDisappears(ctx, acctest.Provider, tfconnect.ResourceQueueQuickConnectAssociation(), resourceName),
				),
				ExpectNonEmptyPlan: true,
			},
		},
	})
}

func testAccCheckQueueQuickConnectAssociationExists(ctx context.Context, resourceName string) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		rs, ok := s.RootModule().Resources[resourceName]
		if !ok {
			return fmt.Errorf("Connect Queue Quick Connect Association not found: %s", resourceName)
		}

		if rs.Primary.ID == "" {
			return fmt.Errorf("Connect Queue Quick Connect Association ID not set")
		}

		instanceID, queueID, err := tfconnect.QueueQuickConnectAssociationParseResourceID(rs.Primary.ID)

		if err != nil {
			return err
		}

		conn := acctest.Provider.Meta().(*conns.AWSClient).ConnectConn(ctx)

		queueQuickConnects, err := tfconnect.GetQueueQuickConnectIDs(ctx, conn, instanceID, queueID)

		if err != nil {
			return fmt.Errorf("error finding Connect Queue Quick Connect Association (%s): %w", rs.Primary.ID, err)
		}

		if queueQuickConnects == nil {
			return fmt.Errorf("error finding Connect Queue Quick Connect Association (%s): not found", rs.Primary.ID)
		}

		return nil
	}
}

func testAccCheckQueueQuickConnectAssociationDestroy(ctx context.Context) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		for _, rs := range s.RootModule().Resources {
			if rs.Type != "aws_connect_queue_quick_connect_association" {
				continue
			}

			conn := acctest.Provider.Meta().(*conns.AWSClient).ConnectConn(ctx)

			instanceID, queueID, err := tfconnect.QueueQuickConnectAssociationParseResourceID(rs.Primary.ID)

			if err != nil {
				return err
			}

			queueQuickConnects, err := tfconnect.GetQueueQuickConnectIDs(ctx, conn, instanceID, queueID)

			if tfawserr.ErrCodeEquals(err, connect.ErrCodeResourceNotFoundException) {
				continue
			}

			if err != nil {
				return fmt.Errorf("error finding Connect Queue Quick Connect Association (%s): %w", rs.Primary.ID, err)
			}

			if queueQuickConnects != nil {
				return fmt.Errorf("Connect Queue Quick Connect Association (%s) still exists", rs.Primary.ID)
			}
		}

		return nil
	}
}

func testAccQueueQuickConnectAssociationConfig_base(rName, rName2, rName3, phoneNumber string) string {
	return fmt.Sprintf(`
resource "aws_connect_instance" "test" {
  identity_management_type = "CONNECT_MANAGED"
  inbound_calls_enabled    = true
  instance_alias           = %[1]q
  outbound_calls_enabled   = true
}

data "aws_connect_hours_of_operation" "test" {
  instance_id = aws_connect_instance.test.id
  name        = "Basic Hours"
}

resource "aws_connect_queue" "test" {
  instance_id           = aws_connect_instance.test.id
  name                  = %[2]q
  description           = "Used to test queue quick connect association resource"
  hours_of_operation_id = data.aws_connect_hours_of_operation.test.hours_of_operation_id

  tags = {
    "Name" = "Test Queue",
  }

	lifecycle {
		ignore_changes = [
			quick_connect_ids
		]
	}
}

resource "aws_connect_quick_connect" "test" {
  instance_id = aws_connect_instance.test.id
  name        = %[3]q
  description = "Used to test queue quick connect association resource"

  quick_connect_config {
    quick_connect_type = "PHONE_NUMBER"

    phone_config {
      phone_number = %[4]q
    }
  }

  tags = {
    "Name" = "Test Quick Connect"
  }
}

data "aws_connect_queue" "test" {
  instance_id = aws_connect_instance.test.id
  queue_id    = aws_connect_queue.test.queue_id
}
`, rName, rName2, rName3, phoneNumber)
}

func testAccQueueQuickConnectAssociationConfig_basic(rName, rName2, rName3, phoneNumber string) string {
	return acctest.ConfigCompose(
		testAccQueueQuickConnectAssociationConfig_base(rName, rName2, rName3, phoneNumber),
		`
resource "aws_connect_queue_quick_connect_association" "test" {
  instance_id       = aws_connect_instance.test.id
	queue_id          = aws_connect_queue.test.queue_id
	quick_connect_ids = [
		aws_connect_quick_connect.test.quick_connect_id
	]
}
`)
}
