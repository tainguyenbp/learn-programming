# Ansible
```

```

# Terraform
```
Terraform Lifecycle Ignore changes
1. Ignore Changes to a Specific Attribute
resource "aws_instance" "example" {
  ami           = "ami-12345678"
  instance_type = "t2.micro"
  tags = {
    Name = "example-instance"
  }
  lifecycle {
    ignore_changes = [
      tags["Name"]  # Bỏ qua thay đổi của thẻ Name
    ]
  }
}

2. Ignore an Entire Attribute
resource "aws_instance" "example" {
  ami           = "ami-12345678"
  instance_type = "t2.micro" 
  root_block_device {
    volume_size = 50
  }
  lifecycle {
    ignore_changes = [
      root_block_device  # Bỏ qua mọi thay đổi trong root_block_device
    ]
  }
}

3. Ignore Multiple Attributes
resource "aws_instance" "example" {
  ami           = "ami-12345678"
  instance_type = "t2.micro"
  tags = {
    Environment = "production"
    Owner       = "team-a"
  }
  lifecycle {
    ignore_changes = [
      tags["Environment"],  # Bỏ qua thay đổi của Environment
      instance_type         # Bỏ qua thay đổi của instance_type
    ]
  }
}

4. Ignore Changes in a Module
module "example" {
  source = "./module-example"
  instance_type = "t2.micro"
}
lifecycle {
  ignore_changes = [
    instance_type  # Ignore changes to instance_type in the module
  ]
}

Import terraform with exsist resource

variable.tf
output.tf
main.tf

module/elasticache
module/eks
module/rds
module/ec2
module/vpc
module/subnet
module/ecs

```
# Tracing, Tracking, Application Performance Management
```
Sentry
APM ElasticSearch
OpenTeleMetry
```

# SAST Tools
```
Bitbucket
Jenkins: every the repository have to Jenkinfiles template, this template will filter and define all service in one Jenkinfiles, The Jenkin Master installed in EKS and Slave use EC2.

Runing CI Parrarel 
       Scan static code -> success -> done
     /                  \ 
CI: |                     failed -> send this log or this ouput to PR - u can use hook or api.
     \
       Build docker image -> scan docker image -> success -> push docker image
                                               \ failed -> send this log or this ouput to PR - u can use hook or api.
Sobelow
Snyk
SonarQube
```
