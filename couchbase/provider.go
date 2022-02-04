package couchbase

import (
	"context"
	"time"

	"github.com/couchbase/gocb/v2"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func Provider() *schema.Provider {
	return &schema.Provider{
		TerraformVersion: terraformVersion,
		Schema: map[string]*schema.Schema{
			providerConnStr: {
				Type:        schema.TypeString,
				Required:    true,
				DefaultFunc: schema.EnvDefaultFunc("CB_CONNECTION_STRING", ""),
			},
			providerUsername: {
				Type:        schema.TypeString,
				Required:    true,
				DefaultFunc: schema.EnvDefaultFunc("CB_USERNAME", ""),
			},
			providerPassword: {
				Type:        schema.TypeString,
				Required:    true,
				DefaultFunc: schema.EnvDefaultFunc("CB_PASSWORD", ""),
				Sensitive:   true,
			},
			providerConnectionTimeout: {
				Type:        schema.TypeInt,
				Required:    false,
				Optional:    true,
				Default:     15,
				DefaultFunc: schema.EnvDefaultFunc("CB_MANAGEMENT_TIMEOUT", ""),
			},
			providerTlsSkipVerify: {
				Type:        schema.TypeBool,
				Required:    false,
				Optional:    true,
				Default:     false,
				DefaultFunc: schema.EnvDefaultFunc("CB_TLS_SKIP_VERIFY", ""),
			},
		},

		ResourcesMap: map[string]*schema.Resource{
			"couchbase_bucket":              resourceBucket(),
			"couchbase_security_group":      resourceSecurityGroup(),
			"couchbase_security_user":       resourceSecurityUser(),
			"couchbase_primary_query_index": resourcePrimaryQueryIndex(),
			"couchbase_query_index":         resourceQueryIndex(),
		},

		ConfigureContextFunc: providerConfigure,
	}
}

func providerConfigure(ctx context.Context, d *schema.ResourceData) (interface{}, diag.Diagnostics) {
	cc := &CouchbaseConnection{
		ConnStr: d.Get(providerConnStr).(string),
		ClusterOptions: gocb.ClusterOptions{
			Username: d.Get(providerUsername).(string),
			Password: d.Get(providerPassword).(string),
			TimeoutsConfig: gocb.TimeoutsConfig{
				ManagementTimeout: time.Duration(d.Get(providerConnectionTimeout).(int)) * time.Second,
			},
			// TODO
			SecurityConfig: gocb.SecurityConfig{
				TLSSkipVerify: d.Get(providerTlsSkipVerify).(bool),
				AllowedSaslMechanisms: []gocb.SaslMechanism{
					gocb.PlainSaslMechanism,
				},
			},
		},
	}

	_, diags := cc.ConnectionValidate()
	if diags != nil {
		return nil, diags
	}

	return cc, diags
}
