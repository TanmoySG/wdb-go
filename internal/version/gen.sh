# version.go generator

BASEDIR=$(dirname "$0")

version=$(yq '.LIBRARY_VERSION' ${BASEDIR}/version.yaml)

gen_go="// This is a Generated File. DO NOT EDIT MANUALLY\npackage version\nvar Version string=\""${version}"\""

echo $gen_go >${BASEDIR}/version.go
