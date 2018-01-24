# Use goreman (`go get -u github.com/mattn/goreman`) to run
#   goreman -f Procfile start
#
# This bootstraps a new cluster. To join/restart remove
# the `--cluster-bootstrap` flag
#
# https://devcenter.heroku.com/articles/procfile
#
journeys1: journeys start cluster1 --id 1 --cluster-port=4747 --cluster-bootstrap
journeys2: journeys start cluster1 --id 2 --cluster-port=4848
journeys3: journeys start cluster1 --id 3 --cluster-port=4949