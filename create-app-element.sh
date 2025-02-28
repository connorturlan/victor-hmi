#!/bin/bash
# get params
type=$1
app=$2
name=$3
shift
shift

basePath="./apps"

# set the path to either containers components depending on the param.
if [ $type = "component" ] 
then
	localPath="${basePath}/${app}/components"
elif [ $type = "container" ] 
then
	localPath="${basePath}/${app}/containers"
else 
	echo Invalid type $type.
	echo "usage: create-element [type:component^|container] [name]"
	exit 1
fi

# validate the name
if [ -z $name ]
then
	echo Name must be specified.
	echo usage: create-element $type [name]
	exit 1
fi

# create the directory structure.
if [ ! -d $localPath ]
then
	mkdir -p $localPath
fi

if [[ -d $localPath/$name ]]
then
	echo $type $name already exists.
	echo "Overwrite? (y/n): "
	read confirm
	if [ $confirm != "y" ]
	then
		echo Component unmodified, exiting...
		exit 1
	fi
	echo Overwriting component $name.
else
	mkdir -p $localPath/$name
fi

# create the index file.
echo 'import { '$name' } from "'$name'.jsx";
export default '$name';' > $localPath/$name/index.js

# create the scss file.
echo ".$name {}" > $localPath/$name/$name.module.scss 

# create the jsx file.
echo \
'import styles from "./'$name'.module.scss";

function '$name'(props) {
	return <></>;
}

export default '$name';
' > "$localPath/$name/$name.jsx"

# notify the user.
echo Component $name created successfully.
exit 0