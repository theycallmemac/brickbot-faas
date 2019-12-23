curl -s  https://raw.githubusercontent.com/redbrick/brickbot/master/package.json > package.json
echo -e "Author: $(jq -r ".author" package.json)\n"
echo -e "Version: $(jq -r ".version" package.json)\n"
echo -e "Description: $(jq -r ".description" package.json)\n"
echo -e "Code: $(jq -r ".homepage" package.json)\n"
