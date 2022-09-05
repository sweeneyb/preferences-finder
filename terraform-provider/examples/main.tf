terraform {
  required_providers {
    pf = {
      version = "0.1.0"
      source  = "briansweeney-dev/dev/preferences-finder"
    }
  }
}

data "pf_works" "all" {}

# Returns all products
output "all_works" {
  value = data.pf_works.all.works
}

