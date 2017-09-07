# manifest

[![GoDoc](
	https://godoc.org/github.com/skycoin/viscript?status.svg)](
		https://godoc.org/github.com/ZSM5J/manifest/manifest)


All functionality in cmd folder. Open read or merge folder and start executable file from command line. 

If you want create snapshot:

1. Open cmd/read/config.json and change folder's path
2. When open terminal and execute ./read

If you need to merge snapshots:

1. Go to cmd/merge
2. When open terminal and execute ./merge
3. In console you will see snap shot list with ID
4. Right into terminal merge "1st snapshot id" "2n snapshot id" and so on. Example "merge 1 5 8" 
