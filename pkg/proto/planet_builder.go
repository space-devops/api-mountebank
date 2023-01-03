package space

import (
	"github.com/space-devops/api-mountebank/pkg/logger"
	"github.com/space-devops/api-mountebank/pkg/objects"
	"github.com/space-devops/api-mountebank/pkg/utils"
)

func BuildPlanetList(raw []byte, correlationId string) (*PlanetList, error) {
	var planetList objects.PlanetList
	if err := utils.JsonObjectToObject(raw, &planetList, correlationId); err != nil {
		logger.LogPanic("Unable to marshall body response", correlationId)
	}

	var list []*PlanetListItem

	for _, planet := range planetList.Planets {
		gplanet := PlanetListItem{
			Id:   int32(planet.Id),
			Name: planet.Name,
		}

		list = append(list, &gplanet)
	}

	gPlanetList := PlanetList{
		PlanetList: list,
	}

	return &gPlanetList, nil
}

func BuildPlanetDetails(raw []byte, correlationId string) (*PlanetDetails, error) {
	var planet objects.PlanetWrapper
	if err := utils.JsonObjectToObject(raw, &planet, correlationId); err != nil {
		logger.LogPanic("Unable to marshall body response", correlationId)
	}

	gPlanetDetails := PlanetDetails{
		Id:   int32(planet.Planet.Id),
		Name: planet.Planet.Name,
		Type: planet.Planet.Type,
		EquatorialRadius: &EquatorialRadius{
			Value:      planet.Planet.Radius.Value,
			MetricUnit: planet.Planet.Radius.Unit,
		},
		Mass: &Mass{
			Value:      planet.Planet.Mass.Value,
			MetricUnit: planet.Planet.Mass.Unit,
		},
		Volume: &Volume{
			Value:      planet.Planet.Volume.Value,
			MetricUnit: planet.Planet.Volume.Unit,
		},
		Density: &Density{
			Value:      planet.Planet.Density.Value,
			MetricUnit: planet.Planet.Density.Unit,
		},
		Satellites: planet.Planet.Satellites,
	}

	return &gPlanetDetails, nil
}
