package client

import (
	"strconv"
	"strings"
	"time"

	log "github.com/sirupsen/logrus"
)

func AutoScale(service string) {

	log.Infof("Autoscaling %s", service)

	//La función entra en un bucle infinito, ejecutándose para monitorear y escalar el servicio continuamente.
	for {
		var avgCpuUsage float64
		//Obtener Info del Servicio
		stats, err := GetInfoByService(service)
		if err != nil {
			log.Errorf("Error getting %s stats: %v", service, err)
			continue
		}

		containersAmount := len(stats)

		//Calcular el Uso Promedio de CPU
		for _, container := range stats {

			stringCPU := strings.Trim(container.CPU, "%")
			intCPU, err := strconv.ParseFloat(stringCPU, 64)
			if err != nil {
				log.Errorf("Error parsing string: %v", err)
				continue
			}

			avgCpuUsage += intCPU
		}

		avgCpuUsage = avgCpuUsage / float64(containersAmount)

		//Escalado hacia arriba
		//Si el uso promedio de CPU es mayor o igual al 60% o el número de contenedores es menor que 2,
		//escala el servicio hacia arriba llamando a ScaleService.
		if avgCpuUsage >= 60 || containersAmount < 2 {
			instances, err := ScaleService(service)
			if err != nil {
				log.Errorf("Error creating %s container: %s", service, err)
				continue
			}

			log.Infof("Scaling up %s to %d instances", service, instances)

			//Escalado hacia abajo
			//Si el uso promedio de CPU es menor que 20% y el número de contenedores es mayor que 2,
			//elimina un contenedor llamando a DeleteContainer.

			err = DeleteContainer(stats[containersAmount-1].Id)
			if err != nil {
				log.Errorf("Error deleting %s container: %s", service, err)
				continue
			}

			log.Infof("Scaling down %s to %d instances", service, containersAmount-1)
		}
		//Duermo la funcion por 20 segundos
		time.Sleep(20 * time.Second)
	}
}
