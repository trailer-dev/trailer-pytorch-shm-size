package recommendation

import (
	"encoding/json"
	"math"
)

const Bytes4GB = 4294967296.0

func environmentHasPytorch(environment Environment) bool {
	for _, pkg := range environment.Packages {
		if pkg.Name == "pytorch" || pkg.Name == "torch" {
			return true
		}
	}
	return false
}

func Match(data []byte) (bool, error) {
	var workspace Workspace

	if err := json.Unmarshal(data, &workspace); err != nil {
		return false, err
	}

	// Check if ANY environment has pytorch
	for _, env := range workspace.Image.Configuration.Environments {
		if environmentHasPytorch(env) {
			return true, nil
		}
	}

	// we got this far, so none of the environments have pytorch, this recommendation is not applicable
	return false, nil
}

func Recommend(data []byte) ([]byte, error) {
	var workspace Workspace

	if err := json.Unmarshal(data, &workspace); err != nil {
		return nil, err
	}

	workspace.Resources.ShmSize = int64(math.Min(Bytes4GB, float64(workspace.Host.TotalMemory/4)))

	// Convert back to JSON
	updatedWorkspace, err := json.Marshal(workspace)
	if err != nil {
		return nil, err
	}

	return updatedWorkspace, nil
}
