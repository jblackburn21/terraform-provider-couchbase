############
# Provider #
############

terraform {
  required_providers {
    couchbase = {
      version = "0.1.0"
      source  = "github.com/jblackburn21/couchbase"
    }
  }
}

provider "couchbase" {
  connection_string  = "couchbase://localhost"
  username           = "Administrator"
#  password           = "password"
  management_timeout = 10
#  tls_skip_verify    = true
}

resource "couchbase_bucket" "tf_example" {
  name                     = "tf_example"
  ram_quota_mb             = 256
  flush_enabled            = false
  max_expire               = 0
  conflict_resolution_type = "seqno"
  compression_mode         = "passive"
  num_replicas             = 1
}


