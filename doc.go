/*
Package kpxcconvert is a CLI tool for converting 1Password export files to KeePassXC import files.

Usage

Say you export your 1Password vault data to

	~/Documents/1Password\ 2018-01-09,\ 08_23\ PM\ \(99\ items\ and\ 0\ folders\).1pif

for a vault called "Primary" and you want to generate a KeePassXC CSV-import compatible CSV called

	out.csv

Then you might run:

	kpxcconvert --group Primary --op ~/Documents/1Password\ 2018-01-09,\ 08_23\ PM\ \(99\ items\ and\ 0\ folders\).1pif/data.1pif --kp out.csv
*/
package main
