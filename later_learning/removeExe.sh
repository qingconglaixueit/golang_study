#!/bin/bash

process(){
	for table in $dir
	do 
		if [ ${table##*.} == "exe" ];then
			echo ${table}
		fi
	done
}

listFiles()
{
	#1st param, the dir name
	for file in `ls $1`;
	do

			if [ -d $1/$file ]; then
				#echo "$1/$file"
				listFiles "$1/$file" 
			else
				tmp=$1/$file
				if [ ${tmp##*.} == "exe" ];then
					rm ${tmp}
					echo "已删除----"${tmp}
				fi
			fi
	done

}
echo -e "\033[33m ===== starting check exe file ...===== \033[0m"
listFiles $1
echo -e "\033[32m ===== check exe file successfully...===== \033[0m"