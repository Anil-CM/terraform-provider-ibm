variable "lb_name" {
  default = "testlb"
}

variable "subnet1" {
  default = "0737-2e916ac5-58d6-4f7b-9843-c7fae08b5953"
}

variable "lb_pool_name" {
  default = "teslbpool"
}

variable "instance_group_name" {
  default = "testacc_ig"
}

variable "template_id" {
  default = "0737-bef876ed-3b51-40eb-9f94-beb26c1694e0"
}

variable "instance_count" {
  default = 5
}

variable "application_port" {
  default = 9000
}

variable "manager_name" {
  default = "testacc_ig_manager"
}

variable "aggregation_window" {
  default = 300
}

variable "cooldown_period" {
  default = 270
}

variable "max_count" {
  default = 3
}

variable "min_count" {
  default = 1
}

variable "policy_name" {
  default = "testpolicy"
}