// Copyright IBM Corp. 2023 All Rights Reserved.
// Licensed under the Mozilla Public License v2.0

package scc_test

import (
	"fmt"
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/acctest"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/v2/terraform"

	acc "github.com/IBM-Cloud/terraform-provider-ibm/ibm/acctest"
	"github.com/IBM-Cloud/terraform-provider-ibm/ibm/conns"
	"github.com/IBM/scc-go-sdk/v5/securityandcompliancecenterapiv3"
)

func TestAccIbmSccControlLibraryBasic(t *testing.T) {
	var conf securityandcompliancecenterapiv3.ControlLibrary
	controlLibraryName := fmt.Sprintf("tf_control_library_name_%d", acctest.RandIntRange(10, 100))
	controlLibraryDescription := fmt.Sprintf("tf_control_library_description_%d", acctest.RandIntRange(10, 100))
	controlLibraryType := "custom"
	controlLibraryNameUpdate := controlLibraryName
	controlLibraryDescriptionUpdate := controlLibraryDescription
	controlLibraryTypeUpdate := controlLibraryType

	resource.Test(t, resource.TestCase{
		PreCheck:     func() { acc.TestAccPreCheck(t) },
		Providers:    acc.TestAccProviders,
		CheckDestroy: testAccCheckIbmSccControlLibraryDestroy,
		Steps: []resource.TestStep{
			resource.TestStep{
				Config: testAccCheckIbmSccControlLibraryConfigBasic(controlLibraryName, controlLibraryDescription, controlLibraryType),
				Check: resource.ComposeAggregateTestCheckFunc(
					testAccCheckIbmSccControlLibraryExists("ibm_scc_control_library.scc_control_library_instance", conf),
					resource.TestCheckResourceAttr("ibm_scc_control_library.scc_control_library_instance", "control_library_name", controlLibraryName),
					resource.TestCheckResourceAttr("ibm_scc_control_library.scc_control_library_instance", "control_library_description", controlLibraryDescription),
					resource.TestCheckResourceAttr("ibm_scc_control_library.scc_control_library_instance", "control_library_type", controlLibraryType),
				),
			},
			resource.TestStep{
				Config: testAccCheckIbmSccControlLibraryConfigBasic(controlLibraryNameUpdate, controlLibraryDescriptionUpdate, controlLibraryTypeUpdate),
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr("ibm_scc_control_library.scc_control_library_instance", "control_library_name", controlLibraryNameUpdate),
					resource.TestCheckResourceAttr("ibm_scc_control_library.scc_control_library_instance", "control_library_description", controlLibraryDescriptionUpdate),
					resource.TestCheckResourceAttr("ibm_scc_control_library.scc_control_library_instance", "control_library_type", controlLibraryTypeUpdate),
				),
			},
		},
	})
}

