terraform {
  cloud {
    organization = "hclvelocity"

    workspaces {
      name = "AKS-Cluster"
    }
  }
}
