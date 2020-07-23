provider "ibm" {
    generation = 2
}

resource "ibm_is_lb" "lb" {
  name    = var.lb_name
  subnets = [var.subnet1]
}

resource "ibm_is_lb_pool" "testacc_pool" {
  name           = var.lb_pool_name
  lb             = ibm_is_lb.lb.id
  algorithm      = "round_robin"
  protocol       = "http"
  health_delay   = 60
  health_retries = 5
  health_timeout = 30
  health_type    = "http"
}

resource "ibm_is_instance_group" "instance_group" {
    name =  var.instance_group_name
    instance_template = var.template_id
    instance_count = var.instance_count
    load_balancer = ibm_is_lb.lb.id
    load_balancer_pool = element(split("/", ibm_is_lb_pool.testacc_pool.id), 1)
    subnets = [var.subnet1]
    application_port = var.application_port
}

resource "ibm_is_instance_group_manager" "instance_group_manager" {
    name = var.manager_name
    aggregation_window = var.aggregation_window
    instance_group = ibm_is_instance_group.instance_group.id
    cooldown = var.cooldown_period
    manager_type = "autoscale"
    enable_manager = false
    max_membership_count = var.max_count
    min_membership_count = var.min_count
}

resource "ibm_is_instance_group_manager_policy" "cpuPolicy" {
    instance_group = ibm_is_instance_group.instance_group.id
    instance_group_manager =  ibm_is_instance_group_manager.instance_group_manager.id
    metric_type = "cpu"
    metric_value = 70
    policy_type = "target"
    name = var.policy_name
}

data "ibm_is_instance_group" "instance_group" {
    name =  "mygroup"
}

data "ibm_is_instance_group_manager" "instance_group_manager" {
    instance_group = ibm_is_instance_group.instance_group.id
}

data "ibm_is_instance_group_manager_policy" "instance_group_manager_policy" {
    instance_group = ibm_is_instance_group.instance_group.id
    instance_group_manager = ibm_is_instance_group_manager.instance_group_manager.id
}