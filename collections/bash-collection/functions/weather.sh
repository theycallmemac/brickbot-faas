#!/bin/bash
curl "wttr.in/$*" 2>&1 | grep -A 6 "Weather report" | sed '1,2d' |sed "s,\x1B\[[0-9;]*[a-zA-Z],,g"| sed -E 's/^.{15}//' > weather.summary
if [[ $(curl "wttr.in/$*" 2>&1 | grep "Oymyakon") == "so we have brought you to Oymyakon," ]];then
    echo -e "Weather report for Oymyakon:\n"
else
    echo -e "Weather report for $*:\n"
fi
summary=$(head -1 weather.summary)
temp=$(head -2 weather.summary | tail -1)
winds=$(head -3 weather.summary | tail -1)
precipitation=$(head -6 weather.summary | tail -1)
echo -e "Summary: $summary\nTemperature: $temp\nWind Speed and Direction: $winds\nPrecipitation: $precipitation\n"
