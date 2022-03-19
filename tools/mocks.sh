#!/bin/sh

INTERFACE_FOLDER="src/interfaces"

echo 'Starting to generate mocks...'

for entry in ${INTERFACE_FOLDER}/*; do
  mockgen -source=$entry -destination=mockgen/mock_`basename $entry` -package=mockgen
  echo 'Generated mock for '$entry
done

echo 'Finished generating mocks!'
