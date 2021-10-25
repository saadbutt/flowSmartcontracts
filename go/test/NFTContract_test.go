package test

import (
	"testing"

	"github.com/onflow/cadence"
	"github.com/stretchr/testify/assert"
)

// First TestCase to check Deployment and Supply Field Initialize Correctly
func TestNFTContractDeployment(t *testing.T) {
	b := newEmulator()

	_, _, _ = NowwhereDeployContracts(b, t)

	t.Run("Should have initialized Supply field correctly", func(t *testing.T) {
		//	supply := executeScriptAndCheck(t, b, NowwhereGenerateGetSupplyScript(fungibleAddr, nowwhereAddr), nil)
		//	assert.EqualValues(t, CadenceUInt64(0), supply)
	})
}

// First TestCase to check Deployment and Collection, Schema and Template Initilization
func TestNFTContractDeploymentCheckFields(t *testing.T) {
	b := newEmulator()

	fungibleAddr, nowwhereAddr, _ := NowwhereDeployContracts(b, t)

	t.Run("Should have initialized Supply field correctly", func(t *testing.T) {
		supply := executeScriptAndCheck(t, b, NowwhereGenerateGetCollectionScript(fungibleAddr, nowwhereAddr), nil)
		firstName := cadence.NewUInt64(1)
		hayward := cadence.NewString("")
		metadata := []cadence.KeyValuePair{{Key: firstName, Value: hayward}}

		assert.EqualValues(t, CadenceDictionary(metadata), supply)
	})
}
