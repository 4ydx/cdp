for i in */ ; do cd $i; echo "running $i"; go build; cd ..; done