func TestAccIbmSccControlLibraryAllArgs(t *testing.T) {
	var conf securityandcompliancecenterapiv3.ControlLibrary
	controlLibraryName := fmt.Sprintf("tf_control_library_name_%d", acctest.RandIntRange(10, 100))
	controlLibraryDescription := fmt.Sprintf("tf_control_library_description_%d", acctest.RandIntRange(10, 100))
	controlLibraryType := "custom"
	versionGroupLabel := "11111111-2222-3333-4444-555555555555"
	controlLibraryVersion := "0.0.1"
	latest := "true"
	controlsCount := "1"

	controlLibraryNameUpdate := controlLibraryName
	controlLibraryDescriptionUpdate := controlLibraryDescription
	controlLibraryTypeUpdate := "custom"
	versionGroupLabelUpdate := versionGroupLabel
	controlLibraryVersionUpdate := "0.0.2"
	latestUpdate := "true"

	resource.Test(t, resource.TestCase{
		PreCheck:     func() { acc.TestAccPreCheck(t) },
		Providers:    acc.TestAccProviders,
		CheckDestroy: testAccCheckIbmSccControlLibraryDestroy,
		Steps: []resource.TestStep{
			resource.TestStep{
				Config: testAccCheckIbmSccControlLibraryConfig(controlLibraryName, controlLibraryDescription, controlLibraryType, versionGroupLabel, controlLibraryVersion, latest),
				Check: resource.ComposeAggregateTestCheckFunc(
					testAccCheckIbmSccControlLibraryExists("ibm_scc_control_library.scc_control_library_instance", conf),
					resource.TestCheckResourceAttr("ibm_scc_control_library.scc_control_library_instance", "control_library_name", controlLibraryName),
					resource.TestCheckResourceAttr("ibm_scc_control_library.scc_control_library_instance", "control_library_description", controlLibraryDescription),
					resource.TestCheckResourceAttr("ibm_scc_control_library.scc_control_library_instance", "control_library_type", controlLibraryType),
					resource.TestCheckResourceAttr("ibm_scc_control_library.scc_control_library_instance", "version_group_label", versionGroupLabel),
					resource.TestCheckResourceAttr("ibm_scc_control_library.scc_control_library_instance", "control_library_version", controlLibraryVersion),
					resource.TestCheckResourceAttr("ibm_scc_control_library.scc_control_library_instance", "latest", latest),
					resource.TestCheckResourceAttr("ibm_scc_control_library.scc_control_library_instance", "controls_count", controlsCount),
				),
			},
			resource.TestStep{
				Config: testAccCheckIbmSccControlLibraryConfig(controlLibraryNameUpdate, controlLibraryDescriptionUpdate, controlLibraryTypeUpdate, versionGroupLabelUpdate, controlLibraryVersionUpdate, latestUpdate),
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr("ibm_scc_control_library.scc_control_library_instance", "control_library_name", controlLibraryNameUpdate),
					resource.TestCheckResourceAttr("ibm_scc_control_library.scc_control_library_instance", "control_library_description", controlLibraryDescriptionUpdate),
					resource.TestCheckResourceAttr("ibm_scc_control_library.scc_control_library_instance", "control_library_type", controlLibraryTypeUpdate),
					resource.TestCheckResourceAttr("ibm_scc_control_library.scc_control_library_instance", "version_group_label", versionGroupLabelUpdate),
					resource.TestCheckResourceAttr("ibm_scc_control_library.scc_control_library_instance", "control_library_version", controlLibraryVersionUpdate),
					resource.TestCheckResourceAttr("ibm_scc_control_library.scc_control_library_instance", "latest", latestUpdate),
					resource.TestCheckResourceAttr("ibm_scc_control_library.scc_control_library_instance", "controls_count", controlsCount),
				),
			},
			resource.TestStep{
				ResourceName:      "ibm_scc_control_library.scc_control_library_instance",
				ImportState:       true,
				ImportStateVerify: true,
			},
		},
	})
}

func testAccCheckIbmSccControlLibraryConfigBasic(controlLibraryName string, controlLibraryDescription string, controlLibraryType string) string {
	return fmt.Sprintf(`
		resource "ibm_scc_control_library" "scc_control_library_instance" {
			control_library_name = "%s"
			control_library_description = "%s"
			control_library_type = "%s"
			version_group_label = "03354ab4-03be-41c0-a469-826fc0262e78"
			latest = true
			controls {
				control_name = "control-name"
				control_id = "1fa45e17-9322-4e6c-bbd6-1c51db08e790"
				control_description = "control_description"
				control_category = "control_category"
				control_tags = [ "control_tags" ]
				control_specifications {
					control_specification_id = "f3517159-889e-4781-819a-89d89b747c85"
					responsibility = "user"
					component_id = "f3517159-889e-4781-819a-89d89b747c85"
					component_name = "f3517159-889e-4781-819a-89d89b747c85"
					environment = "environment"
					control_specification_description = "control_specification_description"
					assessments {
						assessment_id = "rule-a637949b-7e51-46c4-afd4-b96619001bf1"
						assessment_method = "ibm-cloud-rule"
						assessment_type = "automated"
						assessment_description = "assessment_description"
						parameters {
							parameter_display_name = "Sign out due to inactivity in seconds"
                            parameter_name         = "session_invalidation_in_seconds"
							parameter_type = "numeric"
						}
					}
				}
				control_docs {
					control_docs_id = "control_docs_id"
					control_docs_type = "control_docs_type"
				}
				control_requirement = true
				status = "enabled"
			}
		}
	`, controlLibraryName, controlLibraryDescription, controlLibraryType)
}

