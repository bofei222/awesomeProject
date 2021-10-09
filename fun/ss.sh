# if no specify ip ,disconect all for this port

if  [ ! -n "$1" ] ;then
	echo "please input one port"
fi

index=0
for line in $( netstat  -ant | grep $1 | awk '{print $5}')
do
  if [ $index != 0 ]
  then
	string=$line
	array=(${string//:/ })  
       if  [ ! -n "$2" ] ;then
        	ss -K dst ${array[0]} dport =  ${array[1]}
       else	
		if [ $2 == ${array[0]} ];then 		 
		       ss -K dst ${array[0]} dport =  ${array[1]}
		fi
        	 
       fi	       
  fi
  ((index++))
done
