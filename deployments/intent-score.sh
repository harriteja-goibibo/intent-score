#!/usr/bin/env bash
env=$1
INTENT-SCORE_LOG='/logs/intent-score_server.log'

current_pid=`ps -ef | grep intent-score_server.bin | grep "env=$env" | awk '{print $2}'`


function start_server()
{
    current_dir=$pwd
    cd /usr/local/goibibo/intent-score
    nohup ./bin/intent-score_server.bin -env=$env >> $INTENT-SCORE_LOG 2>&1 &
    cd $current_dir
    current_pid=$!
    echo "Intent-Score is started with pid: $current_pid"
}

function stop_server()
{
    echo "Found Intent-Score Process Id: $current_pid"
    kill -9 $current_pid
    echo "Stopped Intent-Score Process Id: $current_pid"

}

case "$2" in
    start)
	   if [ "$current_pid" != "" ];
	   then
		   echo "Intent-Score is already running"
	   else
		   start_server
	   fi
       ;;
    stop)
	   if [ "$current_pid" != "" ];
	   then
		   stop_server
	   else
		   echo "Intent-Score is not running"
	   fi
       ;;
    restart)
	   if [ "$current_pid" != "" ];
	   then
		   stop_server
		   start_server
	   else
		   echo "Intent-Score is not running"
		   start_server
	   fi
       ;;
    status)
       if [ "$current_pid" != "" ];
       then
          echo "Intent-Score is running with Process Id: $current_pid"
       else
          echo "Intent-Score is not running"
       fi
       ;;
esac