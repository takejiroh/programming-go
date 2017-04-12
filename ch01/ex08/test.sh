go build

rm ./kekka.txt
./ex08 gopl.io http://gopl.io > kekka.txt

rm ./added.txt
rm ./dontadd.txt

./ex08 gopl.io > added.txt
cat added.txt | grep "add http:// prefix."

./ex08 http://gopl.io > dontadd.txt
cat dontadd.txt | grep "don't add prefix."





