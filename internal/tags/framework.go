// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package tags

import (
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

// Terraform Plugin Framework variants of tags schemas.

func TagsAttribute() schema.Attribute {
	return schema.MapAttribute{
		CustomType:  MapType,
		ElementType: types.StringType,
		Optional:    true,
	}
}

func TagsAttributeComputedOnly() schema.Attribute {
	return schema.MapAttribute{
		ElementType: types.StringType,
		Computed:    true,
	}
}

var (
	Unknown = types.MapUnknown(types.StringType)
)

var (
	Null = NewMapValueNull()
)
