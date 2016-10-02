# any codes
#
filename=xxxmicro
base_dir=../batch

echo "[-START--] golang build on ${base_dir} as main package folder."
env GOOS=linux GOARCH=amd64 go build -o $filename >/dev/null

if [[ $? != 0 ]]; then
  echo "[-END----] failed golang build."
  exit 1
fi

echo "[-END----] completed golang build."

for d in $(ls $deploy_dir); do
  echo "[-DEPLOY-] copy file to $deploy_dir/$d"
  cp $filename $deploy_dir/$d
done
# any codes
