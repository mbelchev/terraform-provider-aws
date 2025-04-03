// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package connect

import (
	"context"
	"fmt"
	"log"
	"strings"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/service/connect"
	"github.com/hashicorp/aws-sdk-go-base/v2/awsv1shim/v2/tfawserr"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-provider-aws/internal/conns"
	"github.com/hashicorp/terraform-provider-aws/internal/flex"
)

const queueQuickConnectAssociationIDSeparator = ":"

// @SDKResource("aws_connect_queue_quick_connect_association", name="Queue Quick Connect association")
func ResourceQueueQuickConnectAssociation() *schema.Resource {
	return &schema.Resource{
		CreateWithoutTimeout: ResourceQueueQuickConnectAssociationCreate,
		ReadWithoutTimeout:   ResourceQueueQuickConnectAssociationRead,
		UpdateWithoutTimeout: ResourceQueueQuickConnectAssociationUpdate,
		DeleteWithoutTimeout: ResourceQueueQuickConnectAssociationDelete,
		Importer: &schema.ResourceImporter{
			StateContext: schema.ImportStatePassthroughContext,
		},
		Schema: map[string]*schema.Schema{
			"instance_id": {
				Type:     schema.TypeString,
				Required: true,
				ForceNew: true,
			},
			"queue_id": {
				Type:     schema.TypeString,
				Required: true,
				ForceNew: true,
			},
			"quick_connect_ids": {
				Type:     schema.TypeSet,
				Required: true,
				MinItems: 1,
				MaxItems: 50,
				Elem: &schema.Schema{
					Type: schema.TypeString,
				},
			},
		},
	}
}

func ResourceQueueQuickConnectAssociationCreate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	conn := meta.(*conns.AWSClient).ConnectConn(ctx)

	instanceID := d.Get("instance_id").(string)
	queueID := d.Get("queue_id").(string)

	input := &connect.AssociateQueueQuickConnectsInput{
		InstanceId:      aws.String(instanceID),
		QueueId:         aws.String(queueID),
		QuickConnectIds: flex.ExpandStringSet(d.Get("quick_connect_ids").(*schema.Set)),
	}

	_, err := conn.AssociateQueueQuickConnectsWithContext(ctx, input)

	if err != nil {
		return diag.Errorf("creating Connect Queue Quick Connect association (%s): %s", d.Id(), err)
	}

	d.SetId(QueueQuickConnectAssociationCreateResourceID(instanceID, queueID))

	return ResourceQueueQuickConnectAssociationRead(ctx, d, meta)
}

func ResourceQueueQuickConnectAssociationRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	conn := meta.(*conns.AWSClient).ConnectConn(ctx)

	instanceID, queueID, err := QueueQuickConnectAssociationParseResourceID(d.Id())

	if err != nil {
		return diag.FromErr(err)
	}

	queueQuickConnects, err := GetQueueQuickConnectIDs(ctx, conn, instanceID, queueID)

	if !d.IsNewResource() && tfawserr.ErrCodeEquals(err, connect.ErrCodeResourceNotFoundException) {
		log.Printf("[WARN] Connect Queue Quick Connect Association (%s) not found, removing from state", d.Id())
		d.SetId("")
		return nil
	}

	if err != nil {
		return diag.Errorf("getting Queue Quick Connect Association (%s): %s", d.Id(), err)
	}

	d.Set("instance_id", aws.String(instanceID))
	d.Set("queue_id", aws.String(queueID))
	d.Set("quick_connect_ids", flex.FlattenStringSet(queueQuickConnects))

	return nil
}

func ResourceQueueQuickConnectAssociationUpdate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	conn := meta.(*conns.AWSClient).ConnectConn(ctx)

	instanceID, queueID, err := QueueQuickConnectAssociationParseResourceID(d.Id())

	if err != nil {
		return diag.FromErr(err)
	}

	if d.HasChange("quick_connect_ids") {
		o, n := d.GetChange("quick_connect_ids")

		if o == nil {
			o = new(schema.Set)
		}
		if n == nil {
			n = new(schema.Set)
		}

		os := o.(*schema.Set)
		ns := n.(*schema.Set)
		quickConnectIdsUpdateAdd := ns.Difference(os)
		quickConnectIdsUpdateRemove := os.Difference(ns)

		if len(quickConnectIdsUpdateAdd.List()) > 0 {
			_, err = conn.AssociateQueueQuickConnectsWithContext(ctx, &connect.AssociateQueueQuickConnectsInput{
				InstanceId:      aws.String(instanceID),
				QueueId:         aws.String(queueID),
				QuickConnectIds: flex.ExpandStringSet(quickConnectIdsUpdateAdd),
			})
			if err != nil {
				return diag.Errorf("updating Queues Quick Connect IDs, specifically associating quick connects to queue (%s): %s", d.Id(), err)
			}
		}

		if len(quickConnectIdsUpdateRemove.List()) > 0 {
			_, err = conn.DisassociateQueueQuickConnectsWithContext(ctx, &connect.DisassociateQueueQuickConnectsInput{
				InstanceId:      aws.String(instanceID),
				QueueId:         aws.String(queueID),
				QuickConnectIds: flex.ExpandStringSet(quickConnectIdsUpdateRemove),
			})
			if err != nil {
				return diag.Errorf("updating Queues Quick Connect IDs, specifically disassociating quick connects from queue (%s): %s", d.Id(), err)
			}
		}
	}

	return ResourceQueueQuickConnectAssociationRead(ctx, d, meta)
}

func ResourceQueueQuickConnectAssociationDelete(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	conn := meta.(*conns.AWSClient).ConnectConn(ctx)

	instanceID, queueID, err := QueueQuickConnectAssociationParseResourceID(d.Id())

	if err != nil {
		return diag.FromErr(err)
	}

	_, err = conn.DisassociateQueueQuickConnectsWithContext(ctx, &connect.DisassociateQueueQuickConnectsInput{
		InstanceId:      aws.String(instanceID),
		QueueId:         aws.String(queueID),
		QuickConnectIds: flex.ExpandStringSet(d.Get("quick_connect_ids").(*schema.Set)),
	})

	if err != nil {
		return diag.Errorf("deleting Queue Quick Connect Association (%s): %s", d.Id(), err)
	}

	return nil
}

func QueueQuickConnectAssociationParseResourceID(id string) (string, string, error) {
	parts := strings.SplitN(id, queueQuickConnectAssociationIDSeparator, 2)
	if len(parts) != 2 || parts[0] == "" || parts[1] == "" {
		return "", "",
			fmt.Errorf("unexpected format for ID (%q), expected instanceID"+queueQuickConnectAssociationIDSeparator+
				"queueID", id)
	}

	return parts[0], parts[1], nil
}

func QueueQuickConnectAssociationCreateResourceID(instanceID string, queueID string) string {
	parts := []string{instanceID, queueID}
	id := strings.Join(parts, queueQuickConnectAssociationIDSeparator)

	return id
}
