#!/bin/bash
RED='\033[0;31m'
NC='\033[0m' # No Color

path=${SDRAN_CHART_DIR}

# echo -e "${RED}cleaning onos-e2t ...${NC}"
# helm uninstall -n sdran onos-e2t

# echo -e "${RED}deploying onos-e2t ...${NC}"
# helm install -n sdran onos-e2t $path/sdran-helm-charts/onos-e2t --wait

echo -e "${RED}cleaning kpimon previous deployment  ...${NC}"

helm -n sdran uninstall kpimon-test 

echo -e "${RED}building latest docker image  ...${NC}"
make images

echo -e "${RED}deploying kpimon ...${NC}"
helm install -n sdran kpimon-test $path/sd-ran-helm-charts/onos-kpimon --wait
echo


