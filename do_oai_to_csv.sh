#!/bin/bash

#Set your Metha extracts directory
oai_dir="/home/justin/.cache/metha"

printf 'Git pull \n'
#git pull

printf 'Metha-sync \n'
while IFS="" read -r p || [ -n "$p" ]
do
  oai_code=`echo $p | cut -d" " -f 1`
  oai_url=`echo $p | cut -d" " -f 2`
  oai_set=`echo $p | cut -d" " -f 3`
  printf '%s\n' "$p"
  printf "$oai_code" 
  printf '\n'
  printf $oai_url
  printf '\n'
  printf $oai_set
  printf '\n'
  printf '\n'
  printf 'OAI Set \n'
  if [[ ! -z "$oai_set" ]] ; then
    oai_set_url=" -set $oai_set"
  else
    oai_set_url=null
  fi
  printf $oai_set_url

  if [[ $oai_code == "NAME" ]] ; then
    continue
  fi

  printf '\n'
  printf 'Metha-cat \n'
  printf '\n'
  echo $oai_set_url
  printf '\n'
  if [[ ! -z "$oai_set" ]] ; then
  METHA_DIR=$oai_dir metha-cat -set $oai_set $oai_url > input-xml/$oai_code.xml
  else
  METHA_DIR=$oai_dir metha-cat $oai_url > input-xml/$oai_code.xml
  fi
  go run oai-xml-2-csv-$oai_code.go
  printf '\n'
done < provider_oai_list.txt

printf 'Finish \n'
#printf 'Run provider scripts \n'
#./get_provider_info.sh
#./get_provider_xml.sh

#printf 'Git commit and push \n'
#git add oai_extracts/*
#git commit -am 'Auto commit'
#git push
