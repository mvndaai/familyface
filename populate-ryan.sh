#!/bin/bash

host=localhost:9090
#To get an fbid http://findmyfbid.com/

{ #This hides the output

  curl $host/dump #Clear data

  ryan=`curl $host/person/add -d '{"givennames": "Ryan", "surname":"Petersen", "birthday":"05/14/1985", "occupation": "QA", "fbid":"562515263" }'`
  jess=`curl $host/person/add -d '{"givennames": "Jessica", "surname":"Petersen", "birthday":"05/14/1985", "fbid":"519727488"}'`
  curl $host/person/add/marriage -d "{\"person_id\": \"$ryan\", \"spouse_id\":\"$jess\"}"
  mom=`curl $host/person/add -d '{"givennames": "Tracy", "surname":"Petersen", "birthday":"07/07/1954", "occupation": "retired", "pic":"tracy.jpg"}'`
  dad=`curl $host/person/add -d '{"givennames": "Roger", "surname":"Petersen", "birthday":"07/11/1954", "occupation": "Healthcare", "pic":"dad.jpg"}'`
  curl $host/person/add/marriage -d "{\"person_id\": \"$mom\", \"spouse_id\":\"$dad\"}"
  curl $host/person/add/mom -d "{\"child_id\": \"$ryan\", \"mom_id\":\"$mom\"}"
  curl $host/person/add/dad -d "{\"child_id\": \"$ryan\", \"dad_id\":\"$dad\"}"
  ady=`curl $host/person/add -d '{"givennames": "Adalynn", "surname":"Petersen", "birthday":"03/03/2011", "pic":"ady.jpg"}'`
  hayden=`curl $host/person/add -d '{"givennames": "Hayden", "surname":"Petersen", "birthday":"01/17/2014", "pic":"hay.jpg"}'`
  tobin=`curl $host/person/add -d '{"givennames": "Tobin", "surname":"Petersen", "birthday":"10/27/2015", "pic":"bo.jpg"}'`
  curl $host/person/add/mom -d "{\"child_id\": \"$ady\", \"mom_id\":\"$jess\"}"
  curl $host/person/add/dad -d "{\"child_id\": \"$ady\", \"dad_id\":\"$ryan\"}"
  curl $host/person/add/mom -d "{\"child_id\": \"$hayden\", \"mom_id\":\"$jess\"}"
  curl $host/person/add/dad -d "{\"child_id\": \"$hayden\", \"dad_id\":\"$ryan\"}"
  curl $host/person/add/mom -d "{\"child_id\": \"$tobin\", \"mom_id\":\"$jess\"}"
  curl $host/person/add/dad -d "{\"child_id\": \"$tobin\", \"dad_id\":\"$ryan\"}"

  cole=`curl $host/person/add -d '{"givennames": "Coleman", "surname":"Petersen", "birthday":"03/31/1988", "occupation": "PR", "pic":"cole.jpg" }'`
  anne=`curl $host/person/add -d '{"givennames": "Anne", "surname":"Petersen", "birthday":"06/14/1989", "pic":"anne.jpg"}'`
  curl $host/person/add/marriage -d "{\"person_id\": \"$cole\", \"spouse_id\":\"$anne\"}"

  jocy=`curl $host/person/add -d '{"givennames": "Jocelyn", "surname":"Petersen", "birthday":"12/01/2015", "pic":"jocy.jpg"}'`

  curl $host/person/add/mom -d "{\"child_id\": \"$jocy\", \"mom_id\":\"$anne\"}"
  curl $host/person/add/dad -d "{\"child_id\": \"$jocy\", \"dad_id\":\"$cole\"}"

  curl $host/person/add/mom -d "{\"child_id\": \"$cole\", \"mom_id\":\"$mom\"}"
  curl $host/person/add/dad -d "{\"child_id\": \"$cole\", \"dad_id\":\"$dad\"}"





} &> /dev/null
echo "person $mark"
