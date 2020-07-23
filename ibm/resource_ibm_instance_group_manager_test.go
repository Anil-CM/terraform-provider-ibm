package ibm

import (
	"fmt"
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/helper/acctest"
	"github.com/hashicorp/terraform-plugin-sdk/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/terraform"
	"github.ibm.com/ibmcloud/vpc-go-sdk-scoped/vpcv1"
)

func TestAccIBMISInstanceGroupManager_basic(t *testing.T) {
	name := fmt.Sprintf("terraformigmanager%d", acctest.RandIntRange(10, 100))
	aggWindow := 120
	aggWindowUpdate := 150
	coolDown := 300
	coolDownUpdate := 330
	maxCount := 2
	maxCountUpdate := 2
	minCount := 1
	minCountUpdate := 3

	resource.Test(t, resource.TestCase{
		PreCheck:     func() { testAccPreCheck(t) },
		Providers:    testAccProviders,
		CheckDestroy: testAccCheckIBMISInstanceGroupManagerDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccCheckIBMISInstanceGroupManagerConfig(name, aggWindow, coolDown, maxCount, minCount),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(
						"ibm_is_instance_group_manager.instance_group_manager", "name", name),
					resource.TestCheckResourceAttr(
						"ibm_is_instance_group_manager.instance_group_manager", "aggregation_window", fmt.Sprintf("%d", aggWindow)),
					resource.TestCheckResourceAttr(
						"ibm_is_instance_group.instance_group_manager", "cooldown", fmt.Sprintf("%d", coolDown)),
					resource.TestCheckResourceAttr(
						"ibm_is_instance_group.instance_group_manager", "max_count", fmt.Sprintf("%d", maxCount)),
					resource.TestCheckResourceAttr(
						"ibm_is_instance_group.instance_group_manager", "min_count", fmt.Sprintf("%d", minCount)),
				),
			},
			{
				Config: testAccCheckIBMISInstanceGroupManagerConfig(name, aggWindowUpdate, coolDownUpdate, maxCountUpdate, minCountUpdate),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(
						"ibm_is_instance_group_manager.instance_group_manager", "name", name),
					resource.TestCheckResourceAttr(
						"ibm_is_instance_group_manager.instance_group_manager", "aggregation_window", fmt.Sprintf("%d", aggWindowUpdate)),
					resource.TestCheckResourceAttr(
						"ibm_is_instance_group.instance_group_manager", "cooldown", fmt.Sprintf("%d", coolDownUpdate)),
					resource.TestCheckResourceAttr(
						"ibm_is_instance_group.instance_group_manager", "max_count", fmt.Sprintf("%d", maxCountUpdate)),
					resource.TestCheckResourceAttr(
						"ibm_is_instance_group.instance_group_manager", "min_count", fmt.Sprintf("%d", minCountUpdate)),
				),
			},
		},
	})
}

func testAccCheckIBMISInstanceGroupManagerDestroy(s *terraform.State) error {
	sess, _ := testAccProvider.Meta().(ClientSession).VpcV1APIScoped()
	for _, rs := range s.RootModule().Resources {
		if rs.Type != "ibm_is_instance_group_manager" {
			continue
		}

		instanceGroup := rs.Primary.Attributes["instance_group"]
		getInstanceGroupManagerOptions := vpcv1.GetInstanceGroupManagerOptions{
			ID:              &rs.Primary.ID,
			InstanceGroupID: &instanceGroup,
		}
		_, _, err := sess.GetInstanceGroupManager(&getInstanceGroupManagerOptions)
		if err == nil {
			return fmt.Errorf("instance group still exists: %s", rs.Primary.ID)
		}
	}
	return nil
}

func testAccCheckIBMISInstanceGroupManagerConfig(name string, aggregationWindow, cooldown, maxCount, MinCount int) string {
	return fmt.Sprintf(`
	
	resource "ibm_is_lb" "lb" {
		name    = "lbtest"
		subnets = ["0737-2e916ac5-58d6-4f7b-9843-c7fae08b5953"]
	  }
	  
	resource "ibm_is_lb_pool" "testacc_pool" {
		name           = "testpooltest"
		lb             = ibm_is_lb.lb.id
		algorithm      = "round_robin"
		protocol       = "http"
		health_delay   = 60
		health_retries = 5
		health_timeout = 30
		health_type    = "http"
	  }

	resource "ibm_is_instance_group" "instance_group" {
		name =  "testacc"
		instance_template = "0737-bef876ed-3b51-40eb-9f94-beb26c1694e0"
		instance_count = 4
		load_balancer = ibm_is_lb.lb.id
		load_balancer_pool = element(split("/", ibm_is_lb_pool.testacc_pool.id), 1)
		subnets = ["0737-2e916ac5-58d6-4f7b-9843-c7fae08b5953"]
		application_port = 9009
	  }

	resource "ibm_is_instance_group_manager" "instance_group_manager" {
		name = "%s"
		aggregation_window = %d
		instance_group = ibm_is_instance_group.instance_group.id
		cooldown = %d
		manager_type = "autoscale"
		enable_manager = true
		max_membership_count = %d
		min_membership_count = %d
	  }
`, name, aggregationWindow, cooldown, maxCount, MinCount)

}
