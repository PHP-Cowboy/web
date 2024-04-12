killall web

rm -rf nohup.log

nohup ./web >> nohup.log 2>&1 &

echo "success"
