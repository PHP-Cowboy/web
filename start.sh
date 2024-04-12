killall web

rm -rf nohup.log

mv web.0 web

chmod +x web

nohup ./web >> nohup.log 2>&1 &

echo "success"
