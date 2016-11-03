#!/bin/bash

host=localhost:9090
#To get an fbid http://findmyfbid.com/

{ #This hides the output

  curl $host/dump #Clear data

  mark=`curl $host/person/add -d '{"givennames": "Mark", "surname":"Z", "birthday":"05/14/1984", "occupation": "CEO", "fbid":"4" }'`
  spouse=`curl $host/person/add -d '{"givennames": "Priscilla", "surname":"C", "fbid":"140"}'`
  curl $host/person/add/marriage -d "{\"person_id\": \"$mark\", \"spouse_id\":\"$spouse\"}"
  mom=`curl $host/person/add -d '{"givennames": "Momma", "surname":"Z"}'`
  dad=`curl $host/person/add -d '{"givennames": "Papa", "surname":"Z"}'`
  curl $host/person/add/marriage -d "{\"person_id\": \"$mom\", \"spouse_id\":\"$dad\"}"
  curl $host/person/add/mom -d "{\"child_id\": \"$mark\", \"mom_id\":\"$mom\"}"
  curl $host/person/add/dad -d "{\"child_id\": \"$mark\", \"dad_id\":\"$dad\"}"
  baby=`curl $host/person/add -d '{"givennames": "Max", "surname":"Z"}'`
  curl $host/person/add/mom -d "{\"child_id\": \"$baby\", \"mom_id\":\"$spouse\"}"
  curl $host/person/add/dad -d "{\"child_id\": \"$baby\", \"dad_id\":\"$mark\"}"

} &> /dev/null
echo "person $mark"
