package ibm

import (
	"fmt"
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/helper/acctest"
	"github.com/hashicorp/terraform-plugin-sdk/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/terraform"
	"github.ibm.com/ibmcloud/vpc-go-sdk-scoped/vpcv1"
)

func TestAccIBMISInstanceGroup_basic(t *testing.T) {
	name := fmt.Sprintf("terraformigroup%d", acctest.RandIntRange(10, 100))
	instances := acctest.RandIntRange(1, 4)
	port := acctest.RandIntRange(31000, 40000)
	instancesUpdate := 5

	resource.Test(t, resource.TestCase{
		PreCheck:     func() { testAccPreCheck(t) },
		Providers:    testAccProviders,
		CheckDestroy: testAccCheckIBMISInstanceGroupDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccCheckIBMISInstanceGroupConfig(name, instancesUpdate, port),
				Check: resource.ComposeTestCheckFunc(
					//testAccCheckIBMISInstanceGroupExists("ibm_is_instance_group.testacc_instance_group", instancegroup),
					resource.TestCheckResourceAttr(
						"ibm_is_instance_group.instance_group", "name", name),
					resource.TestCheckResourceAttr(
						"ibm_is_instance_group.instance_group", "instance_count", fmt.Sprintf("%d", instances)),
					resource.TestCheckResourceAttr(
						"ibm_is_instance_group.instance_group", "application_port", fmt.Sprintf("%d", port)),
				),
			},
			{
				Config: testAccCheckIBMISInstanceGroupConfig(name, instances, port),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(
						"ibm_is_instance_group.instance_group", "name", name),
					resource.TestCheckResourceAttr(
						"ibm_is_instance_group.instance_group", "instance_count", fmt.Sprintf("%d", instancesUpdate)),
				),
			},
		},
	})
}

func testAccCheckIBMISInstanceGroupDestroy(s *terraform.State) error {
	sess, _ := testAccProvider.Meta().(ClientSession).VpcV1APIScoped()
	for _, rs := range s.RootModule().Resources {
		if rs.Type != "ibm_is_instance_group" {
			continue
		}

		getInstanceGroupOptions := vpcv1.GetInstanceGroupOptions{ID: &rs.Primary.ID}
		_, _, err := sess.GetInstanceGroup(&getInstanceGroupOptions)

		if err == nil {
			return fmt.Errorf("instance group still exists: %s", rs.Primary.ID)
		}
	}
	return nil
}

func testAccCheckIBMISInstanceGroupConfig(name string, instanceCount, applicationPort int) string {
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
		name =  "%s"
		instance_template = "0737-bef876ed-3b51-40eb-9f94-beb26c1694e0"
		instance_count = %d
		load_balancer = ibm_is_lb.lb.id
		load_balancer_pool = element(split("/", ibm_is_lb_pool.testacc_pool.id), 1)
		subnets = ["0737-2e916ac5-58d6-4f7b-9843-c7fae08b5953"]
		application_port = %d
	  }
`, name, instanceCount, applicationPort)

}
