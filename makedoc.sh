#!/bin/bash

cd common
godoc -html ./ > index.html

cd Administration
godoc -html ./ > index.html

cd ../Auth
godoc -html ./ > index.html

cd ../Profile
godoc -html ./ > index.html
