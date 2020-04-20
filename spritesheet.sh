#!/bin/bash

cd assets/tiles/
convert grass.png stone.png dir.png water.png +append ../../src/sprite.png
