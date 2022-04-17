#!/bin/bash
RED='\033[0;31m'
NC='\033[0m' # No Color

echo -e "${RED}cleaning kpimon previous deployment  ...${NC}"

helm -n sdran uninstall kpimon-test 

echo -e "${RED}building latest docker image  ...${NC}"
make images

path=${SDRAN_CHART_DIR}

echo -e "${RED}deploying kpimon ...${NC}"
helm install -n sdran kpimon-test $path/helm-chart-rmet-xapp --wait
echo


