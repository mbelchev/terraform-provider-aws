// Code generated by internal/generate/servicepackages/main.go; DO NOT EDIT.

package workspaces

import (
	"context"

	aws_sdkv2 "github.com/aws/aws-sdk-go-v2/aws"
	workspaces_sdkv2 "github.com/aws/aws-sdk-go-v2/service/workspaces"
	"github.com/hashicorp/terraform-provider-aws/internal/conns"
	"github.com/hashicorp/terraform-provider-aws/internal/types"
	"github.com/hashicorp/terraform-provider-aws/names"
)

type servicePackage struct{}

func (p *servicePackage) FrameworkDataSources(ctx context.Context) []*types.ServicePackageFrameworkDataSource {
	return []*types.ServicePackageFrameworkDataSource{}
}

func (p *servicePackage) FrameworkResources(ctx context.Context) []*types.ServicePackageFrameworkResource {
	return []*types.ServicePackageFrameworkResource{
		{
			Factory: newResourceConnectionAlias,
			Name:    "Connection Alias",
			Tags: &types.ServicePackageResourceTags{
				IdentifierAttribute: names.AttrID,
			},
		},
	}
}

func (p *servicePackage) SDKDataSources(ctx context.Context) []*types.ServicePackageSDKDataSource {
	return []*types.ServicePackageSDKDataSource{
		{
			Factory:  DataSourceBundle,
			TypeName: "aws_workspaces_bundle",
		},
		{
			Factory:  DataSourceDirectory,
			TypeName: "aws_workspaces_directory",
		},
		{
			Factory:  DataSourceImage,
			TypeName: "aws_workspaces_image",
		},
		{
			Factory:  DataSourceWorkspace,
			TypeName: "aws_workspaces_workspace",
		},
	}
}

func (p *servicePackage) SDKResources(ctx context.Context) []*types.ServicePackageSDKResource {
	return []*types.ServicePackageSDKResource{
		{
			Factory:  ResourceDirectory,
			TypeName: "aws_workspaces_directory",
			Name:     "Directory",
			Tags: &types.ServicePackageResourceTags{
				IdentifierAttribute: names.AttrID,
			},
		},
		{
			Factory:  ResourceIPGroup,
			TypeName: "aws_workspaces_ip_group",
			Name:     "IP Group",
			Tags: &types.ServicePackageResourceTags{
				IdentifierAttribute: names.AttrID,
			},
		},
		{
			Factory:  ResourceWorkspace,
			TypeName: "aws_workspaces_workspace",
			Name:     "Workspace",
			Tags: &types.ServicePackageResourceTags{
				IdentifierAttribute: names.AttrID,
			},
		},
	}
}

func (p *servicePackage) ServicePackageName() string {
	return names.WorkSpaces
}

// NewClient returns a new AWS SDK for Go v2 client for this service package's AWS API.
func (p *servicePackage) NewClient(ctx context.Context, config map[string]any) (*workspaces_sdkv2.Client, error) {
	cfg := *(config["aws_sdkv2_config"].(*aws_sdkv2.Config))

	return workspaces_sdkv2.NewFromConfig(cfg, func(o *workspaces_sdkv2.Options) {
		if endpoint := config[names.AttrEndpoint].(string); endpoint != "" {
			o.BaseEndpoint = aws_sdkv2.String(endpoint)
		}
	}), nil
}

func ServicePackage(ctx context.Context) conns.ServicePackage {
	return &servicePackage{}
}