func testAccCheckIbmSccControlLibraryConfig(controlLibraryName string, controlLibraryDescription string, controlLibraryType string, versionGroupLabel string, controlLibraryVersion string, latest string) string {
	return fmt.Sprintf(`

		resource "ibm_scc_control_library" "scc_control_library_instance" {
			control_library_name = "%s"
			control_library_description = "%s"
			control_library_type = "%s"
			version_group_label = "%s"
			control_library_version = "%s"
			latest = %s
			controls {
				control_name = "control-name"
				control_id = "1fa45e17-9322-4e6c-bbd6-1c51db08e790"
				control_description = "control_description"
				control_category = "control_category"
				control_tags = [ "control_tags" ]
				control_specifications {
					control_specification_id = "f3517159-889e-4781-819a-89d89b747c85"
					responsibility = "user"
					component_id = "f3517159-889e-4781-819a-89d89b747c85"
					component_name = "f3517159-889e-4781-819a-89d89b747c85"
					environment = "environment"
					control_specification_description = "control_specification_description"
					assessments {
						assessment_id = "rule-a637949b-7e51-46c4-afd4-b96619001bf1"
						assessment_method = "ibm-cloud-rule"
						assessment_type = "automated"
						assessment_description = "assessment_description"
						parameters {
							parameter_display_name = "Sign out due to inactivity in seconds"
                            parameter_name         = "session_invalidation_in_seconds"
							parameter_type = "numeric"
						}
					}
				}
				control_docs {
					control_docs_id = "control_docs_id"
					control_docs_type = "control_docs_type"
				}
				control_requirement = true
				status = "enabled"
			}
		}
	`, controlLibraryName, controlLibraryDescription, controlLibraryType, versionGroupLabel, controlLibraryVersion, latest)
}

func testAccCheckIbmSccControlLibraryExists(n string, obj securityandcompliancecenterapiv3.ControlLibrary) resource.TestCheckFunc {

	return func(s *terraform.State) error {
		rs, ok := s.RootModule().Resources[n]
		if !ok {
			return fmt.Errorf("Not found: %s", n)
		}

		securityandcompliancecenterapiClient, err := acc.TestAccProvider.Meta().(conns.ClientSession).SecurityAndComplianceCenterV3()
		if err != nil {
			return err
		}

		getControlLibraryOptions := &securityandcompliancecenterapiv3.GetControlLibraryOptions{}

		getControlLibraryOptions.SetControlLibrariesID(rs.Primary.ID)

		controlLibrary, _, err := securityandcompliancecenterapiClient.GetControlLibrary(getControlLibraryOptions)
		if err != nil {
			return err
		}

		obj = *controlLibrary
		return nil
	}
}

func testAccCheckIbmSccControlLibraryDestroy(s *terraform.State) error {
	securityandcompliancecenterapiClient, err := acc.TestAccProvider.Meta().(conns.ClientSession).SecurityAndComplianceCenterV3()
	if err != nil {
		return err
	}
	for _, rs := range s.RootModule().Resources {
		if rs.Type != "ibm_scc_control_library" {
			continue
		}

		getControlLibraryOptions := &securityandcompliancecenterapiv3.GetControlLibraryOptions{}

		getControlLibraryOptions.SetControlLibrariesID(rs.Primary.ID)

		// Try to find the key
		_, response, err := securityandcompliancecenterapiClient.GetControlLibrary(getControlLibraryOptions)

		if err == nil {
			return fmt.Errorf("scc_control_library still exists: %s", rs.Primary.ID)
		} else if response.StatusCode != 404 {
			return fmt.Errorf("Error checking for scc_control_library (%s) has been destroyed: %s", rs.Primary.ID, err)
		}
	}

	return nil
}
