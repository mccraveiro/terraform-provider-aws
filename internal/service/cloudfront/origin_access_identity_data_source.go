package cloudfront

import (
	"fmt"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/arn"
	"github.com/aws/aws-sdk-go/service/cloudfront"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-provider-aws/internal/conns"
)

func DataSourceOriginAccessIdentity() *schema.Resource {
	return &schema.Resource{
		Read: dataSourceOriginAccessIdentityRead,

		Schema: map[string]*schema.Schema{
			"comment": {
				Type:     schema.TypeString,
				Computed: true,
				Default:  nil,
			},
			"caller_reference": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"cloudfront_access_identity_path": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"etag": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"iam_arn": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"id": {
				Type:     schema.TypeString,
				Required: true,
			},
			"s3_canonical_user_id": {
				Type:     schema.TypeString,
				Computed: true,
			},
		},
	}
}

func dataSourceOriginAccessIdentityRead(d *schema.ResourceData, meta interface{}) error {
	conn := meta.(*conns.AWSClient).CloudFrontConn()
	id := d.Get("id").(string)
	params := &cloudfront.GetCloudFrontOriginAccessIdentityInput{
		Id: aws.String(id),
	}

	resp, err := conn.GetCloudFrontOriginAccessIdentity(params)
	if err != nil {
		return err
	}

	// Update attributes from DistributionConfig
	flattenOriginAccessIdentityConfig(d, resp.CloudFrontOriginAccessIdentity.CloudFrontOriginAccessIdentityConfig)
	// Update other attributes outside of DistributionConfig
	d.SetId(aws.StringValue(resp.CloudFrontOriginAccessIdentity.Id))
	d.Set("etag", resp.ETag)
	d.Set("s3_canonical_user_id", resp.CloudFrontOriginAccessIdentity.S3CanonicalUserId)
	d.Set("cloudfront_access_identity_path", fmt.Sprintf("origin-access-identity/cloudfront/%s", *resp.CloudFrontOriginAccessIdentity.Id))
	iamArn := arn.ARN{
		Partition: meta.(*conns.AWSClient).Partition,
		Service:   "iam",
		AccountID: "cloudfront",
		Resource:  fmt.Sprintf("user/CloudFront Origin Access Identity %s", *resp.CloudFrontOriginAccessIdentity.Id),
	}.String()
	d.Set("iam_arn", iamArn)
	return nil
}
