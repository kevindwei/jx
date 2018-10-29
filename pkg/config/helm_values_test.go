package config_test

import (
	"gopkg.in/yaml.v2"
	"testing"

	"io/ioutil"

	"github.com/jenkins-x/jx/pkg/config"
	"github.com/stretchr/testify/assert"
)

func TestEnvironmentExposecontrollerHelmValues(t *testing.T) {
	t.Parallel()

	testFile, err := ioutil.ReadFile("helm_values_test.yaml")
	assert.NoError(t, err)
	helmValuesFromFile := config.HelmValuesConfig{}
	err = yaml.Unmarshal(testFile, &helmValuesFromFile)

	a := make(map[string]string)
	a["helm.sh/hook"] = "post-install,post-upgrade"
	a["helm.sh/hook-delete-policy"] = "hook-succeeded"

	values := config.HelmValuesConfig{
		ExposeController: &config.ExposeController{},
	}

	values.ExposeController.Annotations = a
	values.ExposeController.Config.Exposer = "Ingress"
	values.ExposeController.Config.Domain = "jenkinsx.io"
	values.ExposeController.Config.HTTP = "false"
	values.ExposeController.Config.TLSAcme = "false"
	assert.Equal(t, helmValuesFromFile, values, "expected exposecontroller helm values do not match")
}
